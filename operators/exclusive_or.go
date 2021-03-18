package operators

type ExclusiveDisjunction struct {
	A Expression
	B Expression
}

func (e *ExclusiveDisjunction) SetLeftChild(n Node) {
	e.A = n
}

func (e *ExclusiveDisjunction) SetRightChild(n Node) {
	e.B = n
}

// todo: enable Normalize function
func (e *ExclusiveDisjunction) Normalize() Expression {
	left, right := e.LeftChild().Normalize(), e.RightChild().Normalize()
	return And(
		Or(left, right),
		Not(And(left, right)),
	)
}

func (e *ExclusiveDisjunction) String() string {
	return "⊗"
}

func (e *ExclusiveDisjunction) NodeEquivalent(n Node) bool {
	_, ok := n.(*ExclusiveDisjunction)
	return ok
}

func (e *ExclusiveDisjunction) LeftChild() Node {
	return e.A
}

func (e *ExclusiveDisjunction) RightChild() Node {
	return e.B
}

func (e *ExclusiveDisjunction) ConstEval(a, b Constant) Constant {
	return Cons(a.Value() != b.Value())
}

func (e *ExclusiveDisjunction) Join(a, b Expression) Operator {
	return &ExclusiveDisjunction{a, b}
}
