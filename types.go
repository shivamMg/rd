package rd

import "github.com/shivamMg/ppds/tree"

// Token represents a token received after tokenization.
type Token interface{}

// Tree represents a parse tree node. Symbol is either a terminal or a
// non-terminal. Subtrees are child nodes of the current node.
type Tree struct {
	Symbol   string
	Subtrees []*Tree
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

// Add adds a subtree as a child to t.
func (t *Tree) Add(subtree *Tree) {
	t.Subtrees = append(t.Subtrees, subtree)
}

// Detach removes a subtree as a child of t.
func (t *Tree) Detach(subtree *Tree) {
	for i, st := range t.Subtrees {
		if st == subtree {
			t.Subtrees = append(t.Subtrees[:i], t.Subtrees[i+1:]...)
			break
		}
	}
}

// Sprint returns formatted tree.
func (t *Tree) Sprint() string {
	return tree.SprintHrn(t)
}

// Print prints formatted tree to STDOUT.
func (t *Tree) Print() {
	tree.PrintHrn(t)
}
