package main

import (
	"bufio"
	"strings"
)

// Examples copied from https://en.wikipedia.org/wiki/PL/0#Examples
// Spaces between tokens have been added to avoid writing a lexer for the
// language.
var examples = []string{
	`var x , squ ;

	procedure square ;
	begin
	   squ := x * x
	end ;

	begin
	   x := 1 ;
	   while x <= 10 do
	   begin
		  call square ;
		  x := x + 1
	   end
	end .`,

	`const max = 100 ;
	var arg , ret ;

	procedure isprime ;
	var i ;
	begin
		ret := 1 ;
		i := 2 ;
		while i < arg do
		begin
			if arg / i * i = arg then
			begin
				ret := 0 ;
				i := arg
			end ;
			i := i + 1
		end
	end ;

	procedure primes ;
	begin
		arg := 2 ;
		while arg < max do
		begin
			call isprime ;
			if ret = 1 then ! arg ;
			arg := arg + 1
		end
	end ;

	call primes
	.`,
}

func squareProgram() (tokens []string) {
	return lex(examples[0])
}

func primeProgram() (tokens []string) {
	return lex(examples[1])
}

func lex(input string) (tokens []string) {
	// TODO: write a lexer for the language
	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		tokens = append(tokens, strings.Fields(s.Text())...)
	}
	return
}
