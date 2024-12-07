package parser

import "testing"

func TestParagraphSplitting(t *testing.T) {
	lines := []string{"Line one", "Line two", "", "Second paragraph"}

	root := Parse(lines)

	if len(root.GetChildren()) != 2 {
		t.Fatalf("Should be 2 children instead found %d", len(root.GetChildren()))
	}
}

func TestRawTextInParagraph(t *testing.T) {
	lines := []string{"Line one", "", "Second paragraph"}
	root := Parse(lines)
	// should be root.paragraph.rawText
	firstContent := root.GetChildren()[0].GetChildren()[0].GetContent()

	if firstContent != "Line one" {
		t.Fatalf("Expected \"Line one\", instead got \"%v\"", firstContent)
	}
}
