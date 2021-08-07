package numerics

import "github.com/timbeurskens/gobdd/operators"

// Add constructs an expression such that a + b == c.
// The resolution of a should equal the resolution of b and be 1 less than the resolution of c.
// If needed, a number can be padded with trailing zeroes to match resolutions.
// todo: loosen the res(c) == 1 + res(a) == 1 + res(b) restriction:: res(c) == 1 + max(res(a), res(b))
func Add(a, b, c Number) operators.Expression {
	if len(a) != len(b) || len(a)+1 != len(c) {
		panic("length of a and b and carry should match and be 1 smaller than length of c")
	}

	carry := Variable(len(a))

	exprs := make([]operators.Expression, len(c))
	for i := range c {
		if i == 0 {
			exprs[i] = operators.And(
				operators.Biimplies(c[i], operators.Xor(a[i], b[i])),
				operators.Biimplies(carry[i], operators.And(a[i], b[i])),
			)
		} else if i > 0 && i < len(a) {
			exprs[i] = operators.And(
				operators.Biimplies(c[i], operators.Xor(a[i], b[i], carry[i-1])),
				operators.Biimplies(
					carry[i],
					operators.Or(
						operators.And(a[i], b[i]),
						operators.And(a[i], carry[i-1]),
						operators.And(b[i], carry[i-1]),
					),
				),
			)
		} else { // i == len(a)
			exprs[i] = operators.And(
				operators.Biimplies(c[i], carry[i-1]),
			)
		}
	}

	return operators.And(exprs...)
}

// Mult constructs an expression such that a * b == c.
// The resolution of a and b should match, and the resolution of c should equal the sum of the resolutions of a and b.
// todo: loosen the res(a) == res(b) restriction
func Mult(a, b, c Number) operators.Expression {
	if len(a) != len(b) || len(c) != len(a)+len(b) {
		panic("length of a and b should match and length of c should be a + b")
	}

	// if a[i] == 1 then inter[i] = inter[i-1] + b[i] << i
	// if a[i] == 0 then inter[i] = inter[i-1]
	// inter[0] == 0

	inter := make([]Number, len(a)+1)
	exprs := make([]operators.Expression, len(a))

	for i := range inter {
		if i == 0 {
			// inter[0] == 0
			inter[i] = Constant(0, len(b))
		} else {
			// inter[i; i > 0] == var...
			if i == len(a) {
				inter[i] = c
			} else {
				inter[i] = Variable(len(b) + i)
			}

			// add b << i on the i'th step
			bAdded := Add(inter[i-1], Shift(b, i-1), inter[i])

			// add 0 on the i'th step
			bEq := Equals(inter[i], Pad(inter[i-1], 1))

			exprs[i-1] = operators.And(
				// if the i'th factor of a is 1, then add b << i
				operators.Implies(a[i-1], bAdded),
				// if the i'th factor of a is 0, then add 0
				operators.Implies(operators.Not(a[i-1]), bEq),
			)
		}
	}

	// for every 1 in a, shift b by the position of the 1 in a, then add it to the result
	return operators.And(exprs...)
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
