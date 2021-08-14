package dsl

import (
	"github.com/timbeurskens/goparselib/parser"
	"os"
	"testing"
)

func TestSimpleExpr(t *testing.T) {
	language := Load()

	node, err := parser.ParseString("(a or b) and (not(a and b)).", language.Root())
	if err != nil {
		t.Error(err)
	}
	err = node.Output(os.Stderr)
	if err != nil {
		t.Error(err)
	}
	//node.Walk(func(n goparselib.Node) bool {
	//	t.Log(n)
	//	return true
	//})
}

func TestBinaryVars(t *testing.T) {
	str := "a or b"
	if node, err := parser.ParseString(str, BinaryExpr); err != nil {
		t.Error(err)
	} else if node.Size != int64(len(str)) {
		t.Errorf("Lengths do not match, expected %d, but got %d", len(str), node.Size)
	}
}

func TestBinaryCons(t *testing.T) {
	str := "true or false"
	if node, err := parser.ParseString(str, BinaryExpr); err != nil {
		t.Error(err)
	} else if node.Size != int64(len(str)) {
		t.Errorf("Lengths do not match, expected %d, but got %d", len(str), node.Size)
	}
}

func TestUnaryVar(t *testing.T) {
	str := "not a"
	if node, err := parser.ParseString(str, UnaryExpr); err != nil {
		t.Error(err)
	} else if node.Size != int64(len(str)) {
		t.Errorf("Lengths do not match, expected %d, but got %d", len(str), node.Size)
	}
}

func TestUnaryCons(t *testing.T) {
	str := "not false"
	if node, err := parser.ParseString(str, UnaryExpr); err != nil {
		t.Error(err)
	} else if node.Size != int64(len(str)) {
		t.Errorf("Lengths do not match, expected %d, but got %d", len(str), node.Size)
	}
}

func TestExprBodyIdent(t *testing.T) {
	str := "var"
	if node, err := parser.ParseString(str, ExprBody); err != nil {
		t.Error(err)
	} else if node.Size != int64(len(str)) {
		t.Errorf("Lengths do not match, expected %d, but got %d", len(str), node.Size)
	}
}

func TestExprBodyConst(t *testing.T) {
	str := "0"
	if node, err := parser.ParseString(str, ExprBody); err != nil {
		t.Error(err)
	} else if node.Size != int64(len(str)) {
		t.Errorf("Lengths do not match, expected %d, but got %d", len(str), node.Size)
	}
}

func TestExprBodyExpr1(t *testing.T) {
	str := "(a or b)"
	if node, err := parser.ParseString(str, ExprBody); err != nil {
		t.Error(err)
	} else if node.Size != int64(len(str)) {
		t.Errorf("Lengths do not match, expected %d, but got %d", len(str), node.Size)
	}
}

func TestExprBodyExpr2(t *testing.T) {
	str := "(1 or 1)"
	if node, err := parser.ParseString(str, ExprBody); err != nil {
		t.Error(err)
	} else if node.Size != int64(len(str)) {
		t.Errorf("Lengths do not match, expected %d, but got %d", len(str), node.Size)
	}
}

func TestExprBodyExpr3(t *testing.T) {
	str := "(not 1)"
	if node, err := parser.ParseString(str, ExprBody); err != nil {
		t.Error(err)
	} else if node.Size != int64(len(str)) {
		t.Errorf("Lengths do not match, expected %d, but got %d", len(str), node.Size)
	}
}

func TestExprBodyExpr4(t *testing.T) {
	str := "( ( not 1 ) )"
	if node, err := parser.ParseString(str, ExprBody); err != nil {
		t.Error(err)
	} else if node.Size != int64(len(str)) {
		t.Errorf("Lengths do not match, expected %d, but got %d", len(str), node.Size)
	}
}

func TestExprBodyExpr5(t *testing.T) {
	str := "( ( not ( 1 ) ) )"
	if node, err := parser.ParseString(str, ExprBody); err != nil {
		t.Error(err)
	} else if node.Size != int64(len(str)) {
		t.Errorf("Lengths do not match, expected %d, but got %d", len(str), node.Size)
	}
}

func TestExprIdent1(t *testing.T) {
	str := "(a)"
	if node, err := parser.ParseString(str, ExprBody); err != nil {
		t.Error(err)
	} else if node.Size != int64(len(str)) {
		t.Errorf("Lengths do not match, expected %d, but got %d", len(str), node.Size)
	}
}

func TestExprIdent2(t *testing.T) {
	str := "( a )"
	if node, err := parser.ParseString(str, ExprBody); err != nil {
		t.Error(err)
	} else if node.Size != int64(len(str)) {
		t.Errorf("Lengths do not match, expected %d, but got %d", len(str), node.Size)
	}
}
