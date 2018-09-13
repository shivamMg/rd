package rd

import "github.com/shivamMg/ppds/tree"

type Token interface{}

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
	for _, subtree := range t.Subtrees {
		c = append(c, subtree)
	}
	return
}

func NewTree(symbol string, subtrees ...*Tree) *Tree {
	t := Tree{Symbol: symbol}
	for _, subtree := range subtrees {
		if subtree != nil {
			t.Subtrees = append(t.Subtrees, subtree)
		}
	}
	return &t
}

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

func (t *Tree) Sprint() string {
	return tree.SprintHrn(t)
}

func (t *Tree) Print() {
	tree.PrintHrn(t)
}
