package rd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"github.com/shivamMg/ppds/tree"
)

const (
	BoxVer      = "│"
	BoxHor      = "─"
	BoxVerRight = "├"
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
	// left padding for logs
	padding int
	log     bool
}

// NewParser returns a new Parser to parse tokens. Production functions
// for non-terminals must be added using Rule method.
func NewParser(tokens []Token, log bool) *Parser {
	return &Parser{
		tokens:  tokens,
		st:      stack{},
		current: -1,
		padding: 0,
		log: log,
	}
}

func (p *Parser) Logf(format string, v ...interface{}) {
	if !p.log {
		return
	}
	prefix := strings.Repeat(BoxVer+" ", p.padding)
	newV := []interface{}{prefix}
	newV = append(newV, v...)
	logger.Printf("%s"+format, newV...)
}

func (p *Parser) matchLogf(format string, v ...interface{}) {
	if !p.log {
		return
	}
	prefix := ""
	if p.padding > 0 {
		prefix = strings.Repeat(BoxVer+" ", p.padding-1)
		prefix += BoxVerRight + " "
	}
	newV := []interface{}{prefix}
	newV = append(newV, v...)
	format = "%s"+format
	logger.Printf(format, newV...)
}

func (p *Parser) enterLogf(format string, v ...interface{}) {
	if !p.log {
		return
	}
	prefix := ""
	if p.padding > 0 {
		prefix = strings.Repeat(BoxVer+" ", p.padding-1)
		prefix += BoxVerRight
	}
	newV := []interface{}{prefix}
	newV = append(newV, v...)
	if p.padding > 0 {
		format = BoxHor +  format
	}
	format = "%s"+format
	logger.Printf(format, newV...)
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
		p.matchLogf("nomatch(%v) - no tokens left\n", token)
		return false
	}
	if token != next {
		p.current--
		p.matchLogf("nomatch(%v,%v)\n", token, next)
		return false
	}
	p.Add(token)
	p.matchLogf("match(%v)\n", token)
	return true
}

func (p *Parser) Enter(nonTerm string) {
	t := NewTree(nonTerm)
	p.st.push(ele{
		index:   p.current,
		nonTerm: t,
	})
	p.enterLogf("enter(%s)\n", nonTerm)
	p.padding++
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
	if !*result {
		p.current = e.index
	} else if parent, ok := p.st.peek(); ok && len(p.st) > 0 {
		parent.nonTerm.Add(e.nonTerm)
	}
	p.padding--
	p.Logf("exit(%t)\n", *result)
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
