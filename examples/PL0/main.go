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
	"github.com/shivamMg/ppds/tree"
	"github.com/shivamMg/rd"
	"github.com/shivamMg/rd/examples/PL0/lexer"
	"github.com/shivamMg/rd/examples/PL0/parser"
)

type node struct {
	data string
	c    []*node
}

func (n *node) Data() interface{} {
	return n.data
}

func (n *node) Children() (c []tree.Node) {
	for _, child := range n.c {
		c = append(c, tree.Node(child))
	}
	return
}

func convert(t *rd.Tree) *node {
	n := new(node)
	n.data = t.Symbol
	for _, c := range t.Subtrees {
		n.c = append(n.c, convert(c))
	}
	return n
}

func main() {
	tokens := lexer.Lex(squareProgram)
	fmt.Println(tokens)
	ok, parseTree := parser.Parse(tokens, parser.Program)
	fmt.Println("Match:", ok)
	fmt.Println(sprint(parseTree))
}

func sprint(t *rd.Tree) string {
	root := convert(t)
	return tree.SprintHrn(root)
}
