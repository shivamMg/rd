package rd

import (
	"fmt"
	"log"
)

type ParsingError struct{}

func (e *ParsingError) Error() string {
	// TODO: Must return useful detail. ex. no tokens left
	return "parsing error"
}

// Builder helps in building a recursive descent parser.
// It stores a slice of tokens and an index to the current token.
// It maintains a stack used in building the parse tree (check Tree()).
// It also builds a debug tree that helps in understanding the parsing
// flow (check DebugTree()).
// Enter/Exit methods are used in logging enter and exit of non-terminal
// functions.
// Add/Next/Match/Reset are used while working with terminals.
type Builder struct {
	tokens         []Token
	current        int
	stack          stack
	finalEle       ele
	debugStack     debugStack
	finalDebugTree *debugTree
	finalErr       *ParsingError
}

// NewBuilder returns a new Builder for tokens.
func NewBuilder(tokens []Token) *Builder {
	return &Builder{
		tokens:     tokens,
		current:    -1,
		stack:      stack{},
		debugStack: debugStack{},
	}
}

// Next returns the next token and increments the current index. ok is false if
// no tokens are left, else true.
func (b *Builder) Next() (token Token, ok bool) {
	b.mustEnter("Next")
	if b.current == len(b.tokens)-1 {
		return nil, false
	}
	b.current++
	return b.tokens[b.current], true
}

// Reset resets the current index for the current non-terminal, and discards
// any matches done inside it.
func (b *Builder) Reset() {
	b.mustEnter("Reset")
	e := b.stack.peek()
	b.current = e.index
	e.nonTerm.Subtrees = []*Tree{}
}

// Add adds token as a child subtree of the current non-terminal tree.
func (b *Builder) Add(token Token) {
	b.mustEnter("Add")
	e := b.stack.peek()
	e.nonTerm.Add(NewTree(fmt.Sprint(token)))
}

// Match matches token with the next token obtained through Next method. In
// case of a non-match the current index is decremented to its original value.
// ok is false in case of non-match or no-tokens-left, else true for a match.
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

// Enter pushes non-terminal on the stack making it the current non-terminal.
// Subsequent matches, and calls to other non-terminals are added to the current
// non-terminal until the Exit call. It should be the first statement inside the
// non-terminal's function.
func (b *Builder) Enter(nonTerm string) {
	b.stack.push(ele{
		index:   b.current,
		nonTerm: NewTree(nonTerm),
	})
	b.debugStack.push(newDebugTree(nonTerm))
}

// Exit pops the current non-terminal from the stack. In case of false result
// the current index is reset to where it was before entering the current
// non-terminal. In case of true result: 1) the final parse tree is set if the
// current non-terminal was root 2) else it's added to its parent non-terminal.
// It should be placed right next to Enter in a defer statement.
func (b *Builder) Exit(result *bool) {
	b.mustEnter("Exit")
	if result == nil {
		panic("Exit result cannot be nil")
	}
	e := b.stack.pop()
	if *result {
		if b.stack.isEmpty() {
			b.finalEle = e
		} else {
			parent := b.stack.peek()
			parent.nonTerm.Add(e.nonTerm)
		}
	} else {
		b.current = e.index
	}
	dt := b.debugStack.pop()
	dt.data += fmt.Sprintf("(%t)", *result)
	if b.debugStack.isEmpty() {
		b.finalDebugTree = dt
		if !*result {
			b.finalErr = &ParsingError{}
		}
	} else {
		parent := b.debugStack.peek()
		parent.add(dt)
	}
}

// Tree returns the parse tree. It's set after the root non-terminal exits with
// true result.
func (b *Builder) Tree() *Tree {
	return b.finalEle.nonTerm
}

// DebugTree returns a tree that includes all matches and non-matches, and
// non-terminal results (displayed in parentheses) captured throughout parsing.
// Helps in understanding the parsing flow. It's set after the root non-terminal
// exits. It has methods Print and Sprint.
func (b *Builder) DebugTree() *debugTree {
	return b.finalDebugTree
}

// Err returns the parsing error. It's set after the root non-terminal exits
// with false result.
func (b *Builder) Err() *ParsingError {
	return b.finalErr
}

func (b Builder) mustEnter(operation string) {
	if len(b.stack) == 0 {
		log.Panicf("cannot %s. must Enter a non-terminal first", operation)
	}
}
