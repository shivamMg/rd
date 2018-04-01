package main

// Node represents a node in a parse tree.
type Node struct {
	Symbol   string
	Children []*Node
}

// NewNode returns a node with Symbol set to symbol.
func NewNode(symbol string) *Node {
	return &Node{Symbol: symbol}
}

// Add adds node as a child node to n.
func (n *Node) Add(node *Node) {
	n.Children = append(n.Children, node)
}

// Tree represents a parse tree.
type Tree struct {
	Root *Node
}
