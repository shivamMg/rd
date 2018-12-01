package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/shivamMg/rd"
)

func main() {
	if len(os.Args) != 2 {
		printExit("invalid arguments. pass a domain name as an argument")
	}
	s := NewScanner(strings.NewReader(os.Args[1]))
	var tokens []rd.Token
	for {
		token, lit, err := s.Scan()
		if err != nil {
			panic(err)
		}
		switch token {
		case EOF:
			goto parsing
		case Whitespace, Dot, Hyphen:
			tokens = append(tokens, token)
		case Letter, Digit:
			tokens = append(tokens, lit)
		case Illegal:
			printExit("lexing failed. illegal token: " + string(lit))
		}
	}
parsing:
	fmt.Print("Tokens: ")
	for _, token := range tokens {
		switch v := token.(type) {
		case rune:
			fmt.Printf("%s ", string(v))
		case Token:
			fmt.Printf("%s ", v)
		}
	}
	fmt.Println()
	fmt.Println("Grammar:\n", grammar)

	b = rd.NewBuilder(tokens)
	enter = b.Enter
	exit = b.Exit
	match = b.Match
	next = b.Next
	add = b.Add
	checkOrNotOK = b.CheckOrNotOK
	if ok := domain(); !ok || b.Err() != nil {
		fmt.Print("Debug tree:\n\n", b.DebugTree())
		printExit("parsing failed.", b.Err())
	}
	fmt.Print("Parse tree:\n\n", b.ParseTree())
}

func printExit(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
	os.Exit(1)
}
