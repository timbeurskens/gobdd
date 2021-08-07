package operators

import (
	"fmt"
)

type Node interface {
	NodeEquivalent(n Node) bool
	LeftChild() Node
	RightChild() Node
	SetLeftChild(n Node)
	SetRightChild(n Node)
	Normalize() Expression
	fmt.Stringer
}

type Term interface {
	CNFClause
	Negate() Term
	TermEquivalent(t Term) bool
	Variable() Variable
	Expression
	fmt.Stringer
}

type Expression interface {
	Node
}

type Operator interface {
	Expression
	ConstEval(a, b Constant) Constant
	Join(a, b Expression) Operator
}

type Variable interface {
	Term

	// Leq returns true iff this is less or equal than variable
	// todo: variable reordering
	Leq(variable Variable) bool
}

type Constant interface {
	Term
	Value() bool
}

// Model is a type representing a trace in the bdd
type Model map[Variable]bool

type CNF []CNFClause

func (cnf CNF) Expr() Expression {
	// todo: optimize
	intermediate := make([]Expression, 0, len(cnf))
	for _, clause := range cnf {
		terms := clause.Terms()

		exprs := make([]Expression, len(terms))

		for i := range terms {
			exprs[i] = terms[i]
		}

		intermediate = append(intermediate, Or(exprs...))
	}
	return And(intermediate...)
}

type CNFClause interface {
	fmt.Stringer
	NumTerms() int
	Terms() []Term
	HasTerm(term Term) bool
	Exclude(Term) CNFClause
}

func IsTerm(clause CNFClause) bool {
	_, ok := clause.(Term)
	return ok
}
