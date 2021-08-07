package gobdd

import (
	"fmt"
	"log"
	"math"
	"testing"

	"github.com/timbeurskens/gobdd/algorithm"
	"github.com/timbeurskens/gobdd/bdd_test"
	. "github.com/timbeurskens/gobdd/operators"
	"github.com/timbeurskens/gobdd/operators/bdd"
)

// makeNumber equals the variables in a to the binary representation of n
func makeNumber(a []Variable, n int) Expression {
	repr := make([]Variable, len(a))

	for i := range repr {
		repr[i] = Cons(((n >> i) & 1) == 1)
	}

	return makeEquality(a, repr)
}

func resolveNumber(a []Variable, model Model) int {
	var result = 0
	for i := range a {
		if model[a[i]] {
			result |= 1 << i
		}
	}
	return result
}

func TestNumber(t *testing.T) {
	a := make([]Variable, 8)
	for i := range a {
		a[i] = IVar(i)
	}
	expr := makeNumber(a, 97)
	t.Log(dotExpressionTreeRec(expr))

	expr = algorithm.PruneUnary(expr)
	tree := algorithm.FromExpression(expr)

	t.Log("Size:", Size(expr), "->", Size(tree))

	if model, ok := bdd.FindModel(tree); ok {
		t.Log(model)
		result := model.Variables(true)
		t.Log(result)
		verification := resolveNumber(a, model)
		t.Log(verification)
		if verification != 97 {
			t.Fatal("Incorrect number resolution")
		}
	}
}

// makeAddition equals the variables in c to the addition of a and b
func makeAddition(a, b []Variable, c []Variable) Expression {
	if len(a) != len(b) || len(a)+1 != len(c) {
		panic("length of a and b and carry should match and be 1 smaller than length of c")
	}

	carry := IncVarCollection(len(a))

	exprs := make([]Expression, len(c))
	for i := range c {
		if i == 0 {
			exprs[i] = And(
				Biimplies(c[i], Xor(a[i], b[i])),
				Biimplies(carry[i], And(a[i], b[i])),
			)
		} else if i > 0 && i < len(a) {
			exprs[i] = And(
				Biimplies(c[i], Xor(a[i], b[i], carry[i-1])),
				Biimplies(carry[i], Or(And(a[i], b[i]), And(a[i], carry[i-1]), And(b[i], carry[i-1]))),
			)
		} else { // i == len(a)
			exprs[i] = And(
				Biimplies(c[i], carry[i-1]),
			)
		}
	}

	return And(exprs...)
}

func TestAddition(t *testing.T) {
	a_size, b_size := 3, 3
	a_num, b_num := 3, 6

	a, b, c := IncVarCollection(a_size), IncVarCollection(b_size), IncVarCollection(a_size+1)
	a_expr := makeNumber(a, a_num)
	b_expr := makeNumber(b, b_num)
	c_expr := makeAddition(a, b, c)

	expr := And(a_expr, b_expr, c_expr)

	t.Log(PrintExpressiontree(expr))
	t.Log(dotExpressionTreeRec(expr))

	expr = algorithm.PruneUnary(expr)
	tree := algorithm.FromExpression(expr)

	t.Log("Size:", Size(expr), "->", Size(tree))

	DotSubtree(tree)

	if model, ok := bdd.FindModel(tree); ok {
		t.Log(model)
		result := model.Variables(true)
		t.Log(result)

		t.Log(resolveNumber(a, model), resolveNumber(b, model), resolveNumber(c, model))

		if resolveNumber(c, model) != a_num+b_num {
			t.Fatal("c should be", a_num+b_num)
		}
	}
}

func binaryShift(a []Variable, n int) []Variable {
	res := make([]Variable, len(a)+n)
	for i := range res {
		if i < n {
			res[i] = Cons(false)
		} else {
			res[i] = a[i-n]
		}
	}
	return res
}

func zeroPadding(a []Variable, n int) []Variable {
	res := make([]Variable, len(a)+n)
	for i := range res {
		if i < len(a) {
			res[i] = a[i]
		} else {
			res[i] = Cons(false)
		}
	}
	return res
}

