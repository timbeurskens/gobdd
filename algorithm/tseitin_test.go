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
		operators.Implies(operators.Biimplies(a, b), operators.And(operators.Or(b, c), a)),
		operators.Or(a, b, c),
		operators.And(a, b, c),
	}
)

func TestCNF(t *testing.T) {
	be := bdd_test.Bench{T: t}
	for _, e := range expressions {
		t.Run(fmt.Sprintf("Expression %s is equal to CNF of expression", e.String()), func(t *testing.T) {
			be.AssertEquivalent("expression is equal to CNF of expression", FromExpression(e), FromExpression(CNF(e).Expr()))
		})
	}
}

func TestNNF(t *testing.T) {
	be := bdd_test.Bench{T: t}

	for _, e := range expressions {
		t.Run(fmt.Sprintf("Expression %s is equal to NNF of expression", e.String()), func(t *testing.T) {
			be.AssertEquivalent("expression is equal to NNF of expression", FromExpression(e), FromExpression(NNF(e)))
		})
	}
}
