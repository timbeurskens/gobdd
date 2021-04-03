package main

import (
	"fmt"

	. "gobdd"
	"gobdd/algorithm"
	"gobdd/operators"
	"gobdd/operators/bdd"
)

func main() {
	p, q, r, s := operators.Var("p"), operators.Var("q"), operators.Var("r"), operators.Var("s")

	expr1 := operators.Implies(p, q)
	expr2 := operators.Or(operators.Not(p), q)

	expr1 = algorithm.PruneUnary(expr1)
	expr2 = algorithm.PruneUnary(expr2)

	expr1Tree := algorithm.FromExpression(expr1)
	expr2Tree := algorithm.FromExpression(expr2)

	DotSubtree(expr1Tree)
	DotSubtree(expr2Tree)

	DotSubtree(algorithm.FromExpression(operators.Xor(p, q)))

	fmt.Println(bdd.Sat(expr1Tree))
	fmt.Println(bdd.Sat(algorithm.FromExpression(algorithm.PruneUnary(operators.Not(operators.Implies(p, operators.Cons(true)))))))

	cnf := algorithm.FromExpression(algorithm.PruneUnary(operators.And(operators.Or(p, operators.Not(q)), operators.Or(q, r), operators.Or(q, s), operators.Or(p, operators.Not(s)), operators.Or(operators.Not(q), s), operators.Or(operators.Not(p), operators.Not(r)))))
	fmt.Println(bdd.FindModel(cnf))

	DotSubtree(cnf)
	// DotExpressionTree(operators.And(operators.Or(p, operators.Not(q)), operators.Or(q, r), operators.Or(q, s), operators.Or(p, operators.Not(s)), operators.Or(operators.Not(q), s), operators.Or(operators.Not(p), operators.Not(r))))

	DotSubtree(algorithm.FromExpression(operators.And(p, q)))
}
