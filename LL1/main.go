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
	terms := []string{"+", "*", "id", "(", ")"}
	nonTerms := []string{"E", "E'", "T", "T'", "F"}
	start := "E"
	rules := map[string][]string{
		"E":  []string{"T", "E'"},
		"E'": []string{"+", "T", "E'", "|", "ε"},
		"T":  []string{"F", "T'"},
		"T'": []string{"*", "F", "T'", "|", "ε"},
		"F":  []string{"id", "|", "(", "E", ")"},
	}

	p := NewParser(terms, nonTerms, start, rules)
	fmt.Println(p.ParseTable())
}
