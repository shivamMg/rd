package main

import (
	"bufio"
	"strings"
)

// Inputs copied from https://en.wikipedia.org/wiki/PL/0#Examples
// Spaces between tokens have been added to avoid writing a lexer for the
// language.
var inputs = []string{
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

	`const max = 100;
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
	.`,
}

func squareProgram() (tokens []string) {
	s := bufio.NewScanner(strings.NewReader(inputs[0]))
	for s.Scan() {
		tokens = append(tokens, strings.Fields(s.Text())...)
	}
	return
}
