package parser

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/billyedmoore/Slowdown/utils"
)

type BlockNodeBuilder interface {
	isValidStart(string) bool
	parse(int, []string, RootNode) (int, BlockNode, error)
}

type ATXHeadingBuilder struct {
}

func (builder ATXHeadingBuilder) isValidStart(s string) bool {
	allowedSpaces := 3
	headingLevel := 0

	for _, c := range s {
		if unicode.IsSpace(c) {
			if headingLevel > 0 {
				return true
			}
			if allowedSpaces > 0 {
				allowedSpaces -= 1
			} else {
				return false
			}
		}

		if c == '#' {
			allowedSpaces = 0
			headingLevel += 1
			if headingLevel > 6 {
				return false
			}
		}
	}

	return false
}

func (builder ATXHeadingBuilder) parse(lineIndex int, lines []string, root RootNode) (int, BlockNode, error) {
	allowedSpaces := 3
	headingLevel := 0

	contentStartIndex := -1

	// Read in the #s before the content
	for characterIndex, c := range lines[lineIndex] {
		if unicode.IsSpace(c) {
			if headingLevel > 0 {
				contentStartIndex = characterIndex
				break
			}
			if allowedSpaces > 0 {
				allowedSpaces -= 1
			} else {
				return -1, nil, fmt.Errorf("Couldn't make header, line starts with more than 3 spaces.")
			}
		}

		if c == '#' {
			headingLevel++
			if headingLevel > 6 {
				return -1, nil, fmt.Errorf("Couldn't make header, header level should be 6 or less.")
			}
		}
	}

	if headingLevel == 0 {
		return -1, nil, fmt.Errorf("Couldn't make header, there isn't a good start here.")
	}

	lineWithoutPrefix := lines[lineIndex][contentStartIndex:]
	lineWithoutPrefix = strings.TrimSpace(lineWithoutPrefix)

	for []rune(lineWithoutPrefix)[len(lineWithoutPrefix)-1] == '#' {
		lineWithoutPrefix = lineWithoutPrefix[:len(lineWithoutPrefix)-1]
	}

	lineWithoutPrefix = strings.TrimSpace(lineWithoutPrefix)

	// UNTESTED WHETHER THIS WORKS

	newChildNode := UnparsedInlineNode{content: lineWithoutPrefix, root: root}

	newNode := HeadingNode{children: []Node{newChildNode}, root: root, headingLevel: headingLevel}
	return lineIndex + 1, newNode, nil
}

type ParagraphBuilder struct {
}

func (builder ParagraphBuilder) isValidStart(s string) bool {
	// As long as the line isnt empty it can be the start of a
	// paragraph (paragraph is the most simple case)
	return (len(s) > 0)
}

func (builder ParagraphBuilder) parse(start int, lines []string, root RootNode) (int, BlockNode, error) {
	var end int = -1

	broken := false
	for i := start; i < len(lines) && !broken; i++ {
		for _, builder := range paragraphBreakingBuilders {
			if broken {
				break
			}
			if builder.isValidStart(lines[i]) {
				// Check it doesn't fail to parse
				_, _, err := builder.parse(i, lines, root)
				if err == nil {
					end = i
					broken = true
				} else {
				}
			}
		}
		if len(lines[i]) == 0 {
			end = i
			break
		}
	}

	if end == -1 {
		end = len(lines)
	}

	newChildContent := strings.Join(lines[start:end], "\n")
	newChildNode := UnparsedInlineNode{content: newChildContent, root: root}

	newNode := ParagraphNode{children: []Node{newChildNode}, root: root}

	return end, newNode, nil
}

type ThematicBreakBuilder struct {
}

func (builder ThematicBreakBuilder) isValidStart(s string) bool {
	allowedSpaces := 3

	// Check starts with 3 or less spaces
	for _, c := range s {
		if unicode.IsSpace(c) {
			if allowedSpaces > 0 {
				allowedSpaces -= 1
			} else {
				return false
			}
		} else {
			break
		}
	}

	s = utils.RemoveWhitespace(s)

	var c rune = '!'
	count := 0

	for _, char := range []rune(s) {
		if c == '!' {
			c = char
		}

		if c != char {
			return false
		} else {
			count += 1
		}
	}

	return count >= 3 && (c == '-' || c == '_' || c == '*')

}

func (builder ThematicBreakBuilder) parse(lineIndex int, lines []string, root RootNode) (int, BlockNode, error) {
	if builder.isValidStart(lines[lineIndex]) {
		return lineIndex + 1, ThematicBreakNode{root: root}, nil
	} else {
		return -1, nil, fmt.Errorf("Couldn't make thematic break.")
	}
}
