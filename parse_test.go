package main

import (
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
