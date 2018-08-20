package main

import (
	"github.com/alecthomas/chroma"
)

var Lexer = chroma.MustNewLexer(
	&chroma.Config{
		Name: "PL0",
		Aliases: []string{"pl0"},
		Filenames: []string{"*.pl0"},
		MimeTypes: []string{"text/x-pl0src"},
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

