package bdd

import "github.com/timbeurskens/gobdd/operators"

// Equivalent returns true iff subtree a and b are equivalent
func Equivalent(a, b operators.Node) bool {
	if a == nil || b == nil {
		return a == b
	}

	// a is equivalent to b if the nodes are equivalent and their respective left and right child are equivalent
	return a.NodeEquivalent(b) &&
		Equivalent(a.LeftChild(), b.LeftChild()) &&
		Equivalent(a.RightChild(), b.RightChild())
}

// Sat returns true iff there is a satisfying assignment for subtree n
// note that a subtree is satisfiable if it is not unsatisfiable, thus subtree n is not equivalent to subtree "false"
func Sat(n operators.Node) bool {
	return !Unsat(n)
}

// Unsat returns true iff there is no satisfying assignment for subtree n
// note that a subtree is unsatisfiable if it is not satisfiable, thus subtree n is equivalent to subtree "false"
func Unsat(n operators.Node) bool {
	return Equivalent(n, operators.Cons(false))
}

// Tautology returns true iff the expression yields true for every possible assignment
func Tautology(n operators.Node) bool {
	return Equivalent(n, operators.Cons(true))
}

// FindModel searches for a satisfying assignment
func FindModel(n operators.Node) (operators.Model, bool) {
	assignment := make(operators.Model)

	if SubtreeSearch(n, assignment, operators.Cons(true)) {
		return assignment, true
	}
	return nil, false
}

// FindCounterExample searches for an assignment such that the expression returns false
func FindCounterExample(n operators.Node) (operators.Model, bool) {
	assignment := make(operators.Model)
	if SubtreeSearch(n, assignment, operators.Cons(false)) {
		return assignment, true
	}
	return nil, false
}

func SubtreeSearch(root operators.Node, assignment operators.Model, search operators.Constant) bool {
	switch root.(type) {
	case *operators.Choice:
		// try the left subtree
		assignment[root.(*operators.Choice).Var] = true
		if SubtreeSearch(root.LeftChild(), assignment, search) {
			return true
		}

		// try the right subtree
		assignment[root.(*operators.Choice).Var] = false
		if SubtreeSearch(root.RightChild(), assignment, search) {
			return true
		}

		// no possible solution
		delete(assignment, root.(*operators.Choice).Var)

		return false
	case operators.Constant:
		// in the leaf a result can either be true or false
		return search == root
	}

	// if the tree has nodes other than choice and constant, fail immediately
	return false
}
