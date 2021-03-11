package operators

type Biimplication struct {
	A Expression
	B Expression
}

func (bi *Biimplication) SetLeftChild(n Node) {
	bi.A = n
}

func (bi *Biimplication) SetRightChild(n Node) {
	bi.B = n
}

// todo: enable normalize function
// func (bi *Biimplication) Normalize() Expression {
//     return And(Implies(bi.LeftChild(), bi.RightChild()).(Operator).Normalize(), Implies(bi.RightChild(), bi.LeftChild()).(Operator).Normalize())
// }

func (bi *Biimplication) String() string {
	return "⟷"
}

func (bi *Biimplication) NodeEquivalent(n Node) bool {
	_, ok := n.(*Biimplication)
	return ok
}

func (bi *Biimplication) LeftChild() Node {
	return bi.A
}

func (bi *Biimplication) RightChild() Node {
	return bi.B
}

func (bi *Biimplication) ConstEval(a, b Constant) Constant {
	return Cons(a.Value() == b.Value())
}

func (bi *Biimplication) Join(a, b Expression) Operator {
	return &Biimplication{a, b}
}
