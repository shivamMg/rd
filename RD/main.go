package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	t "github.com/shivammg/parsers/types"
)

/*
Recursive Descent parser for the following grammar:
	E  -> TE'
	E' -> +TE'|ε
	T  -> FT'
	T' -> *FT'|ε
	F  -> id|(E)

ε represents empty string.
*/

type Parser struct {
	input    []string
	curIndex int
}

func NewParser(input []string) *Parser {
	return &Parser{input: input, curIndex: -1}
}

func (p Parser) next() string {
	p.curIndex++
	return p.input[p.curIndex]
}

func (p Parser) match(symbol string) bool {
	return symbol == p.next()
}

// E parses: E  -> TE'
func (p Parser) E() (*t.Tree, error) {
	cur := p.curIndex

	t1, err := p.T()
	if err != nil {
		p.curIndex = cur
		return nil, err
	}
	t2, err := p.EPrime()
	if err != nil {
		p.curIndex = cur
		return nil, err
	}
	return t.NewTree("E", t1, t2), nil
}

// EPrime parses: E' -> +TE'|ε
func (p Parser) EPrime() (*t.Tree, error) {
	cur := p.curIndex

	if !p.match("+") {
		// epsilon exists for the rule
		return nil, nil
	}
	t1, err := p.T()
	if err != nil {
		p.curIndex = cur
		return nil, err
	}
	t2, err := p.EPrime()
	if err != nil {
		p.curIndex = cur
		return nil, err
	}
	return t.NewTree("E'", t.NewTree("+"), t1, t2), nil
}

// T parses: T  -> FT'
func (p Parser) T() (*t.Tree, error) {
	cur := p.curIndex

	t1, err := p.F()
	if err != nil {
		p.curIndex = cur
		return nil, err
	}
	t2, err := p.TPrime()
	if err != nil {
		p.curIndex = cur
		return nil, err
	}
	return t.NewTree("T", t1, t2), nil
}

// TPrime parses: T' -> *FT'|ε
func (p Parser) TPrime() (*t.Tree, error) {
	cur := p.curIndex

	if !p.match("*") {
		// epsilon exists for the rule
		return nil, nil
	}
	t1, err := p.F()
	if err != nil {
		p.curIndex = cur
		return nil, err
	}
	t2, err := p.TPrime()
	if err != nil {
		p.curIndex = cur
		return nil, err
	}
	return t.NewTree("T'", t.NewTree("*"), t1, t2), nil
}

// F parses: F  -> id|(E)
func (p Parser) F() (*t.Tree, error) {
	cur := p.curIndex

	if p.match("id") {
		return t.NewTree("F", t.NewTree("id")), nil
	}
	if p.match("(") {
		t1, err := p.E()
		if err != nil {
			p.curIndex = cur
			return nil, err
		}
		if p.match(")") {
			return t.NewTree("F", t.NewTree("("), t1, t.NewTree(")")), nil
		}
	}
	p.curIndex = cur
	return nil, errors.New("No match")
}

func main() {
	p := NewParser([]string{"id"})
	t, err := p.E()
	if err != nil {
		log.Fatalln(err)
	}
	b, _ := json.MarshalIndent(t, "", "  ")
	fmt.Println(string(b))
}
