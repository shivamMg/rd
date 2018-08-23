package main

import (
	"github.com/shivamMg/rd/examples/PL0/lexer"
	"github.com/shivamMg/rd/examples/PL0/parser"
	"testing"
)

func TestPrimeProgram(t *testing.T) {
	const output = `program
├─ block
│  ├─ const
│  ├─ ident
│  │  └─ max
│  ├─ =
│  ├─ number
│  │  └─ 100
│  ├─ ;
│  ├─ var
│  ├─ ident
│  │  └─ arg
│  ├─ ,
│  ├─ ident
│  │  └─ ret
│  ├─ ;
│  ├─ procedure
│  ├─ ident
│  │  └─ isprime
│  ├─ ;
│  ├─ block
│  │  ├─ var
│  │  ├─ ident
│  │  │  └─ i
│  │  ├─ ;
│  │  └─ statement
│  │     ├─ begin
│  │     ├─ statement
│  │     │  ├─ ident
│  │     │  │  └─ ret
│  │     │  ├─ :=
│  │     │  └─ expression
│  │     │     └─ term
│  │     │        └─ factor
│  │     │           └─ number
│  │     │              └─ 1
│  │     ├─ ;
│  │     ├─ statement
│  │     │  ├─ ident
│  │     │  │  └─ i
│  │     │  ├─ :=
│  │     │  └─ expression
│  │     │     └─ term
│  │     │        └─ factor
│  │     │           └─ number
│  │     │              └─ 2
│  │     ├─ ;
│  │     ├─ statement
│  │     │  ├─ while
│  │     │  ├─ condition
│  │     │  │  ├─ expression
│  │     │  │  │  └─ term
│  │     │  │  │     └─ factor
│  │     │  │  │        └─ ident
│  │     │  │  │           └─ i
│  │     │  │  ├─ <
│  │     │  │  └─ expression
│  │     │  │     └─ term
│  │     │  │        └─ factor
│  │     │  │           └─ ident
│  │     │  │              └─ arg
│  │     │  ├─ do
│  │     │  └─ statement
│  │     │     ├─ begin
│  │     │     ├─ statement
│  │     │     │  ├─ if
│  │     │     │  ├─ condition
│  │     │     │  │  ├─ expression
│  │     │     │  │  │  └─ term
│  │     │     │  │  │     ├─ factor
│  │     │     │  │  │     │  └─ ident
│  │     │     │  │  │     │     └─ arg
│  │     │     │  │  │     ├─ /
│  │     │     │  │  │     ├─ factor
│  │     │     │  │  │     │  └─ ident
│  │     │     │  │  │     │     └─ i
│  │     │     │  │  │     ├─ *
│  │     │     │  │  │     └─ factor
│  │     │     │  │  │        └─ ident
│  │     │     │  │  │           └─ i
│  │     │     │  │  ├─ =
│  │     │     │  │  └─ expression
│  │     │     │  │     └─ term
│  │     │     │  │        └─ factor
│  │     │     │  │           └─ ident
│  │     │     │  │              └─ arg
│  │     │     │  ├─ then
│  │     │     │  └─ statement
│  │     │     │     ├─ begin
│  │     │     │     ├─ statement
│  │     │     │     │  ├─ ident
│  │     │     │     │  │  └─ ret
│  │     │     │     │  ├─ :=
│  │     │     │     │  └─ expression
│  │     │     │     │     └─ term
│  │     │     │     │        └─ factor
│  │     │     │     │           └─ number
│  │     │     │     │              └─ 0
│  │     │     │     ├─ ;
│  │     │     │     ├─ statement
│  │     │     │     │  ├─ ident
│  │     │     │     │  │  └─ i
│  │     │     │     │  ├─ :=
│  │     │     │     │  └─ expression
│  │     │     │     │     └─ term
│  │     │     │     │        └─ factor
│  │     │     │     │           └─ ident
│  │     │     │     │              └─ arg
│  │     │     │     └─ end
│  │     │     ├─ ;
│  │     │     ├─ statement
│  │     │     │  ├─ ident
│  │     │     │  │  └─ i
│  │     │     │  ├─ :=
│  │     │     │  └─ expression
│  │     │     │     ├─ term
│  │     │     │     │  └─ factor
│  │     │     │     │     └─ ident
│  │     │     │     │        └─ i
│  │     │     │     ├─ +
│  │     │     │     └─ term
│  │     │     │        └─ factor
│  │     │     │           └─ number
│  │     │     │              └─ 1
│  │     │     └─ end
│  │     └─ end
│  ├─ ;
│  ├─ procedure
│  ├─ ident
│  │  └─ primes
│  ├─ ;
│  ├─ block
│  │  └─ statement
│  │     ├─ begin
│  │     ├─ statement
│  │     │  ├─ ident
│  │     │  │  └─ arg
│  │     │  ├─ :=
│  │     │  └─ expression
│  │     │     └─ term
│  │     │        └─ factor
│  │     │           └─ number
│  │     │              └─ 2
│  │     ├─ ;
│  │     ├─ statement
│  │     │  ├─ while
│  │     │  ├─ condition
│  │     │  │  ├─ expression
│  │     │  │  │  └─ term
│  │     │  │  │     └─ factor
│  │     │  │  │        └─ ident
│  │     │  │  │           └─ arg
│  │     │  │  ├─ <
│  │     │  │  └─ expression
│  │     │  │     └─ term
│  │     │  │        └─ factor
│  │     │  │           └─ ident
│  │     │  │              └─ max
│  │     │  ├─ do
│  │     │  └─ statement
│  │     │     ├─ begin
│  │     │     ├─ statement
│  │     │     │  ├─ call
│  │     │     │  └─ ident
│  │     │     │     └─ isprime
│  │     │     ├─ ;
│  │     │     ├─ statement
│  │     │     │  ├─ if
│  │     │     │  ├─ condition
│  │     │     │  │  ├─ expression
│  │     │     │  │  │  └─ term
│  │     │     │  │  │     └─ factor
│  │     │     │  │  │        └─ ident
│  │     │     │  │  │           └─ ret
│  │     │     │  │  ├─ =
│  │     │     │  │  └─ expression
│  │     │     │  │     └─ term
│  │     │     │  │        └─ factor
│  │     │     │  │           └─ number
│  │     │     │  │              └─ 1
│  │     │     │  ├─ then
│  │     │     │  └─ statement
│  │     │     │     ├─ !
│  │     │     │     └─ expression
│  │     │     │        └─ term
│  │     │     │           └─ factor
│  │     │     │              └─ ident
│  │     │     │                 └─ arg
│  │     │     ├─ ;
│  │     │     ├─ statement
│  │     │     │  ├─ ident
│  │     │     │  │  └─ arg
│  │     │     │  ├─ :=
│  │     │     │  └─ expression
│  │     │     │     ├─ term
│  │     │     │     │  └─ factor
│  │     │     │     │     └─ ident
│  │     │     │     │        └─ arg
│  │     │     │     ├─ +
│  │     │     │     └─ term
│  │     │     │        └─ factor
│  │     │     │           └─ number
│  │     │     │              └─ 1
│  │     │     └─ end
│  │     └─ end
│  ├─ ;
│  └─ statement
│     ├─ call
│     └─ ident
│        └─ primes
└─ .
`
	tokens := lexer.Lex(primeProgram)
	ok, tree := parser.Parse(tokens, parser.Program)
	if !ok {
		t.Error("Match failed")
	}
	if sprint(tree) != output {
		t.Error("Incorrect parse tree")
	}
}

