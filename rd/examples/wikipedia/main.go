/*
Recursive Descent parser for the following grammar.
The grammar has been copied from https://en.wikipedia.org/wiki/Recursive_descent_parser#Example_parser

program = block "." .

block =
    ["const" ident "=" number {"," ident "=" number} ";"]
    ["var" ident {"," ident} ";"]
    {"procedure" ident ";" block ";"} statement .

statement =
    ident ":=" expression
    | "call" ident
    | "begin" statement {";" statement } "end"
    | "if" condition "then" statement
    | "while" condition "do" statement .

condition =
    "odd" expression
    | expression ("="|"#"|"<"|"<="|">"|">=") expression .

expression = ["+"|"-"] term {("+"|"-") term} .

term = factor {("*"|"/") factor} .

factor =
    ident
    | number
    | "(" expression ")" .
*/
package main

import (
	"encoding/json"
	"fmt"

	"github.com/shivammg/parsers/rd"
)

// Non-terminals
const (
	Program    = "program"
	Block      = "block"
	Statement  = "statement"
	Condition  = "condition"
	Expression = "expression"
	Term       = "term"
	Factor     = "factor"
	Ident      = "ident"
	Number     = "number"
)

func main() {
	p := rd.NewParser([]string{"const", "id", "=", "9", ";", "id", ":=", "9", "."})

	p.Register(Program, func() bool {
		return p.Match(Block) && p.Match(".")
	})

	p.Register(Block, func() bool {
		if p.Match("const") {
			if !(p.Match(Ident) && p.Match("=") && p.Match(Number)) {
				return false
			}
			for p.Match(",") {
				if !(p.Match(Ident) && p.Match("=") && p.Match(Number)) {
					return false
				}
			}
			if !p.Match(";") {
				return false
			}
		}
		if p.Match("var") {
			if !p.Match(Ident) {
				return false
			}
			for p.Match(",") {
				if !p.Match(Ident) {
					return false
				}
			}
			if !p.Match(";") {
				return false
			}
		}
		for p.Match("procedure") {
			if !(p.Match(Ident) && p.Match(";") && p.Match(Block) && p.Match(";")) {
				return false
			}
		}
		return p.Match(Statement)
	})

	p.Register(Statement, func() bool {
		switch {
		case p.Match(Ident):
			if !(p.Match(":=") && p.Match(Expression)) {
				return false
			}
		case p.Match("call"):
			if !p.Match(Ident) {
				return false
			}
		case p.Match("begin"):
			if !p.Match(Statement) {
				return false
			}
			for p.Match(";") {
				if !p.Match(Statement) {
					return false
				}
			}
			if !p.Match("end") {
				return false
			}
		case p.Match("if"):
			if !(p.Match(Condition) && p.Match("then") && p.Match(Statement)) {
				return false
			}
		case p.Match("while"):
			if !(p.Match(Condition) && p.Match("do") && p.Match(Statement)) {
				return false
			}
		default:
			return false
		}
		return true
	})

	p.Register(Condition, func() bool {
		switch {
		case p.Match("odd"):
			return p.Match(Expression)
		case p.Match(Expression):
			switch {
			case p.Match("="):
			case p.Match("#"):
			case p.Match("<"):
			case p.Match("<="):
			case p.Match(">"):
			case p.Match(">="):
			default:
				return false
			}
			return p.Match(Expression)
		default:
			return false
		}
	})

	p.Register(Expression, func() bool {
		switch {
		case p.Match("+"):
		case p.Match("-"):
		}
		if !p.Match(Term) {
			return false
		}
		for p.Match("+") || p.Match("-") {
			if !p.Match(Term) {
				return false
			}
		}
		return true
	})

	p.Register(Term, func() bool {
		if !p.Match(Factor) {
			return false
		}
		for p.Match("*") || p.Match("/") {
			if !p.Match(Factor) {
				return false
			}
		}
		return true
	})

	p.Register(Factor, func() bool {
		switch {
		case p.Match(Ident):
		case p.Match(Number):
		case p.Match("(") && p.Match(Expression) && p.Match(")"):
		default:
			return false
		}
		return true
	})

	p.Register(Ident, func() bool {
		return p.Match("id")
	})

	p.Register(Number, func() bool {
		return p.Match("9")
	})

	fmt.Println(p.Match(Program))
	d, _ := json.MarshalIndent(p.Tree(), "", "  ")
	fmt.Println(string(d))
}
