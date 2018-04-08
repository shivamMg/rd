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
	"fmt"
	"regexp"

	"github.com/DiSiqueira/GoTree"
	"github.com/shivammg/parsers/rd"
	"github.com/shivammg/parsers/types"
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

var NonTerminals []string

func init() {
	NonTerminals = append(NonTerminals, Program, Block, Statement, Condition,
		Expression, Term, Factor, Ident, Number)
}

func main() {
	p := rd.NewParser(squareProgram())

	p.Rule(Program, func() bool {
		return p.Match(Block) && p.Match(".")
	})

	p.Rule(Block, func() bool {
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

	p.Rule(Statement, func() bool {
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

	p.Rule(Condition, func() bool {
		switch {
		case p.Match("odd"):
			return p.Match(Expression)
		case p.Match(Expression):
			if p.Match("=") || p.Match("#") || p.Match("<") || p.Match("<=") || p.Match(">") || p.Match(">=") {
				return p.Match(Expression)
			}
			return false
		default:
			return false
		}
	})

	p.Rule(Expression, func() bool {
		if p.Match("+") || p.Match("-") {
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

	p.Rule(Term, func() bool {
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

	p.Rule(Factor, func() bool {
		switch {
		case p.Match(Ident):
		case p.Match(Number):
		case p.Match("(") && p.Match(Expression) && p.Match(")"):
		default:
			return false
		}
		return true
	})

	p.Rule(Ident, func() bool {
		next := p.NextToken()
		if ok, _ := regexp.MatchString(`[[:alpha:]]`, next); !ok {
			p.Retract()
			return false
		}
		// must not be a reserved word
		reserved := NonTerminals
		reserved = append(reserved, "const", "var", "procedure", "begin", "end", "call", "if", "then", "while", "do", "odd")
		for _, sym := range reserved {
			if next == sym {
				p.Retract()
				return false
			}
		}
		p.Add(next)
		return true
	})

	p.Rule(Number, func() bool {
		next := p.NextToken()
		if ok, _ := regexp.MatchString(`[[:digit:]]`, next); !ok {
			p.Retract()
			return false
		}
		p.Add(next)
		return true
	})

	fmt.Println("Match:", p.Match(Program))
	print(p.Tree())
}

func print(t *types.Tree) {
	goTree := *createGoTree(t)
	fmt.Println(goTree.Print())
}

func createGoTree(root *types.Tree) *gotree.Tree {
	t := gotree.New(root.Symbol)
	for _, child := range root.Children {
		t.AddTree(*createGoTree(child))
	}
	return &t
}
