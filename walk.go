package rd

type WalkStatus int

const (
	// GoToNext is the default traversal of every node.
	GoToNext WalkStatus = iota
	// SkipChildren tells walker to skip all children of current node.
	SkipChildren
	// Terminate tells walker to terminate the traversal.
	Terminate
)

// TreeVisitor is a callback to be called when traversing the syntax tree.
// Called twice for every node: once with entering=true when the branch is
// first visited, then with entering=false after all the children are done.
type TreeVisitor interface {
	Visit(tree *Tree, entering bool) WalkStatus
}

// TreeVisitorFunc casts a function to match TreeVisitor interface.
type TreeVisitorFunc func(tree *Tree, entering bool) WalkStatus

// Walk traverses tree recursively.
func Walk(tree *Tree, visitor TreeVisitor) WalkStatus {

	status := visitor.Visit(tree, true) // entering
	if status == Terminate {
		visitor.Visit(tree, false)
		return status
	}
	if status != SkipChildren {
		for _, t := range tree.Subtrees {
			status = Walk(t, visitor)
			if status == Terminate {
				return status
			}
		}
	}
	status = visitor.Visit(tree, false) // exiting
	if status == Terminate {
		return status
	}
	return GoToNext
}

// Visit calls visitor function.
func (f TreeVisitorFunc) Visit(tree *Tree, entering bool) WalkStatus { return f(tree, entering) }

// WalkFunc is like Walk but accepts just a callback function.
func WalkFunc(tree *Tree, f TreeVisitorFunc) {
	visitor := TreeVisitorFunc(f)
	Walk(tree, visitor)
}
