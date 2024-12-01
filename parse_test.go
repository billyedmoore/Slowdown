package main

import (
	"strconv"
	"testing"

	"github.com/billyedmoore/Slowdown/parser"
)

func TestParagraphSplitting(t *testing.T) {
	lines := []string{"Line one", "Line two", "", "Second paragraph"}

	root := parser.Parse(lines)

	if len(root.GetChildren()) != 2 {
		t.Fatalf("Should be 2 children instead found %d", len(root.GetChildren()))
	}
}

func TestEmptyFileParsing(t *testing.T) {
	lines := []string{}

	root := parser.Parse(lines)

	if len(root.GetChildren()) != 0 {
		t.Fatalf("Should be 0 children instead found %d", len(root.GetChildren()))
	}
}

func TestAllInlineBlocksBeingParsed(t *testing.T) {
	lines := []string{"#Top Level", "Line one", "Line two", "", "Second paragraph"}
	root := parser.Parse(lines)

	var queue []parser.Node = make([]parser.Node, 0)

	queue = append(queue, root)

	var current_node parser.Node = nil
	for len(queue) > 0 {
		current_node, queue = queue[0], queue[1:]
		queue = append(queue, current_node.GetChildren()...)
		if current_node.GetNodeType() == "UNPARSED_INLINE" {
			t.Fatalf("Found an unparsed inline block after parsing")
		}
	}
}

func TestRawTextInParagraph(t *testing.T) {
	lines := []string{"Line one", "", "Second paragraph"}
	root := parser.Parse(lines)
	// should be root.paragraph.rawText
	firstContent := root.GetChildren()[0].GetChildren()[0].GetContent()

	if firstContent != "Line one" {
		t.Fatalf("Expected \"Line one\", instead got \"%v\"", firstContent)
	}
}

func TestATXHeadingWithWhitespace(t *testing.T) {
	lines := []string{"  #                    Line one        ", "", "Second paragraph"}
	root := parser.Parse(lines)

	firstBlock := root.GetChildren()[0]
	firstContent := firstBlock.GetChildren()[0]

	if firstContent.GetContent() != "Line one" {
		t.Fatalf("Expected \"Line one\", instead got %v \"%v\"", firstBlock.GetNodeType(), firstContent.GetContent())
	}
}

func TestATXHeadingWithMoreThan3Spaces(t *testing.T) {
	lines := []string{"     #                    Line one        ", "", "Second paragraph"}
	root := parser.Parse(lines)

	firstContent := root.GetChildren()[0].GetChildren()[0]

	if firstContent.GetNodeType() == "HEADING_BLOCK" {
		t.Fatal("Expected a \"PARAGRAPH_BLOCK\" but got a \"HEADING_BLOCK\" instead")
	}

}

func TestATXHeadingWithMoreThanSix(t *testing.T) {
	lines := []string{"####### Line one", "", "Second paragraph"}
	root := parser.Parse(lines)

	firstContent := root.GetChildren()[0].GetChildren()[0]

	if firstContent.GetNodeType() == "HEADING_BLOCK" {
		t.Fatal("Expected a \"PARAGRAPH_BLOCK\" but got a \"HEADING_BLOCK\" instead")
	}

}

func TestATXHeadingBreakingParagraph(t *testing.T) {
	lines := []string{"Line one", "Line two", "# Heading", "Second paragraph"}
	root := parser.Parse(lines)

	firstContent := root.GetChildren()[0].GetChildren()[0].GetContent()

	if firstContent != "Line one\nLine two" {
		t.Fatalf("Expected a \"Line one\\nLine two\" but got a %v instead", strconv.Quote(firstContent))
	}

}