// makeMultiplication
func makeMultiplication(a, b []Variable, c []Variable) Expression {
	if len(a) != len(b) || len(c) != len(a)+len(b) {
		panic("length of a and b should match and length of c should be a + b")
	}

	// if a[i] == 1 then inter[i] = inter[i-1] + b[i] << i
	// if a[i] == 0 then inter[i] = inter[i-1]
	// inter[0] == 0

	inter := make([][]Variable, len(a)+1)
	exprs := make([]Expression, len(a))

	for i := range inter {
		if i == 0 {
			// inter[0] == 0
			inter[i] = ConsCollection(len(b), false)
		} else {
			// inter[i; i > 0] == var...
			if i == len(a) {
				inter[i] = c
			} else {
				inter[i] = IncVarCollection(len(b) + i)
			}

			// add b << i on the i'th step
			bAdded := makeAddition(inter[i-1], binaryShift(b, i-1), inter[i])

			// add 0 on the i'th step
			bEq := makeEquality(inter[i], zeroPadding(inter[i-1], 1))

			exprs[i-1] = And(
				// if the i'th factor of a is 1, then add b << i
				Implies(a[i-1], bAdded),
				// if the i'th factor of a is 0, then add 0
				Implies(Not(a[i-1]), bEq),
			)
		}
	}

	// for every 1 in a, shift b by the position of the 1 in a, then add it to the result
	return And(exprs...)
}

func TestMultiplication(t *testing.T) {
	a_num, b_num := 3, 3
	a_size, b_size := 2, 2

	a, b, c := IncVarCollection(a_size), IncVarCollection(b_size), IncVarCollection(a_size+b_size)
	a_expr := makeNumber(a, a_num)
	b_expr := makeNumber(b, b_num)
	c_expr := makeMultiplication(a, b, c)

	expr := And(a_expr, b_expr, c_expr)

	expr = algorithm.PruneUnary(expr)

	t.Log(PrintExpressiontree(expr))
	t.Log(dotExpressionTreeRec(expr))

	tree := algorithm.FromExpression(expr)

	t.Log("Size:", Size(expr), "->", Size(tree))

	DotSubtree(tree)

	if model, ok := bdd.FindModel(tree); ok {
		t.Log(model)
		result := model.Variables(true)
		t.Log(result)

		t.Log(resolveNumber(a, model), resolveNumber(b, model), resolveNumber(c, model))

		// for i, num := range inter {
		// 	if i > 0 {
		// 		t.Log("inter", i, len(num), model[a[i-1]], resolveNumber(num, model))
		// 	} else {
		// 		t.Log("inter", i, len(num), false, resolveNumber(num, model))
		// 	}
		// }

		if resolveNumber(c, model) != a_num*b_num {
			t.Fatal("c should be", a_num*b_num)
		}
	}
}

// makeEquality constructs an expression such that a and b should be equal
func makeEquality(a, b []Variable) Expression {
	if len(a) != len(b) {
		panic("length of a and b should match")
	}

	exprs := make([]Expression, len(a))
	for i := range exprs {
		exprs[i] = Biimplies(a[i], b[i])
	}

	return And(exprs...)
}

func makePrimeTest(prime int) (Expression, []Variable, []Variable) {
	// estimate the number of bits needed
	bits := int(math.Ceil(math.Log2(float64(prime))))
	bits = 2 + (bits-(bits/2))*2

	log.Printf("Using %d bits for prime computation", bits)

	// construct two arbitrary numbers: a and b
	a, b := IncVarCollection(bits/2), IncVarCollection(bits/2)
	// construct a number c
	c := IncVarCollection(bits)

	exprEq := And(
		// c is the number to test for prime
		makeNumber(c, prime),
		// a != 1
		Not(makeNumber(a, 1)),
		// b != 1
		Not(makeNumber(b, 1)),
	)

	// c == a * b
	exprMul := makeMultiplication(a, b, c)

	// a * b == c and c == number to test
	return And(exprEq, exprMul), a, b
}

func TestIsPrimeCDCL(t *testing.T) {
	prime := 15
	expr, a, b := makePrimeTest(prime)

	// convert to NNF
	nnf := algorithm.NNF(expr)
	cnf := algorithm.TransformTseitin(nnf)

	sat, m := algorithm.CDCL(cnf)

	t.Log(sat, m, a, b)
}

func TestIsPrimeBDD(t *testing.T) {
	bench := bdd_test.Bench{T: t}

	prime := 15
	expr, a, b := makePrimeTest(prime)

	// prepare the expression tree
	expr = algorithm.PruneUnary(expr)

	t.Log(dotExpressionTreeRec(expr))

	// run bdd algorithm
	tree := algorithm.FromExpression(expr)

	t.Log("Size:", Size(expr), "->", Size(tree))

	bench.AssertUnsat(fmt.Sprintf("%d is prime", prime), tree)

	if bdd.Sat(tree) {
		if model, ok := bdd.FindModel(tree); ok {
			aResolv := resolveNumber(a, model)
			bResolv := resolveNumber(b, model)
			t.Fatalf("Decomposition of %d: %d x %d", prime, aResolv, bResolv)
		} else {
			t.Fatal("Could not construct model of non-prime number")
		}
	}
}
