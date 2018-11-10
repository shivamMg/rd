package rd

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

type debugStack []*DebugTree

func (ds debugStack) isEmpty() bool {
	return len(ds) == 0
}

func (ds debugStack) peek() *DebugTree {
	return ds[len(ds)-1]
}

func (ds *debugStack) pop() *DebugTree {
	l := len(*ds)
	dt := (*ds)[l-1]
	*ds = (*ds)[:l-1]
	return dt
}

func (ds *debugStack) push(dt *DebugTree) {
	*ds = append(*ds, dt)
}
