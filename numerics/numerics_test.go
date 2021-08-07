package numerics

import (
	"fmt"
	"github.com/timbeurskens/gobdd/algorithm"
	"github.com/timbeurskens/gobdd/bdd_test"
	"github.com/timbeurskens/gobdd/operators"
	"github.com/timbeurskens/gobdd/operators/bdd"
	"log"
	"math"
	"testing"
)

func TestResize(t *testing.T) {
	v1 := Variable(10)

	v2 := Resize(v1, 4)
	if len(v2) != 4 {
		t.Fatal("Expected length to be 4")
	}

	v3 := Resize(v2, 13)
	if len(v3) != 13 {
		t.Fatal("Expected length to be 13")
	}
}

func makePrimeTest(prime uint) (operators.Expression, Number, Number) {
	// estimate the number of bits needed
	bits := int(math.Ceil(math.Log2(float64(prime))))
	bits = 2 + (bits-(bits/2))*2

	log.Printf("Using %d bits for prime computation", bits)

	// construct two arbitrary numbers: a and b
	a, b := Variable(bits/2), Variable(bits/2)
	// construct a number c
	c := Constant(prime, bits)

	one := Constant(1, bits/2)

	exprEq := operators.And(
		// a != 1
		operators.Not(Equals(a, one)),
		// b != 1
		operators.Not(Equals(b, one)),
	)

	// c == a * b
	exprMul := Mult(a, b, c)

	// a * b == c and c == number to test
	return operators.And(exprEq, exprMul), a, b
}

func TestDivision(t *testing.T) {
	bench := bdd_test.Bench{T: t}

	var src, div uint = 56, 4

	a, b, c := Variable(4), Constant(div, 4), Constant(src, 8)

	// c div b = a; a * b = c
	expr := Mult(a, b, c)

	// prepare the expression tree
	expr = algorithm.PruneUnary(expr)

	t.Log("Size before: ", operators.Size(expr))

	// run bdd algorithm
	tree := algorithm.FromExpression(expr)

	t.Log("Size after:", operators.Size(tree))

	bench.AssertSat("square root of 100 exists", tree)

	if bdd.Sat(tree) {
		if model, ok := bdd.FindModel(tree); ok {
			aResolv, err := a.Resolve(model)
			if err != nil {
				t.Error(err)
			}
			t.Log("56 / 4 =", aResolv)

			if aResolv != src/div {
				t.Error("division did not produce the right result")
			}
		}
	}
}

func TestSqrt(t *testing.T) {
	bench := bdd_test.Bench{T: t}

	var src uint = 100

	a, b := Constant(src, 8), Variable(4)

	// b = sqrt(a); b² = a
	expr := Mult(b, b, a)

	// prepare the expression tree
	expr = algorithm.PruneUnary(expr)

	t.Log("Size before: ", operators.Size(expr))

	// run bdd algorithm
	tree := algorithm.FromExpression(expr)

	t.Log("Size after:", operators.Size(tree))

	bench.AssertSat("square root of 100 exists", tree)

	if bdd.Sat(tree) {
		if model, ok := bdd.FindModel(tree); ok {
			bResolv, err := b.Resolve(model)
			if err != nil {
				t.Error(err)
			}
			t.Log("sqrt(100) =", bResolv)

			if bResolv*bResolv != src {
				t.Error("sqrt did not produce the right result")
			}
		}
	}
}

func TestPrimeDecomposition(t *testing.T) {
	bench := bdd_test.Bench{T: t}

	// prime1 and prime2 are invisible to the solver
	var prime1, prime2 uint = 13, 7

	// feed the composite number
	combined := prime1 * prime2

	// estimate the number of bits needed
	bits := int(math.Ceil(math.Log2(float64(combined))))
	bits = 2 + (bits-(bits/2))*2

	log.Printf("Using %d bits for prime computation", bits)

	// construct two arbitrary numbers: a and b
	a, b := Variable(bits/2), Variable(bits/2)
	// construct a number c
	c := Constant(combined, bits)

	one := Constant(1, bits/2)

	exprEq := operators.And(
		// a != 1
		operators.Not(Equals(a, one)),
		// b != 1
		operators.Not(Equals(b, one)),
	)

	// c == a * b
	exprMul := Mult(a, b, c)

	// a * b == c and c == number to test
	expr := operators.And(exprEq, exprMul)

	// prepare the expression tree
	expr = algorithm.PruneUnary(expr)

	// run bdd algorithm
	tree := algorithm.FromExpression(expr)

	bench.AssertSat(fmt.Sprintf("%d is a composed number", combined), tree)

	if bdd.Sat(tree) {
		if model, ok := bdd.FindModel(tree); ok {
			aResolv, err := a.Resolve(model)
			if err != nil {
				t.Error(err)
			}
			bResolv, err := b.Resolve(model)
			if err != nil {
				t.Error(err)
			}
			t.Logf("Prime-decomposition of %d: %d x %d", combined, aResolv, bResolv)

			if aResolv*bResolv != combined {
				t.Error("a and b are not a valid decomposition of the original number")
			}
		} else {
			t.Fatal("Could not construct model of non-prime number")
		}
	}
}

func TestIsPrime(t *testing.T) {
	bench := bdd_test.Bench{T: t}

	var prime uint = 17

	expr, a, b := makePrimeTest(prime)

	// prepare the expression tree
	expr = algorithm.PruneUnary(expr)

	t.Log("Size before: ", operators.Size(expr))

	// run bdd algorithm
	tree := algorithm.FromExpression(expr)

	t.Log("Size after:", operators.Size(tree))

	bench.AssertUnsat(fmt.Sprintf("%d is prime", prime), tree)

	if bdd.Sat(tree) {
		if model, ok := bdd.FindModel(tree); ok {
			aResolv, err := a.Resolve(model)
			if err != nil {
				t.Error(err)
			}
			bResolv, err := b.Resolve(model)
			if err != nil {
				t.Error(err)
			}
			t.Fatalf("Decomposition of %d: %d x %d", prime, aResolv, bResolv)
		} else {
			t.Fatal("Could not construct model of non-prime number")
		}
	}
}
