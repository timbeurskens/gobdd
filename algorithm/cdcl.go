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
	term := s.Clauses[last].(operators.Term).Negate()
	s.Clauses = append(s.Clauses[:last], term)
	log.Println("backtrack", *s)
	return term
}

// CDCL implements the conflict-driven-clause-learning algorithm
func CDCL(cnf operators.CNF) operators.Expression {
	stack := NewCDCLStack(cnf)

	variables := operators.Variables(cnf)

	return operators.Cons(recursiveCDCL(nil, variables, stack))
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

	vNeg := stack.Backtrack()
	if recursiveCDCL(vNeg, remaining, stack) {
		// yeah, return
		// sat
		return true
	}

	// helaas
	// unsat
	return false
}
