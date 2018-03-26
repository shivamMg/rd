package main

import (
	"reflect"
	"testing"
)

type input struct {
	terms, nonTerms []string
	start           string
	rules           map[string][]string
}

var testInputs []input

func init() {
	testInputs = append(testInputs,
		input{
			terms:    []string{"a", "b", "c", "d", "e"},
			nonTerms: []string{"S", "A", "B", "C", "D", "E"},
			start:    "S",
			rules: map[string][]string{
				"S": []string{"A", "B", "C", "D", "E"},
				"A": []string{"a", "|", "ε"},
				"B": []string{"b", "|", "ε"},
				"C": []string{"c"},
				"D": []string{"d", "|", "ε"},
				"E": []string{"e", "|", "ε"},
			},
		},
		input{
			terms:    []string{"a", "b", "c", "d"},
			nonTerms: []string{"S", "B", "C"},
			start:    "S",
			rules: map[string][]string{
				"S": []string{"B", "b", "|", "C", "d"},
				"B": []string{"a", "B", "|", "ε"},
				"C": []string{"c", "C", "|", "ε"},
			},
		},
		input{
			terms:    []string{"+", "*", "id", "(", ")"},
			nonTerms: []string{"E", "E'", "T", "T'", "F"},
			start:    "E",
			rules: map[string][]string{
				"E":  []string{"T", "E'"},
				"E'": []string{"+", "T", "E'", "|", "ε"},
				"T":  []string{"F", "T'"},
				"T'": []string{"*", "F", "T'", "|", "ε"},
				"F":  []string{"id", "|", "(", "E", ")"},
			},
		},
		input{
			terms:    []string{"a", "b", "d", "g", "h"},
			nonTerms: []string{"S", "A", "B", "C"},
			start:    "S",
			rules: map[string][]string{
				"S": []string{"A", "C", "B", "|", "C", "b", "B", "|", "B", "a"},
				"A": []string{"d", "a", "|", "B", "C"},
				"B": []string{"g", "|", "ε"},
				"C": []string{"h", "|", "ε"},
			},
		},
		input{
			terms:    []string{"a", "+", "(", ")"},
			nonTerms: []string{"S", "F"},
			start:    "S",
			rules: map[string][]string{
				"S": []string{"F", "|", "(", "S", "+", "F", ")"},
				"F": []string{"a"},
			},
		},
	)
}

func TestFirstAll(t *testing.T) {
	ins := make([]input, len(testInputs))
	copy(ins, testInputs)

	tests := []struct {
		in   input
		want map[string]Set
	}{
		{
			in: ins[0],
			want: map[string]Set{
				"S": *NewSet([]string{"a", "b", "c"}),
				"A": *NewSet([]string{"a", "ε"}),
				"B": *NewSet([]string{"b", "ε"}),
				"C": *NewSet([]string{"c"}),
				"D": *NewSet([]string{"d", "ε"}),
				"E": *NewSet([]string{"e", "ε"}),
			},
		},
		{
			in: ins[1],
			want: map[string]Set{
				"S": *NewSet([]string{"a", "b", "c", "d"}),
				"B": *NewSet([]string{"a", "ε"}),
				"C": *NewSet([]string{"c", "ε"}),
			},
		},
		{
			in: ins[2],
			want: map[string]Set{
				"E":  *NewSet([]string{"id", "("}),
				"E'": *NewSet([]string{"+", "ε"}),
				"T":  *NewSet([]string{"id", "("}),
				"T'": *NewSet([]string{"*", "ε"}),
				"F":  *NewSet([]string{"id", "("}),
			},
		},
		{
			in: ins[3],
			want: map[string]Set{
				"S": *NewSet([]string{"d", "g", "h", "b", "a", "ε"}),
				"A": *NewSet([]string{"d", "g", "h", "ε"}),
				"B": *NewSet([]string{"g", "ε"}),
				"C": *NewSet([]string{"h", "ε"}),
			},
		},
	}

	for _, test := range tests {
		p := NewParser(test.in.terms, test.in.nonTerms, test.in.start, test.in.rules)
		got := p.FirstAll()
		if !reflect.DeepEqual(test.want, got) {
			t.Errorf("Expected: %v\nGot: %v\n", test.want, got)
		}
	}
}

func TestFollowAll(t *testing.T) {
	ins := make([]input, len(testInputs))
	copy(ins, testInputs)

	tests := []struct {
		in   input
		want map[string]Set
	}{
		{
			in: ins[0],
			want: map[string]Set{
				"S": *NewSet([]string{"$"}),
				"A": *NewSet([]string{"b", "c"}),
				"B": *NewSet([]string{"c"}),
				"C": *NewSet([]string{"d", "e", "$"}),
				"D": *NewSet([]string{"e", "$"}),
				"E": *NewSet([]string{"$"}),
			},
		},
		{
			in: ins[1],
			want: map[string]Set{
				"S": *NewSet([]string{"$"}),
				"B": *NewSet([]string{"b"}),
				"C": *NewSet([]string{"d"}),
			},
		},
		{
			in: ins[2],
			want: map[string]Set{
				"E":  *NewSet([]string{"$", ")"}),
				"E'": *NewSet([]string{"$", ")"}),
				"T":  *NewSet([]string{"+", "$", ")"}),
				"T'": *NewSet([]string{"+", "$", ")"}),
				"F":  *NewSet([]string{"*", "+", "$", ")"}),
			},
		},
		{
			in: ins[3],
			want: map[string]Set{
				"S": *NewSet([]string{"$"}),
				"A": *NewSet([]string{"h", "g", "$"}),
				"B": *NewSet([]string{"$", "a", "h", "g"}),
				"C": *NewSet([]string{"g", "$", "b", "h"}),
			},
		},
	}

	for _, test := range tests {
		p := NewParser(test.in.terms, test.in.nonTerms, test.in.start, test.in.rules)
		got := p.FollowAll()
		if !reflect.DeepEqual(test.want, got) {
			t.Errorf("Expected: %v\nGot: %v\n", test.want, got)
		}
	}
}

func TestParseTable(t *testing.T) {
	in := testInputs[2]
	want := map[string]map[string][]string{
		"T'": map[string][]string{
			"*": []string{"*", "F", "T'"},
			"+": []string{"ε"},
			"$": []string{"ε"},
			")": []string{"ε"},
		},
		"F": map[string][]string{
			"id": []string{"id"},
			"(":  []string{"(", "E", ")"},
		},
		"E": map[string][]string{
			"id": []string{"T", "E'"},
			"(":  []string{"T", "E'"},
		},
		"E'": map[string][]string{
			"+": []string{"+", "T", "E'"},
			"$": []string{"ε"},
			")": []string{"ε"},
		},
		"T": map[string][]string{
			"id": []string{"F", "T'"},
			"(":  []string{"F", "T'"},
		},
	}
	p := NewParser(in.terms, in.nonTerms, in.start, in.rules)
	got := p.ParseTable()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected: %v\nGot: %v\n", want, got)
	}
}

func TestParse(t *testing.T) {
	in := testInputs[4]
	want := [][]string{
		[]string{"S", "(", "S", "+", "F", ")"},
		[]string{"S", "F"},
		[]string{"F", "a"},
		[]string{"F", "a"},
	}
	p := NewParser(in.terms, in.nonTerms, in.start, in.rules)
	got := p.Parse([]string{"(", "a", "+", "a", ")"})
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected: %v\nGot: %v\n", want, got)
	}
}
