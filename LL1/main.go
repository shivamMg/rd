package main

import "fmt"

/*
var Input = map[string]Rule{
	"E": Rule{"E", "TG"},
	"G": Rule{"G", "+TG|ε"},
	"T": Rule{"T", "FR"},
	"R": Rule{"R", "*FR|ε"},
	"F": Rule{"F", "(E)|i"},
}
*/

func main() {
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

	p := NewParser(terms, nonTerms, start, rules)
	fmt.Println(p.FollowAll())
	fmt.Println(p.FirstAll())
}
