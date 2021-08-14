package dsl

import . "github.com/timbeurskens/goparselib"

var (
	SymOr  = Union{CTerminal("or"), CTerminal("∨")}
	SymAnd = Union{CTerminal("and"), CTerminal("∧")}
	SymNot = Union{CTerminal("not"), CTerminal("¬")}
	Stop   = CTerminal("\\.")

	ConstTrue  = Union{CTerminal("true"), CTerminal("1"), CTerminal("⊤")}
	ConstFalse = Union{CTerminal("false"), CTerminal("0"), CTerminal("⊥")}

	Const = Union{ConstTrue, ConstFalse}

	UnsafeExpr = new(Symbol)

	Layout = Union{Blank, LParen, RParen, Stop, nil}

	Expr = Concat{LParen, R(UnsafeExpr), RParen}

	BinOp = Union{SymOr, SymAnd}
	UnOp  = Union{SymNot}

	ExprBody = Union{Expr, Const, Ident}

	UnaryExpr  = Decorate(Concat{UnOp, ExprBody}, Layout)
	BinaryExpr = Decorate(Concat{ExprBody, BinOp, ExprBody}, Layout)

	Root = Concat{R(UnsafeExpr), Stop}
)

func init() {
	Define(UnsafeExpr, Union{UnaryExpr, BinaryExpr, Const, Ident})
}

func Load() Plugin {
	return MakePlugin(Root, Layout)
}
