package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	t "github.com/shivammg/parsers/types"
)

// Parser represents a Recursive Descent parser.
type Parser struct {
	input []string
	// Maintains input indexes as you keep deriving productions.
	// Used in backtracking to a previous input index.
	stack []int
	// Input's index where we're currently at.
	current int
	// map[non-terminal] -> production
	rules map[string]func() (*t.Tree, error)
}

// NewParser returns a new Parser.
func NewParser(input []string) *Parser {
	return &Parser{
		input:   input,
		stack:   []int{-1},
		current: -1,
		rules:   make(map[string]func() (*t.Tree, error)),
	}
}

// Match matches terminal with the next token in input.
func (p *Parser) Match(term string) bool {
	if p.current >= len(p.input)-1 {
		return false
	}
	p.current++
	return term == p.input[p.current]
}

// Backtrack resets current to where it was after last derivation.
func (p *Parser) Backtrack() {
	l := len(p.stack)
	last := p.stack[l-1]
	p.stack = p.stack[:l-1]
	p.current = last
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
	p.stack = append(p.stack, p.current)
	tree, err := f()
	if err != nil {
		p.Backtrack()
	}
	return tree, err
}

func main() {
	p := NewParser([]string{"a", "c"})

	p.Register("E", func() (*t.Tree, error) {
		if p.Match("a") {
			t1, err := p.Run("F")
			if err == nil {
				return t.NewTree("E", t1), nil
			}
		}
		p.Backtrack()
		t1, err := p.Run("G")
		if err == nil {
			return t.NewTree("E", t1), nil
		}
		return nil, errors.New("No match")
	})

	p.Register("F", func() (*t.Tree, error) {
		if p.Match("b") {
			return t.NewTree("F", t.NewTree("b")), nil
		}
		return nil, errors.New("No match")
	})

	p.Register("G", func() (*t.Tree, error) {
		if p.Match("c") {
			return t.NewTree("F", t.NewTree("c")), nil
		}
		return nil, errors.New("No match")
	})

	tree, err := p.Run("E")
	if err != nil {
		log.Fatalln(err)
	}
	b, _ := json.MarshalIndent(tree, "", "  ")
	fmt.Println(string(b))
}
