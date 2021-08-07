package operators

const IncVarMask int = 573829

var (
	varCount = 0
)

func IncVar() Term {
	result := IVar(varCount ^ IncVarMask)
	varCount += 1
	return result
}

func IncVarCollection(n int) []Term {
	res := make([]Term, n)
	for i := range res {
		res[i] = IncVar()
	}
	return res
}

func ConsCollection(n int, b bool) []Term {
	consts := make([]Term, n)
	for i := range consts {
		consts[i] = Cons(b)
	}
	return consts
}
