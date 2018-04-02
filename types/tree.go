package types

// Tree represents a parse tree.
type Tree struct {
	Symbol   string
	Children []*Tree
}

// NewTree returns a Tree with Symbol set to symbol, and adds
// subtrees as it's children.
func NewTree(symbol string, subtrees ...*Tree) *Tree {
	t := Tree{Symbol: symbol}
	for _, subtree := range subtrees {
		if subtree != nil {
			t.Children = append(t.Children, subtree)
		}
	}
	return &t
}

// Add adds tree as a child to t.
func (t *Tree) Add(tree *Tree) {
	t.Children = append(t.Children, tree)
}
