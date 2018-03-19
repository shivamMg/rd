package main

import "fmt"

func main() {
	/*
		terms := []string{"a", "b", "c", "d", "e"}
		nonTerms := []string{"S", "A", "B", "C", "D", "E"}
		start := "S"
		rules := map[string][]string{
			"S": []string{"A", "B", "C", "D", "E"},
			"A": []string{"a", "|", "ε"},
			"B": []string{"b", "|", "ε"},
			"C": []string{"c"},
			"D": []string{"d", "|", "ε"},
			"E": []string{"e", "|", "ε"},
		}
	*/
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
