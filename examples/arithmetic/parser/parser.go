package parser

import (
	"fmt"
	"regexp"

	"github.com/shivamMg/rd"
	. "github.com/shivamMg/rd/examples/arithmetic/tokens"
)

const Grammar = `
	Expr   = Term Expr'
	Expr'  = "+" Expr | "-" Expr | ε
	Term   = Factor Term'
	Term'  = "*" Term | "/" Term | ε
	Factor = "(" Expr ")" | "-" Factor | Number
`

var numberRegex = regexp.MustCompile(`^(\d*\.\d+|\d+)$`)

func Expr(b *rd.Builder) (ok bool) {
	b.Enter("Expr")
	defer b.Exit(&ok)

	return Term(b) && ExprPrime(b)
}

func ExprPrime(b *rd.Builder) (ok bool) {
	b.Enter("Expr'")
	defer b.Exit(&ok)

	if b.Match(Plus) {
		return Expr(b)
	}
	if b.Match(Minus) {
		return Expr(b)
	}
	b.Add(Epsilon)
	return true
}

func Term(b *rd.Builder) (ok bool) {
	b.Enter("Term")
	defer b.Exit(&ok)

	return Factor(b) && TermPrime(b)
}

func TermPrime(b *rd.Builder) (ok bool) {
	b.Enter("Term'")
	defer b.Exit(&ok)

	if b.Match(Star) {
		return Term(b)
	}
	if b.Match(Slash) {
		return Term(b)
	}
	b.Add(Epsilon)
	return true
}

func Factor(b *rd.Builder) (ok bool) {
	b.Enter("Factor")
	defer b.Exit(&ok)

	if b.Match(OpenParen) {
		return Expr(b) && b.Match(CloseParen)
	}
	if b.Match(Minus) {
		return Factor(b)
	}
	return Number(b)
}

func Number(b *rd.Builder) (ok bool) {
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
	return false
}

func Parse(tokens []rd.Token) (parseTree *rd.Tree, debugTree *rd.DebugTree, err error) {
	b := rd.NewBuilder(tokens)
	ok := Expr(b)
	// it's possible that there are tokens left even after parsing.
	// in which case ok will be true and b.Err() will not be nil.
	if ok && b.Err() == nil {
		return b.ParseTree(), b.DebugTree(), nil
	}
	return nil, b.DebugTree(), b.Err()
}
