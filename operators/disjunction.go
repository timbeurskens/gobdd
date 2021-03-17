package operators

type Disjunction struct {
	A Expression
	B Expression
}

func (d *Disjunction) SetLeftChild(n Node) {
	d.A = n
}

func (d *Disjunction) SetRightChild(n Node) {
	d.B = n
}

func (d *Disjunction) Normalize() Expression {
	d.SetLeftChild(d.LeftChild().Normalize())
	d.SetRightChild(d.RightChild().Normalize())
	return d
}

func (d *Disjunction) String() string {
	return "∨"
}

func (d *Disjunction) Join(a, b Expression) Operator {
	return &Disjunction{a, b}
}

func (d *Disjunction) ConstEval(a, b Constant) Constant {
	return Cons(a.Value() || b.Value())
}

func (d *Disjunction) NodeEquivalent(n Node) bool {
	_, ok := n.(*Disjunction)
	return ok
}

func (d *Disjunction) LeftChild() Node {
	return d.A
}

func (d *Disjunction) RightChild() Node {
	return d.B
}
