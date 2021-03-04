package operators

type ExclusiveDisjunction struct {
	A Expression
	B Expression
}

// todo: enable Normalize function
// func (e *ExclusiveDisjunction) Normalize() Expression {
//     return And(Or(e.LeftChild(), e.RightChild()), Not(And(e.LeftChild(), e.RightChild())).(Operator).Normalize())
// }

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
