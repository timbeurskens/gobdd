package operators

type BoolConst bool

func (b *BoolConst) NumTerms() int {
	return 1
}

func (b *BoolConst) Terms() []Term {
	return []Term{b}
}

func (b *BoolConst) HasTerm(term Term) bool {
	return b.NodeEquivalent(term)
}

func (b *BoolConst) Exclude(term Term) CNFClause {
	panic("implement me")
}

func (b *BoolConst) Negate() Term {
	return Cons(!bool(*b))
}

func (b *BoolConst) TermEquivalent(t Term) bool {
	return b.NodeEquivalent(t)
}

func (b *BoolConst) Variable() Variable {
	panic("a const is not a variable")
	return nil
}

func (b *BoolConst) SetLeftChild(n Node) {
	if n != nil {
		panic("const has no left child")
	}
}

func (b *BoolConst) SetRightChild(n Node) {
	if n != nil {
		panic("const has no right child")
	}
}

func (b *BoolConst) Normalize() Expression {
	return b
}

func (b *BoolConst) String() string {
	if b.Value() {
		return "true"
	} else {
		return "false"
	}
}

var (
	TrueConst  BoolConst = true
	FalseConst BoolConst = false
)

func (b *BoolConst) NodeEquivalent(n Node) bool {
	other, ok := n.(*BoolConst)
	return ok && other.Value() == b.Value()
}

func (b *BoolConst) LeftChild() Node {
	return nil
}

func (b *BoolConst) RightChild() Node {
	return nil
}

func (b *BoolConst) Value() bool {
	return bool(*b)
}
