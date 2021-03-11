package operators

import "fmt"

type Node interface {
	NodeEquivalent(n Node) bool
	LeftChild() Node
	RightChild() Node
	SetLeftChild(n Node)
	SetRightChild(n Node)
	fmt.Stringer
}

type Term interface {
	CNFClause
	Negate() Term
	TermEquivalent(t Term) bool
	Expression
	fmt.Stringer
}

type Expression interface {
	Node
	// todo: enable Normalize function for tseitin transformations
	// Normalize() Expression
}

type Operator interface {
	Expression
	ConstEval(a, b Constant) Constant
	Join(a, b Expression) Operator
}

type Variable interface {
	Term

	// todo: variable reordering
	Less(variable Variable) bool
}

type Constant interface {
	Expression
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
	NumTerms() int
	Terms() []Term
	HasTerm(term Term) bool
	Exclude(Term) CNFClause
}

func IsTerm(clause CNFClause) bool {
	_, ok := clause.(Term)
	return ok
}
