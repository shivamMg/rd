package main

import (
	"fmt"
	"os"

	"io/ioutil"

	"github.com/shivamMg/rd/examples/pl0/lexer"
	"github.com/shivamMg/rd/examples/pl0/parser"
)

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

	fmt.Println("\nGrammar:", parser.Grammar)

	parseTree, debugTree, err := parser.Parse(tokens)
	if err != nil {
		fmt.Println("parsing failed.", err)
		fmt.Println("debug tree:\n", debugTree)
		os.Exit(1)
	}

	fmt.Println("Parse Tree:")
	parseTree.Print()
}
