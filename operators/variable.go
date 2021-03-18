package operators

import "fmt"

type StringVariable string

func (s *StringVariable) Variable() Variable {
	return s
}

func (s *StringVariable) SetLeftChild(n Node) {
	if n != nil {
		panic("string variable has no left child")
	}
}

func (s *StringVariable) SetRightChild(n Node) {
	if n != nil {
		panic("string variable has no right child")
	}
}

func (s *StringVariable) Normalize() Expression {
	return s
}

func (s *StringVariable) Terms() []Term {
	return []Term{s}
}

func (s *StringVariable) HasTerm(term Term) bool {
	return s.TermEquivalent(term)
}

func (s *StringVariable) Exclude(term Term) CNFClause {
	if s.HasTerm(term) {
		return nil
	} else {
		return s
	}
}

func (s *StringVariable) NumTerms() int {
	return 1
}

func (s *StringVariable) Negate() Term {
	return &Negation{s}
}

func (s *StringVariable) TermEquivalent(t Term) bool {
	other, ok := t.(*StringVariable)
	return ok && s.NodeEquivalent(other)
}

func (s *StringVariable) Less(variable Variable) bool {
	return s.String() <= variable.String()
}

func (s *StringVariable) NodeEquivalent(n Node) bool {
	other, ok := n.(*StringVariable)
	return ok && other.String() == s.String()
}

func (s *StringVariable) LeftChild() Node {
	return nil
}

func (s *StringVariable) RightChild() Node {
	return nil
}

func (s *StringVariable) String() string {
	return string(*s)
}

type IntVariable int

func (c *IntVariable) Variable() Variable {
	return c
}

func (c *IntVariable) SetLeftChild(n Node) {
	if n != nil {
		panic("int variable has no left child")
	}
}

func (c *IntVariable) SetRightChild(n Node) {
	if n != nil {
		panic("int variable has no right child")
	}
}

func (c *IntVariable) Normalize() Expression {
	return c
}

func (c *IntVariable) Terms() []Term {
	return []Term{c}
}

func (c *IntVariable) HasTerm(term Term) bool {
	return c.TermEquivalent(term)
}

func (c *IntVariable) Exclude(term Term) CNFClause {
	if c.HasTerm(term) {
		return nil
	} else {
		return c
	}
}

func (c *IntVariable) NumTerms() int {
	return 1
}

func (c *IntVariable) Negate() Term {
	return &Negation{c}
}

func (c *IntVariable) TermEquivalent(t Term) bool {
	other, ok := t.(*IntVariable)
	return ok && c.NodeEquivalent(other)
}

func (c *IntVariable) NodeEquivalent(n Node) bool {
	other, ok := n.(*IntVariable)
	return ok && *other == *c
}

func (c *IntVariable) LeftChild() Node {
	return nil
}

func (c *IntVariable) RightChild() Node {
	return nil
}

func (c *IntVariable) String() string {
	return fmt.Sprintf("%d", *c)
}

func (c *IntVariable) Less(variable Variable) bool {
	other, ok := variable.(*IntVariable)
	return ok && *c <= *other
}
