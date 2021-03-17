package operators

// Not is the negation of e
//func Not(e Expression) Expression {
//    if term, ok := e.(Term); ok {
//        return term.Negate()
//    }
//    return Implies(e, &FalseConst)
//}

func Not(e Expression) Expression {
	if term, ok := e.(Term); ok {
		return term.Negate()
	}
	return &Negation{e}
}

// Implies returns a -> b
func Implies(a, b Expression) Expression {
	return &Implication{a, b}
}

func Biimplies(a, b Expression) Expression {
	return &Biimplication{a, b}
}

func And(expr ...Expression) Expression {
	if len(expr) > 2 {
		return And(expr[0], And(expr[1:]...))
	}
	return &Conjunction{expr[0], expr[1]}
}

func Or(expr ...Expression) Expression {
	if len(expr) > 2 {
		return Or(expr[0], Or(expr[1:]...))
	}
	return &Disjunction{expr[0], expr[1]}
}

func Xor(expr ...Expression) Expression {
	if len(expr) > 2 {
		return Xor(expr[0], Xor(expr[1:]...))
	}
	return &ExclusiveDisjunction{expr[0], expr[1]}
}

func Var(name string) Variable {
	v := StringVariable(name)
	return &v
}

func IVar(i int) Variable {
	v := IntVariable(i)
	return &v
}

func Cons(b bool) Constant {
	if b {
		return &TrueConst
	}
	return &FalseConst
}
