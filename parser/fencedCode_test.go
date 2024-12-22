package parser

import (
	"testing"
)

func TestFencedCodeBlockPresence(t *testing.T) {
	lines := []string{"````", "def main():", "	return \"hello\"", "````"}
	root := Parse(lines)

	firstBlock := root.GetChildren()[0]

	if firstBlock.GetNodeType() != "CODE_BLOCK" {
		t.Fatalf("Expected \"CODE_BLOCK\", instead got \"%v\"", firstBlock.GetNodeType())
	}
}

func TestFencedCodeBlockDifferentOpeningAndClosing(t *testing.T) {
	lines := []string{"````", "~~~~"}
	root := Parse(lines)

	firstBlock := root.GetChildren()[0]
	textBlock := firstBlock.GetChildren()[0]

	if firstBlock.GetNodeType() != "CODE_BLOCK" || textBlock.GetContent() != "~~~~" {
		t.Fatalf("Expected \"CODE_BLOCK\" with content \"~~~~\", instead got \"%v\" with content \"%v\"",
			firstBlock.GetNodeType(), textBlock.GetContent())
	}
}

func TestFencedCodeBlockNoClosing(t *testing.T) {
	lines := []string{"````", "print(\"Hello World\")"}
	root := Parse(lines)

	firstBlock := root.GetChildren()[0]
	textBlock := firstBlock.GetChildren()[0]

	if firstBlock.GetNodeType() != "CODE_BLOCK" || textBlock.GetContent() != "print(\"Hello World\")" {
		t.Fatalf("Expected \"CODE_BLOCK\" with content \"print(\"Hello World\")\", instead got \"%v\" with content \"%v\"",
			firstBlock.GetNodeType(), textBlock.GetContent())
	}
}

func TestFencedCodeBlockClosingWithTooManySpaces(t *testing.T) {
	lines := []string{"````", "    ````"}
	root := Parse(lines)

	firstBlock := root.GetChildren()[0]
	textBlock := firstBlock.GetChildren()[0]

	if firstBlock.GetNodeType() != "CODE_BLOCK" || textBlock.GetContent() != "    ````" {
		t.Fatalf("Expected \"CODE_BLOCK\" with content \"    ````\", instead got \"%v\" with content \"%v\"",
			firstBlock.GetNodeType(), textBlock.GetContent())
	}
}

func TestFencedCodeBlockClosingWithInfoString(t *testing.T) {
	lines := []string{"````", "```` info", "````"}
	root := Parse(lines)

	firstBlock := root.GetChildren()[0]
	textBlock := firstBlock.GetChildren()[0]

	if firstBlock.GetNodeType() != "CODE_BLOCK" || textBlock.GetContent() != "```` info" {
		t.Fatalf("Expected \"CODE_BLOCK\" with content \"```` info\", instead got \"%v\" with content \"%v\"",
			firstBlock.GetNodeType(), textBlock.GetContent())
	}
}

func TestFencedCodeBlockWithSpaceInFence(t *testing.T) {
	lines := []string{"````", "```` ``"}
	root := Parse(lines)

	firstBlock := root.GetChildren()[0]
	textBlock := firstBlock.GetChildren()[0]

	if firstBlock.GetNodeType() != "CODE_BLOCK" || textBlock.GetContent() != "```` ``" {
		t.Fatalf("Expected \"CODE_BLOCK\" with content \"```` ``\", instead got \"%v\" with content \"%v\"",
			firstBlock.GetNodeType(), textBlock.GetContent())
	}
}

func TestFencedCodeBlockContent(t *testing.T) {
	lines := []string{"````", "this is content on this row", "````"}
	root := Parse(lines)

	firstBlock := root.GetChildren()[0]
	firstBlock = firstBlock.GetChildren()[0]

	if firstBlock.GetContent() != "this is content on this row" {
		t.Fatalf("Expected \"this is content on this row\", instead got \"%v\"", firstBlock.GetContent())
	}
}

func TestFencedCodeBlockUnindenting(t *testing.T) {
	lines := []string{"   ````", "   this is content on this row", "````"}
	root := Parse(lines)

	firstBlock := root.GetChildren()[0]
	firstBlock = firstBlock.GetChildren()[0]

	if firstBlock.GetContent() != "this is content on this row" {
		t.Fatalf("Expected \"this is content on this row\", instead got \"%v\"", firstBlock.GetContent())
	}
}

func TestFencedCodeBlockUnindentingNonComplete(t *testing.T) {
	lines := []string{"   ````", "    this is content on this row", "````"}
	root := Parse(lines)

	firstBlock := root.GetChildren()[0]
	firstBlock = firstBlock.GetChildren()[0]

	if firstBlock.GetContent() != " this is content on this row" {
		t.Fatalf("Expected \" this is content on this row\", instead got \"%v\"", firstBlock.GetContent())
	}
}
