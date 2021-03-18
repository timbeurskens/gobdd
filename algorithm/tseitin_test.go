package algorithm

import (
	"fmt"
	"gobdd/operators/bdd"
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
		operators.And(a, operators.Not(a)),
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
			pruned := PruneUnary(e)
			nnf := NNF(e)
			cnf := TransformTseitin(nnf)
			cdclSat, _ := CDCL(cnf)
			resBdd := FromExpression(pruned)
			satBdd := bdd.Sat(resBdd)
			be.Assert("cdcl and bdd are SAT equivalent", satBdd == cdclSat)
			//be.Assert("cdcl and bdd are Tautology equivalent", bdd.Tautology(resCdcl) == bdd.Tautology(resBdd))
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
