package main

const (
	// Epsilon represents an empty string
	Epsilon = "ε"
	// Or separator
	Or = "|"
)

type Parser struct {
	terms    *Set
	nonTerms *Set
	start    string
	// Non-terminal -> [][]Symbols
	rules map[string][][]string
}

func NewParser(terms, nonTerms []string, start string, rules map[string][]string) *Parser {
	// TODO: Validate before returning p
	p := Parser{rules: make(map[string][][]string)}
	p.terms = NewSet(terms)
	p.nonTerms = NewSet(nonTerms)
	p.start = start

	// split rules on Or separator
	for nt, rhs := range rules {
		p.rules[nt] = [][]string{}

		prev := 0
		for i, sym := range rhs {
			if sym == Or {
				p.rules[nt] = append(p.rules[nt], rhs[prev:i])
				prev = i + 1
			}
		}
		p.rules[nt] = append(p.rules[nt], rhs[prev:])
	}

	return &p
}

func (p Parser) IsTerm(symbol string) bool {
	return p.terms.Has(symbol)
}

func (p Parser) IsNonTerm(symbol string) bool {
	return p.nonTerms.Has(symbol)
}

func (p Parser) FirstAll() map[string]Set {
	f := map[string]Set{}
	for nt, segs := range p.rules {
		f[nt] = p.first(segs)
	}

	return f
}

func (p Parser) First(nt string) Set {
	return p.first(p.rules[nt])
}

// segs is a slice of segments. Segmentation is done using Or as the
// separator.
func (p Parser) first(segs [][]string) Set {
	set := NewSet([]string{})

	for _, seg := range segs {
	endSeg:
		for i, sym := range seg {
			switch {
			case sym == Epsilon || p.IsTerm(sym):
				set.Add(sym)
				break endSeg
			case p.IsNonTerm(sym):
				f := p.first(p.rules[sym])
				if !f.Has(Epsilon) {
					// we need not continue with the same segment
					set.Merge(f)
					break endSeg
				}

				if i != len(seg)-1 {
					f.Delete(Epsilon)
				}
				set.Merge(f)
			}
		}
	}

	return *set
}

func (p Parser) FollowAll() map[string]Set {
	f := map[string]Set{}
	for nt := range p.rules {
		f[nt] = p.follow(nt)
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

	for lhs, segs := range p.rules {
		// search for nt in rhs
		for _, seg := range segs {
			for i, sym := range seg {
				if sym == nt {
					if i == len(seg)-1 {
						if sym != lhs {
							set.Merge(p.follow(lhs))
						}
						continue
					}

					fw := [][]string{}
					fw = append(fw, seg[i+1:])
					f := p.first(fw)
					if f.Has(Epsilon) {
						f.Delete(Epsilon)
						set.Merge(f)
						set.Merge(p.follow(lhs))
					} else {
						set.Merge(f)
					}
				}
			}
		}
	}

	return *set
}
