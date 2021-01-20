package algorithm

import (
    "testing"

    "gobdd/bdd_test"
    "gobdd/operators"
)

func TestSat(t *testing.T) {
    be := bdd_test.Bench{T: t}
    a, b, c := operators.Var("a"), operators.Var("b"), operators.Var("c")
    be.AssertSat("a or not b and c or a is sat", CDCL(operators.CNF{operators.NClause{a, b.Negate()}, operators.NClause{c, a}}))
}

func TestUnsat(t *testing.T) {
    b := bdd_test.Bench{T: t}
    a := operators.Var("a")
    b.AssertUnsat("a and not a is unsat", CDCL(operators.CNF{operators.NClause{a}, operators.NClause{a.Negate()}}))
}
