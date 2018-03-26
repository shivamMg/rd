package main

// Segment represents a part of input rule's RHS; obtained after
// splitting RHS using Or.
type Segment []string

// Index returns index of symbol in s, else it returns -1.
func (s Segment) Index(symbol string) int {
	for i, sym := range s {
		if sym == symbol {
			return i
		}
	}
	return -1
}

// Rule defines a production rule. LHS is a non-terminal that derives
// RHS, a set of symbols that has been split into segments.
type Rule struct {
	LHS string
	RHS []Segment
}

// Add adds a segment to r's RHS.
func (r *Rule) Add(s Segment) {
	r.RHS = append(r.RHS, s)
}

// Rules helps manage multiple production rules.
// TODO: Move it into parser.go. It'll need validation for
// Add, Get and other methods.
type Rules struct {
	// use map for easier lookup
	rules map[string]Rule
}

// NewRules returns a new Rules.
func NewRules() *Rules {
	return &Rules{rules: make(map[string]Rule)}
}

// Add adds a rule to existing rules.
func (rs *Rules) Add(r Rule) {
	rs.rules[r.LHS] = r
}

// Get returns Rule associated with a non-terminal (LHS).
func (rs Rules) Get(nonTerm string) (Rule, bool) {
	r, ok := rs.rules[nonTerm]
	return r, ok
}

// List returns a slice of all the rules.
func (rs Rules) List() []Rule {
	rules := []Rule{}
	for _, r := range rs.rules {
		rules = append(rules, r)
	}
	return rules
}
