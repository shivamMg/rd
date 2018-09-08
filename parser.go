package rd

import (
	"fmt"
	"log"
	"os"
	"github.com/shivamMg/ppds/tree"
)


type P interface {
	// returns false if no tokens left to match
	Match(token Token) (ok bool)
	// ok is false if no token left
	Next() (token Token, ok bool)
	// panics if no node to attach token (empty stack)
	// always returns true
	Add(token Token)
	Reset()
	Enter(nonTerm string)
	Exit(result *bool)
	Tree() *Tree
}

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

// Parser represents a Recursive Descent parser.
type Parser struct {
	tokens  []Token
	st      stack
	current int
	flowTreeStack []*flowTree
}

// NewParser returns a new Parser to parse tokens. Production functions
// for non-terminals must be added using Rule method.
func NewParser(tokens []Token, log bool) *Parser {
	return &Parser{
		tokens:  tokens,
		st:      stack{},
		current: -1,
	}
}

func (p *Parser) PrintFlowTree() {
	tree.PrintHrn(p.flowTreeStack[0])
}

// NextToken returns next token from after incrementing the
// current index. bool signifies if tokens are finished.
func (p *Parser) Next() (token Token, ok bool) {
	if p.current == len(p.tokens)-1 {
		return nil, false
	}
	p.current++
	return p.tokens[p.current], true
}

func (p *Parser) Reset() {
	e, ok := p.st.peek()
	if !ok {
		panic("can't reset")
	}
	p.current = e.index
	e.nonTerm.Subtrees = []*Tree{}
}

// Add adds terminal token term to the non-terminal that is
// being expanded.
func (p *Parser) Add(token Token) {
	e, ok := p.st.peek()
	if !ok {
		panic("no non-terminal to attach to")
	}
	e.nonTerm.Add(NewTree(fmt.Sprint(token)))
}

func (p *Parser) Match(token Token) (ok bool) {
	next, ok := p.Next()
	if !ok {
		ft := p.flowTreeStack[len(p.flowTreeStack)-1]
		data := fmt.Sprintf("%v ≠ <no tokens left>", token)
		ft.children = append(ft.children, NewFlowTree(data))
		return false
	}
	if token != next {
		p.current--
		ft := p.flowTreeStack[len(p.flowTreeStack)-1]
		data := fmt.Sprintf("%v ≠ %v", next, token)
		ft.children = append(ft.children, NewFlowTree(data))
		return false
	}
	p.Add(token)
	ft := p.flowTreeStack[len(p.flowTreeStack)-1]
	data := fmt.Sprintf("%v = %v", next, token)
	ft.children = append(ft.children, NewFlowTree(data))
	return true
}

func (p *Parser) Enter(nonTerm string) {
	t := NewTree(nonTerm)
	p.st.push(ele{
		index:   p.current,
		nonTerm: t,
	})
	ft := NewFlowTree(nonTerm)
	p.flowTreeStack = append(p.flowTreeStack, ft)
}

func (p *Parser) Exit(result *bool) {
	if result == nil {
		panic("result cannot be nil")
	}
	var e ele
	var ok bool
	// don't pop root
	if len(p.st) > 1 {
		e, ok = p.st.pop()
		if !ok {
			panic("nothing to pop")
		}
	}

	ft := p.flowTreeStack[len(p.flowTreeStack)-1]
	ft.data += fmt.Sprintf("(%t)", *result)
	if len(p.flowTreeStack) > 1 {
		p.flowTreeStack = p.flowTreeStack[:len(p.flowTreeStack)-1]
		last := p.flowTreeStack[len(p.flowTreeStack)-1]
		last.children = append(last.children, ft)
	}

	if !*result {
		p.current = e.index
	} else if parent, ok := p.st.peek(); ok && e.nonTerm != nil {
		parent.nonTerm.Add(e.nonTerm)
	}
}

// Tree retrieves the parse tree for the last production.
func (p Parser) Tree() *Tree {
	if e, ok := p.st.peek(); ok {
		return e.nonTerm
	}
	return nil
}

func (p *Parser) PrintTree() {
	tree.PrintHrn(p.Tree())
}
