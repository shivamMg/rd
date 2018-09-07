/*
Grammar without left recursion. Needs arbitrary lookahead.
	Expr    = Term AddExpr | Term SubExpr
	AddExpr = "+" Expr | ε
	SubExpr = "-" Expr | ε
	Term    = Factor MulExpr | Factor DivExpr
	MulExpr = "*" Term | ε
	DivExpr = "/" Term | ε
	Factor  = "(" Expr ")" | "-" Factor | Number
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

	if Term() && AddExpr() {
		return true
	}
	p.Reset()
	return Term() && SubExpr()
}

func AddExpr() (ok bool) {
	p.Enter("AddExpr")
	defer p.Exit(&ok)

	if p.Match(Plus) && Expr() {
		return true
	}
	p.Reset()
	p.Add(Epsilon)
	return true
}

func SubExpr() (ok bool) {
	p.Enter("SubExpr")
	defer p.Exit(&ok)

	if p.Match(Minus) && Expr() {
		return true
	}
	p.Reset()
	p.Add(Epsilon)
	return true
}

func Term() (ok bool) {
	p.Enter("Term")
	defer p.Exit(&ok)

	if Factor() && MulExpr() {
		return true
	}
	p.Reset()
	return Factor() && DivExpr()
}

func MulExpr() (ok bool) {
	p.Enter("MulExpr")
	defer p.Exit(&ok)

	if p.Match(Star) && Term() {
		return true
	}
	p.Reset()
	p.Add(Epsilon)
	return true
}

func DivExpr() (ok bool) {
	p.Enter("MulExpr")
	defer p.Exit(&ok)

	if p.Match(Slash) && Term() {
		return true
	}
	p.Reset()
	p.Add(Epsilon)
	return true
}

func Factor() (ok bool) {
	p.Enter("Factor")
	defer p.Exit(&ok)

	if p.Match(OpenParen) && Expr() && p.Match(CloseParen) {
		return true
	}
	p.Reset()
	if p.Match(Minus) && Factor() {
		return true
	}
	p.Reset()
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
