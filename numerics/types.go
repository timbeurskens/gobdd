package numerics

import (
	"fmt"
	"github.com/timbeurskens/gobdd/operators"
)

type Number []operators.Term

func (n Number) Resolve(model operators.Model) (uint, error) {
	var result uint = 0
	for i := range n {
		if v, ok := n[i].(operators.Variable); !ok {
			return 0, fmt.Errorf("could not cast %+v to type Variable", n[i])
		} else if s, ok := model[v]; !ok {
			return 0, fmt.Errorf("variable %+v could not be found in the model", v)
		} else if s {
			result |= 1 << i
		}
	}
	return result, nil
}

// Constant creates a Number with resolution n, equal to the value v
func Constant(v uint, n int) Number {
	repr := make(Number, n)

	for i := range repr {
		repr[i] = operators.Cons(((v >> i) & 1) == 1)
	}

	return repr
}

// Variable creates a Number with resolution n
func Variable(n int) Number {
	return operators.IncVarCollection(n)
}

// NamedVariable creates a Number with resolution n, with a given prefix
func NamedVariable(prefix string, n int) Number {
	res := make(Number, n)
	for i := range res {
		res[i] = operators.Var(fmt.Sprintf("%s_%d", prefix, i))
	}
	return res
}

// Resize either pads, or slices the number a, such that its resulting size is n
func Resize(a Number, n int) Number {
	if delta := n - len(a); delta > 0 {
		return Pad(a, delta)
	} else {
		return a[:n]
	}
}

// Pad returns the number a, padded by n trailing zeroes
func Pad(a Number, n int) Number {
	res := make([]operators.Term, len(a)+n)
	for i := range res {
		if i < len(a) {
			res[i] = a[i]
		} else {
			res[i] = operators.Cons(false)
		}
	}
	return res
}

// Shift returns the number a, left-shifted by n bits, where n can be less than 0
func Shift(a Number, n int) Number {
	// n leq 0: take a sub-slice from the original number
	if n <= 0 {
		return a[-n:]
	}

	// append n leading zeroes
	res := make([]operators.Term, len(a)+n)
	for i := range res {
		if i < n {
			res[i] = operators.Cons(false)
		} else {
			res[i] = a[i-n]
		}
	}

	return res
}
