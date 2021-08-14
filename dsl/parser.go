package dsl

import (
	"fmt"
	"github.com/timbeurskens/gobdd/operators"
	"github.com/timbeurskens/goparselib"
	"github.com/timbeurskens/goparselib/parser"
	"os"
	"reflect"
)

func SingleChildReduce(root goparselib.Node) (goparselib.Node, error) {
	if root.Children == nil || len(root.Children) == 0 {
		return root, nil
	}

	if len(root.Children) == 1 {
		if result, err := SingleChildReduce(root.Children[0]); err != nil {
			return goparselib.Node{}, err
		} else {
			return result, nil
		}
	} else {
		result := make([]goparselib.Node, len(root.Children))
		for i, child := range root.Children {
			var err error
			if result[i], err = SingleChildReduce(child); err != nil {
				return goparselib.Node{}, err
			}
		}
		return goparselib.Node{
			Start:    root.Start,
			Size:     root.Size,
			Contents: root.Contents,
			Type:     root.Type,
			Children: result,
		}, nil
	}
}

func ZeroSizeReduce(root goparselib.Node) (goparselib.Node, error) {
	if root.Children == nil || len(root.Children) == 0 {
		return root, nil
	}

	result := make([]goparselib.Node, 0, len(root.Children))
	for _, child := range root.Children {
		if child.Size > 0 {
			if cred, err := ZeroSizeReduce(child); err == nil {
				result = append(result, cred)
			} else {
				return goparselib.Node{}, err
			}
		}
	}

	return goparselib.Node{
		Start:    root.Start,
		Size:     root.Size,
		Contents: root.Contents,
		Type:     root.Type,
		Children: result,
	}, nil
}

func binaryToExpression(root goparselib.Node) (operators.Expression, error) {
	var left, right operators.Expression

	return operators.And(left, right), nil
}

func unaryToExpression(root goparselib.Node) (operators.Expression, error) {
	return nil, fmt.Errorf("not implemented")
}

func identToExpression(node goparselib.Node) (operators.Expression, error) {
	return nil, fmt.Errorf("not implemented")
}

func constToExpression(node goparselib.Node) (operators.Expression, error) {
	return nil, fmt.Errorf("not implemented")
}

func unsafeToExpression(root goparselib.Node) (operators.Expression, error) {
	for _, child := range root.Children {
		if reflect.DeepEqual(child.Type, BinaryExpr) {
			return binaryToExpression(child)
		} else if reflect.DeepEqual(child.Type, UnaryExpr) {
			return unaryToExpression(child)
		} else if reflect.DeepEqual(child.Type, goparselib.Ident) {
			return identToExpression(child)
		} else if reflect.DeepEqual(child.Type, Const) {
			return constToExpression(child)
		}
	}

	return nil, fmt.Errorf("no valid sub-expression found")
}

func ToExpression(input string) (operators.Expression, error) {
	language := Load()

	if sst, err := parser.ParseString(input, language.Root()); err != nil {
		return nil, err
	} else {
		if ast, err := sst.Reduce(language.Layout()...); err != nil {
			return nil, err
		} else if rast, err := ZeroSizeReduce(ast); err != nil {
			return nil, err
		} else if sast, err := SingleChildReduce(rast); err != nil {
			return nil, err
		} else if expr, err := unsafeToExpression(sast); err != nil {
			sast.Output(os.Stderr)
			return nil, err
		} else {
			return expr, nil
		}
	}
}
