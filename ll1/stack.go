package main

import (
	t "github.com/shivammg/parsers/types"
)

type stackEle struct {
	*t.Tree
	parent *t.Tree
}

type stack struct {
	s []stackEle
}

func (st *stack) Push(node, parent *t.Tree) {
	st.s = append(st.s, stackEle{Tree: node, parent: parent})
}

func (st *stack) Pop() *stackEle {
	l := len(st.s)
	if l == 0 {
		return nil
	}

	s := st.s[l-1]
	st.s = st.s[:l-1]
	return &s
}

func (st stack) isEmpty() bool {
	return len(st.s) == 0
}
