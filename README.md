# GoBDD: Binary Decision Diagram implementation in Go

[![Go](https://github.com/timbeurskens/gobdd/actions/workflows/go.yml/badge.svg)](https://github.com/timbeurskens/gobdd/actions/workflows/go.yml)

This project attempts to implement a simple boolean satifiability solver in Go using Reduced-ordered Binary Decision Diagrams.
Users can define their boolean equations and apply the BDD transformation to get a ROBDD (in Graphviz Dot format).
Additional testing methods are available to test for tautologies, contradictions and equivalence.
Equations in CNF can be solved by applying the Conflict-Driven Clause-Learning (CDCL) method.
The Tseitin transformation can be applied to non-CNF formulas to get an equation in CNF which is satisfiable iff the original equation is satisfiable.
Tautology and contradiction testing for the CDCL output is not (yet) supported.

The current (BDD) based method is very inefficient and contains some problems. The N-queens problem can be solved for N=6 in roughly 6 minutes.
Model-search in both ROBDD-based and CDCL methods is supported by this framework.

### Example: tautology test for p or not p

```go
b := Bench{T: t}
p := Var("p")

b.AssertTautology(
    "p or not p is a tautology",
    algorithm.FromExpression(
        Or(p, Not(p)),
    ),
)
```

### Example: N-queens graphical BDD model

![n-queens model](doc/images/bdd.dot.png)

### Example: CDCL SAT solving

```go
be := bdd_test.Bench{T: t}
a := operators.Var("a")

sat, _ := CDCL(operators.CNF{
    operators.NClause{a.Negate(), a.Negate()},
    operators.NClause{a, a},
})

be.Assert("a xor a is unsat", !sat)
```

### Example: CDCL after applying the Tseitin transformation

```go
be := bdd_test.Bench{T: t}

a, b := op.Var("a"), op.Var("b")
e := op.Xor(a, b),

pruned := PruneUnary(e)
nnf := NNF(e)
cnf := TransformTseitin(nnf)

cdclSat, _ := CDCL(cnf)
resBdd := FromExpression(pruned)

satBdd := bdd.Sat(resBdd)
be.Assert("cdcl and bdd are SAT equivalent", satBdd == cdclSat)
```