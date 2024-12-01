package main

import (
	"testing"

	"github.com/billyedmoore/Slowdown/parser"
)

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
