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
	"fmt"
	"regexp"

	"github.com/shivamMg/rd"
)

var (
	numberRegex = regexp.MustCompile(`^(\d*\.\d+|\d+)$`)
	b           *rd.Builder
)

func Expr() (ok bool) {
	b.Enter("Expr")
	defer b.Exit(&ok)

	return Term() && ExprPrime()
}

func ExprPrime() (ok bool) {
	b.Enter("Expr'")
	defer b.Exit(&ok)

	if b.Match(Plus) {
		return Expr()
	}
	if b.Match(Minus) {
		return Expr()
	}
	b.Add(Epsilon)
	return true
}

func Term() (ok bool) {
	b.Enter("Term")
	defer b.Exit(&ok)

	return Factor() && TermPrime()
}

func TermPrime() (ok bool) {
	b.Enter("Term'")
	defer b.Exit(&ok)

	if b.Match(Star) {
		return Term()
	}
	if b.Match(Slash) {
		return Term()
	}
	b.Add(Epsilon)
	return true
}

func Factor() (ok bool) {
	b.Enter("Factor")
	defer b.Exit(&ok)

	if b.Match(OpenParen) {
		return Expr() && b.Match(CloseParen)
	}
	if b.Match(Minus) {
		return Factor()
	}
	return Number()
}

func Number() (ok bool) {
	b.Enter("Number")
	defer b.Exit(&ok)

	token, ok := b.Next()
	if !ok {
		return false
	}
	if numberRegex.MatchString(fmt.Sprint(token)) {
		b.Add(token)
		return true
	}
	b.Reset()
	return false
}

func Parse(tokens []rd.Token) (parseTree *rd.Tree, err error) {
	b = rd.NewBuilder(tokens)
	if ok := Expr(); !ok {
		return nil, errors.New("parsing failed")
	}
	return b.Tree(), nil
}
