package operators

// IsConstant returns true iff n is constant
func IsConstant(n Node) bool {
	_, ok := n.(Constant)
	return ok
}

// FindOperatorPropagation is a major step for the Apply algorithm for operators
func FindOperatorPropagation(a, b Node, op Operator) (v *Choice, left, right Operator) {
	// either a or b not constant
	cha, choka := a.(*Choice)
	chb, chokb := b.(*Choice)

	if !(choka || chokb) {
		panic("no choice on top")
	}

	// find smallest variable
	if choka && chokb {
		if cha.Var.Leq(chb.Var) {
			if chb.Var == cha.Var {
				return cha, op.Join(a.LeftChild(), b.LeftChild()), op.Join(a.RightChild(), b.RightChild())
			} else {
				return cha, op.Join(a.LeftChild(), b), op.Join(a.RightChild(), b)
			}
		} else {
			return chb, op.Join(a, b.LeftChild()), op.Join(a, b.RightChild())
		}
	} else if choka {
		return cha, op.Join(a.LeftChild(), b), op.Join(a.RightChild(), b)
	} else {
		return chb, op.Join(a, b.LeftChild()), op.Join(a, b.RightChild())
	}
}

func (model Model) Variables(value bool) []Variable {
	result := make([]Variable, 0, len(model))
	for variable, v := range model {
		if value == v {
			result = append(result, variable)
		}
	}
	return result
}

// Size computes the size (number of nodes) of subgraph n
// the function always returns a value >= 1
func Size(n Node) int {
	visited := make(map[Node]bool)
	return sizeRecursive(n, visited)
}

func sizeRecursive(root Node, visited map[Node]bool) int {
	visited[root] = true
	result := 1

	if child := root.LeftChild(); child != nil && !visited[child] {
		result += sizeRecursive(child, visited)
	}

	if child := root.RightChild(); child != nil && !visited[child] {
		result += sizeRecursive(child, visited)
	}

	return result
}
