package main

import (
	"errors"
	"github.com/alecthomas/chroma"
	"github.com/shivamMg/rd"
	"strconv"
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
	token := iter()
	for token != nil {
		switch token.Type {
		case chroma.Operator, chroma.Punctuation:
			tokens = append(tokens, token.Value)
		case chroma.NumberInteger:
			num, err := strconv.ParseInt(token.Value, 10, 64)
			if err != nil {
				return nil, errors.New("token couldn't be converted to 64bit decimal")
			}
			tokens = append(tokens, num)
		case chroma.NumberFloat:
			num, err := strconv.ParseFloat(token.Value, 64)
			if err != nil {
				return nil, errors.New("token couldn't be converted to 64bit float")
			}
			tokens = append(tokens, num)
		case chroma.Error:
			return nil, errors.New("invalid token")
		}
		token = iter()
	}
	return
}
