package gobdd

import (
    "fmt"
    "reflect"
    "unsafe"

    "gobdd/operators"
)

// todo: cleanup

func PrintSubtree(n operators.Node) {
    if n == nil {
        return
    }

    fmt.Print("(")
    PrintSubtree(n.LeftChild())
    switch n.(type) {
    case *operators.Choice:
        fmt.Print("<-", n.String(), "->")
    case operators.Constant:
        fmt.Print(n.(operators.Constant).Value())
    }
    PrintSubtree(n.RightChild())
    fmt.Print(")")
}

func DotExpressionTree(n operators.Expression)  {
    fmt.Println("digraph G {")
    dotExpressionTreeRec(n)
    fmt.Println("}")
}

func dotExpressionTreeRec(n operators.Expression) string {
    switch n.(type) {
    case operators.Operator:
        vname := fmt.Sprintf("%d", reflect.ValueOf(n).Pointer())
        vlabel := n.String()

        fmt.Printf("%s [label=\"%s\"]", vname, vlabel)
        fmt.Println()

        vtrue := dotExpressionTreeRec(n.LeftChild())
        vfalse := dotExpressionTreeRec(n.RightChild())
        fmt.Println(vname, "->", vtrue)
        fmt.Println(vname, "->", vfalse, "[style=dotted]")

        return vname

    case operators.Constant, operators.Variable:
        vname := fmt.Sprintf("%d", reflect.ValueOf(n).Pointer())
        vlabel := n.String()

        fmt.Printf("%s [label=\"%s\"]", vname, vlabel)
        fmt.Println()

        return vname
    }
    return ""
}

func DotSubtree(n operators.Node) {
    fmt.Println("digraph G {")
    fmt.Println(operators.Cons(true).String())
    fmt.Println(operators.Cons(false).String())
    dotSubtreeRec(n)
    fmt.Println("}")
}

func dotSubtreeRec(n operators.Node) string {
    switch n.(type) {
    case *operators.Choice:
        vname := fmt.Sprintf("%d", unsafe.Pointer(n.(*operators.Choice)))
        vlabel := n.String()

        fmt.Printf("%s [label=\"%s\"]", vname, vlabel)
        fmt.Println()

        vtrue := dotSubtreeRec(n.LeftChild())
        vfalse := dotSubtreeRec(n.RightChild())
        fmt.Println(vname, "->", vtrue)
        fmt.Println(vname, "->", vfalse, "[style=dotted]")

        return vname
    case operators.Constant:
        return n.String()
    }
    return ""
}
