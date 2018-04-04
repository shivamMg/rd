package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	terms := []string{"a", "+", "(", ")"}
	nonTerms := []string{"S", "F"}
	start := "S"
	rules := map[string][]string{
		"S": []string{"F", "|", "(", "S", "+", "F", ")"},
		"F": []string{"a"},
	}

	p := NewParser(terms, nonTerms, start, rules)
	tree, _ := p.Parse([]string{"(", "a", "+", "a", ")"})
	data, _ := json.MarshalIndent(tree, "", "  ")
	fmt.Println(string(data))
}
