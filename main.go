package main

import (
	"fmt"

	"github.com/billyedmoore/Slowdown/parser"
)

func main() {
	lines := []string{"Line one", "Line two", "", "Second paragraph"}

	root := parser.Parse(lines)

	fmt.Printf("%v\n", root)
}
