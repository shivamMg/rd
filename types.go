package rd

import "github.com/shivamMg/ppds/tree"

// Token represents a token received after tokenization.
type Token interface{}

// Tree is a parse tree node. Symbol can either be a terminal (Token) or a non-terminal
// (see Builder's Enter method). Tokens matched using Builder's Match method or added
// using Builder's Add method, can be retrieved by type asserting Symbol.
// Subtrees are child nodes of the current node.
type Tree struct {
	Symbol   interface{}
	Subtrees []*Tree
}

func NewTree(symbol interface{}, subtrees ...*Tree) *Tree {
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

func (t *Tree) String() string {
	return tree.SprintHrn(t)
}

// DebugTree is a debug tree node. Can be printed to help tracing the
// parsing flow.
type DebugTree struct {
	data     string
	subtrees []*DebugTree
}

func newDebugTree(data string) *DebugTree {
	return &DebugTree{
		data:     data,
		subtrees: []*DebugTree{},
	}
}

func (dt *DebugTree) add(subtree *DebugTree) {
	dt.subtrees = append(dt.subtrees, subtree)
}

func (dt *DebugTree) Data() interface{} {
	return dt.data
}

func (dt *DebugTree) Children() (c []tree.Node) {
	for _, child := range dt.subtrees {
		c = append(c, child)
	}
	return
}

func (dt *DebugTree) String() string {
	return tree.SprintHrn(dt)
}
