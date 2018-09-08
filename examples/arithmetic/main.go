package main

import (
	"fmt"
	"log"
)

func main() {
	expr := "2.8+ (3 - .733)/ 23"
	// expr := "2 - 3"
	tokens, err := Lex(expr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%q\n", expr)
	fmt.Println(tokens)
	p, err := Parse(tokens)
	if err != nil {
		log.Fatal(err)
	}
	p.PrintTree()
	// p.PrintFlowTree()
}
