package gobdd

import (
	"github.com/timbeurskens/gobdd/algorithm"
	"github.com/timbeurskens/gobdd/bdd_test"
	. "github.com/timbeurskens/gobdd/operators"
	"github.com/timbeurskens/gobdd/operators/bdd"
	"testing"
)

func TestFalseEquivalence(t *testing.T) {
	b := bdd_test.Bench{T: t}
	b.AssertEquivalent("false equals false", Cons(false), Cons(false))
}

func TestConjNeg(t *testing.T) {
	b := bdd_test.Bench{T: t}

	p := Var("p")

	falseTree := algorithm.FromExpression(Cons(false))
	pTree := algorithm.FromExpression(algorithm.PruneUnary(And(p, Not(p))))

	t.Log(pTree)

	// test for conjunction and negation
	b.AssertEquivalent("P and not P should be equal to false", falseTree, pTree)
}

func TestComplicated(t *testing.T) {
	// b := bdd_test.Bench{T: t}

	p, q, r := Var("p"), Var("q"), Var("r")
	expr := And(p, q, r, Or(p, q, r))
	exprTree := algorithm.FromExpression(expr)

	t.Log(expr)
	PrintSubtree(exprTree)
	DotSubtree(exprTree)
}

func TestSat(t *testing.T) {
	b := bdd_test.Bench{T: t}

	b.AssertSat("p has a satisfiable assignment", algorithm.FromExpression(Var("p")))
}

func TestUnsat(t *testing.T) {
	b := bdd_test.Bench{T: t}

	b.AssertUnsat("false is unsatisfiable", algorithm.FromExpression(Cons(false)))
}

func TestUnsat2(t *testing.T) {
	b := bdd_test.Bench{T: t}

	p := Var("p")

	expr := algorithm.PruneUnary(And(p, p.Negate()))

	b.AssertUnsat("false is unsatisfiable", algorithm.FromExpression(expr))
}

func TestEquivalence1(t *testing.T) {
	b := bdd_test.Bench{T: t}

	b.AssertEquivalent("false is equivalent to false", algorithm.FromExpression(Cons(false)), algorithm.FromExpression(Cons(false)))
	b.AssertNotEquivalent("false is not equivalent to true", algorithm.FromExpression(Cons(false)), algorithm.FromExpression(Cons(true)))
}

func TestEquivalence(t *testing.T) {
	b := bdd_test.Bench{T: t}

	p, q := Var("p"), Var("q")

	b.AssertEquivalent("p xor q is equivalent to not(p biimplication q)", algorithm.FromExpression(Xor(p, q)), algorithm.FromExpression(algorithm.PruneUnary(Not(Biimplies(p, q)))))
}

func TestImplication(t *testing.T) {
	b := bdd_test.Bench{T: t}
	p := Var("p")
	b.AssertUnsat("not p implies true is unsatisfiable", algorithm.FromExpression(algorithm.PruneUnary(Not(Implies(p, Cons(true))))))
}

func TestTautology(t *testing.T) {
	b := bdd_test.Bench{T: t}
	p := Var("p")

	b.AssertTautology(
		"p or not p is a tautology",
		algorithm.FromExpression(
			algorithm.PruneUnary(Or(p, Not(p))),
		),
	)
}

func TestModel(t *testing.T) {
	b := bdd_test.Bench{T: t}

	p := Var("p")

	model, ok := bdd.FindModel(algorithm.FromExpression(p))

	t.Log(model)

	b.Assert("p has model", ok)

	b.Assert("model should have p=true", model[p])
}

func TestSize(t *testing.T) {
	b := bdd_test.Bench{T: t}

	p := Var("p")
	q := Var("q")

	b.AssertSize("simple choice should have size 3", 3, algorithm.FromExpression(p))

	b.AssertSize("conjunction should have size 4", 4, algorithm.FromExpression(And(p, q)))
}
