package rd_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/shivamMg/rd"
)

const (
	Plus       = "+"
	Minus      = "-"
	Star       = "*"
	Slash      = "/"
	OpenParen  = "("
	CloseParen = ")"
	Epsilon    = "ε"
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

func TestArithmeticExpressionsGrammar(t *testing.T) {
	tokens := []rd.Token{"2.8", "+", "(", "3", "-", ".733", ")", "/", "23"}
	b = rd.NewBuilder(tokens)
	ok := Expr()
	if !ok {
		t.Error("expected parsing to pass")
	}

	debugTree := `Expr(true)
├─ Term(true)
│  ├─ Factor(true)
│  │  ├─ 2.8 ≠ (
│  │  ├─ 2.8 ≠ -
│  │  └─ Number(true)
│  └─ Term'(true)
│     ├─ + ≠ *
│     └─ + ≠ /
└─ Expr'(true)
   ├─ + = +
   └─ Expr(true)
      ├─ Term(true)
      │  ├─ Factor(true)
      │  │  ├─ ( = (
      │  │  ├─ Expr(true)
      │  │  │  ├─ Term(true)
      │  │  │  │  ├─ Factor(true)
      │  │  │  │  │  ├─ 3 ≠ (
      │  │  │  │  │  ├─ 3 ≠ -
      │  │  │  │  │  └─ Number(true)
      │  │  │  │  └─ Term'(true)
      │  │  │  │     ├─ - ≠ *
      │  │  │  │     └─ - ≠ /
      │  │  │  └─ Expr'(true)
      │  │  │     ├─ - ≠ +
      │  │  │     ├─ - = -
      │  │  │     └─ Expr(true)
      │  │  │        ├─ Term(true)
      │  │  │        │  ├─ Factor(true)
      │  │  │        │  │  ├─ .733 ≠ (
      │  │  │        │  │  ├─ .733 ≠ -
      │  │  │        │  │  └─ Number(true)
      │  │  │        │  └─ Term'(true)
      │  │  │        │     ├─ ) ≠ *
      │  │  │        │     └─ ) ≠ /
      │  │  │        └─ Expr'(true)
      │  │  │           ├─ ) ≠ +
      │  │  │           └─ ) ≠ -
      │  │  └─ ) = )
      │  └─ Term'(true)
      │     ├─ / ≠ *
      │     ├─ / = /
      │     └─ Term(true)
      │        ├─ Factor(true)
      │        │  ├─ 23 ≠ (
      │        │  ├─ 23 ≠ -
      │        │  └─ Number(true)
      │        └─ Term'(true)
      │           ├─ <no tokens left> ≠ *
      │           └─ <no tokens left> ≠ /
      └─ Expr'(true)
         ├─ <no tokens left> ≠ +
         └─ <no tokens left> ≠ -
`
	got := b.SprintDebugTree()
	if got != debugTree {
		t.Errorf("invalid debug tree. expected: %s\ngot: %s\n", debugTree, got)
	}

	parseTree := `Expr
├─ Term
│  ├─ Factor
│  │  └─ Number
│  │     └─ 2.8
│  └─ Term'
│     └─ ε
└─ Expr'
   ├─ +
   └─ Expr
      ├─ Term
      │  ├─ Factor
      │  │  ├─ (
      │  │  ├─ Expr
      │  │  │  ├─ Term
      │  │  │  │  ├─ Factor
      │  │  │  │  │  └─ Number
      │  │  │  │  │     └─ 3
      │  │  │  │  └─ Term'
      │  │  │  │     └─ ε
      │  │  │  └─ Expr'
      │  │  │     ├─ -
      │  │  │     └─ Expr
      │  │  │        ├─ Term
      │  │  │        │  ├─ Factor
      │  │  │        │  │  └─ Number
      │  │  │        │  │     └─ .733
      │  │  │        │  └─ Term'
      │  │  │        │     └─ ε
      │  │  │        └─ Expr'
      │  │  │           └─ ε
      │  │  └─ )
      │  └─ Term'
      │     ├─ /
      │     └─ Term
      │        ├─ Factor
      │        │  └─ Number
      │        │     └─ 23
      │        └─ Term'
      │           └─ ε
      └─ Expr'
         └─ ε
`
	got = b.Tree().Sprint()
	if got != parseTree {
		t.Errorf("invalid parse tree. expected: %s\ngot: %s\n", parseTree, got)
	}
}
