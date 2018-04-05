package rd

import (
	"errors"

	t "github.com/shivammg/parsers/types"
)

const (
	// ErrNoMatch is error text to signify that no matches were
	// found for input token.
	ErrNoMatch = "No match"
	// ErrRuleNotFound is error text to signify that no production
	// rule for was found for the non-terminal.
	ErrRuleNotFound = "Rule not found"
)

type stack []int

func (st stack) peek() int {
	l := len(st)
	if l == 0 {
		return -1
	}
	return st[l-1]
}

func (st *stack) pop() int {
	l := len(*st)
	if l == 0 {
		return -1
	}
	ele := (*st)[l-1]
	*st = (*st)[:l-1]
	return ele
}

func (st *stack) push(ele int) {
	*st = append(*st, ele)
}

// Parser represents a Recursive Descent parser.
type Parser struct {
	// Stores input tokens.
	input []string
	// Maintains input indexes as you keep deriving productions.
	// Used in backtracking to a previous input index.
	st stack
	// Input's index where we're currently at.
	current int
	// map[non-terminal] -> production function.
	rules map[string]func() (*t.Tree, error)
}

// NewParser returns a new Parser to parse input. Grammar production
// rule functions must be registered to parse the input.
func NewParser(input []string) *Parser {
	return &Parser{
		input:   input,
		st:      stack{},
		current: -1,
		rules:   make(map[string]func() (*t.Tree, error)),
	}
}

// Backtrack resets the current position inside input. It is reset to where
// it was during the beginning of the production function, the one that it's
// called inside.
func (p *Parser) Backtrack() {
	p.current = p.st.peek()
}

// Register saves a production function for a non-terminal. This function can
// then be called using Run method.
func (p *Parser) Register(nonTerm string, f func() (*t.Tree, error)) {
	p.rules[nonTerm] = f
}

// Match matches terminal with the next token in input, and returns a boolean
// accordingly. If the match was unsuccessful then Backtrack is called,
// implicitly - user won't have to call it inside production function.
func (p *Parser) Match(term string) bool {
	if p.current >= len(p.input)-1 {
		return false
	}
	p.current++
	if term != p.input[p.current] {
		p.Backtrack()
		return false
	}
	return true
}

// Run calls the production function for a non-terminal that was saved using
// Register.
func (p *Parser) Run(nonTerm string) (*t.Tree, error) {
	f, ok := p.rules[nonTerm]
	if !ok {
		return nil, errors.New(ErrRuleNotFound)
	}
	p.st.push(p.current)
	tree, err := f()
	if err != nil {
		p.Backtrack()
	}
	p.st.pop()
	return tree, err
}
