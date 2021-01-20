package algorithm

import "gobdd/operators"

// CNF should yield 3-clause elements, to be used by Tseitin transformation
func CNF(e operators.Expression) operators.CNF {
    panic("not implemented")
}

// NNF converts a given expression to negation-normal-form by replacing every operator to an equivalent disjunction/conjunction/negation
// and applying demorgan on the expressions to push negation to leafs
func NNF(e operators.Expression) operators.Expression {
    op, ok := e.(operators.Operator)
    if ok {
        return op.Normalize()
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
