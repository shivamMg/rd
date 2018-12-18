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

var numberRegex = regexp.MustCompile(`^(\d*\.\d+|\d+)$`)

func Expr(b *rd.Builder) (ok bool) {
	defer b.Enter("Expr").Exit(&ok)

	if Term(b) && b.Match(Plus) && Expr(b) {
		return true
	}
	b.Backtrack()
	if Term(b) && b.Match(Minus) && Expr(b) {
		return true
	}
	b.Backtrack()
	return Term(b)
}

func Term(b *rd.Builder) (ok bool) {
	defer b.Enter("Term").Exit(&ok)

	if Factor(b) && b.Match(Star) && Term(b) {
		return true
	}
	b.Backtrack()
	if Factor(b) && b.Match(Slash) && Term(b) {
		return true
	}
	b.Backtrack()
	return Factor(b)
}

func Factor(b *rd.Builder) (ok bool) {
	defer b.Enter("Factor").Exit(&ok)

	if b.Match(OpenParen) && Expr(b) && b.Match(CloseParen) {
		return true
	}
	b.Backtrack()
	if b.Match(Minus) && Factor(b) {
		return true
	}
	b.Backtrack()
	return Number(b)
}

func Number(b *rd.Builder) (ok bool) {
	defer b.Enter("Number").Exit(&ok)

	token, ok := b.Next()
	if !ok {
		return false
	}
	if numberRegex.MatchString(fmt.Sprint(token)) {
		b.Add(token)
		return true
	}
	b.Backtrack()
	return false
}

func Parse(tokens []rd.Token) (parseTree *rd.Tree, debugTree *rd.DebugTree, err error) {
	b := rd.NewBuilder(tokens)
	if ok := Expr(b); ok && b.Err() == nil {
		return b.ParseTree(), b.DebugTree(), nil
	}
	return nil, b.DebugTree(), b.Err()
}
