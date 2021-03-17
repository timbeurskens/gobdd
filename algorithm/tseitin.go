package algorithm

import "gobdd/operators"

// TransformTseitin uses the Tseitin transformation to convert an arbitrary expression into CNF
// assume e is in NNF (negation normal form)
func TransformTseitin(e operators.Expression) operators.CNF {
	nMax := operators.Size(e)
	result := make(operators.CNF, 1, nMax)
	queue := make([][2]operators.Expression, 1, nMax)

	i := 0
	start := operators.IVar(i)

	var work [2]operators.Expression

	queue[0] = [2]operators.Expression{start, e}
	result[0] = start

	for len(queue) > 0 {
		work, queue = queue[0], queue[1:]

		left := work[1].LeftChild()
		if _, ok := left.(operators.Operator); left != nil && ok {
			i++
			v := operators.IVar(i)
			queue = append(queue, [2]operators.Expression{v, left})
			left = v
		}

		right := work[1].RightChild()
		if _, ok := right.(operators.Operator); right != nil && ok {
			i++
			v := operators.IVar(i)
			queue = append(queue, [2]operators.Expression{v, right})
			right = v
		}

		// convert work[1] to single operator expression
		work[1].SetLeftChild(left)
		work[1].SetRightChild(right)

		// convert simple clause to cnf
		result = append(result, CNF(operators.Biimplies(work[0], work[1]))...)
	}

	panic("not yet functional")
	return result
}
