package main

import (
	"testing"

	"github.com/shivamMg/rd"
)

func TestArithmeticExpressionsGrammar(t *testing.T) {
	tokens := []rd.Token{"2.8", "+", "(", "3", "-", ".733", ")", "/", "23"}
	b = rd.NewBuilder(tokens)
	ok := Expr()
	if !ok {
		t.Error("parsing failed")
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
	got := b.DebugTree().Sprint()
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
		t.Errorf("invalid parse tree. want: %s\ngot: %s\n", parseTree, got)
	}
}
