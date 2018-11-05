package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/shivamMg/rd"
	"github.com/shivamMg/rd/examples/arithmetic/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("invalid arguments. pass arithmetic expression as argument")
		os.Exit(1)
	}

	expr := strings.Join(os.Args[1:], " ")
	tokens, err := Lex(expr)
	if err != nil {
		fmt.Println("Lexing failed.", err)
		os.Exit(1)
	}
	printTokens(tokens)

	fmt.Println("Grammar in EBNF:")
	fmt.Println(parser.Grammar)

	parseTree, debugTree, err := parser.Parse(tokens)
	if err != nil {
		fmt.Println("Parsing failed.", err)
		fmt.Print("Debug Tree:\n", debugTree, "\n")
		os.Exit(1)
	}
	fmt.Print("Parse Tree:\n\n")
	parseTree.Print()
}

func printTokens(tokens []rd.Token) {
	fmt.Print("Tokens: ")
	var b strings.Builder
	for _, token := range tokens {
		b.WriteString(fmt.Sprint(token, " "))
	}
	fmt.Println(b.String())
}
