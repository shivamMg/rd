package rd

import "github.com/shivamMg/ppds/tree"

// Tree represents a parse tree.
type Tree struct {
	Symbol   string
	Subtrees []*Tree
}

func (t *Tree) Data() interface{} {
	if t == nil {
		return ""
	}
	return t.Symbol
}

func (t *Tree) Children() (c []tree.Node) {
	if t == nil {
		return
	}
	for _, subtree := range t.Subtrees {
		c = append(c, subtree)
	}
	return
}

// NewTree returns a Tree with Symbol set to symbol, and adds
// subtrees as it's children.
func NewTree(symbol string, subtrees ...*Tree) *Tree {
	t := Tree{Symbol: symbol}
	for _, subtree := range subtrees {
		if subtree != nil {
			t.Subtrees = append(t.Subtrees, subtree)
		}
	}
	return &t
}

// Add adds subtree to t.
func (t *Tree) Add(subtree *Tree) {
	t.Subtrees = append(t.Subtrees, subtree)
}

func (t *Tree) Detach(newSubtree *Tree) {
	for i, subtree := range t.Subtrees {
		if subtree == newSubtree {
			t.Subtrees = append(t.Subtrees[:i], t.Subtrees[i+1:]...)
			break
		}
	}
}
