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
		printExit("invalid arguments. pass arithmetic expression as argument")
	}

	expr := strings.Join(os.Args[1:], " ")
	tokens, err := Lex(expr)
	if err != nil {
		printExit("Lexing failed.", err)
	}
	printTokens(tokens)

	fmt.Print("Grammar:")
	fmt.Print(parser.Grammar)

	parseTree, debugTree, err := parser.Parse(tokens)
	if err != nil {
		fmt.Print("Debug Tree:\n\n", debugTree)
		printExit("Parsing failed.", err)
	}
	fmt.Print("Parse Tree:\n\n", parseTree)
}

func printTokens(tokens []rd.Token) {
	fmt.Print("Tokens: ")
	var b strings.Builder
	for _, token := range tokens {
		b.WriteString(fmt.Sprint(token, " "))
	}
	fmt.Println(b.String())
}

func printExit(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
	os.Exit(1)
}
