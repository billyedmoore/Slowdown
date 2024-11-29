package parser

type Node interface {
	GetChildren() []Node
	IsLeaf() bool
	GetNodeType() string
}

// This means we can easily add requirements to these as and when we
// need to do so.
type BlockNode interface {
	Node
}

type InlineNode interface {
	Node
}

type RootNode struct {
	Node
	children []Node
	links    map[string]string
}

func (node RootNode) GetChildren() []Node {
	return node.children
}

func (node RootNode) IsLeaf() bool {
	return false
}

func (node RootNode) GetNodeType() string {
	return "ROOT"
}

type UnparsedInlineNode struct {
	content string
	root    RootNode
}

func (node UnparsedInlineNode) GetChildren() []Node {
	return make([]Node, 0)
}

func (node UnparsedInlineNode) IsLeaf() bool {
	return false
}

func (node UnparsedInlineNode) GetNodeType() string {
	return "UNPARSED_INLINE"
}

type ParagraphNode struct {
	children []Node
	root     RootNode
}

func (node ParagraphNode) GetChildren() []Node {
	return node.children
}

func (node ParagraphNode) IsLeaf() bool {
	// All children must be InlineBlocks
	return true
}

func (node ParagraphNode) GetNodeType() string {
	return "PARAGRAPH_BLOCK"
}
