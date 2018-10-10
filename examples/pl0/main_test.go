package main_test

import (
	"testing"

	"github.com/shivamMg/rd/examples/pl0/lexer"
	"github.com/shivamMg/rd/examples/pl0/parser"
)

const (
	squareProgram = `VAR x, squ;

PROCEDURE square;
BEGIN
   squ:= x * x
END;

BEGIN
   x := 1;
   WHILE x <= 10 DO
   BEGIN
      CALL square;
      ! squ;
      x := x + 1
   END
END.`
	primeProgram = `const max = 100;
var arg, ret;

procedure isprime;
var i;
begin
	ret := 1;
	i := 2;
	while i < arg do
	begin
		if arg / i * i = arg then
		begin
			ret := 0;
			i := arg
		end;
		i := i + 1
	end
end;

procedure primes;
begin
	arg := 2;
	while arg < max do
	begin
		call isprime;
		if ret = 1 then ! arg;
		arg := arg + 1
	end
end;

call primes
.`
)

func TestPrimeProgram(t *testing.T) {
	const output = `Program
├─ Block
│  ├─ const
│  ├─ Ident
│  │  └─ max
│  ├─ =
│  ├─ Number
│  │  └─ 100
│  ├─ ;
│  ├─ var
│  ├─ Ident
│  │  └─ arg
│  ├─ ,
│  ├─ Ident
│  │  └─ ret
│  ├─ ;
│  ├─ procedure
│  ├─ Ident
│  │  └─ isprime
│  ├─ ;
│  ├─ Block
│  │  ├─ var
│  │  ├─ Ident
│  │  │  └─ i
│  │  ├─ ;
│  │  └─ Statement
│  │     ├─ begin
│  │     ├─ Statement
│  │     │  ├─ Ident
│  │     │  │  └─ ret
│  │     │  ├─ :=
│  │     │  └─ Expression
│  │     │     └─ Term
│  │     │        └─ Factor
│  │     │           └─ Number
│  │     │              └─ 1
│  │     ├─ ;
│  │     ├─ Statement
│  │     │  ├─ Ident
│  │     │  │  └─ i
│  │     │  ├─ :=
│  │     │  └─ Expression
│  │     │     └─ Term
│  │     │        └─ Factor
│  │     │           └─ Number
│  │     │              └─ 2
│  │     ├─ ;
│  │     ├─ Statement
│  │     │  ├─ while
│  │     │  ├─ Condition
│  │     │  │  ├─ Expression
│  │     │  │  │  └─ Term
│  │     │  │  │     └─ Factor
│  │     │  │  │        └─ Ident
│  │     │  │  │           └─ i
│  │     │  │  ├─ <
│  │     │  │  └─ Expression
│  │     │  │     └─ Term
│  │     │  │        └─ Factor
│  │     │  │           └─ Ident
│  │     │  │              └─ arg
│  │     │  ├─ do
│  │     │  └─ Statement
│  │     │     ├─ begin
│  │     │     ├─ Statement
│  │     │     │  ├─ if
│  │     │     │  ├─ Condition
│  │     │     │  │  ├─ Expression
│  │     │     │  │  │  └─ Term
│  │     │     │  │  │     ├─ Factor
│  │     │     │  │  │     │  └─ Ident
│  │     │     │  │  │     │     └─ arg
│  │     │     │  │  │     ├─ /
│  │     │     │  │  │     ├─ Factor
│  │     │     │  │  │     │  └─ Ident
│  │     │     │  │  │     │     └─ i
│  │     │     │  │  │     ├─ *
│  │     │     │  │  │     └─ Factor
│  │     │     │  │  │        └─ Ident
│  │     │     │  │  │           └─ i
│  │     │     │  │  ├─ =
│  │     │     │  │  └─ Expression
│  │     │     │  │     └─ Term
│  │     │     │  │        └─ Factor
│  │     │     │  │           └─ Ident
│  │     │     │  │              └─ arg
│  │     │     │  ├─ then
│  │     │     │  └─ Statement
│  │     │     │     ├─ begin
│  │     │     │     ├─ Statement
│  │     │     │     │  ├─ Ident
│  │     │     │     │  │  └─ ret
│  │     │     │     │  ├─ :=
│  │     │     │     │  └─ Expression
│  │     │     │     │     └─ Term
│  │     │     │     │        └─ Factor
│  │     │     │     │           └─ Number
│  │     │     │     │              └─ 0
│  │     │     │     ├─ ;
│  │     │     │     ├─ Statement
│  │     │     │     │  ├─ Ident
│  │     │     │     │  │  └─ i
│  │     │     │     │  ├─ :=
│  │     │     │     │  └─ Expression
│  │     │     │     │     └─ Term
│  │     │     │     │        └─ Factor
│  │     │     │     │           └─ Ident
│  │     │     │     │              └─ arg
│  │     │     │     └─ end
│  │     │     ├─ ;
│  │     │     ├─ Statement
│  │     │     │  ├─ Ident
│  │     │     │  │  └─ i
│  │     │     │  ├─ :=
│  │     │     │  └─ Expression
│  │     │     │     ├─ Term
│  │     │     │     │  └─ Factor
│  │     │     │     │     └─ Ident
│  │     │     │     │        └─ i
│  │     │     │     ├─ +
│  │     │     │     └─ Term
│  │     │     │        └─ Factor
│  │     │     │           └─ Number
│  │     │     │              └─ 1
│  │     │     └─ end
│  │     └─ end
│  ├─ ;
│  ├─ procedure
│  ├─ Ident
│  │  └─ primes
│  ├─ ;
│  ├─ Block
│  │  └─ Statement
│  │     ├─ begin
│  │     ├─ Statement
│  │     │  ├─ Ident
│  │     │  │  └─ arg
│  │     │  ├─ :=
│  │     │  └─ Expression
│  │     │     └─ Term
│  │     │        └─ Factor
│  │     │           └─ Number
│  │     │              └─ 2
│  │     ├─ ;
│  │     ├─ Statement
│  │     │  ├─ while
│  │     │  ├─ Condition
│  │     │  │  ├─ Expression
│  │     │  │  │  └─ Term
│  │     │  │  │     └─ Factor
│  │     │  │  │        └─ Ident
│  │     │  │  │           └─ arg
│  │     │  │  ├─ <
│  │     │  │  └─ Expression
│  │     │  │     └─ Term
│  │     │  │        └─ Factor
│  │     │  │           └─ Ident
│  │     │  │              └─ max
│  │     │  ├─ do
│  │     │  └─ Statement
│  │     │     ├─ begin
│  │     │     ├─ Statement
│  │     │     │  ├─ call
│  │     │     │  └─ Ident
│  │     │     │     └─ isprime
│  │     │     ├─ ;
│  │     │     ├─ Statement
│  │     │     │  ├─ if
│  │     │     │  ├─ Condition
│  │     │     │  │  ├─ Expression
│  │     │     │  │  │  └─ Term
│  │     │     │  │  │     └─ Factor
│  │     │     │  │  │        └─ Ident
│  │     │     │  │  │           └─ ret
│  │     │     │  │  ├─ =
│  │     │     │  │  └─ Expression
│  │     │     │  │     └─ Term
│  │     │     │  │        └─ Factor
│  │     │     │  │           └─ Number
│  │     │     │  │              └─ 1
│  │     │     │  ├─ then
│  │     │     │  └─ Statement
│  │     │     │     ├─ !
│  │     │     │     └─ Expression
│  │     │     │        └─ Term
│  │     │     │           └─ Factor
│  │     │     │              └─ Ident
│  │     │     │                 └─ arg
│  │     │     ├─ ;
│  │     │     ├─ Statement
│  │     │     │  ├─ Ident
│  │     │     │  │  └─ arg
│  │     │     │  ├─ :=
│  │     │     │  └─ Expression
│  │     │     │     ├─ Term
│  │     │     │     │  └─ Factor
│  │     │     │     │     └─ Ident
│  │     │     │     │        └─ arg
│  │     │     │     ├─ +
│  │     │     │     └─ Term
│  │     │     │        └─ Factor
│  │     │     │           └─ Number
│  │     │     │              └─ 1
│  │     │     └─ end
│  │     └─ end
│  ├─ ;
│  └─ Statement
│     ├─ call
│     └─ Ident
│        └─ primes
└─ .
`
	tokens, err := lexer.Lex(primeProgram)
	if err != nil {
		t.Error("lexing failed.", err)
	}
	// TODO: Validate debugTree as well
	parseTree, _, err := parser.Parse(tokens)
	if err != nil {
		t.Error("parsing failed.", err)
	}
	if parseTree.Sprint() != output {
		t.Errorf("invalid parse tree. want: %s. got: %s.", output, parseTree.Sprint())
	}
}

