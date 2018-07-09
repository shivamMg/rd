/*
Recursive Descent parser for PL/0 programming language.
The grammar has been copied from https://en.wikipedia.org/wiki/PL/0#Grammar


program = block "." .

block =
    ["const" ident "=" number {"," ident "=" number} ";"]
    ["var" ident {"," ident} ";"]
    {"procedure" ident ";" block ";"} statement .

statement =
	ident ":=" expression
	| "!" expression
	| "?" ident
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
	"github.com/shivamMg/rd"
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

// Terminals
const (
	Period     = "."
	Const      = "const"
	Comma      = ","
	Semicolon  = ";"
	Var        = "var"
	Procedure  = "procedure"
	Assignment = ":="
	Exclam     = "!"
	Ques       = "?"
	Call       = "call"
	Begin      = "begin"
	End        = "end"
	If         = "if"
	Then       = "then"
	While      = "while"
	Do         = "do"
	Odd        = "odd"
	Equal      = "="
	Hash       = "#"
	LT         = "<"
	LTE        = "<="
	GT         = ">"
	GTE        = ">="
	Plus       = "+"
	Minus      = "-"
	Mul        = "*"
	Div        = "/"
	OpenParen  = "("
	CloseParen = ")"
)

// Symbols stores all non-terminals + all terminals.
var Symbols []string

// NonTerminals stores all non-terminals.
var NonTerminals []string

// Terminals stores all terminals.
var Terminals []string

func init() {
	NonTerminals = []string{Program, Block, Statement, Condition,
		Expression, Term, Factor, Ident, Number}
	Terminals = []string{Period, Const, Comma, Semicolon, Var,
		Procedure, Assignment, Call, Begin, End, If, Then, While,
		Do, Odd, Equal, Hash, LT, LTE, GT, GTE, Plus, Minus, Mul,
		Div, OpenParen, CloseParen, Exclam}
	Symbols = append(NonTerminals, Terminals...)
}

func main() {
	p := rd.NewParser(primeProgram())

	p.Rule(Program, func() bool {
		return p.Match(Block) && p.Match(Period)
	})

	p.Rule(Block, func() bool {
		if p.Match(Const) {
			if !(p.Match(Ident) && p.Match(Equal) && p.Match(Number)) {
				return false
			}
			for p.Match(Comma) {
				if !(p.Match(Ident) && p.Match(Equal) && p.Match(Number)) {
					return false
				}
			}
			if !p.Match(Semicolon) {
				return false
			}
		}
		if p.Match(Var) {
			if !p.Match(Ident) {
				return false
			}
			for p.Match(Comma) {
				if !p.Match(Ident) {
					return false
				}
			}
			if !p.Match(Semicolon) {
				return false
			}
		}
		for p.Match(Procedure) {
			if !(p.Match(Ident) && p.Match(Semicolon) && p.Match(Block) && p.Match(Semicolon)) {
				return false
			}
		}
		return p.Match(Statement)
	})

	p.Rule(Statement, func() bool {
		switch {
		case p.Match(Ident):
			if !(p.Match(Assignment) && p.Match(Expression)) {
				return false
			}
		case p.Match(Exclam):
			if !p.Match(Expression) {
				return false
			}
		case p.Match(Ques):
			if !p.Match(Ident) {
				return false
			}
		case p.Match(Call):
			if !p.Match(Ident) {
				return false
			}
		case p.Match(Begin):
			if !p.Match(Statement) {
				return false
			}
			for p.Match(Semicolon) {
				if !p.Match(Statement) {
					return false
				}
			}
			if !p.Match(End) {
				return false
			}
		case p.Match(If):
			if !(p.Match(Condition) && p.Match(Then) && p.Match(Statement)) {
				return false
			}
		case p.Match(While):
			if !(p.Match(Condition) && p.Match(Do) && p.Match(Statement)) {
				return false
			}
		default:
			return false
		}
		return true
	})

	p.Rule(Condition, func() bool {
		switch {
		case p.Match(Odd):
			return p.Match(Expression)
		case p.Match(Expression):
			if p.Match(Equal) || p.Match(Hash) || p.Match(LT) || p.Match(LTE) || p.Match(GT) || p.Match(GTE) {
				return p.Match(Expression)
			}
			return false
		default:
			return false
		}
	})

	p.Rule(Expression, func() bool {
		if p.Match(Plus) || p.Match(Minus) {
		}
		if !p.Match(Term) {
			return false
		}
		for p.Match(Plus) || p.Match(Minus) {
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
		for p.Match(Mul) || p.Match(Div) {
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
		case p.Match(OpenParen) && p.Match(Expression) && p.Match(CloseParen):
		default:
			return false
		}
		return true
	})

	p.Rule(Ident, func() bool {
		// for this purpose of this demonstration a non-reserved alphabetical symbol
		// is a valid identifier
		next, _ := p.NextToken()
		for _, sym := range Symbols {
			if next == sym {
				p.Retract()
				return false
			}
		}
		if ok, _ := regexp.MatchString(`[[:alpha:]]`, next); !ok {
			p.Retract()
			return false
		}
		p.Add(next)
		return true
	})

	p.Rule(Number, func() bool {
		next, _ := p.NextToken()
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

func print(t *rd.Tree) {
	goTree := createGoTree(t)
	fmt.Println(goTree.Print())
}

func createGoTree(root *rd.Tree) gotree.Tree {
	t := gotree.New(root.Symbol)
	for _, c := range root.Children {
		ct := createGoTree(c)
		t.AddTree(ct)
	}
	return t
}
