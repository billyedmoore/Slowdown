package parser

import "strings"

type BlockNodeBuilder interface {
	isValidStart(string) bool
	parse(int, []string, RootNode) (int, BlockNode, error)
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
