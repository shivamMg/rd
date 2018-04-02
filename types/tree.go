package types

// Tree represents a node in a parse tree.
type Tree struct {
	Symbol   string
	Children []*Tree
}

// NewTree returns a node with Symbol set to symbol.
func NewTree(symbol string) *Tree {
	return &Tree{Symbol: symbol}
}

// Add adds subtree as a child to n.
func (t *Tree) Add(tree *Tree) {
	t.Children = append(t.Children, tree)
}
