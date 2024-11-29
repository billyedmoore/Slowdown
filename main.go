package main

import (
	"github.com/billyedmoore/Slowdown/parser"
)

func main() {
	lines := []string{"Line one", "Line two", "", "Second paragraph"}

	root := parser.Parse(lines)
}
