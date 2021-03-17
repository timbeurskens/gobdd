package operators

type Implication struct {
	A Expression
	B Expression
}

func (i *Implication) SetLeftChild(n Node) {
	i.A = n
}

func (i *Implication) SetRightChild(n Node) {
	i.B = n
}

// todo: enable normalize function
func (i *Implication) Normalize() Expression {
	left := Not(i.LeftChild()).Normalize()
	right := i.RightChild().Normalize()
	return Or(left, right)
}

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
