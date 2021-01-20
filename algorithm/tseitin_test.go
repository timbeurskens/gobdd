package algorithm

import (
    "testing"

    "gobdd/bdd_test"
    "gobdd/operators"
)

var (
    a = operators.Var("a")
    b = operators.Var("b")
    c = operators.Var("c")
    e = operators.Implies(operators.Biimplies(a,b), operators.And(operators.Or(b, c), a))
)

func TestCNF(t *testing.T) {
    be := bdd_test.Bench{T: t}

    be.AssertEquivalent("expression is equal to CNF of expression", FromExpression(e), FromExpression(CNF(e).Expr()))
}

func TestNNF(t *testing.T) {
    be := bdd_test.Bench{T: t}

    be.AssertEquivalent("expression is equal to NNF of expression", FromExpression(e), FromExpression(NNF(e)))
}