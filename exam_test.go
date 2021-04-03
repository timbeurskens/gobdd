package gobdd

import (
	"github.com/timbeurskens/gobdd/algorithm"
	"github.com/timbeurskens/gobdd/bdd_test"
	. "github.com/timbeurskens/gobdd/operators"
	bdd2 "github.com/timbeurskens/gobdd/operators/bdd"
	"testing"
)

func TestJan2020ex1b(t *testing.T) {
	b := bdd_test.Bench{T: t}
	p, q, r, s := Var("p"), Var("q"), Var("r"), Var("s")

	// this example shows potential improvements to the tool: sharing for term s is possible
	expr := And(Biimplies(s, q), Or(r, Not(p)))

	bdd := algorithm.FromExpression(algorithm.PruneUnary(expr))
	DotSubtree(bdd)

	b.AssertSat("(s <-> q) && (r || -p)", bdd)

	model, ok := bdd2.FindModel(bdd)
	t.Log(ok, model)
	b.Assert("bdd has model", ok)

	counter, ok := bdd2.FindCounterExample(bdd)
	t.Log(ok, counter)
	b.Assert("bdd also has counterexample", ok)
}
