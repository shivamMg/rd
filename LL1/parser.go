package main

import (
	"errors"
)

const (
	// Epsilon symbol represents an empty string.
	Epsilon = "ε"
	// Or symbol used for spliting input rules into segments.
	Or = "|"
	// Dollar represents the end of an input.
	Dollar = "$"
)

// Parser contains terminals, non-terminals, start symbol and production
// rules for LL(1) parsing. Besides the Parse method, it provides methods
// for intermediate stages of parsing.
type Parser struct {
	terms    *Set
	nonTerms *Set
	start    string
	rules    *Rules
}

// NewParser returns a Parser after validating inputs.
func NewParser(terms, nonTerms []string, start string, rules map[string][]string) *Parser {
	// TODO: Add validation before returning p
	p := Parser{}
	p.terms = NewSet(terms)
	p.nonTerms = NewSet(nonTerms)
	p.start = start
	p.rules = NewRules()

	// split rhs on Or, and store these as segments
	for lhs, rhs := range rules {
		r := Rule{LHS: lhs}
		prev := 0
		for i, sym := range rhs {
			if sym == Or {
				r.Add(Segment(rhs[prev:i]))
				prev = i + 1
			}
		}
		r.Add(Segment(rhs[prev:]))
		p.rules.Add(r)
	}
	return &p
}

// IsTerm returns true if symbol is a terminal, else false.
func (p Parser) IsTerm(symbol string) bool { return p.terms.Has(symbol) }

// IsNonTerm returns true if symbol is a non-terminal, else false.
func (p Parser) IsNonTerm(symbol string) bool { return p.nonTerms.Has(symbol) }

// IsSymbol returns true if symbol is an epsilon, terminal, or non-terminal.
func (p Parser) IsSymbol(symbol string) bool {
	return symbol == Epsilon || p.terms.Has(symbol) || p.nonTerms.Has(symbol)
}

// Parse returns parse tree for input.
func (p Parser) Parse(input []string) (*Tree, error) {
	root := NewNode(p.start)
	st := new(stack)
	st.Push(NewNode(Dollar), nil)
	st.Push(root, nil)

	input = append(input, Dollar)
	table := p.ParseTable()
	// index of current string in input
	pos := 0
	for !st.isEmpty() {
		s, sym := st.Pop(), input[pos]
		switch {
		case s.Symbol == Dollar || p.IsTerm(s.Symbol):
			if s.Symbol != sym {
				return nil, errors.New("Bad symbol: " + sym)
			}
			if s.Symbol == Dollar {
				break
			}
			s.parent.Add(s.Node)
			pos++
		case p.IsNonTerm(s.Symbol):
			seg := table[s.Symbol][sym]
			for i := len(seg) - 1; i >= 0; i-- {
				st.Push(NewNode(seg[i]), s.Node)
			}
			if root != s.Node {
				s.parent.Add(s.Node)
			}
		}
	}
	return &Tree{Root: root}, nil
}

// ParseTable returns parsing table for the rules.
func (p Parser) ParseTable() map[string]map[string][]string {
	table := map[string]map[string][]string{}
	// populate table with non-terminals as rows and terminals as columns
	for _, nt := range p.nonTerms.List() {
		table[nt] = make(map[string][]string)
		table[nt][Dollar] = nil
		for _, t := range p.terms.List() {
			table[nt][t] = nil
		}
	}

	for _, r := range p.rules.List() {
		for _, seg := range r.RHS {
			set := *NewSet(nil)
			if seg[0] == Epsilon {
				set = p.follow(r.LHS)
			} else {
				set = p.firstSeg(Segment(seg))
			}

			for _, t := range set.List() {
				table[r.LHS][t] = []string(seg)
			}
		}
	}

	return table
}

// FirstAll returns FIRST set for every non-terminal.
func (p Parser) FirstAll() map[string]Set {
	f := map[string]Set{}
	for _, r := range p.rules.List() {
		f[r.LHS] = p.first(r.LHS)
	}
	return f
}

// First returns FIRST set for a symbol.
func (p Parser) First(symbol string) Set {
	if p.IsSymbol(symbol) {
		return p.first(symbol)
	}
	return *NewSet(nil)
}

func (p Parser) first(symbol string) Set {
	set := NewSet(nil)
	if symbol == Epsilon || p.IsTerm(symbol) {
		set.Add(symbol)
		return *set
	}

	// it's a non-terminal
	r, _ := p.rules.Get(symbol)
	for _, seg := range r.RHS {
		set.Merge(p.firstSeg(seg))
	}
	return *set
}

// firstSeg returns FIRST set by iterating over seg.
func (p Parser) firstSeg(seg Segment) Set {
	set := NewSet(nil)
	for i, sym := range seg {
		f := p.first(sym)
		if !f.Has(Epsilon) {
			set.Merge(f)
			break
		}

		// don't include epsilon except for the last symbol
		if i != len(seg)-1 {
			f.Delete(Epsilon)
		}
		set.Merge(f)
	}
	return *set
}

// FollowAll returns FOLLOW set for every non-terminal.
func (p Parser) FollowAll() map[string]Set {
	f := map[string]Set{}
	for _, r := range p.rules.List() {
		f[r.LHS] = p.follow(r.LHS)
	}
	return f
}

// Follow returns FOLLOW set for a non-terminal.
func (p Parser) Follow(nonTerm string) Set {
	if p.IsNonTerm(nonTerm) {
		return p.follow(nonTerm)
	}
	return *NewSet(nil)
}

func (p Parser) follow(nonTerm string) Set {
	set := NewSet(nil)
	if nonTerm == p.start {
		set.Add(Dollar)
	}

	for _, r := range p.rules.List() {
		// search for nonTerm in RHS
		for _, seg := range r.RHS {
			i := seg.Index(nonTerm)
			if i == -1 {
				continue
			}

			// if nonTerm is the last symbol then follow LHS
			// of the rule unless it's the same as nonTerm
			if i == len(seg)-1 {
				if nonTerm != r.LHS {
					set.Merge(p.follow(r.LHS))
				}
				continue
			}

			f := p.firstSeg(Segment(seg[i+1:]))
			if f.Has(Epsilon) {
				f.Delete(Epsilon)
				// follow LHS of the rule unless it's the same as nonTerm
				if nonTerm != r.LHS {
					set.Merge(p.follow(r.LHS))
				}
			}
			set.Merge(f)
		}
	}

	return *set
}
