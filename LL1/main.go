package main

import "fmt"

func main() {
	terms := []string{"a", "+", "(", ")"}
	nonTerms := []string{"S", "F"}
	start := "S"
	rules := map[string][]string{
		"S": []string{"F", "|", "(", "S", "+", "F", ")"},
		"F": []string{"a"},
	}

	p := NewParser(terms, nonTerms, start, rules)
	fmt.Println(p.Parse([]string{"(", "a", "+", "a", ")"}))
}
