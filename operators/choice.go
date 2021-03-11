package operators

type Choice struct {
	True  Node
	Var   Variable
	False Node
}

func (c *Choice) SetLeftChild(n Node) {
	c.True = n
}

func (c *Choice) SetRightChild(n Node) {
	c.False = n
}

func (c *Choice) String() string {
	return c.Var.String()
}

func (c *Choice) NodeEquivalent(n Node) bool {
	other, ok := n.(*Choice)
	return ok && c.Var.NodeEquivalent(other.Var)
}

func (c *Choice) LeftChild() Node {
	return c.True
}

func (c *Choice) RightChild() Node {
	return c.False
}

func NewChoice(v Variable) Node {
	return JoinByChoice(v, Cons(true), Cons(false))
}

func JoinByChoice(v Variable, trueTree, falseTree Node) Node {
	return &Choice{
		True:  trueTree,
		Var:   v,
		False: falseTree,
	}
}

func (c *Choice) SwapTrue(n Node) (old Node) {
	old, c.True = c.True, n
	return
}

func (c *Choice) SwapFalse(n Node) (old Node) {
	old, c.False = c.False, n
	return
}
