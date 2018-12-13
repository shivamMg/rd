package main

import (
	"fmt"
	"os"
	"strings"

	"flag"

	"github.com/shivamMg/rd"
	"github.com/shivamMg/rd/examples/arithmetic/backtrackingparser"
	"github.com/shivamMg/rd/examples/arithmetic/parser"
)

var (
	useBacktrackingParser = flag.Bool("backtrackingparser", false, "use backtracking parser")
	expr                  = flag.String("expr", "", "arithmetic expression to be parsed")
)

func main() {
	flag.Parse()
	if *expr == "" {
		printExit("empty expr flag. pass arithmetic expression as expr. ex. -expr='1+2'")
	}

	tokens, err := Lex(*expr)
	if err != nil {
		printExit("Lexing failed.", err)
	}
	printTokens(tokens)

	fmt.Print("Grammar:")
	fmt.Print(parser.Grammar)

	var parseTree *rd.Tree
	var debugTree *rd.DebugTree
	if *useBacktrackingParser {
		parseTree, debugTree, err = backtrackingparser.Parse(tokens)
	} else {
		parseTree, debugTree, err = parser.Parse(tokens)
	}
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
