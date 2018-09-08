package main

import (
	"fmt"
	"log"
)

func main() {
	expr := "2.8 + (3 - .733) / 23"
	tokens, err := Lex(expr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tokens:", tokens)
	parseTree, err := Parse(tokens)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Parse tree:")
	parseTree.Print()
}
