package operators

const IncVarMask int = 573829

var (
	varCount = 0
)

func IncVar() Variable {
	result := IVar(varCount ^ IncVarMask)
	varCount += 1
	return result
}

func IncVarCollection(n int) []Variable {
	res := make([]Variable, n)
	for i := range res {
		res[i] = IncVar()
	}
	return res
}

func ConsCollection(n int, b bool) []Variable {
	consts := make([]Variable, n)
	for i := range consts {
		consts[i] = Cons(b)
	}
	return consts
}
