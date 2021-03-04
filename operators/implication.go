package operators

type Implication struct {
	A Expression
	B Expression
}

// todo: enable normalize function
// func (i *Implication) Normalize() Expression {
//     lnorm := i.LeftChild().(Expression).Normalize()
//
//     neg := Not().Normalize()
//     return Or(neg, i.RightChild().(Expression).Normalize())
// }

func (i *Implication) String() string {
	return "→"
}

func (i *Implication) Join(a, b Expression) Operator {
	return &Implication{a, b}
}

func (i *Implication) ConstEval(a, b Constant) Constant {
	if a.Value() {
		return Cons(b.Value())
	}
	return Cons(true)
}

func (i *Implication) NodeEquivalent(n Node) bool {
	_, ok := n.(*Implication)
	return ok
}

func (i *Implication) LeftChild() Node {
	return i.A
}

func (i *Implication) RightChild() Node {
	return i.B
}
