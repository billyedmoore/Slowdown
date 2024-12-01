package parser

import (
	"fmt"
	"strings"
	"unicode"
)

type BlockNodeBuilder interface {
	isValidStart(string) bool
	parse(int, []string, RootNode) (int, BlockNode, error)
}

type ATXHeadingBuilder struct {
}

func (builder ATXHeadingBuilder) isValidStart(s string) bool {
	// Pretty quick and dirty check if it sort of looks like an ATX header
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
	for i := start; i < len(lines); i++ {
		atxBuilder := ATXHeadingBuilder{}
		if atxBuilder.isValidStart(lines[i]) {
			// This doesn't matter for atxHeadings but when other blocks that can
			// break up paragraphs are here isValidStart won't mean it will sucessfully
			// parse
			_, _, err := atxBuilder.parse(i, lines, root)
			if err == nil {
				end = i
				break
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
