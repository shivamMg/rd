package rd

import (
	"fmt"
	"log"
)

// ParsingError is error returned by Builder's Err method in case an error occurs
// during parsing.
type ParsingError struct {
	errString string
}

func (e *ParsingError) Error() string {
	return e.errString
}

func newParsingError(errString string) *ParsingError {
	return &ParsingError{errString: errString}
}

// Builder stores details about tokens, index to current token, etc. and provides
// methods to build recursive descent parsers conveniently. It keeps a track of
// entry/exit from non-terminal functions, and terminal matches done inside them.
// Results from non-terminal function calls help create the parse tree. A debug
// tree is also created to help trace flow across non-terminal functions.
type Builder struct {
	tokens         []Token
	current        int
	stack          stack
	finalEle       ele
	debugStack     debugStack
	finalDebugTree *DebugTree
	finalErr       *ParsingError
	skip           bool
}

// NewBuilder returns a new Builder for the tokens.
func NewBuilder(tokens []Token) *Builder {
	return &Builder{
		tokens:     tokens,
		current:    -1,
		stack:      stack{},
		debugStack: debugStack{},
	}
}

// Peek returns the ith token without updating the current index. i must be
// relative to the current index.
//
// ex. if current index points to tkn3:
//  tokens:           tkn1 tkn2 tkn3 tkn4 tkn5
//  original indexes:  0    1    2    3    4
//  relative indexes: -2   -1    0    1    2
// you can use:
//  Peek(-2) to get tkn1,
//  Peek(-1) to get tkn2,
//  Peek(1) to get tkn4,
//  Peek(2) to get tkn5.
//
// ok is false if i lies outside original index range, else true.
func (b *Builder) Peek(i int) (token Token, ok bool) {
	b.mustEnter("Peek")
	j := b.current + i
	if j < 0 || j >= len(b.tokens) {
		return nil, false
	}
	return b.tokens[j], true
}

// Check is a convenience function over Peek. It calls Peek to check if returned
// token is same as token, and returned ok is true.
func (b *Builder) Check(token Token, i int) bool {
	b.mustEnter("Check")
	peekedToken, ok := b.Peek(i)
	return peekedToken == token && ok
}

// CheckOrNotOK is a convenience function over Peek. It calls Peek to check if
// returned token is same as token, or returned ok is false.
func (b *Builder) CheckOrNotOK(token Token, i int) bool {
	b.mustEnter("CheckOrNotOK")
	peekedToken, ok := b.Peek(i)
	return peekedToken == token || !ok
}

// Next increments the current index to return the next token. ok is false if
// no tokens are left, else true.
func (b *Builder) Next() (token Token, ok bool) {
	b.mustEnter("Next")
	return b.next()
}

func (b *Builder) next() (token Token, ok bool) {
	if b.current == len(b.tokens)-1 {
		return nil, false
	}
	b.current++
	return b.tokens[b.current], true
}

// Backtrack resets the current index for the non-terminal function it's called inside,
// and sets it to the value it was before entering this function. It also discards any
// matches done inside the function.
func (b *Builder) Backtrack() {
	b.mustEnter("Backtrack")
	e := b.stack.peek()
	b.current = e.index
	e.nonTerm.Subtrees = []*Tree{}
}

// Add adds token as a symbol in the parse tree. It's added under the current
// non-terminal subtree.
func (b *Builder) Add(token Token) {
	b.mustEnter("Add")
	e := b.stack.peek()
	e.nonTerm.Add(NewTree(token))
}

// Match matches the next token to token. In case of a non-match the current index
// is decremented to its original value. ok is false in case of non-match or if no
// tokens are left, else true for a match.
//
// Internally Match calls Next to grab the next token. In case of a match it adds
// it by calling Add. Debug info is also added to the debug tree.
func (b *Builder) Match(token Token) (ok bool) {
	b.mustEnter("Match")
	debugMsg := ""
	defer func() {
		dt := b.debugStack.peek()
		dt.add(newDebugTree(debugMsg))
	}()

	next, ok := b.Next()
	if !ok {
		debugMsg = fmt.Sprint("<no tokens left> ≠ ", token)
		return false
	}
	if next != token {
		b.current--
		debugMsg = fmt.Sprint(next, " ≠ ", token)
		return false
	}
	b.Add(token)
	debugMsg = fmt.Sprint(next, " = ", token)
	return true
}

// Skip removes the current non-terminal from the parse tree regardless of the
// exit result. It's helpful in case of null productions - where non-terminals
// don't contribute to the parse tree.
func (b *Builder) Skip() {
	b.skip = true
}

// Enter adds non-terminal to the parse tree making it the current non-terminal.
// Subsequent terminal matches and calls to non-terminal functions add symbols
// under this non-terminal.
//
// Enter should be called right after entering the non-terminal function.
func (b *Builder) Enter(nonTerm Token) *Builder {
	b.stack.push(ele{
		index:   b.current,
		nonTerm: NewTree(nonTerm),
	})
	b.debugStack.push(newDebugTree(fmt.Sprint(nonTerm)))
	return b
}

// Exit registers exit from a non-terminal function. result indicates if it had a
// successful production or not. result must not be nil. In case of a false result
// or a call to Skip, the current index is reset to where it was before entering
// the non-terminal. In case of a true result:
//  1. Parse tree is set (see ParseTree) if the current non-terminal was root.
//  2. Else it's added as a subtree to its parent non-terminal.
//
// The convenient way to call Exit is by using a named boolean return for the
// non-terminal function, and passing it's address to a deferred Exit.
func (b *Builder) Exit(result *bool) {
	b.mustEnter("Exit")
	if result == nil {
		panic("Exit result cannot be nil")
	}
	e := b.stack.pop()
	resetCurrent := false
	switch {
	case b.skip:
		resetCurrent = true
		b.skip = false
	case *result && b.stack.isEmpty():
		if _, ok := b.next(); ok {
			b.finalErr = newParsingError("not all tokens consumed")
		} else {
			b.finalEle = e
		}
	case *result:
		parent := b.stack.peek()
		parent.nonTerm.Add(e.nonTerm)
	case b.stack.isEmpty():
		// TODO: add additional info to the error message
		b.finalErr = newParsingError("parsing error")
		resetCurrent = true
	default:
		resetCurrent = true
	}
	if resetCurrent {
		b.current = e.index
	}

	dt := b.debugStack.pop()
	dt.data += fmt.Sprintf("(%t)", *result)
	if b.debugStack.isEmpty() {
		b.finalDebugTree = dt
	} else {
		parent := b.debugStack.peek()
		parent.add(dt)
	}
}

// ParseTree returns the parse tree. It's set after the root non-terminal exits with
// true result. Returns nil otherwise.
func (b *Builder) ParseTree() *Tree {
	return b.finalEle.nonTerm
}

// DebugTree returns the debug tree which includes all matches and non-matches, and
// non-terminal results (displayed in parentheses) captured throughout parsing. It
// helps in tracing the parsing flow. It's set after the root non-terminal exits.
// Returns nil otherwise.
func (b *Builder) DebugTree() *DebugTree {
	return b.finalDebugTree
}

// Err returns the parsing error. It's set after the root non-terminal exits with a
// false result. Returns nil otherwise.
func (b *Builder) Err() *ParsingError {
	return b.finalErr
}

func (b Builder) mustEnter(operation string) {
	if len(b.stack) == 0 {
		log.Panicf("cannot %s. must Enter a non-terminal first", operation)
	}
}
