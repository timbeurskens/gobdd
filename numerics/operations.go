package numerics

import "github.com/timbeurskens/gobdd/operators"

// Add constructs an expression such that a + b == c.
// The resolution of a should equal the resolution of b and be 1 less than the resolution of c.
// If needed, a number can be padded with trailing zeroes to match resolutions.
// todo: loosen the res(c) == 1 + res(a) == 1 + res(b) restriction:: res(c) == 1 + max(res(a), res(b))
func Add(a, b, c Number) operators.Expression {
	panic("Not implemented")
	return nil
}

// Mult constructs an expression such that a * b == c.
// The resolution of a and b should match, and the resolution of c should equal the sum of the resolutions of a and b.
// todo: loosen the res(a) == res(b) restriction
func Mult(a, b, c Number) operators.Expression {
	panic("Not implemented")
	return nil
}

// Equals constructs an expression such that a == b.
// The resolution of a and b should be equal.
func Equals(a, b Number) operators.Expression {
	if len(a) != len(b) {
		panic("length of a and b should match")
	}

	exprs := make([]operators.Expression, len(a))
	for i := range exprs {
		exprs[i] = operators.Biimplies(a[i], b[i])
	}

	return operators.And(exprs...)
}
