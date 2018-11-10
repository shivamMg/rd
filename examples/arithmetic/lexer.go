package main

import (
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
			{`\d*\.\d+`, chroma.NumberFloat, nil},
			{`\d+`, chroma.NumberInteger, nil},
		},
	},
)

func Lex(expr string) (tokens []rd.Token, err error) {
	iter, err := lexer.Tokenise(nil, expr)
	if err != nil {
		return nil, err
	}
	for _, token := range iter.Tokens() {
		if token.Type == chroma.Error {
			return nil, fmt.Errorf("invalid token: %v", token)
		}
		tokens = append(tokens, token.Value)
	}
	return tokens, nil
}
