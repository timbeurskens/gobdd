package operators

type BoolConst bool

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
    TrueConst BoolConst = true
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