func TestSquareProgram(t *testing.T) {
	const output = `program
├─ block
│  ├─ var
│  ├─ ident
│  │  └─ x
│  ├─ ,
│  ├─ ident
│  │  └─ squ
│  ├─ ;
│  ├─ procedure
│  ├─ ident
│  │  └─ square
│  ├─ ;
│  ├─ block
│  │  └─ statement
│  │     ├─ begin
│  │     ├─ statement
│  │     │  ├─ ident
│  │     │  │  └─ squ
│  │     │  ├─ :=
│  │     │  └─ expression
│  │     │     └─ term
│  │     │        ├─ factor
│  │     │        │  └─ ident
│  │     │        │     └─ x
│  │     │        ├─ *
│  │     │        └─ factor
│  │     │           └─ ident
│  │     │              └─ x
│  │     └─ end
│  ├─ ;
│  └─ statement
│     ├─ begin
│     ├─ statement
│     │  ├─ ident
│     │  │  └─ x
│     │  ├─ :=
│     │  └─ expression
│     │     └─ term
│     │        └─ factor
│     │           └─ number
│     │              └─ 1
│     ├─ ;
│     ├─ statement
│     │  ├─ while
│     │  ├─ condition
│     │  │  ├─ expression
│     │  │  │  └─ term
│     │  │  │     └─ factor
│     │  │  │        └─ ident
│     │  │  │           └─ x
│     │  │  ├─ <=
│     │  │  └─ expression
│     │  │     └─ term
│     │  │        └─ factor
│     │  │           └─ number
│     │  │              └─ 10
│     │  ├─ do
│     │  └─ statement
│     │     ├─ begin
│     │     ├─ statement
│     │     │  ├─ call
│     │     │  └─ ident
│     │     │     └─ square
│     │     ├─ ;
│     │     ├─ statement
│     │     │  ├─ !
│     │     │  └─ expression
│     │     │     └─ term
│     │     │        └─ factor
│     │     │           └─ ident
│     │     │              └─ squ
│     │     ├─ ;
│     │     ├─ statement
│     │     │  ├─ ident
│     │     │  │  └─ x
│     │     │  ├─ :=
│     │     │  └─ expression
│     │     │     ├─ term
│     │     │     │  └─ factor
│     │     │     │     └─ ident
│     │     │     │        └─ x
│     │     │     ├─ +
│     │     │     └─ term
│     │     │        └─ factor
│     │     │           └─ number
│     │     │              └─ 1
│     │     └─ end
│     └─ end
└─ .
`

	tokens := lexer.Lex(squareProgram)
	ok, tree := parser.Parse(tokens, parser.Program)
	if !ok {
		t.Error("Match failed")
	}
	if sprint(tree) != output {
		t.Error("Incorrect parse tree")
	}
}
