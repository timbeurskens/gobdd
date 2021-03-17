package algorithm

import (
	"gobdd/operators"
	"log"
)

// CNF should yield 3-clause elements, to be used by Tseitin transformation
func CNF(e operators.Expression) operators.CNF {
	log.Println("cnf(", e.LeftChild(), e, e.RightChild().LeftChild(), e.RightChild(), e.RightChild().RightChild(), ")")
	//panic("not implemented")
	return nil
}

// NNF converts a given expression to negation-normal-form by replacing every operator to an equivalent disjunction/conjunction/negation
// and applying demorgan on the expressions to push negation to leafs
func NNF(e operators.Expression) operators.Expression {
	op, ok := e.(operators.Operator)
	if ok {
		// todo: enable NNF Normalization
		// return op.Normalize()
	}
	return op
}

// a <-> (b -> c)
// -------
// (-a || (b -> c))
// (-a || -b || c)
// ---+---
// (-(b->c) || a)
// (-(-b || c) || a)
// ((b && -c) || a)
// (b || a) && (-c || a)

// a <-> (b || c)

// a <-> (b && c)
