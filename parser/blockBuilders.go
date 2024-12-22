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

	newChildNode := UnparsedInlineNode{content: lineWithoutPrefix, root: root}

	newNode := HeadingNode{children: []Node{newChildNode}, root: root, headingLevel: headingLevel}
	return lineIndex + 1, newNode, nil
}

type ParagraphBuilder struct {
}

func (builder ParagraphBuilder) isValidStart(s string) bool {
	// As long as the line isnt empty it can be the start of a
	// paragraph (paragraph is the default case)
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

// Check if string is valid fenced code block opening and return the
// length of the opening fence if it is (else -1)
func parseCodeBlockFence(s string, opening bool) (int, int, rune, error) {
	if len(s) < 3 {
		return -1, -1, rune(0), fmt.Errorf("Line shorter than 3 chars")
	}

	var indent int = utils.HowManySpacesDoesLineStartWith(s)

	if indent > 3 {
		return -1, -1, rune(0), fmt.Errorf("Line doesn't start with three or less spaces")
	}
	s = strings.TrimLeft(s, " ")
	openingAsArr := []rune(s)

	c := openingAsArr[0]

	if c != '`' && c != '~' {
		return -1, -1, rune(0), fmt.Errorf("First char isn't ` or ~")
	}

	if openingAsArr[1] != c || openingAsArr[2] != c {
		return -1, -1, rune(0), fmt.Errorf("First 3 chars of fence aren't the same")
	}

	numberCharsInOpeningFence := 1
	for numberCharsInOpeningFence < len(openingAsArr) {
		if openingAsArr[numberCharsInOpeningFence] == c {
			numberCharsInOpeningFence++
		} else {
			break
		}
	}

	if !opening {
		s = strings.TrimRightFunc(s, unicode.IsSpace)
		if len(s) != numberCharsInOpeningFence {
			return -1, -1, rune(0), fmt.Errorf("Closing fence cannot have info strings")
		}
	} else {
		if c == '`' {
			for _, char := range openingAsArr[numberCharsInOpeningFence:] {
				if char == '`' {
					return -1, -1, rune(0), fmt.Errorf("Cannot have ` in info string where ` is the fence char.")

				}
			}
		}
	}
	return numberCharsInOpeningFence, indent, c, nil
}

func parseOpeningCodeBlockFence(s string) (int, int, rune, string, error) {
	i, indentLength, c, err := parseCodeBlockFence(s, true)

	if err != nil {
		return -1, -1, rune(0), "", err
	}

	infoString := s[i:len(s)]

	return i, indentLength, c, infoString, nil
}

func parseClosingCodeBlockFence(s string, openingFenceLength int, fenceChar rune) error {
	i, _, c, err := parseCodeBlockFence(s, false)

	if err != nil {
		return err
	}

	if i < openingFenceLength {
		return fmt.Errorf("Closing fence must be as long or longer than opening fence")
	}

	if c != fenceChar {
		return fmt.Errorf("Closing fence and opening fence must be made of the same char")
	}

	return nil
}

type FencedCodeBuilder struct {
}

func (builder FencedCodeBuilder) isValidStart(s string) bool {
	_, _, _, _, err := parseOpeningCodeBlockFence(s)
	return err == nil
}

func (builder FencedCodeBuilder) parse(lineIndex int, lines []string, root RootNode) (int, BlockNode, error) {
	fenceLength, indentLength, fenceChar, infoString, err := parseOpeningCodeBlockFence(lines[lineIndex])

	if err != nil {
		return -1, nil, fmt.Errorf("Invalid opening fence")
	}

	//closed := false

	closingFenceIndex := lineIndex + 1
	for closingFenceIndex < len(lines) {
		err = parseClosingCodeBlockFence(lines[closingFenceIndex], fenceLength, fenceChar)
		if err == nil {
			//		closed = true
			break
		}
		closingFenceIndex += 1
	}

	contentLines := lines[lineIndex+1 : closingFenceIndex]
	unindentedLines := make([]string, len(contentLines))
	for i, line := range contentLines {
		spaces := utils.HowManySpacesDoesLineStartWith(line)
		unindentedLines[i] = line[min(spaces, indentLength):]
	}

	content := RawTextInlineNode{root: root, content: strings.Join(unindentedLines, "\n")}
	node := CodeNode{root: root, children: []Node{content}, infoString: infoString}
	if false {
		return -1, nil, fmt.Errorf("Code block not closed")
	} else {
		return closingFenceIndex + 1, node, nil
	}
}
