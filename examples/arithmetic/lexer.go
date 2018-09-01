package main

import (
	"errors"
	"log"

	"fmt"

	"github.com/alecthomas/chroma"
	"github.com/shivamMg/rd"
)

var lexer = chroma.MustNewLexer(
	&chroma.Config{
		Name: "Arithmetic Expressions",
	},
	chroma.Rules{
		"root": {
			{`\s+`, chroma.Text, nil},
			{`[+\-*/]`, chroma.Operator, nil},
			{`[()]`, chroma.Punctuation, nil},
			{`(\d*\.\d+|\d+)`, chroma.Number, nil},
		},
	},
)

func Lex(expr string) (tokens []rd.Token, err error) {
	iter, err := lexer.Tokenise(nil, expr)
	if err != nil {
		return nil, err
	}
	token := iter()
	for token != nil {
		switch token.Type {
		case chroma.Operator, chroma.Punctuation, chroma.Number:
			tokens = append(tokens, token.Value)
		case chroma.Error:
			return nil, errors.New("invalid token")
		}
		token = iter()
	}
	return
}

func main() {
	expr := "2.8+ (3 - .733)/ 23"
	tokens, err := Lex(expr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%q\n", expr)
	for _, token := range tokens {
		fmt.Println(token)
	}
	fmt.Println("parsing")
	Parse(tokens)
}
