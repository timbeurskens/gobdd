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

func DeMorgan(e operators.Expression) operators.Expression {
	if neg, ok := e.(*operators.Negation); ok {
		child := neg.LeftChild()
		switch child.(type) {
		case *operators.Conjunction:
			return &operators.Disjunction{
				A: DeMorgan(operators.Not(child.LeftChild())),
				B: DeMorgan(operators.Not(child.RightChild())),
			}
		case *operators.Disjunction:
			return &operators.Conjunction{
				A: DeMorgan(operators.Not(child.LeftChild())),
				B: DeMorgan(operators.Not(child.RightChild())),
			}
		case operators.Constant:
			// boolean constants are easily negated
			return child.(operators.Constant).Negate()
		case operators.Variable:
			return e
		default:
			panic("unsupported type for DeMorgan, should be normalized")
		}
	} else {
		if e.LeftChild() != nil {
			e.SetLeftChild(DeMorgan(e.LeftChild()))
		}
		if e.RightChild() != nil {
			e.SetRightChild(DeMorgan(e.RightChild()))
		}
		return e
	}
}

// NNF converts a given expression to negation-normal-form by replacing every operator to an equivalent disjunction/conjunction/negation
// and applying demorgan on the expressions to push negation to leafs
func NNF(e operators.Expression) operators.Expression {
	// convert to conjunctions, negations and disjunctions
	normal := e.Normalize()

	return DeMorgan(normal)
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
