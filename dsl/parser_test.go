package dsl

import (
	"github.com/timbeurskens/goparselib/parser"
	"testing"
)

func TestToExpression(t *testing.T) {
	t.Log(ToExpression("(a or b) and (c)"))
}

func TestLayoutReductionOnRoot(t *testing.T) {
	exprStr := "(a or b) and (c)."
	language := Load()

	if sst, err := parser.ParseString(exprStr, language.Root()); err != nil {
		t.Error(err)
	} else {
		if len(sst.Children) != 2 {
			t.Error("Expected 2 children before reduction")
		} else if sst.Children[1].Contents != "." {
			t.Errorf("Expected the contents of child[1] to be \".\". Instead got \"%s\" (char0: %d; at: %d)", sst.Children[1].Contents, sst.Children[1].Contents[0], sst.Children[1].Start)
		}

		if ast, err := sst.Reduce(language.Layout()...); err != nil {
			t.Error(err)
		} else {
			if len(ast.Children) != 1 {
				t.Error("Expected 1 child after reduction")
			}
		}
	}
}

func TestLayoutReductionOnFirstChild(t *testing.T) {
	//exprStr := "(a or b) and (c)."
	//language := Load()
	//
	//if sst, err := parser.ParseString(exprStr, language.Root()); err != nil {
	//	t.Error(err)
	//} else {
	//	if len(sst.Children) != 2 {
	//		t.Error("Expected 2 children before reduction")
	//	}
	//
	//	if ast, err := sst.Reduce(language.Layout()...); err != nil {
	//		t.Error(err)
	//	} else {
	//		if len(ast.Children) != 1 {
	//			t.Error("Expected 1 child after reduction")
	//		}
	//	}
	//}
}
