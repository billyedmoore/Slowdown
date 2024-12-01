package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/billyedmoore/Slowdown/parser"
)

func main() {
	lines := []string{"Line one", "Line two", "# Heading", "Second paragraph"}
	//lines := []string{"  #                    Line one        ", "", "Second paragraph"}

	root := parser.Parse(lines)
	traverse(root, 0)
}

func traverse(node parser.Node, depth int) {
	print(strings.Repeat("\t", depth))

	fmt.Printf("Node: %v ", node.GetNodeType())
	if len(node.GetContent()) > 0 {
		fmt.Printf("Content: %v", strconv.Quote(node.GetContent()))
	}
	print("\n")

	for _, child := range node.GetChildren() {
		traverse(child, depth+1)
	}
}
