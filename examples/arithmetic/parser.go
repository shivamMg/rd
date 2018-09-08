/*
Grammar without recursion and left-factored. Needs single lookahead. Suitable for R.D. parsing.

	Expr   = Term Expr'
	Expr'  = "+" Expr | "-" Expr | ε
	Term   = Factor Term'
	Term'  = "*" Term | "/" Term | ε
	Factor = "(" Expr ")" | "-" Factor | Number
*/
package main

import (
	"errors"
	"github.com/shivamMg/rd"
)

var p rd.P

func Expr() (ok bool) {
	p.Enter("Expr")
	defer p.Exit(&ok)

	return Term() && ExprPrime()
}

func ExprPrime() (ok bool) {
	p.Enter("Expr'")
	defer p.Exit(&ok)

	if p.Match(Plus) {
		return Expr()
	}
	if p.Match(Minus) {
		return Expr()
	}
	p.Add(Epsilon)
	return true
}

func Term() (ok bool) {
	p.Enter("Term")
	defer p.Exit(&ok)

	return Factor() && TermPrime()
}

func TermPrime() (ok bool) {
	p.Enter("Term'")
	defer p.Exit(&ok)

	if p.Match(Star) {
		return Term()
	}
	if p.Match(Slash) {
		return Term()
	}
	p.Add(Epsilon)
	return true
}

func Factor() (ok bool) {
	p.Enter("Factor")
	defer p.Exit(&ok)

	if p.Match(OpenParen) {
		return Expr() && p.Match(CloseParen)
	}
	if p.Match(Minus) {
		return Factor()
	}
	return Number()
}

func Number() (ok bool) {
	p.Enter("Number")
	defer p.Exit(&ok)

	token, ok := p.Next()
	if !ok {
		return false
	}
	switch token.(type) {
	case int64, float64:
		p.Add(token)
		return true
	}
	p.Reset()
	return false
}

func Parse(tokens []rd.Token) (*rd.Parser, error) {
	parser := rd.NewParser(tokens, true)
	p = parser
	if ok := Expr(); !ok {
		return nil, errors.New("parsing failed")
	}
	return parser, nil
}
