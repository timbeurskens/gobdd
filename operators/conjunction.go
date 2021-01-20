package operators

type Conjunction struct {
    A Expression
    B Expression
}

func (c *Conjunction) Normalize() Expression {
    return c
}

func (c *Conjunction) String() string {
    return "∧"
}

func (c *Conjunction) Join(a, b Expression) Operator {
    return &Conjunction{a, b}
}

func (c *Conjunction) ConstEval(a, b Constant) Constant {
    return Cons(a.Value() && b.Value())
}

func (c *Conjunction) NodeEquivalent(n Node) bool {
    _, ok := n.(*Conjunction)
    return ok
}

func (c *Conjunction) LeftChild() Node {
    return c.A
}

func (c *Conjunction) RightChild() Node {
    return c.B
}