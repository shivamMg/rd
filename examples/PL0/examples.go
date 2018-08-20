// Examples copied from https://en.wikipedia.org/wiki/PL/0#Examples
package main

import (
	"strings"
	"github.com/alecthomas/chroma"
)

func squareProgram() (tokens []string) {
	code := `VAR x, squ;

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
	return lex(code)
}

func primeProgram() (tokens []string) {
	code := `const max = 100;
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
		if ret = 1 then write arg;
		arg := arg + 1
	end
end;

call primes
.`
	return lex(code)
}

func lex(code string) (tokens []string) {
	iter, err := Lexer.Tokenise(nil, code)
	if err != nil {
		panic(err)
	}
	for _, token := range iter.Tokens() {
		if token.Type != chroma.Text {
			tokens = append(tokens, strings.ToLower(token.Value))
		}
	}
	return tokens
}
