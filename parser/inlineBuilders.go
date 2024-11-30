package parser

import (
	"unicode"
)

type InlineNodeBuilder interface {
	isValidStart(rune) bool
	parse(int, string, RootNode) (int, InlineNode, error)
}

type RawTextBuilder struct{}

func (builder RawTextBuilder) isValidStart(c rune) bool {
	return !unicode.IsSpace(c)
}

func (builder RawTextBuilder) parse(start int, s string, root RootNode) (int, InlineNode, error) {
	prev_char_new_line := false
	end := -1
	for i, c := range s[start:] {
		if c == '\n' {
			if prev_char_new_line {
				end = i - 1
				break
			} else {
				prev_char_new_line = true
			}
		} else {
			prev_char_new_line = false
		}
	}

	if end == -1 {
		end = len(s)
	}

	newNode := RawTextInlineNode{content: s[start:end], root: root}

	return end, newNode, nil
}
