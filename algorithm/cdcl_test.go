package algorithm

import (
	bdd2 "gobdd/operators/bdd"
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

func TestCNFXorSat(t *testing.T) {
	be := bdd_test.Bench{T: t}
	a := operators.Var("a")
	b := operators.Var("b")

	be.AssertSat("a xor b is sat", CDCL(operators.CNF{
		operators.NClause{a.Negate(), b.Negate()},
		operators.NClause{a, b},
	}))
}

func TestCNFSat(t *testing.T) {
	be := bdd_test.Bench{T: t}
	a := operators.Var("a")
	b := operators.Var("b")
	c := operators.Var("c")

	be.AssertSat("a, b, c is sat", CDCL(operators.CNF{
		operators.NClause{a.Negate(), b.Negate(), c.Negate()},
		operators.NClause{a, b, c},
	}))
}

func TestCNFXorUnsat(t *testing.T) {
	be := bdd_test.Bench{T: t}
	a := operators.Var("a")

	be.AssertUnsat("a xor a is unsat", CDCL(operators.CNF{
		operators.NClause{a.Negate(), a.Negate()},
		operators.NClause{a, a},
	}))
}

func testCDCLTseitin(t *testing.T) {
	be := bdd_test.Bench{T: t}

	p, q, r, s := operators.Var("p"), operators.Var("q"), operators.Var("r"), operators.Var("s")

	// this example shows potential improvements to the tool: sharing for term s is possible
	expr := operators.And(operators.Biimplies(s, q), operators.Or(r, operators.Not(p)))

	// convert to NNF
	nnf := NNF(expr)

	cnf := TransformTseitin(nnf)

	sat := CDCL(cnf)

	be.AssertSat("(s <-> q) && (r || -p)", sat)

	model, ok := bdd2.FindModel(sat)
	t.Log(ok, model)
	be.Assert("bdd has model", ok)

	counter, ok := bdd2.FindCounterExample(sat)
	t.Log(ok, counter)
	be.Assert("bdd also has counterexample", ok)
}
