package parser

var blockNodeBuilders = []BlockNodeBuilder{ParagraphBuilder{}}

func Parse(lines []string) RootNode {
	return parseBlocks(lines)
}

func parseBlocks(lines []string) RootNode {

	root := RootNode{links: make(map[string]string)}

	for i := 0; i < len(lines); {
		for _, builder := range blockNodeBuilders {
			if builder.isValidStart(lines[i]) {
				new_i, node, err := builder.parse(i, lines, root)

				if err == nil {
					i = new_i
					root.children = append(root.children, node)
				}
			} else {
				i++
			}
		}
	}
	return root
}
