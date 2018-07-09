package rd

type ele struct {
	index   int
	nonTerm *Tree
}

// stack stores trees as non-terminals are expanded. index stores
// the tokens' index for which the nonTerm was expanded. The last
// ele is the non-terminal that is currently being derived.
type stack []ele

func (st stack) peek() (ele, bool) {
	l := len(st)
	if l == 0 {
		return ele{}, false
	}
	return st[l-1], true
}

func (st *stack) pop() (ele, bool) {
	l := len(*st)
	if l == 0 {
		return ele{}, false
	}
	e := (*st)[l-1]
	*st = (*st)[:l-1]
	return e, true
}

func (st *stack) push(e ele) {
	*st = append(*st, e)
}

// Parser represents a Recursive Descent parser.
type Parser struct {
	// Stores tokens tokens.
	tokens  []string
	st      stack
	current int
	rules   map[string]func() bool
}

// NewParser returns a new Parser to parse tokens. Production functions
// for non-terminals must be added using Rule method.
func NewParser(tokens []string) *Parser {
	return &Parser{
		tokens:  tokens,
		st:      stack{},
		current: -1,
		rules:   make(map[string]func() bool),
	}
}

// Rule saves a production function for a non-terminal.
func (p *Parser) Rule(nonTerm string, f func() bool) {
	p.rules[nonTerm] = f
}

// Reset resets parser p to parse newer tokens.
func (p *Parser) Reset(tokens []string) {
	p.tokens = tokens
	p.st = stack{}
	p.current = -1
}

// CurrentIndex returns the current token's index.
func (p Parser) CurrentIndex() int {
	return p.current
}

// Current returns the token where we are currently at.
func (p *Parser) Current() string {
	return p.tokens[p.current]
}

// NextToken returns next token from after incrementing the
// current index. bool signifies if tokens are finished.
func (p *Parser) NextToken() (string, bool) {
	if p.current >= len(p.tokens)-1 {
		return "", false
	}
	p.current++
	return p.tokens[p.current], true
}

// Retract decrements the current index to bring current
// to the previous token.
func (p *Parser) Retract() {
	p.current--
}

// Add adds terminal token term to the non-terminal that is
// being expanded.
func (p *Parser) Add(term string) {
	e, _ := p.st.peek()
	e.nonTerm.Add(NewTree(term))
}

// Match first makes out if symbol is a terminal or a non-terminal - by checking
// if a production rule exists against it. If it's a terminal then it's matched by
// the next token. If it's a non-terminal then the production function for it is
// called.
// Match also handles backtracking. In case of a terminal non-match, Retract is called.
// In case of a failed non-terminal production, current is put to where it was before
// the production.
func (p *Parser) Match(symbol string) bool {
	f, ok := p.rules[symbol]
	if !ok {
		// it's a terminal
		next, ok := p.NextToken()
		if !ok {
			p.Retract()
			return false
		}
		if symbol != next {
			p.Retract()
			return false
		}
		p.Add(symbol)
		return true
	}

	// it's a non-terminal
	t := NewTree(symbol)
	// if it's not the first production attach t
	// to the last non-terminal that was expanded
	if e, ok := p.st.peek(); ok {
		e.nonTerm.Add(t)
	}

	p.st.push(ele{index: p.current, nonTerm: t})
	isMatch := f()
	// don't pop the first production - the root
	if len(p.st) > 1 {
		p.st.pop()
	}

	if !isMatch {
		// detach t from the last non-terminal that
		// was expanded
		e, _ := p.st.peek()
		e.nonTerm.Detach(t)
		p.current = e.index
	}
	return isMatch
}

// Tree retrieves the parse tree for the last production.
func (p Parser) Tree() *Tree {
	if e, ok := p.st.peek(); ok {
		return e.nonTerm
	}
	return nil
}
