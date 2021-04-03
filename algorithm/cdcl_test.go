package algorithm

import (
	"github.com/timbeurskens/gobdd/bdd_test"
	"github.com/timbeurskens/gobdd/operators"
	bdd2 "github.com/timbeurskens/gobdd/operators/bdd"
	"testing"
)

func TestSat(t *testing.T) {
	be := bdd_test.Bench{T: t}
	a, b, c := operators.Var("a"), operators.Var("b"), operators.Var("c")
	sat, _ := CDCL(operators.CNF{operators.NClause{a, b.Negate()}, operators.NClause{c, a}})
	be.Assert("a or not b and c or a is sat", sat)
}

func TestUnsat(t *testing.T) {
	b := bdd_test.Bench{T: t}
	a := operators.Var("a")
	sat, _ := CDCL(operators.CNF{operators.NClause{a}, operators.NClause{a.Negate()}})
	b.Assert("a and not a is unsat", !sat)
}

func TestCNFXorSat(t *testing.T) {
	be := bdd_test.Bench{T: t}
	a := operators.Var("a")
	b := operators.Var("b")

	sat, _ := CDCL(operators.CNF{
		operators.NClause{a.Negate(), b.Negate()},
		operators.NClause{a, b},
	})

	be.Assert("a xor b is sat", sat)
}

func TestCNFSat(t *testing.T) {
	be := bdd_test.Bench{T: t}
	a := operators.Var("a")
	b := operators.Var("b")
	c := operators.Var("c")

	sat, _ := CDCL(operators.CNF{
		operators.NClause{a.Negate(), b.Negate(), c.Negate()},
		operators.NClause{a, b, c},
	})

	be.Assert("a, b, c is sat", sat)
}

func TestCNFXorUnsat(t *testing.T) {
	be := bdd_test.Bench{T: t}
	a := operators.Var("a")

	sat, _ := CDCL(operators.CNF{
		operators.NClause{a.Negate(), a.Negate()},
		operators.NClause{a, a},
	})

	be.Assert("a xor a is unsat", !sat)
}

func testCDCLTseitin(t *testing.T) {
	be := bdd_test.Bench{T: t}

	p, q, r, s := operators.Var("p"), operators.Var("q"), operators.Var("r"), operators.Var("s")

	// this example shows potential improvements to the tool: sharing for term s is possible
	expr := operators.And(operators.Biimplies(s, q), operators.Or(r, operators.Not(p)))

	// convert to NNF
	nnf := NNF(expr)

	cnf := TransformTseitin(nnf)

	sat, m := CDCL(cnf)

	be.Assert("(s <-> q) && (r || -p)", sat)

	model, ok := bdd2.FindModel(m)
	t.Log(ok, model)
	be.Assert("bdd has model", ok)

	counter, ok := bdd2.FindCounterExample(m)
	t.Log(ok, counter)
	be.Assert("bdd also has counterexample", ok)
}
