package algorithm

import (
	"fmt"
	"gobdd/operators/bdd"
	"testing"

	"gobdd/bdd_test"
	op "gobdd/operators"
)

var (
	a = op.Var("a")
	b = op.Var("b")
	c = op.Var("c")
	d = op.Var("d")

	expressions = []op.Expression{
		op.Implies(a, b),
		op.Biimplies(a, b),
		op.And(a, op.Not(a)),
		op.Biimplies(a, op.Implies(b, c)),
		op.Xor(a, b),
		op.Implies(op.Biimplies(a, b), op.And(op.Or(b, c), a)),
		op.Or(a, b, c),
		op.And(a, b, c),
		op.Not(op.Or(a, b)),
		op.Not(op.Xor(a, b)),
		op.And(op.Or(&op.TrueConst, a), op.And(b, &op.FalseConst)),
		op.And(op.Xor(c, a), op.Or(b, op.Xor(a, op.Xor(b, c)))),
		op.Biimplies(a, op.Biimplies(b, op.Biimplies(c, d))),
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
