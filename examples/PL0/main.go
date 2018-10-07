package main

import (
	"fmt"
	"os"

	"github.com/shivamMg/rd"
	"github.com/shivamMg/rd/examples/PL0/lexer"
	"github.com/shivamMg/rd/examples/PL0/parser"
	"io/ioutil"
)

// Grammar is PL/0's grammar in EBNF. Copied from https://en.wikipedia.org/wiki/PL/0#Grammar
const Grammar = `
	program = block "." .

	block =
		["const" ident "=" number {"," ident "=" number} ";"]
		["var" ident {"," ident} ";"]
		{"procedure" ident ";" block ";"} statement .

	statement =
		ident ":=" expression
		| "!" expression
		| "?" ident
		| "call" ident
		| "begin" statement {";" statement } "end"
		| "if" condition "then" statement
		| "while" condition "do" statement .

	condition =
		"odd" expression
		| expression ("="|"#"|"<"|"<="|">"|">=") expression .

	expression = ["+"|"-"] term {("+"|"-") term} .

	term = factor {("*"|"/") factor} .

	factor =
		ident
		| number
		| "(" expression ")" .
`

func main() {
	if len(os.Args) != 2 {
		fmt.Println("invalid arguments. pass PL/0 program file as an argument")
		os.Exit(1)
	}
	code, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("could not open file %q. err: %v", os.Args[1], err)
		os.Exit(1)
	}

	tokens, err := lexer.Lex(string(code))
	if err != nil {
		fmt.Println("lexing failed.", err)
		os.Exit(1)
	}
	fmt.Println("Tokens:", tokens)

	fmt.Println("\nGrammar:", Grammar)

	parseTree, err := parser.Parse(tokens)
	if err != nil {
		fmt.Println("parsing failed.", err)
		if e, ok := err.(*rd.ParsingError); ok {
			fmt.Println("debug tree:")
			e.PrintDebugTree()
		}
		os.Exit(1)
	}

	fmt.Println("Parse Tree:")
	parseTree.Print()
}
