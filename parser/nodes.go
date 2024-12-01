package parser

import "fmt"

type Node interface {
	GetChildren() []Node
	SetChildren([]Node)
	AreChildrenBlocks() bool
	IsLeaf() bool
	GetContent() string // unsure about this because only RAWINLINETEXT should ever have content
	GetNodeType() string
	GetRoot() RootNode
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

func (node RootNode) SetChildren(newChildren []Node) {
	copy(node.children, newChildren)
}

func (node RootNode) IsLeaf() bool {
	return false
}

func (node RootNode) AreChildrenBlocks() bool {
	return true
}

func (node RootNode) GetNodeType() string {
	return "ROOT"
}

func (node RootNode) GetContent() string {
	return ""
}

func (node RootNode) GetRoot() RootNode {
	return node
}

//--INLINE NODES--

type UnparsedInlineNode struct {
	content string
	root    RootNode
}

func (node UnparsedInlineNode) GetContent() string {
	return node.content
}

func (node UnparsedInlineNode) GetChildren() []Node {
	return make([]Node, 0)
}

func (node UnparsedInlineNode) IsLeaf() bool {
	return true
}

func (node UnparsedInlineNode) AreChildrenBlocks() bool {
	return false
}

func (node UnparsedInlineNode) GetNodeType() string {
	return "UNPARSED_INLINE"
}

func (node UnparsedInlineNode) GetRoot() RootNode {
	return node.root
}
func (node UnparsedInlineNode) SetChildren(newChildren []Node) {
	fmt.Printf("TRYING TO SET THE CHILDREN OF A NODE WITHOUT CHILDREN\n")
}

type RawTextInlineNode struct {
	content string
	root    RootNode
}

func (node RawTextInlineNode) GetChildren() []Node {
	return make([]Node, 0)
}

func (node RawTextInlineNode) IsLeaf() bool {
	return true
}

func (node RawTextInlineNode) AreChildrenBlocks() bool {
	return false
}

func (node RawTextInlineNode) GetNodeType() string {
	return "RAW_TEXT_INLINE"
}

func (node RawTextInlineNode) GetContent() string {
	return node.content
}

func (node RawTextInlineNode) GetRoot() RootNode {
	return node.root
}

func (node RawTextInlineNode) SetChildren(newChildren []Node) {
	fmt.Printf("TRYING TO SET THE CHILDREN OF A NODE WITHOUT CHILDREN\n")
}

//--BLOCK NODES

type ParagraphNode struct {
	children []Node
	root     RootNode
}

func (node ParagraphNode) GetChildren() []Node {
	return node.children
}

func (node ParagraphNode) IsLeaf() bool {
	return false
}

func (node ParagraphNode) AreChildrenBlocks() bool {
	return false
}

func (node ParagraphNode) GetNodeType() string {
	return "PARAGRAPH_BLOCK"
}

func (node ParagraphNode) GetContent() string {
	return ""
}

func (node ParagraphNode) GetRoot() RootNode {
	return node.root
}

func (node ParagraphNode) SetChildren(newChildren []Node) {
	copy(node.children, newChildren)
}

type HeadingNode struct {
	children     []Node
	headingLevel int
	root         RootNode
}

func (node HeadingNode) GetChildren() []Node {
	return node.children
}

func (node HeadingNode) IsLeaf() bool {
	return false
}

func (node HeadingNode) AreChildrenBlocks() bool {
	return false
}

func (node HeadingNode) GetNodeType() string {
	return "HEADING_BLOCK"
}

func (node HeadingNode) GetContent() string {
	return ""
}

func (node HeadingNode) GetRoot() RootNode {
	return node.root
}

func (node HeadingNode) SetChildren(newChildren []Node) {
	copy(node.children, newChildren)
}
