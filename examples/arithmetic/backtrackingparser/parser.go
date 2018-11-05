package backtrackingparser

import (
	"fmt"
	"regexp"

	"github.com/shivamMg/rd"
	. "github.com/shivamMg/rd/examples/arithmetic/tokens"
)

const Grammar = `
	Expr   = Term "+" Expr | Term "-" Expr | Term
	Term   = Factor "*" Term | Factor "/" Term | Factor
	Factor = "(" Expr ")" | "-" Factor | Number
`

var (
	numberRegex = regexp.MustCompile(`^(\d*\.\d+|\d+)$`)
	b           *rd.Builder
)

func Expr() (ok bool) {
	b.Enter("Expr")
	defer b.Exit(&ok)

	if Term() && b.Match(Plus) && Expr() {
		return true
	}
	b.Reset()
	if Term() && b.Match(Minus) && Expr() {
		return true
	}
	b.Reset()
	return Term()
}

func Term() (ok bool) {
	b.Enter("Term")
	defer b.Exit(&ok)

	if Factor() && b.Match(Star) && Term() {
		return true
	}
	b.Reset()
	if Factor() && b.Match(Slash) && Term() {
		return true
	}
	b.Reset()
	return Factor()
}

func Factor() (ok bool) {
	b.Enter("Factor")
	defer b.Exit(&ok)

	if b.Match(OpenParen) && Expr() && b.Match(CloseParen) {
		return true
	}
	b.Reset()
	if b.Match(Minus) && Factor() {
		return true
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

func Parse(tokens []rd.Token) (parseTree *rd.Tree, debugTree string, err error) {
	b = rd.NewBuilder(tokens)
	if ok := Expr(); !ok {
		return nil, b.DebugTree().Sprint(), b.Err()
	}
	return b.Tree(), b.DebugTree().Sprint(), nil
}
