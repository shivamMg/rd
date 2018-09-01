package rd

import "fmt"

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
}

// NewParser returns a new Parser to parse tokens. Production functions
// for non-terminals must be added using Rule method.
func NewParser(tokens []Token) *Parser {
	return &Parser{
		tokens:  tokens,
		st:      stack{},
		current: -1,
	}
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
		return false
	}
	if token != next {
		p.current--
		return false
	}
	p.Add(token)
	return true
}

func (p *Parser) Enter(nonTerm string) {
	fmt.Println("Enter", nonTerm)
	t := NewTree(nonTerm)
	p.st.push(ele{
		index: p.current,
		nonTerm: t,
	})
}

func (p *Parser) Exit(result *bool) {
	if result == nil {
		panic("result cannot be nil")
	}
	e, ok := p.st.pop()
	if !ok {
		panic("nothing to pop")
	}
	if !*result {
		p.current = e.index
	} else {
		parent, ok := p.st.peek()
		if !ok {
			panic("nothing to peek")
		}
		parent.nonTerm.Add(e.nonTerm)
	}
	fmt.Println("Exit", e.nonTerm.Symbol)
}

// Tree retrieves the parse tree for the last production.
func (p Parser) Tree() *Tree {
	if e, ok := p.st.peek(); ok {
		return e.nonTerm
	}
	return nil
}
