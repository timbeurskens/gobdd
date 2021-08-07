package algorithm

import "github.com/timbeurskens/gobdd/operators"

// TransformTseitin uses the Tseitin transformation to convert an arbitrary expression into CNF
// assume e is in NNF (negation normal form)
func TransformTseitin(e operators.Expression) operators.CNF {
	nMax := operators.Size(e)
	result := make(operators.CNF, 1, nMax)
	queue := make([][2]operators.Expression, 1, nMax)

	start := operators.IncVar()

	var work [2]operators.Expression

	var leftVar, rightVar operators.Variable
	var exprSplit operators.Expression

	queue[0] = [2]operators.Expression{start, e}
	result[0] = start

	for len(queue) > 0 {
		work, queue = queue[0], queue[1:]

		leftVar = operators.IncVar().(operators.Variable)
		rightVar = operators.IncVar().(operators.Variable)

		switch work[1].(type) {
		case operators.Constant:
			exprSplit = work[1]
		case operators.Variable:
			exprSplit = work[1]
		case *operators.Negation:
			// negations can only occur on a variable
			exprSplit = work[1]
		case *operators.Conjunction:
			exprSplit = &operators.Conjunction{
				A: leftVar,
				B: rightVar,
			}
			queue = append(queue, [2]operators.Expression{leftVar, work[1].LeftChild()})
			queue = append(queue, [2]operators.Expression{rightVar, work[1].RightChild()})
		case *operators.Disjunction:
			exprSplit = &operators.Disjunction{
				A: leftVar,
				B: rightVar,
			}
			queue = append(queue, [2]operators.Expression{leftVar, work[1].LeftChild()})
			queue = append(queue, [2]operators.Expression{rightVar, work[1].RightChild()})
		default:
			panic("unrecognized operator type in Tseitin transformation")
		}

		// convert simple clause to cnf
		cnf := CNF(operators.Biimplies(work[0], exprSplit))
		result = append(result, cnf...)
	}

	return result
}
