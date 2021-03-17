package operators

type Negation struct {
	T Expression
}

func (n *Negation) Variable() Variable {
	return n.Negate().Variable()
}

func (n *Negation) SetLeftChild(node Node) {
	n.T = node
}

func (n *Negation) SetRightChild(node Node) {
	if n != nil {
		panic("negation has no right child")
	}
}

// Normalize of a negation is still a negation
func (n *Negation) Normalize() Expression {
	return &Negation{
		T: n.LeftChild().Normalize(),
	}
}

func (n *Negation) NodeEquivalent(o Node) bool {
	other, ok := o.(*Negation)
	return ok && other.T.NodeEquivalent(n.T)
}

func (n *Negation) LeftChild() Node {
	return n.T
}

func (n *Negation) RightChild() Node {
	return nil
}

func (n *Negation) Terms() []Term {
	return []Term{n}
}

func (n *Negation) HasTerm(term Term) bool {
	return n.TermEquivalent(term)
}

func (n *Negation) Exclude(term Term) CNFClause {
	if n.HasTerm(term) {
		return nil
	} else {
		return n
	}
}

func (n *Negation) String() string {
	return "¬" + n.T.String()
}

func (n *Negation) NumTerms() int {
	return 1
}

func (n *Negation) Negate() Term {
	return n.T.(Term)
}

func (n *Negation) TermEquivalent(t Term) bool {
	other, ok := t.(*Negation)
	return ok && n.NodeEquivalent(other)
}

func ToNormalTerm(t Term) Term {
	// either a term or a negation
	if _, ok := t.(*Negation); ok {
		return t.Negate()
	}
	return t
}
