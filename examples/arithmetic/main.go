package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/shivamMg/rd"
)

// Grammar without recursion. Left-factored. Needs single lookahead. Suitable for R.D. parsing.
const Grammar = `
	Expr   = Term Expr'
	Expr'  = "+" Expr | "-" Expr | ε
	Term   = Factor Term'
	Term'  = "*" Term | "/" Term | ε
	Factor = "(" Expr ")" | "-" Factor | Number
`

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
	fmt.Println(Grammar)

	parseTree, err := Parse(tokens)
	if err != nil {
		fmt.Println("Parsing failed.", err)
		if e, ok := err.(*rd.ParsingError); ok {
			fmt.Print("Debug Tree:\n\n")
			e.PrintDebugTree()
		}
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
