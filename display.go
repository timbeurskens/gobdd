package gobdd

import (
	"fmt"
	"io"
	"os"
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

func PrintExpressiontree(e operators.Expression) string {
	if e == nil {
		return ""
	}
	if _, ok := e.(operators.Variable); ok {
		return e.String()
	}
	if _, ok := e.(operators.Constant); ok {
		return e.String()
	}
	return fmt.Sprintf("(%s %s %s)", PrintExpressiontree(e.LeftChild()), e.String(), PrintExpressiontree(e.RightChild()))
}

func DotExpressionTree(n operators.Expression) {
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

type DeferError struct {
	Err error
}

func (d *DeferError) Do(err error) error {
	if d.Err != nil {
		d.Err = err
	}
	return d.Err
}

func DotSubtreeWriter(n operators.Node, writer io.StringWriter) error {
	err := DeferError{}
	var e error

	_, e = writer.WriteString("digraph G {\n")
	_ = err.Do(e)
	_, e = writer.WriteString(fmt.Sprintf("%s\n", operators.Cons(true).String()))
	_ = err.Do(e)
	_, e = writer.WriteString(fmt.Sprintf("%s\n", operators.Cons(false).String()))
	_ = err.Do(e)

	dotSubtreeRecWriter(n, writer)

	_, e = writer.WriteString("}\n")
	_ = err.Do(e)

	return err.Do(nil)
}

func dotSubtreeRecWriter(n operators.Node, writer io.StringWriter) string {
	err := DeferError{}
	var e error

	switch n.(type) {
	case *operators.Choice:
		vname := fmt.Sprintf("%d", unsafe.Pointer(n.(*operators.Choice)))
		vlabel := n.String()

		_, e = writer.WriteString(fmt.Sprintf("%s [label=\"%s\"]\n", vname, vlabel))
		_ = err.Do(e)

		vtrue := dotSubtreeRecWriter(n.LeftChild(), writer)
		vfalse := dotSubtreeRecWriter(n.RightChild(), writer)

		_, e = writer.WriteString(fmt.Sprintf("%s -> %s\n", vname, vtrue))
		_ = err.Do(e)
		_, e = writer.WriteString(fmt.Sprintf("%s -> %s [style=dotted]\n", vname, vfalse))
		_ = err.Do(e)

		return vname
	case operators.Constant:
		return n.String()
	}
	return ""
}

func DotSubtree(n operators.Node) error {
	return DotSubtreeWriter(n, os.Stdout)
}
