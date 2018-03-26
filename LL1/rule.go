package main

type Segment []string

type Rule struct {
	LHS string
	RHS []Segment
}

func (r *Rule) Add(s Segment) {
	r.RHS = append(r.RHS, s)
}

type Rules struct {
	// Use map for easier lookup.
	rules map[string]Rule
}

func NewRules() *Rules {
	return &Rules{rules: make(map[string]Rule)}
}

func (rs *Rules) Add(r Rule) {
	rs.rules[r.LHS] = r
}

func (rs Rules) Get(nonTerm string) (Rule, bool) {
	r, ok := rs.rules[nonTerm]
	return r, ok
}

func (rs Rules) List() []Rule {
	rules := []Rule{}
	for _, r := range rs.rules {
		rules = append(rules, r)
	}
	return rules
}