func TestSquareProgram(t *testing.T) {
	const output = `Program
├─ Block
│  ├─ var
│  ├─ Ident
│  │  └─ x
│  ├─ ,
│  ├─ Ident
│  │  └─ squ
│  ├─ ;
│  ├─ procedure
│  ├─ Ident
│  │  └─ square
│  ├─ ;
│  ├─ Block
│  │  └─ Statement
│  │     ├─ begin
│  │     ├─ Statement
│  │     │  ├─ Ident
│  │     │  │  └─ squ
│  │     │  ├─ :=
│  │     │  └─ Expression
│  │     │     └─ Term
│  │     │        ├─ Factor
│  │     │        │  └─ Ident
│  │     │        │     └─ x
│  │     │        ├─ *
│  │     │        └─ Factor
│  │     │           └─ Ident
│  │     │              └─ x
│  │     └─ end
│  ├─ ;
│  └─ Statement
│     ├─ begin
│     ├─ Statement
│     │  ├─ Ident
│     │  │  └─ x
│     │  ├─ :=
│     │  └─ Expression
│     │     └─ Term
│     │        └─ Factor
│     │           └─ Number
│     │              └─ 1
│     ├─ ;
│     ├─ Statement
│     │  ├─ while
│     │  ├─ Condition
│     │  │  ├─ Expression
│     │  │  │  └─ Term
│     │  │  │     └─ Factor
│     │  │  │        └─ Ident
│     │  │  │           └─ x
│     │  │  ├─ <=
│     │  │  └─ Expression
│     │  │     └─ Term
│     │  │        └─ Factor
│     │  │           └─ Number
│     │  │              └─ 10
│     │  ├─ do
│     │  └─ Statement
│     │     ├─ begin
│     │     ├─ Statement
│     │     │  ├─ call
│     │     │  └─ Ident
│     │     │     └─ square
│     │     ├─ ;
│     │     ├─ Statement
│     │     │  ├─ !
│     │     │  └─ Expression
│     │     │     └─ Term
│     │     │        └─ Factor
│     │     │           └─ Ident
│     │     │              └─ squ
│     │     ├─ ;
│     │     ├─ Statement
│     │     │  ├─ Ident
│     │     │  │  └─ x
│     │     │  ├─ :=
│     │     │  └─ Expression
│     │     │     ├─ Term
│     │     │     │  └─ Factor
│     │     │     │     └─ Ident
│     │     │     │        └─ x
│     │     │     ├─ +
│     │     │     └─ Term
│     │     │        └─ Factor
│     │     │           └─ Number
│     │     │              └─ 1
│     │     └─ end
│     └─ end
└─ .
`
	tokens, err := lexer.Lex(squareProgram)
	if err != nil {
		t.Error("lexing failed.", err)
	}
	// TODO: Validate debugTree as well
	parseTree, _, err := parser.Parse(tokens)
	if err != nil {
		t.Error("parsing failed.", err)
	}
	if parseTree.Sprint() != output {
		t.Errorf("invalid parse tree. want: %s. got: %s.", output, parseTree.Sprint())
	}
}
