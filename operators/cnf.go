package operators

import "fmt"

type NClause []Term

func (N NClause) String() string {
	return fmt.Sprint([]Term(N))
}

func (N NClause) Terms() []Term {
	return N
}

func (N NClause) NumTerms() int {
	return len(N)
}

func (N NClause) HasTerm(term Term) bool {
	for _, t := range N {
		if t.TermEquivalent(term) {
			return true
		}
	}
	return false
}

func (N NClause) Exclude(term Term) CNFClause {
	result := make(NClause, 0, len(N))
	excluded := false
	for _, t := range N {
		if t.TermEquivalent(term) {
			excluded = true
		} else {
			result = append(result, t)
		}
	}
	if excluded {
		if len(result) > 0 {
			return result
		} else {
			return nil
		}
	} else {
		return N
	}
}

func JoinCNF(a, b CNF) CNF {
	result := make(CNF, 0, len(a)+len(b))
	result = append(result, a...)
	result = append(result, b...)
	return result
}

func Variables(cnf CNF) []Term {
	result := make(NClause, 0)

	for _, clause := range cnf {
		terms := clause.Terms()
		for _, term := range terms {
			// term.Variable may return nil
			if normal := term.Variable(); normal != nil && !result.HasTerm(normal) {
				result = append(result, normal)
			}
		}
	}

	return result
}
