package main

//go:generate stringer -type Token -linecomment

type Token int

const (
	Plus       Token = iota // +
	Minus                   // -
	Mul                     // *
	Div                     // /
	OpenParen               // (
	CloseParen              // )
)

var TokenStrings = map[string]Token{
	Plus.String():       Plus,
	Minus.String():      Minus,
	Mul.String():        Mul,
	Div.String():        Div,
	OpenParen.String():  OpenParen,
	CloseParen.String(): CloseParen,
}
