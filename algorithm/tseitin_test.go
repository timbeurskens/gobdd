package algorithm

import (
	"fmt"
	"testing"

	"gobdd/bdd_test"
	"gobdd/operators"
)

var (
	a = operators.Var("a")
	b = operators.Var("b")
	c = operators.Var("c")

	expressions = []operators.Expression{
		operators.Implies(a, b),
		operators.Biimplies(a, b),
		operators.Biimplies(a, operators.Implies(b, c)),
		operators.Xor(a, b),
		operators.Implies(operators.Biimplies(a, b), operators.And(operators.Or(b, c), a)),
		operators.Or(a, b, c),
		operators.And(a, b, c),
		operators.Not(operators.Or(a, b)),
	}
)

func TestTransformTseitin(t *testing.T) {
	for i, e := range expressions {
		t.Run(fmt.Sprintf("Expression %d is equal to CNF of expression", i), func(t *testing.T) {
			be := bdd_test.Bench{T: t}
			nnf := NNF(e)
			cnf := TransformTseitin(nnf)
			be.AssertEquivalent("expression is equal to CNF of expression", FromExpression(e), FromExpression(cnf.Expr()))
		})
	}
}

func TestCNF(t *testing.T) {
	for i, e := range expressions {
		t.Run(fmt.Sprintf("Expression %d is equal to CNF of expression", i), func(t *testing.T) {
			be := bdd_test.Bench{T: t}
			be.AssertEquivalent("expression is equal to CNF of expression", FromExpression(e), FromExpression(CNF(e).Expr()))
		})
	}
}

func TestNNF(t *testing.T) {
	for i, e := range expressions {
		t.Run(fmt.Sprintf("Expression %d is equal to NNF of expression", i), func(t *testing.T) {
			nnf := NNF(e)
			a, b := FromExpression(PruneUnary(e)), FromExpression(PruneUnary(nnf))
			be := bdd_test.Bench{T: t}
			be.AssertEquivalent("expression is equal to NNF of expression", a, b)
		})
	}
}
