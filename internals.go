package rd

import "github.com/shivamMg/ppds/tree"

type ele struct {
	index   int
	nonTerm *Tree
}

type stack []ele

func (st stack) isEmpty() bool {
	return len(st) == 0
}

func (st stack) peek() ele {
	return st[len(st)-1]
}

func (st *stack) pop() ele {
	l := len(*st)
	e := (*st)[l-1]
	*st = (*st)[:l-1]
	return e
}

func (st *stack) push(e ele) {
	*st = append(*st, e)
}

type debugTree struct {
	data     string
	subtrees []*debugTree
}

func newDebugTree(data string) *debugTree {
	return &debugTree{
		data:     data,
		subtrees: []*debugTree{},
	}
}

func (dt *debugTree) add(subtree *debugTree) {
	dt.subtrees = append(dt.subtrees, subtree)
}

func (dt *debugTree) Data() interface{} {
	return dt.data
}

func (dt *debugTree) Children() (c []tree.Node) {
	for _, child := range dt.subtrees {
		c = append(c, child)
	}
	return
}

func (dt *debugTree) Print() {
	tree.PrintHrn(dt)
}

func (dt *debugTree) Sprint() string {
	return tree.SprintHrn(dt)
}

type debugStack []*debugTree

func (ds debugStack) isEmpty() bool {
	return len(ds) == 0
}

func (ds debugStack) peek() *debugTree {
	return ds[len(ds)-1]
}

func (ds *debugStack) pop() *debugTree {
	l := len(*ds)
	dt := (*ds)[l-1]
	*ds = (*ds)[:l-1]
	return dt
}

func (ds *debugStack) push(dt *debugTree) {
	*ds = append(*ds, dt)
}
