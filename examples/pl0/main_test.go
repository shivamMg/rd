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

	squareProgramParseTree = `Program
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

	squareProgramDebugTree = `Program(true)
├─ Block(true)
│  ├─ var ≠ const
│  ├─ var = var
│  ├─ Ident(true)
│  ├─ , = ,
│  ├─ Ident(true)
│  ├─ ; ≠ ,
│  ├─ ; = ;
│  ├─ procedure = procedure
│  ├─ Ident(true)
│  ├─ ; = ;
│  ├─ Block(true)
│  │  ├─ begin ≠ const
│  │  ├─ begin ≠ var
│  │  ├─ begin ≠ procedure
│  │  └─ Statement(true)
│  │     ├─ Ident(false)
│  │     ├─ begin ≠ !
│  │     ├─ begin ≠ ?
│  │     ├─ begin ≠ call
│  │     ├─ begin = begin
│  │     ├─ Statement(true)
│  │     │  ├─ Ident(true)
│  │     │  ├─ := = :=
│  │     │  └─ Expression(true)
│  │     │     ├─ x ≠ +
│  │     │     ├─ x ≠ -
│  │     │     ├─ Term(true)
│  │     │     │  ├─ Factor(true)
│  │     │     │  │  └─ Ident(true)
│  │     │     │  ├─ * = *
│  │     │     │  ├─ Factor(true)
│  │     │     │  │  └─ Ident(true)
│  │     │     │  ├─ end ≠ *
│  │     │     │  └─ end ≠ /
│  │     │     ├─ end ≠ +
│  │     │     └─ end ≠ -
│  │     ├─ end ≠ ;
│  │     └─ end = end
│  ├─ ; = ;
│  ├─ begin ≠ procedure
│  └─ Statement(true)
│     ├─ Ident(false)
│     ├─ begin ≠ !
│     ├─ begin ≠ ?
│     ├─ begin ≠ call
│     ├─ begin = begin
│     ├─ Statement(true)
│     │  ├─ Ident(true)
│     │  ├─ := = :=
│     │  └─ Expression(true)
│     │     ├─ 1 ≠ +
│     │     ├─ 1 ≠ -
│     │     ├─ Term(true)
│     │     │  ├─ Factor(true)
│     │     │  │  ├─ Ident(false)
│     │     │  │  └─ Number(true)
│     │     │  ├─ ; ≠ *
│     │     │  └─ ; ≠ /
│     │     ├─ ; ≠ +
│     │     └─ ; ≠ -
│     ├─ ; = ;
│     ├─ Statement(true)
│     │  ├─ Ident(false)
│     │  ├─ while ≠ !
│     │  ├─ while ≠ ?
│     │  ├─ while ≠ call
│     │  ├─ while ≠ begin
│     │  ├─ while ≠ if
│     │  ├─ while = while
│     │  ├─ Condition(true)
│     │  │  ├─ x ≠ odd
│     │  │  ├─ Expression(true)
│     │  │  │  ├─ x ≠ +
│     │  │  │  ├─ x ≠ -
│     │  │  │  ├─ Term(true)
│     │  │  │  │  ├─ Factor(true)
│     │  │  │  │  │  └─ Ident(true)
│     │  │  │  │  ├─ <= ≠ *
│     │  │  │  │  └─ <= ≠ /
│     │  │  │  ├─ <= ≠ +
│     │  │  │  └─ <= ≠ -
│     │  │  ├─ <= ≠ =
│     │  │  ├─ <= ≠ #
│     │  │  ├─ <= ≠ <
│     │  │  ├─ <= = <=
│     │  │  └─ Expression(true)
│     │  │     ├─ 10 ≠ +
│     │  │     ├─ 10 ≠ -
│     │  │     ├─ Term(true)
│     │  │     │  ├─ Factor(true)
│     │  │     │  │  ├─ Ident(false)
│     │  │     │  │  └─ Number(true)
│     │  │     │  ├─ do ≠ *
│     │  │     │  └─ do ≠ /
│     │  │     ├─ do ≠ +
│     │  │     └─ do ≠ -
│     │  ├─ do = do
│     │  └─ Statement(true)
│     │     ├─ Ident(false)
│     │     ├─ begin ≠ !
│     │     ├─ begin ≠ ?
│     │     ├─ begin ≠ call
│     │     ├─ begin = begin
│     │     ├─ Statement(true)
│     │     │  ├─ Ident(false)
│     │     │  ├─ call ≠ !
│     │     │  ├─ call ≠ ?
│     │     │  ├─ call = call
│     │     │  └─ Ident(true)
│     │     ├─ ; = ;
│     │     ├─ Statement(true)
│     │     │  ├─ Ident(false)
│     │     │  ├─ ! = !
│     │     │  └─ Expression(true)
│     │     │     ├─ squ ≠ +
│     │     │     ├─ squ ≠ -
│     │     │     ├─ Term(true)
│     │     │     │  ├─ Factor(true)
│     │     │     │  │  └─ Ident(true)
│     │     │     │  ├─ ; ≠ *
│     │     │     │  └─ ; ≠ /
│     │     │     ├─ ; ≠ +
│     │     │     └─ ; ≠ -
│     │     ├─ ; = ;
│     │     ├─ Statement(true)
│     │     │  ├─ Ident(true)
│     │     │  ├─ := = :=
│     │     │  └─ Expression(true)
│     │     │     ├─ x ≠ +
│     │     │     ├─ x ≠ -
│     │     │     ├─ Term(true)
│     │     │     │  ├─ Factor(true)
│     │     │     │  │  └─ Ident(true)
│     │     │     │  ├─ + ≠ *
│     │     │     │  └─ + ≠ /
│     │     │     ├─ + = +
│     │     │     ├─ Term(true)
│     │     │     │  ├─ Factor(true)
│     │     │     │  │  ├─ Ident(false)
│     │     │     │  │  └─ Number(true)
│     │     │     │  ├─ end ≠ *
│     │     │     │  └─ end ≠ /
│     │     │     ├─ end ≠ +
│     │     │     └─ end ≠ -
│     │     ├─ end ≠ ;
│     │     └─ end = end
│     ├─ end ≠ ;
│     └─ end = end
└─ . = .
`

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

	primeProgramParseTree = `Program
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

	primeProgramDebugTree = `Program(true)
├─ Block(true)
│  ├─ const = const
│  ├─ Ident(true)
│  ├─ = = =
│  ├─ Number(true)
│  ├─ ; ≠ ,
│  ├─ ; = ;
│  ├─ var = var
│  ├─ Ident(true)
│  ├─ , = ,
│  ├─ Ident(true)
│  ├─ ; ≠ ,
│  ├─ ; = ;
│  ├─ procedure = procedure
│  ├─ Ident(true)
│  ├─ ; = ;
│  ├─ Block(true)
│  │  ├─ var ≠ const
│  │  ├─ var = var
│  │  ├─ Ident(true)
│  │  ├─ ; ≠ ,
│  │  ├─ ; = ;
│  │  ├─ begin ≠ procedure
│  │  └─ Statement(true)
│  │     ├─ Ident(false)
│  │     ├─ begin ≠ !
│  │     ├─ begin ≠ ?
│  │     ├─ begin ≠ call
│  │     ├─ begin = begin
│  │     ├─ Statement(true)
│  │     │  ├─ Ident(true)
│  │     │  ├─ := = :=
│  │     │  └─ Expression(true)
│  │     │     ├─ 1 ≠ +
│  │     │     ├─ 1 ≠ -
│  │     │     ├─ Term(true)
│  │     │     │  ├─ Factor(true)
│  │     │     │  │  ├─ Ident(false)
│  │     │     │  │  └─ Number(true)
│  │     │     │  ├─ ; ≠ *
│  │     │     │  └─ ; ≠ /
│  │     │     ├─ ; ≠ +
│  │     │     └─ ; ≠ -
│  │     ├─ ; = ;
│  │     ├─ Statement(true)
│  │     │  ├─ Ident(true)
│  │     │  ├─ := = :=
│  │     │  └─ Expression(true)
│  │     │     ├─ 2 ≠ +
│  │     │     ├─ 2 ≠ -
│  │     │     ├─ Term(true)
│  │     │     │  ├─ Factor(true)
│  │     │     │  │  ├─ Ident(false)
│  │     │     │  │  └─ Number(true)
│  │     │     │  ├─ ; ≠ *
│  │     │     │  └─ ; ≠ /
│  │     │     ├─ ; ≠ +
│  │     │     └─ ; ≠ -
│  │     ├─ ; = ;
│  │     ├─ Statement(true)
│  │     │  ├─ Ident(false)
│  │     │  ├─ while ≠ !
│  │     │  ├─ while ≠ ?
│  │     │  ├─ while ≠ call
│  │     │  ├─ while ≠ begin
│  │     │  ├─ while ≠ if
│  │     │  ├─ while = while
│  │     │  ├─ Condition(true)
│  │     │  │  ├─ i ≠ odd
│  │     │  │  ├─ Expression(true)
│  │     │  │  │  ├─ i ≠ +
│  │     │  │  │  ├─ i ≠ -
│  │     │  │  │  ├─ Term(true)
│  │     │  │  │  │  ├─ Factor(true)
│  │     │  │  │  │  │  └─ Ident(true)
│  │     │  │  │  │  ├─ < ≠ *
│  │     │  │  │  │  └─ < ≠ /
│  │     │  │  │  ├─ < ≠ +
│  │     │  │  │  └─ < ≠ -
│  │     │  │  ├─ < ≠ =
│  │     │  │  ├─ < ≠ #
│  │     │  │  ├─ < = <
│  │     │  │  └─ Expression(true)
│  │     │  │     ├─ arg ≠ +
│  │     │  │     ├─ arg ≠ -
│  │     │  │     ├─ Term(true)
│  │     │  │     │  ├─ Factor(true)
│  │     │  │     │  │  └─ Ident(true)
│  │     │  │     │  ├─ do ≠ *
│  │     │  │     │  └─ do ≠ /
│  │     │  │     ├─ do ≠ +
│  │     │  │     └─ do ≠ -
│  │     │  ├─ do = do
│  │     │  └─ Statement(true)
│  │     │     ├─ Ident(false)
│  │     │     ├─ begin ≠ !
│  │     │     ├─ begin ≠ ?
│  │     │     ├─ begin ≠ call
│  │     │     ├─ begin = begin
│  │     │     ├─ Statement(true)
│  │     │     │  ├─ Ident(false)
│  │     │     │  ├─ if ≠ !
│  │     │     │  ├─ if ≠ ?
│  │     │     │  ├─ if ≠ call
│  │     │     │  ├─ if ≠ begin
│  │     │     │  ├─ if = if
│  │     │     │  ├─ Condition(true)
│  │     │     │  │  ├─ arg ≠ odd
│  │     │     │  │  ├─ Expression(true)
│  │     │     │  │  │  ├─ arg ≠ +
│  │     │     │  │  │  ├─ arg ≠ -
│  │     │     │  │  │  ├─ Term(true)
│  │     │     │  │  │  │  ├─ Factor(true)
│  │     │     │  │  │  │  │  └─ Ident(true)
│  │     │     │  │  │  │  ├─ / ≠ *
│  │     │     │  │  │  │  ├─ / = /
│  │     │     │  │  │  │  ├─ Factor(true)
│  │     │     │  │  │  │  │  └─ Ident(true)
│  │     │     │  │  │  │  ├─ * = *
│  │     │     │  │  │  │  ├─ Factor(true)
│  │     │     │  │  │  │  │  └─ Ident(true)
│  │     │     │  │  │  │  ├─ = ≠ *
│  │     │     │  │  │  │  └─ = ≠ /
│  │     │     │  │  │  ├─ = ≠ +
│  │     │     │  │  │  └─ = ≠ -
│  │     │     │  │  ├─ = = =
│  │     │     │  │  └─ Expression(true)
│  │     │     │  │     ├─ arg ≠ +
│  │     │     │  │     ├─ arg ≠ -
│  │     │     │  │     ├─ Term(true)
│  │     │     │  │     │  ├─ Factor(true)
│  │     │     │  │     │  │  └─ Ident(true)
│  │     │     │  │     │  ├─ then ≠ *
│  │     │     │  │     │  └─ then ≠ /
│  │     │     │  │     ├─ then ≠ +
│  │     │     │  │     └─ then ≠ -
│  │     │     │  ├─ then = then
│  │     │     │  └─ Statement(true)
│  │     │     │     ├─ Ident(false)
│  │     │     │     ├─ begin ≠ !
│  │     │     │     ├─ begin ≠ ?
│  │     │     │     ├─ begin ≠ call
│  │     │     │     ├─ begin = begin
│  │     │     │     ├─ Statement(true)
│  │     │     │     │  ├─ Ident(true)
│  │     │     │     │  ├─ := = :=
│  │     │     │     │  └─ Expression(true)
│  │     │     │     │     ├─ 0 ≠ +
│  │     │     │     │     ├─ 0 ≠ -
│  │     │     │     │     ├─ Term(true)
│  │     │     │     │     │  ├─ Factor(true)
│  │     │     │     │     │  │  ├─ Ident(false)
│  │     │     │     │     │  │  └─ Number(true)
│  │     │     │     │     │  ├─ ; ≠ *
│  │     │     │     │     │  └─ ; ≠ /
│  │     │     │     │     ├─ ; ≠ +
│  │     │     │     │     └─ ; ≠ -
│  │     │     │     ├─ ; = ;
│  │     │     │     ├─ Statement(true)
│  │     │     │     │  ├─ Ident(true)
│  │     │     │     │  ├─ := = :=
│  │     │     │     │  └─ Expression(true)
│  │     │     │     │     ├─ arg ≠ +
│  │     │     │     │     ├─ arg ≠ -
│  │     │     │     │     ├─ Term(true)
│  │     │     │     │     │  ├─ Factor(true)
│  │     │     │     │     │  │  └─ Ident(true)
│  │     │     │     │     │  ├─ end ≠ *
│  │     │     │     │     │  └─ end ≠ /
│  │     │     │     │     ├─ end ≠ +
│  │     │     │     │     └─ end ≠ -
│  │     │     │     ├─ end ≠ ;
│  │     │     │     └─ end = end
│  │     │     ├─ ; = ;
│  │     │     ├─ Statement(true)
│  │     │     │  ├─ Ident(true)
│  │     │     │  ├─ := = :=
│  │     │     │  └─ Expression(true)
│  │     │     │     ├─ i ≠ +
│  │     │     │     ├─ i ≠ -
│  │     │     │     ├─ Term(true)
│  │     │     │     │  ├─ Factor(true)
│  │     │     │     │  │  └─ Ident(true)
│  │     │     │     │  ├─ + ≠ *
│  │     │     │     │  └─ + ≠ /
│  │     │     │     ├─ + = +
│  │     │     │     ├─ Term(true)
│  │     │     │     │  ├─ Factor(true)
│  │     │     │     │  │  ├─ Ident(false)
│  │     │     │     │  │  └─ Number(true)
│  │     │     │     │  ├─ end ≠ *
│  │     │     │     │  └─ end ≠ /
│  │     │     │     ├─ end ≠ +
│  │     │     │     └─ end ≠ -
│  │     │     ├─ end ≠ ;
│  │     │     └─ end = end
│  │     ├─ end ≠ ;
│  │     └─ end = end
│  ├─ ; = ;
│  ├─ procedure = procedure
│  ├─ Ident(true)
│  ├─ ; = ;
│  ├─ Block(true)
│  │  ├─ begin ≠ const
│  │  ├─ begin ≠ var
│  │  ├─ begin ≠ procedure
│  │  └─ Statement(true)
│  │     ├─ Ident(false)
│  │     ├─ begin ≠ !
│  │     ├─ begin ≠ ?
│  │     ├─ begin ≠ call
│  │     ├─ begin = begin
│  │     ├─ Statement(true)
│  │     │  ├─ Ident(true)
│  │     │  ├─ := = :=
│  │     │  └─ Expression(true)
│  │     │     ├─ 2 ≠ +
│  │     │     ├─ 2 ≠ -
│  │     │     ├─ Term(true)
│  │     │     │  ├─ Factor(true)
│  │     │     │  │  ├─ Ident(false)
│  │     │     │  │  └─ Number(true)
│  │     │     │  ├─ ; ≠ *
│  │     │     │  └─ ; ≠ /
│  │     │     ├─ ; ≠ +
│  │     │     └─ ; ≠ -
│  │     ├─ ; = ;
│  │     ├─ Statement(true)
│  │     │  ├─ Ident(false)
│  │     │  ├─ while ≠ !
│  │     │  ├─ while ≠ ?
│  │     │  ├─ while ≠ call
│  │     │  ├─ while ≠ begin
│  │     │  ├─ while ≠ if
│  │     │  ├─ while = while
│  │     │  ├─ Condition(true)
│  │     │  │  ├─ arg ≠ odd
│  │     │  │  ├─ Expression(true)
│  │     │  │  │  ├─ arg ≠ +
│  │     │  │  │  ├─ arg ≠ -
│  │     │  │  │  ├─ Term(true)
│  │     │  │  │  │  ├─ Factor(true)
│  │     │  │  │  │  │  └─ Ident(true)
│  │     │  │  │  │  ├─ < ≠ *
│  │     │  │  │  │  └─ < ≠ /
│  │     │  │  │  ├─ < ≠ +
│  │     │  │  │  └─ < ≠ -
│  │     │  │  ├─ < ≠ =
│  │     │  │  ├─ < ≠ #
│  │     │  │  ├─ < = <
│  │     │  │  └─ Expression(true)
│  │     │  │     ├─ max ≠ +
│  │     │  │     ├─ max ≠ -
│  │     │  │     ├─ Term(true)
│  │     │  │     │  ├─ Factor(true)
│  │     │  │     │  │  └─ Ident(true)
│  │     │  │     │  ├─ do ≠ *
│  │     │  │     │  └─ do ≠ /
│  │     │  │     ├─ do ≠ +
│  │     │  │     └─ do ≠ -
│  │     │  ├─ do = do
│  │     │  └─ Statement(true)
│  │     │     ├─ Ident(false)
│  │     │     ├─ begin ≠ !
│  │     │     ├─ begin ≠ ?
│  │     │     ├─ begin ≠ call
│  │     │     ├─ begin = begin
│  │     │     ├─ Statement(true)
│  │     │     │  ├─ Ident(false)
│  │     │     │  ├─ call ≠ !
│  │     │     │  ├─ call ≠ ?
│  │     │     │  ├─ call = call
│  │     │     │  └─ Ident(true)
│  │     │     ├─ ; = ;
│  │     │     ├─ Statement(true)
│  │     │     │  ├─ Ident(false)
│  │     │     │  ├─ if ≠ !
│  │     │     │  ├─ if ≠ ?
│  │     │     │  ├─ if ≠ call
│  │     │     │  ├─ if ≠ begin
│  │     │     │  ├─ if = if
│  │     │     │  ├─ Condition(true)
│  │     │     │  │  ├─ ret ≠ odd
│  │     │     │  │  ├─ Expression(true)
│  │     │     │  │  │  ├─ ret ≠ +
│  │     │     │  │  │  ├─ ret ≠ -
│  │     │     │  │  │  ├─ Term(true)
│  │     │     │  │  │  │  ├─ Factor(true)
│  │     │     │  │  │  │  │  └─ Ident(true)
│  │     │     │  │  │  │  ├─ = ≠ *
│  │     │     │  │  │  │  └─ = ≠ /
│  │     │     │  │  │  ├─ = ≠ +
│  │     │     │  │  │  └─ = ≠ -
│  │     │     │  │  ├─ = = =
│  │     │     │  │  └─ Expression(true)
│  │     │     │  │     ├─ 1 ≠ +
│  │     │     │  │     ├─ 1 ≠ -
│  │     │     │  │     ├─ Term(true)
│  │     │     │  │     │  ├─ Factor(true)
│  │     │     │  │     │  │  ├─ Ident(false)
│  │     │     │  │     │  │  └─ Number(true)
│  │     │     │  │     │  ├─ then ≠ *
│  │     │     │  │     │  └─ then ≠ /
│  │     │     │  │     ├─ then ≠ +
│  │     │     │  │     └─ then ≠ -
│  │     │     │  ├─ then = then
│  │     │     │  └─ Statement(true)
│  │     │     │     ├─ Ident(false)
│  │     │     │     ├─ ! = !
│  │     │     │     └─ Expression(true)
│  │     │     │        ├─ arg ≠ +
│  │     │     │        ├─ arg ≠ -
│  │     │     │        ├─ Term(true)
│  │     │     │        │  ├─ Factor(true)
│  │     │     │        │  │  └─ Ident(true)
│  │     │     │        │  ├─ ; ≠ *
│  │     │     │        │  └─ ; ≠ /
│  │     │     │        ├─ ; ≠ +
│  │     │     │        └─ ; ≠ -
│  │     │     ├─ ; = ;
│  │     │     ├─ Statement(true)
│  │     │     │  ├─ Ident(true)
│  │     │     │  ├─ := = :=
│  │     │     │  └─ Expression(true)
│  │     │     │     ├─ arg ≠ +
│  │     │     │     ├─ arg ≠ -
│  │     │     │     ├─ Term(true)
│  │     │     │     │  ├─ Factor(true)
│  │     │     │     │  │  └─ Ident(true)
│  │     │     │     │  ├─ + ≠ *
│  │     │     │     │  └─ + ≠ /
│  │     │     │     ├─ + = +
│  │     │     │     ├─ Term(true)
│  │     │     │     │  ├─ Factor(true)
│  │     │     │     │  │  ├─ Ident(false)
│  │     │     │     │  │  └─ Number(true)
│  │     │     │     │  ├─ end ≠ *
│  │     │     │     │  └─ end ≠ /
│  │     │     │     ├─ end ≠ +
│  │     │     │     └─ end ≠ -
│  │     │     ├─ end ≠ ;
│  │     │     └─ end = end
│  │     ├─ end ≠ ;
│  │     └─ end = end
│  ├─ ; = ;
│  ├─ call ≠ procedure
│  └─ Statement(true)
│     ├─ Ident(false)
│     ├─ call ≠ !
│     ├─ call ≠ ?
│     ├─ call = call
│     └─ Ident(true)
└─ . = .
`
)

func TestPrograms(t *testing.T) {
	tests := []struct {
		inputProgram      string
		expectedParseTree string
		expectedDebugTree string
	}{
		{squareProgram, squareProgramParseTree, squareProgramDebugTree},
		{primeProgram, primeProgramParseTree, primeProgramDebugTree},
	}

	for _, test := range tests {
		tokens, err := lexer.Lex(test.inputProgram)
		if err != nil {
			t.Error("lexing failed.", err)
		}
		parseTree, debugTree, err := parser.Parse(tokens)
		if err != nil {
			t.Error("parsing failed.", err)
		}
		if got := parseTree.String(); got != test.expectedParseTree {
			t.Errorf("invalid parse tree. expected: %s. got: %s.", test.expectedParseTree, got)
		}
		if got := debugTree.String(); got != test.expectedDebugTree {
			t.Errorf("invalid debug tree. expected: %s. got: %s.", test.expectedDebugTree, got)
		}
	}
}
