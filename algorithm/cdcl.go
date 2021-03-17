package algorithm

import (
	"log"

	"gobdd/operators"
)

// steps:
// a choice adds a CDCLDecide item on the "stack". Every clause following from this decision will be added in the CDCLDecide node
// any choice made afterwards is added to this node as well
// The backtrack action removes the nearest CDCLDecide node and adds the negation of choice to the parent CDCLNode
// when backtrack is performed on an empty stack, FAIL must follow

type CDCLStack struct {
	Clauses operators.CNF
	Indexes []int
}

func NewCDCLStack(cnf operators.CNF) *CDCLStack {
	s := CDCLStack{
		Clauses: make(operators.CNF, len(cnf)),
		Indexes: make([]int, 0),
	}

	if copy(s.Clauses, cnf) < len(cnf) {
		return nil
	}

	return &s
}

func (s CDCLStack) peek() operators.Term {
	return s.Clauses[s.Indexes[len(s.Indexes)-1]].(operators.Term)
}

func (s *CDCLStack) Decide(term operators.Term) {
	position := len(s.Clauses)
	s.Clauses = append(s.Clauses, term)
	s.Indexes = append(s.Indexes, position)
	log.Println("decide", term, *s)
}

func (s *CDCLStack) UnitPropagate(clause operators.CNFClause) {
	s.Clauses = append(s.Clauses, clause)
	log.Println("propagate", clause, *s)
}

func (s *CDCLStack) Backtrack() operators.Term {
	var last int
	s.Indexes, last = s.Indexes[:len(s.Indexes)-1], s.Indexes[len(s.Indexes)-1]
	term := s.Clauses[last].(operators.Term)
	s.Clauses = s.Clauses[:last]
	log.Println("backtrack", *s)
	return term
}

func (s *CDCLStack) IsTrue(term operators.Term) bool {
	for _, clause := range s.Clauses {
		if clause.NumTerms() == 1 && clause.Exclude(term) == nil {
			return true
		}
	}
	return false
}

func ModelFromCDCLStack(stack *CDCLStack, variables []operators.Term) (model operators.Node) {
	if len(variables) == 0 {
		return &operators.TrueConst
	}

	choiceVar, remaining := variables[0].Variable(), variables[1:]

	var trueTree, falseTree operators.Node

	if stack.IsTrue(choiceVar) && !stack.IsTrue(choiceVar.Negate()) {
		trueTree = ModelFromCDCLStack(stack, remaining)
		falseTree = &operators.FalseConst
	} else if stack.IsTrue(choiceVar.Negate()) && !stack.IsTrue(choiceVar) {
		trueTree = &operators.FalseConst
		falseTree = ModelFromCDCLStack(stack, remaining)
	} else {
		return &operators.FalseConst
	}

	return &operators.Choice{
		True:  trueTree,
		Var:   choiceVar,
		False: falseTree,
	}
}

// CDCL implements the conflict-driven-clause-learning algorithm
func CDCL(cnf operators.CNF) (model operators.Node) {
	stack := NewCDCLStack(cnf)

	variables := operators.Variables(cnf)

	_ = operators.Cons(recursiveCDCL(nil, variables, stack))

	model = ModelFromCDCLStack(stack, variables)

	return
}

func recursiveCDCL(v operators.Term, variables []operators.Term, stack *CDCLStack) bool {
	if v != nil {
		neg := v.Negate()

		for i := 0; len(stack.Clauses[i:]) > 0; i = i + 1 {
			clause := stack.Clauses[i]

			if clause.HasTerm(neg) {
				if excl := clause.Exclude(neg); excl == nil {
					// negation!
					return false
				} else {
					stack.UnitPropagate(excl)
				}
			}
		}
	}

	// start recursive process
	if len(variables) == 0 {
		return true
	}

	v, remaining := variables[0], variables[1:]
	stack.Decide(v)
	if recursiveCDCL(v, remaining, stack) {
		// yeah, return
		// sat
		return true
	}

	term := stack.Backtrack()
	vNeg := term.Negate()
	stack.Decide(term.Negate())
	if recursiveCDCL(vNeg, remaining, stack) {
		// yeah, return
		// sat
		return true
	}

	stack.Backtrack()

	// helaas
	// unsat
	return false
}
