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
