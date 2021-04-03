package algorithm

import (
	"github.com/timbeurskens/gobdd/operators"
)

// CNF should yield 3-clause elements, to be used by Tseitin transformation
// assume e is an expression consisting of a bi-implication with left-child = variable & right-child = single-operator expression
func CNF(e operators.Expression) (cnf operators.CNF) {
	cnf = make(operators.CNF, 0, 4)

	leftVar, ok := e.LeftChild().(operators.Variable)
	if !ok {
		panic("cnf: left is not a variable")
	}

	switch e.RightChild().(type) {
	case operators.Variable:
		rightVar := e.RightChild().(operators.Term)
		cnf = append(cnf, operators.CNF{
			operators.NClause{leftVar, rightVar.Negate()},
			operators.NClause{leftVar.Negate(), rightVar},
		}...)
	case *operators.Negation:
		rightVar := e.RightChild().(operators.Term)
		cnf = append(cnf, operators.CNF{
			operators.NClause{leftVar, rightVar.Negate()},
			operators.NClause{leftVar.Negate(), rightVar},
		}...)
	case *operators.Disjunction:
		dis := e.RightChild().(*operators.Disjunction)
		l, r := dis.LeftChild().(operators.Term), dis.RightChild().(operators.Term)
		cnf = append(cnf, operators.CNF{
			operators.NClause{leftVar.Negate(), l, r},
			operators.NClause{leftVar, l.Negate()},
			operators.NClause{leftVar, r.Negate()},
		}...)
	case *operators.Conjunction:
		dis := e.RightChild().(*operators.Conjunction)
		l, r := dis.LeftChild().(operators.Term), dis.RightChild().(operators.Term)
		cnf = append(cnf, operators.CNF{
			operators.NClause{leftVar, l.Negate(), r.Negate()},
			operators.NClause{leftVar.Negate(), l},
			operators.NClause{leftVar.Negate(), r},
		}...)
	case operators.Constant:
		c := e.RightChild().(operators.Constant)
		var t operators.Term

		if c.Value() {
			t = leftVar
		} else {
			t = leftVar.Negate()
		}
		cnf = append(cnf, operators.CNF{
			t,
		}...)
	default:
		panic("unrecognized right side")
	}

	return
}

func DeMorgan(e operators.Expression) operators.Expression {
	if neg, ok := e.(*operators.Negation); ok {
		child := neg.RightChild()
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
