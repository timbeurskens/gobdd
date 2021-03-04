package algorithm

import (
	"gobdd/operators"
	"gobdd/operators/bdd"
)

// todo: optimization:
// The new Choice object in algorithm Apply is currently not optimal:
// the garbage collector has a hard time managing the objects and takes ~33% of computation time

// robdd(false) = false
// robdd(true) = true
// robdd(p) = p(1, 0)
// robdd(-phi) = robdd(phi -> false)
// robdd(phi # rho) = apply(robdd(phi), robdd(rho), #)

func buildTree(e operators.Expression) (root operators.Node) {
	if cons, ok := e.(operators.Constant); ok {
		// every constant remains a constant
		return cons
	} else if v, ok := e.(operators.Variable); ok {
		// for every variable p: introduce choice p(true, false)
		return operators.NewChoice(v)
	} else if v, ok := e.(*operators.Negation); ok {
		// 	special case for negations, due to a compatibility issue for CNF terms, negations need to be converted
		return buildTree(operators.Implies(v.Negate(), operators.Cons(false)))
	} else if op, ok := e.(operators.Operator); ok {
		// first make sure the subtrees are complete
		a, b := buildTree(e.LeftChild()), buildTree(e.RightChild())

		// do an apply step on the two subtrees with the given expression e
		return Apply(a, b, op)
	}
	return e
}

// not working at the moment due to equivalence issue
func reduceTree(root operators.Node) operators.Node {
	switch root.(type) {
	case operators.Constant:
		return root
	case *operators.Choice:
		left := reduceTree(root.LeftChild())
		right := reduceTree(root.RightChild())

		if bdd.Equivalent(left, right) {
			return left
		} else {
			return operators.JoinByChoice(root.(*operators.Choice).Var, left, right)
		}
	}
	return root
}

// FromExpression builds a bdd from a given expression
func FromExpression(e operators.Expression) operators.Node {
	root := buildTree(e)
	return reduceTree(root)
}

// apply(T, U, #) = #(T, U)

// #(T, U) = T # U if T,U in {true,false}
// #(T, U) =
// let p be the smallest variable in TuU
// if p is on top of both T and U: #(p(T1, T2), p(U1, U2)) = p(#(T1, U1), #(T2, U2))
// if p is on top of T but does not occur in U: #(p(T1, T2), U) = p(#(T1, U), #(T2, U))
// if p is on top of U but does not occur in T: #(T, p(U1, U2)) = p(#(T, U1), #(T, U2))

func Apply(a, b operators.Node, op operators.Operator) operators.Node {
	// a and b both constant: evaluate according to truth table
	if operators.IsConstant(a) && operators.IsConstant(b) {
		return op.ConstEval(a.(operators.Constant), b.(operators.Constant))
	}

	// propagate operator to level below
	v, left, right := operators.FindOperatorPropagation(a, b, op)

	// v.SwapTrue(FromExpression(left))
	// v.SwapFalse(FromExpression(right))

	// todo: use choice operator from variable v
	// return v
	return operators.JoinByChoice(v.Var, FromExpression(left), FromExpression(right))
}

// todo: introduce simplifications for implication, biimplication, xor, nor to CNF
// using a unary negation operation (throw negation to leaf, simplify not true -> false, not false -> true)
// reduce using tseitin transformation
