package rd

import (
	"fmt"
	"log"
	"os"
	"github.com/shivamMg/ppds/tree"
)

type flowTree struct {
	data string
	children []*flowTree
}

func NewFlowTree(data string) *flowTree {
	return &flowTree{
		data: data,
		children: []*flowTree{},
	}
}

func (ft *flowTree) Data() interface{} {
	return ft.data
}

func (ft *flowTree) Children() (c []tree.Node) {
	for _, child := range ft.children {
		c = append(c, child)
	}
	return
}


var logger = log.New(os.Stdout, "", 0)

type ele struct {
	index   int
	nonTerm *Tree
}

// stack stores trees as non-terminals are expanded. index stores
// the tokens' index for which the nonTerm was expanded. The last
// ele is the non-terminal that is currently being derived.
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

// Builder represents a Recursive Descent parser.
type Builder struct {
	tokens  []Token
	st      stack
	current int
	flowTreeStack []*flowTree
}

// NewBuilder returns a new Builder to parse tokens. Production functions
// for non-terminals must be added using Rule method.
func NewBuilder(tokens []Token) *Builder {
	return &Builder{
		tokens:  tokens,
		st:      stack{},
		current: -1,
	}
}

func (b *Builder) PrintFlowTree() {
	tree.PrintHrn(b.flowTreeStack[0])
}

// NextToken returns next token from after incrementing the
// current index. bool signifies if tokens are finished.
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
		panic("can't reset")
	}
	b.current = e.index
	e.nonTerm.Subtrees = []*Tree{}
}

// Add adds terminal token term to the non-terminal that is
// being expanded.
func (b *Builder) Add(token Token) {
	e, ok := b.st.peek()
	if !ok {
		panic("no non-terminal to attach to")
	}
	e.nonTerm.Add(NewTree(fmt.Sprint(token)))
}

func (b *Builder) Match(token Token) (ok bool) {
	next, ok := b.Next()
	if !ok {
		ft := b.flowTreeStack[len(b.flowTreeStack)-1]
		data := fmt.Sprintf("%v ≠ <no tokens left>", token)
		ft.children = append(ft.children, NewFlowTree(data))
		return false
	}
	if token != next {
		b.current--
		ft := b.flowTreeStack[len(b.flowTreeStack)-1]
		data := fmt.Sprintf("%v ≠ %v", next, token)
		ft.children = append(ft.children, NewFlowTree(data))
		return false
	}
	b.Add(token)
	ft := b.flowTreeStack[len(b.flowTreeStack)-1]
	data := fmt.Sprintf("%v = %v", next, token)
	ft.children = append(ft.children, NewFlowTree(data))
	return true
}

func (b *Builder) Enter(nonTerm string) {
	t := NewTree(nonTerm)
	b.st.push(ele{
		index:   b.current,
		nonTerm: t,
	})
	ft := NewFlowTree(nonTerm)
	b.flowTreeStack = append(b.flowTreeStack, ft)
}

func (b *Builder) Exit(result *bool) {
	if result == nil {
		panic("result cannot be nil")
	}
	var e ele
	var ok bool
	// don't pop root
	if len(b.st) > 1 {
		e, ok = b.st.pop()
		if !ok {
			panic("nothing to pop")
		}
	}

	ft := b.flowTreeStack[len(b.flowTreeStack)-1]
	ft.data += fmt.Sprintf("(%t)", *result)
	if len(b.flowTreeStack) > 1 {
		b.flowTreeStack = b.flowTreeStack[:len(b.flowTreeStack)-1]
		last := b.flowTreeStack[len(b.flowTreeStack)-1]
		last.children = append(last.children, ft)
	}

	if !*result {
		b.current = e.index
	} else if parent, ok := b.st.peek(); ok && e.nonTerm != nil {
		parent.nonTerm.Add(e.nonTerm)
	}
}

// Tree retrieves the parse tree for the last production.
func (b Builder) Tree() *Tree {
	if e, ok := b.st.peek(); ok {
		return e.nonTerm
	}
	return nil
}

