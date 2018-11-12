package rd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExit_NilResult(t *testing.T) {
	b := NewBuilder(nil)
	b.Enter("")
	assert.Panics(t, func() {
		b.Exit(nil)
	}, "Exit result must be non-nil")
}

func TestExit_FinalEleAndDebugTree(t *testing.T) {
	b := NewBuilder(nil)
	b.Enter("root")
	assert.Equal(t, ele{}, b.finalEle)
	assert.Nil(t, b.finalDebugTree)
	root := b.stack.peek()
	rootDebugTree := b.debugStack.peek()
	result := true
	b.Exit(&result)
	assert.Equal(t, root, b.finalEle)
	assert.Equal(t, rootDebugTree, b.finalDebugTree)
}

func TestExit_FinalErr(t *testing.T) {
	b := NewBuilder(nil)
	b.Enter("root")
	assert.Nil(t, b.finalErr)
	result := false
	b.Exit(&result)
	assert.NotNil(t, b.finalErr)
}

func TestExit_AddToParent(t *testing.T) {
	b := NewBuilder(nil)
	b.Enter("root")
	root := b.stack.peek()
	rootDebugTree := b.debugStack.peek()
	b.Enter("child")
	child := b.stack.peek()
	childDebugTree := b.debugStack.peek()
	result := true
	b.Exit(&result)
	assert.Contains(t, root.nonTerm.Subtrees, child.nonTerm)
	assert.Contains(t, rootDebugTree.subtrees, childDebugTree)
}

func TestExit_FalseResult(t *testing.T) {
	b := NewBuilder(nil)
	b.current = 1
	b.Enter("")
	b.current = 2
	result := false
	b.Exit(&result)
	assert.Equal(t, 1, b.current, "current must be reset")
}
