# GoBDD: Binary Decision Diagram implementation in Go

This project attempts to implement a simple boolean satifiability solver in Go using Reduced-ordered Binary Decision Diagrams.
Users can define their boolean equations and apply the BDD transformation to get a ROBDD (in Graphviz Dot format).
Additional testing methods are available to test for tautologies, contradictions and equivalence.

Future versions should extend the SAT solver to also support CNF formulas using the Conflict-Driven Clause-Learning method (CDCL).
In this variant, a Tseitin transformation can be applied to convert any boolean formula to a CNF.

The current (BDD) based method is very inefficient and contains some problems. The N-queens problem can be solved for N=6 in roughly ... minutes.
Model finding in ROBDD-based methods is trivial and supported by this tool.

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

### Example: N-queens BDD model

![n-queens model](doc/images/bdd.dot.png)