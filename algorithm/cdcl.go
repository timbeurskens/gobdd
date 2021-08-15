package algorithm

import (
	"github.com/timbeurskens/gobdd/display"
	"github.com/timbeurskens/gobdd/operators"
	"strings"
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

func (s *CDCLStack) Strings() []string {
	strs := make([]string, len(s.Clauses))

	for i, clause := range s.Clauses {
		strs[i] = display.PrintExpressiontree(clause.Expr())
	}

	return strs
}

func (s *CDCLStack) String() string {
	return strings.Join(s.Strings(), "\n")
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
	//log.Println("decide", term, *s)
}

func (s *CDCLStack) Push(clause operators.CNFClause) {
	s.Clauses = append(s.Clauses, clause)
	//log.Println("propagate", clause, *s)
}

// UnitPropagate assumes v and updates all clauses in the stack accordingly
// returns true if the resulting stack is satisfiable (or Undetermined), false if it is definitely unsatisfiable.
func (s *CDCLStack) UnitPropagate(v operators.Term) bool {
	if v == nil {
		return true
	}

	neg := v.Negate()

	n := len(s.Clauses)
	count := 0
	//units := make([]operators.Term, 0, n)

	for i := n - 1; i >= 0; i-- {
		clause := s.Clauses[i]

		if clause.HasTerm(neg) {
			if excl := clause.Exclude(neg); excl == nil {
				// negation
				return false
			} else {
				// UnitPropagate 1
				s.Push(excl)
				count++

				// todo: add unit term if it is not yet available in the stack
			}
		}
	}

	return true

	// recursively do unitpropagate until no new units are discovered
}

func (s *CDCLStack) Backtrack() operators.Term {
	var last int
	s.Indexes, last = s.Indexes[:len(s.Indexes)-1], s.Indexes[len(s.Indexes)-1]
	term := s.Clauses[last].(operators.Term)
	s.Clauses = s.Clauses[:last]
	//log.Println("backtrack", *s)
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
func CDCL(cnf operators.CNF) (sat bool, model operators.Node) {
	stack := NewCDCLStack(cnf)

	variables := operators.Variables(cnf)

	sat = recursiveCDCL(nil, variables, stack)

	model = ModelFromCDCLStack(stack, variables)

	return
}

// recursiveCDCL recursively searches for a satisfying assignment of term v, such that the expression in the stack is satisfiable.
func recursiveCDCL(v operators.Term, variables []operators.Term, stack *CDCLStack) bool {

	if !stack.UnitPropagate(v) {
		// if v holds, the expression is unsatisfiable
		return false
	}

	// start recursive process
	if len(variables) == 0 {
		// sat
		return true
	}

	// pick the top variable
	// todo: variable selection based on heuristics / ordering
	v, remaining := variables[0], variables[1:]

	// assume v holds
	stack.Decide(v)

	if recursiveCDCL(v, remaining, stack) {
		// sat
		return true
	}

	// unsat when v holds, so either v does not hold, or this expression is unsat
	// assume v does not hold
	term := stack.Backtrack()
	vNeg := term.Negate()
	stack.Decide(term.Negate())
	if recursiveCDCL(vNeg, remaining, stack) {
		// sat
		return true
	}

	// this expression is unsat
	stack.Backtrack()

	// unsat
	return false
}
