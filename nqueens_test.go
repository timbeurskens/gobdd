package gobdd

import (
    "fmt"
    "testing"

    "gobdd/algorithm"
    "gobdd/bdd_test"
    . "gobdd/operators"
    bdd2 "gobdd/operators/bdd"
)

func TestNQueens(t *testing.T) {
    const n = 4

    b := bdd_test.Bench{T: t}

    field := [n][n]Variable{}

    for i := range field {
        for j := range field[i] {
            field[i][j] = Var(fmt.Sprintf("p_%d_%d", i, j))
        }
    }

    var expr Expression = Cons(true)

    // every row must have at least one queen
    for i := 0; i < n; i++ {
        var disj Expression = Cons(false)
        for j := 0; j < n; j++ {
            disj = Or(disj, field[i][j])
        }
        expr = And(expr, disj)
    }

    // every column must have at least one queen
    for i := 0; i < n; i++ {
        var disj Expression = Cons(false)
        for j := 0; j < n; j++ {
            disj = Or(disj, field[j][i])
        }
        expr = And(expr, disj)
    }

    // every row must have at most one queen
    for i := 0; i < n; i++ {
        for j1 := 1; j1 < n; j1++ {
            for j2 := 0; j2 < j1; j2++ {
                expr = And(expr, Or(Not(field[i][j1]), Not(field[i][j2])))
            }
        }
    }

    // every column must have at most one queen
    for i := 0; i < n; i++ {
        for j1 := 1; j1 < n; j1++ {
            for j2 := 0; j2 < j1; j2++ {
                expr = And(expr, Or(Not(field[j1][i]), Not(field[j2][i])))
            }
        }
    }

    // no two queens on a single diagonal
    for i := 1; i < n*n; i++ {
        i1, j1 := i / n, i % n
        for j := 0; j < i; j++ {
            i2, j2 := j / n, j % n
            if i1-j1 == i2-j2 || i1+j1 == i2+j2 {
                expr = And(expr, Or(Not(field[i1][j1]), Not(field[i2][j2])))
            }
        }
    }

    // gobdd.DotExpressionTree(expr)

    bdd := algorithm.FromExpression(expr)

    t.Log("Size:", Size(expr), "->", Size(bdd))

    DotExpressionTree(expr)
    DotSubtree(bdd)

    b.AssertSat("n-queens is satisfiable", bdd)

    if model, ok := bdd2.FindModel(bdd); ok {
        t.Log(model)

        queens := model.Variables(true)
        b.AssertInfo("there are n queens", len(queens) == n, queens)
    }
}