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
		printExit("invalid arguments. pass PL/0 program file as an argument")
	}
	code, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		printExit("could not open file", os.Args[1], "err:", err)
	}

	tokens, err := lexer.Lex(string(code))
	if err != nil {
		printExit("lexing failed.", err)
	}
	fmt.Println("Tokens:", tokens)

	fmt.Println("\nGrammar:", parser.Grammar)

	parseTree, debugTree, err := parser.Parse(tokens)
	if err != nil {
		fmt.Print("Debug Tree:\n\n", debugTree.Sprint())
		printExit("parsing failed.", err)
	}

	fmt.Println("Parse Tree:")
	parseTree.Print()
}

func printExit(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
	os.Exit(1)
}
