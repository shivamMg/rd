package rd

import (
	t "github.com/shivammg/parsers/types"
)

const (
	// ErrRuleNotFound is error text to signify that no production
	// rule for was found for the non-terminal.
	ErrRuleNotFound = "Rule not found"
)

type ele struct {
	index   int
	nonTerm *t.Tree
}

type stack []ele

func (st stack) peek() ele {
	l := len(st)
	if l == 0 {
		return ele{}
	}
	return st[l-1]
}

func (st *stack) pop() ele {
	l := len(*st)
	if l == 0 {
		return ele{}
	}
	e := (*st)[l-1]
	*st = (*st)[:l-1]
	return e
}

func (st *stack) push(e ele) {
	*st = append(*st, e)
}

func (st *stack) isEmpty() bool {
	return len(*st) == 0
}

// Parser represents a Recursive Descent parser.
type Parser struct {
	// Stores input tokens.
	input []string
	st    stack
	// Input's index where we're currently at.
	current int
	// map[non-terminal] -> production function.
	rules map[string]func() bool
}

// NewParser returns a new Parser to parse input. Grammar production
// rule functions must be registered to parse the input.
func NewParser(input []string) *Parser {
	return &Parser{
		input:   input,
		st:      stack{},
		current: -1,
		rules:   make(map[string]func() bool),
	}
}

func (p *Parser) Backtrack() {
	e := p.st.peek()
	p.current = e.index
}

// Register saves a production function for a non-terminal.
func (p *Parser) Register(nonTerm string, f func() bool) {
	p.rules[nonTerm] = f
}

func (p *Parser) Match(symbol string) bool {
	f, ok := p.rules[symbol]
	if !ok {
		// it's a terminal
		if p.current >= len(p.input)-1 {
			return false
		}
		p.current++
		if symbol != p.input[p.current] {
			p.current--
			return false
		}

		p.st.peek().nonTerm.Add(t.NewTree(symbol))
		return true
	}

	// it's a non-terminal
	tree := t.NewTree(symbol)
	// if it's not the first production
	if !p.st.isEmpty() {
		p.st.peek().nonTerm.Add(tree)
	}

	p.st.push(ele{index: p.current, nonTerm: tree})
	isMatch := f()
	// don't pop if it's the last element
	if len(p.st) > 1 {
		p.st.pop()
	}

	if !isMatch {
		p.st.peek().nonTerm.Detach(tree)
		p.Backtrack()
	}
	return isMatch
}

func (p Parser) Tree() *t.Tree {
	if p.st.isEmpty() {
		return nil
	}
	return p.st.peek().nonTerm
}

/*
func main() {
	p := NewParser([]string{"a", "f", "h"})

	p.Register("A", func() bool {
		if p.Match("a") &&
			p.Match("B") &&
			p.Match("c") {
			return true
		}
		return p.Match("E")
	})

	p.Register("B", func() bool {
		return p.Match("f") && p.Match("g")
	})

	p.Register("E", func() bool {
		return p.Match("a") && p.Match("F")
	})

	p.Register("F", func() bool {
		return p.Match("f") && p.Match("h")
	})

	fmt.Println(p.Match("A"))
	tree := p.st.pop().nonTerm
	b, _ := json.MarshalIndent(tree, "", "  ")
	fmt.Println(string(b))
}

*/
