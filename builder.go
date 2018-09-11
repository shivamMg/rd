package rd

import (
	"fmt"
	"github.com/shivamMg/ppds/tree"
)

// derivation tree
type derivTree struct {
	data     string
	subtrees []*derivTree
}

func newDerivTree(data string) *derivTree {
	return &derivTree{
		data:     data,
		subtrees: []*derivTree{},
	}
}

func (dt *derivTree) add(subtree *derivTree) {
	dt.subtrees = append(dt.subtrees, subtree)
}

func (dt *derivTree) Data() interface{} {
	return dt.data
}

func (dt *derivTree) Children() (c []tree.Node) {
	for _, child := range dt.subtrees {
		c = append(c, child)
	}
	return
}

type derivStack []*derivTree

func (ds derivStack) peek() (*derivTree, bool) {
	l := len(ds)
	if l == 0 {
		return nil, false
	}
	return ds[l-1], true
}

func (ds *derivStack) pop() (*derivTree, bool) {
	l := len(*ds)
	if l == 0 {
		return nil, false
	}
	dt := (*ds)[l-1]
	*ds = (*ds)[:l-1]
	return dt, true
}

func (ds *derivStack) push(dt *derivTree) {
	*ds = append(*ds, dt)
}

type ele struct {
	index   int
	nonTerm *Tree
}

type stack []ele

func (st stack) peek() (ele, bool) {
	l := len(st)
	if l == 0 {
		return ele{}, false
	}
	return st[l-1], true
}

func (st *stack) pop() (ele, bool) {
	l := len(*st)
	if l == 0 {
		return ele{}, false
	}
	e := (*st)[l-1]
	*st = (*st)[:l-1]
	return e, true
}

func (st *stack) push(e ele) {
	*st = append(*st, e)
}

type Builder struct {
	tokens     []Token
	current    int
	st         stack
	finalEle   ele
	derivStack derivStack
	finalDerivTree *derivTree
}

func NewBuilder(tokens []Token) *Builder {
	return &Builder{
		tokens:  tokens,
		current: -1,
		st:      stack{},
		derivStack: derivStack{},
	}
}

func (b *Builder) PrintDerivation() {
	tree.PrintHrn(b.finalDerivTree)
}

func (b *Builder) Next() (token Token, ok bool) {
	if b.current == len(b.tokens)-1 {
		return nil, false
	}
	b.current++
	return b.tokens[b.current], true
}

func (b *Builder) Reset() {
	e, ok := b.st.peek()
	if !ok {
		panic("cannot reset. didn't enter any non-terminal")
	}
	b.current = e.index
	e.nonTerm.Subtrees = []*Tree{}
}

func (b *Builder) Add(token Token) {
	e, ok := b.st.peek()
	if !ok {
		panic("cannot add. didn't enter any non-terminal")
	}
	e.nonTerm.Add(NewTree(fmt.Sprint(token)))
}

func (b *Builder) Match(token Token) (ok bool) {
	next, ok := b.Next()
	if !ok {
		if dt, ok := b.derivStack.peek(); ok {
			dt.add(newDerivTree(fmt.Sprint("<no tokens left> ≠ ", token)))
		}
		return false
	}
	if token != next {
		b.current--
		if dt, ok := b.derivStack.peek(); ok {
			dt.add(newDerivTree(fmt.Sprint(next, " ≠ ", token)))
		}
		return false
	}
	b.Add(token)
	if dt, ok := b.derivStack.peek(); ok {
		dt.add(newDerivTree(fmt.Sprint(next, " = ", token)))
	}
	return true
}

func (b *Builder) Enter(nonTerm string) {
	b.st.push(ele{
		index:   b.current,
		nonTerm: NewTree(nonTerm),
	})
	b.derivStack.push(newDerivTree(nonTerm))
}

func (b *Builder) Exit(result *bool) {
	if result == nil {
		panic("result cannot be nil")
	}
	if e, ok := b.st.pop(); ok {
		if *result {
			if parent, ok := b.st.peek(); ok {
				parent.nonTerm.Add(e.nonTerm)
			} else {
				// no eles left
				b.finalEle = e
			}
		} else {
			b.current = e.index
		}
	}

	if dt, ok := b.derivStack.peek(); ok {
		dt.data += fmt.Sprintf("(%t)", *result)
	}
	if dt, ok := b.derivStack.pop(); ok {
		if parent, ok := b.derivStack.peek(); ok {
			parent.add(dt)
		} else {
			b.finalDerivTree = dt
		}
	}
}

// Tree retrieves the parse tree for the last production.
func (b Builder) Tree() *Tree {
	return b.finalEle.nonTerm
}
