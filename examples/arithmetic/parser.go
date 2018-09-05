/*
Grammar with left recursion. Not suitable for RD parsing.
	Expr   = Expr "+" Term | Expr "-" Term | Term
	Term   = Term "*" Factor | Term "/" Factor | Factor
	Factor = "(" Expr ")" | "-" Factor | Number

Grammar without left recursion. Needs arbitrary lookahead.
	Expr    = Term AddExpr | Term SubExpr
	AddExpr = "+" Expr | ε
	SubExpr = "-" Expr | ε
	Term    = Factor MulExpr | Factor DivExpr
	MulExpr = "*" Term | ε
	DivExpr = "/" Term | ε
	Factor  = "(" Expr ")" | "-" Factor | Number

LL(1) grammar (factored and without recursion). Needs single lookahead.
	Expr   = Term Expr'
	Expr'  = "+" Expr | "-" Expr | ε
	Term   = Factor Term'
	Term'  = "*" Term | "/" Term | ε
	Factor = "(" Expr ")" | "-" Factor | Number
*/
package main

import (
	"fmt"
	"github.com/shivamMg/rd"
	"reflect"
)

var p P

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
	fmt.Println("#Number", token, reflect.TypeOf(token))
	p.Reset()
	return false
}

func Parse(tokens []rd.Token) (*rd.Tree, error) {
	p = rd.NewParser(tokens)
	ok := Expr()
	fmt.Println(ok)
	return nil, nil
}

type P interface {
	// returns false if no tokens left to match
	Match(token rd.Token) (ok bool)
	// ok is false if no token left
	Next() (token rd.Token, ok bool)
	// panics if no node to attach token (empty stack)
	// always returns true
	Add(token rd.Token)
	Reset()
	Enter(nonTerm string)
	Exit(result *bool)
}
