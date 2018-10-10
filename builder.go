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

type Builder struct {
	tokens         []Token
	current        int
	stack          stack
	finalEle       ele
	debugStack     debugStack
	finalDebugTree *debugTree
	finalErr       *ParsingError
}

func NewBuilder(tokens []Token) *Builder {
	return &Builder{
		tokens:     tokens,
		current:    -1,
		stack:      stack{},
		debugStack: debugStack{},
	}
}

func (b *Builder) Next() (token Token, ok bool) {
	b.mustEnter("Next")
	if b.current == len(b.tokens)-1 {
		return nil, false
	}
	b.current++
	return b.tokens[b.current], true
}

func (b *Builder) Reset() {
	b.mustEnter("Reset")
	e := b.stack.peek()
	b.current = e.index
	e.nonTerm.Subtrees = []*Tree{}
}

func (b *Builder) Add(token Token) {
	b.mustEnter("Add")
	e := b.stack.peek()
	e.nonTerm.Add(NewTree(fmt.Sprint(token)))
}

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

func (b *Builder) Enter(nonTerm string) {
	b.stack.push(ele{
		index:   b.current,
		nonTerm: NewTree(nonTerm),
	})
	b.debugStack.push(newDebugTree(nonTerm))
}

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
		b.finalErr = &ParsingError{}
	} else {
		parent := b.debugStack.peek()
		parent.add(dt)
	}
}

func (b *Builder) Tree() *Tree {
	return b.finalEle.nonTerm
}

func (b *Builder) DebugTree() *debugTree {
	return b.finalDebugTree
}

func (b *Builder) Err() *ParsingError {
	return b.finalErr
}

func (b Builder) mustEnter(operation string) {
	if len(b.stack) == 0 {
		log.Panicf("cannot %s. must Enter a non-terminal first", operation)
	}
}
