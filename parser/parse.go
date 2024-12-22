package parser

var blockNodeBuilders = []BlockNodeBuilder{FencedCodeBuilder{}, ATXHeadingBuilder{}, ThematicBreakBuilder{}, ParagraphBuilder{}}
var inlineNodeBuilders = []InlineNodeBuilder{RawTextBuilder{}}

var paragraphBreakingBuilders = []BlockNodeBuilder{FencedCodeBuilder{}, ATXHeadingBuilder{}, ThematicBreakBuilder{}}

func Parse(lines []string) Node {
	AST := parseBlocks(lines)
	return parseInlines(AST)
}

func parseBlocks(lines []string) RootNode {

	root := RootNode{links: make(map[string]string)}

	for i := 0; i < len(lines); {
		consumed := false
		for _, builder := range blockNodeBuilders {
			if builder.isValidStart(lines[i]) {
				new_i, node, err := builder.parse(i, lines, root)

				if err == nil {
					consumed = true
					i = new_i
					root.children = append(root.children, node)
					break
				}
			}
		}
		if !consumed {
			i++
		}
	}
	return root
}

func parseInlines(node Node) Node {
	parse := !node.AreChildrenBlocks() && !node.IsLeaf()

	var newChildren []Node = make([]Node, 0)
	// Recursively search all the nodes for UNPARSED_INLINE nodes then parse them
	for _, child := range node.GetChildren() {
		if parse {
			if child.GetNodeType() == "UNPARSED_INLINE" {
				newChildren = append(newChildren, parseInline(child.GetContent(), node.GetRoot())...)
			} else {
				newChildren = append(newChildren, child)
			}
		}
		parseInlines(child)
	}

	if parse {
		node.SetChildren(newChildren)
	}
	return node
}

func parseInline(contentAsString string, root RootNode) []Node {
	result := make([]Node, 0)

	content := []rune(contentAsString)
	for i := 0; i < len(content); {
		consumed := false
		for _, builder := range inlineNodeBuilders {
			if builder.isValidStart(content[i]) {
				new_i, node, err := builder.parse(i, string(content), root)
				if err == nil {
					consumed = true
					result = append(result, node)
					i = new_i
					break
				}
			}
		}
		if !consumed {
			i++
		}
	}
	return result
}
