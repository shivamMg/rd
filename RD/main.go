package main

import (
	"errors"

	t "github.com/shivammg/parsers/types"
)

// Parser represents a Recursive Descent parser.
type Parser struct {
	input     []string
	backtrack int
	current   int
	rules     map[string]func() (*t.Tree, error)
}

// NewParser returns a new Parser.
func NewParser(input []string) *Parser {
	return &Parser{
		input: input,
		// Input's index stored for backtracking.
		backtrack: -1,
		// Input's index where we're currently at.
		current: -1,
		// map[non-terminal] -> production
		rules: make(map[string]func() (*t.Tree, error)),
	}
}

// Match matches symbol with the next token in input.
func (p *Parser) Match(symbol string) bool {
	if p.current >= len(p.input)-1 {
		return false
	}
	p.current++
	return symbol == p.input[p.current]
}

// Backtrack resets current to where it was after last derivation.
func (p *Parser) Backtrack() {
	p.current = p.backtrack
}

// Register registers a production function for a non-terminal.
func (p *Parser) Register(nonTerm string, f func() (*t.Tree, error)) {
	p.rules[nonTerm] = f
}

// Run calls the production function for a non-terminal.
func (p *Parser) Run(nonTerm string) (*t.Tree, error) {
	f, ok := p.rules[nonTerm]
	if !ok {
		return nil, errors.New("Rule does not exist")
	}
	p.backtrack = p.current
	tree, err := f()
	if err != nil {
		p.Backtrack()
	}
	return tree, err
}

func main() {
}
