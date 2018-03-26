package main

const (
	Epsilon = "ε"
	// Or separator
	Or = "|"
)

type Parser struct {
	terms    *Set
	nonTerms *Set
	start    string
	// Non-terminal -> [][]Symbols
	rules *Rules
}

func NewParser(terms, nonTerms []string, start string, rules map[string][]string) *Parser {
	// TODO: Validate before returning p
	p := Parser{}
	p.terms = NewSet(terms)
	p.nonTerms = NewSet(nonTerms)
	p.start = start
	p.rules = NewRules()

	// split rules on Or separator
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

func (p Parser) IsTerm(symbol string) bool {
	return p.terms.Has(symbol)
}

func (p Parser) IsNonTerm(symbol string) bool {
	return p.nonTerms.Has(symbol)
}

// Parse returns rules for left-most derivation of input
func (p Parser) Parse(input []string) [][]string {
	input = append(input, "$")
	derivation := [][]string{}
	stack := []string{"$", p.start}
	table := p.ParseTable()
	pos := 0
abort:
	for len(stack) > 0 {
		s := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		sym := input[pos]
		switch {
		case p.IsTerm(s) || s == "$":
			if s == sym {
				pos += 1
				if sym == "$" {
					// input accepted
				}
			} else {
				panic("Bad symbol: " + sym)
				break abort
			}
		case p.IsNonTerm(s):
			rule := table[s][sym]
			for i := len(rule) - 1; i >= 0; i-- {
				stack = append(stack, rule[i])
			}

			d := []string{s}
			d = append(d, rule...)
			derivation = append(derivation, d)
		}
	}

	return derivation
}

func (p Parser) ParseTable() map[string]map[string][]string {
	// map[non-term][term] = segment (rule)
	table := map[string]map[string][]string{}

	for _, r := range p.rules.List() {
		for _, seg := range r.RHS {
			set := *NewSet([]string{})
			if seg[0] == Epsilon {
				set = p.follow(r.LHS)
			} else {
				set = p.firstSeg(Segment(seg))
			}

			for _, sym := range set.List() {
				if m, ok := table[r.LHS]; ok {
					m[sym] = seg
					table[r.LHS] = m
				} else {
					table[r.LHS] = make(map[string][]string)
					table[r.LHS][sym] = seg
				}
			}
		}
	}

	return table
}

func (p Parser) FirstAll() map[string]Set {
	f := map[string]Set{}
	for _, r := range p.rules.List() {
		f[r.LHS] = p.first(r.LHS)
	}

	return f
}

func (p Parser) First(nt string) Set {
	return p.first(nt)
}

func (p Parser) firstSeg(seg Segment) Set {
	set := NewSet([]string{})
	for i, sym := range seg {
		f := p.first(sym)
		if !f.Has(Epsilon) {
			set.Merge(f)
			break
		}

		if i != len(seg)-1 {
			f.Delete(Epsilon)
		}
		set.Merge(f)
	}
	return *set
}

// segs is a slice of segments. Segmentation is done using Or as the
// separator.
func (p Parser) first(symbol string) Set {
	set := NewSet([]string{})
	if symbol == Epsilon || p.IsTerm(symbol) {
		set.Add(symbol)
		return *set
	}

	// it's non terminal
	r, ok := p.rules.Get(symbol)
	if !ok {
		return *set
	}

	for _, seg := range r.RHS {
	endSeg:
		for i, sym := range seg {
			f := p.first(sym)
			if !f.Has(Epsilon) {
				set.Merge(f)
				break endSeg
			}

			if i < len(seg)-1 {
				f.Delete(Epsilon)
			}
			set.Merge(f)
		}
	}

	return *set
}

func (p Parser) FollowAll() map[string]Set {
	f := map[string]Set{}
	for _, r := range p.rules.List() {
		f[r.LHS] = p.follow(r.LHS)
	}

	return f
}

func (p Parser) Follow(nt string) Set {
	return p.follow(nt)
}

func (p Parser) follow(nt string) Set {
	set := NewSet([]string{})
	if nt == p.start {
		set.Add("$")
	}

	for _, r := range p.rules.List() {
		// search for nt in rhs
		for _, seg := range r.RHS {
			for i, sym := range seg {
				if sym == nt {
					if i == len(seg)-1 {
						if sym != r.LHS {
							set.Merge(p.follow(r.LHS))
						}
						continue
					}

					f := p.firstSeg(Segment(seg[i+1:]))
					if f.Has(Epsilon) {
						f.Delete(Epsilon)
						set.Merge(f)
						if sym != r.LHS {
							set.Merge(p.follow(r.LHS))
						}
					} else {
						set.Merge(f)
					}
					break
				}
			}
		}
	}

	return *set
}
