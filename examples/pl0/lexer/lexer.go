package lexer

import (
	"fmt"

	"strings"

	"github.com/alecthomas/chroma"
	"github.com/shivamMg/rd"
	pl0Tokens "github.com/shivamMg/rd/examples/pl0/tokens"
)

var lexer = chroma.MustNewLexer(
	&chroma.Config{
		Name:            "PL0",
		Aliases:         []string{"pl0"},
		Filenames:       []string{"*.pl0"},
		MimeTypes:       []string{"text/x-pl0src"},
		CaseInsensitive: true,
	},
	chroma.Rules{
		"root": {
			{`\n`, chroma.Text, nil},
			{`\s+`, chroma.Text, nil},
			{`(var|procedure|const)\b`, chroma.KeywordDeclaration, nil},
			{chroma.Words(``, `\b`, `call`, `begin`, `end`, `if`, `then`, `while`,
				`do`, `odd`), chroma.KeywordReserved, nil},
			{`[.,;()]`, chroma.Punctuation, nil},
			{`(:=|#|<=|>=|<|>|=|[+\-*/]|[!?])`, chroma.Operator, nil},
			{`(0|[1-9]\d*)`, chroma.Number, nil},
			{`[^\W\d]\w*`, chroma.NameVariable, nil},
		},
	},
)

func Lex(code string) ([]rd.Token, error) {
	var tokens []rd.Token
	iter, err := lexer.Tokenise(nil, code)
	if err != nil {
		return nil, err
	}
	for _, token := range iter.Tokens() {
		switch token.Type {
		case chroma.Text:
		case chroma.Number, chroma.NameVariable:
			tokens = append(tokens, token.Value)
		case chroma.Error:
			return nil, fmt.Errorf("invalid token: %v", token)
		default:
			pl0Token, ok := pl0Tokens.TokenFromString(strings.ToLower(token.Value))
			if !ok {
				return nil, fmt.Errorf("invalid token: %v", token)
			}
			tokens = append(tokens, pl0Token)
		}
	}
	return tokens, nil
}
