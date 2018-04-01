package main

type stackEle struct {
	*Node
	parent *Node
}

type stack struct {
	s []stackEle
}

func (st *stack) Push(node, parent *Node) {
	st.s = append(st.s, stackEle{Node: node, parent: parent})
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
