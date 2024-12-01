package parser

import (
	"strconv"
	"testing"
)

func TestATXHeadingWithWhitespace(t *testing.T) {
	lines := []string{"  #                    Line one        ", "", "Second paragraph"}
	root := Parse(lines)

	firstBlock := root.GetChildren()[0]
	firstContent := firstBlock.GetChildren()[0]

	if firstContent.GetContent() != "Line one" {
		t.Fatalf("Expected \"Line one\", instead got %v \"%v\"", firstBlock.GetNodeType(), firstContent.GetContent())
	}
}

func TestATXHeadingWithMoreThan3Spaces(t *testing.T) {
	lines := []string{"     #                    Line one        ", "", "Second paragraph"}
	root := Parse(lines)

	firstContent := root.GetChildren()[0].GetChildren()[0]

	if firstContent.GetNodeType() == "HEADING_BLOCK" {
		t.Fatal("Expected a \"PARAGRAPH_BLOCK\" but got a \"HEADING_BLOCK\" instead")
	}

}

func TestATXHeadingWithMoreThanSix(t *testing.T) {
	lines := []string{"####### Line one", "", "Second paragraph"}
	root := Parse(lines)

	firstContent := root.GetChildren()[0].GetChildren()[0]

	if firstContent.GetNodeType() == "HEADING_BLOCK" {
		t.Fatal("Expected a \"PARAGRAPH_BLOCK\" but got a \"HEADING_BLOCK\" instead")
	}

}

func TestATXHeadingBreakingParagraph(t *testing.T) {
	lines := []string{"Line one", "Line two", "# Heading", "Second paragraph"}
	root := Parse(lines)

	firstContent := root.GetChildren()[0].GetChildren()[0].GetContent()

	if firstContent != "Line one\nLine two" {
		t.Fatalf("Expected a \"Line one\\nLine two\" but got a %v instead", strconv.Quote(firstContent))
	}

}
