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

## Contents

- [Operators](#operators)
  - [Boolean](#boolean)
  - [Numeric](#numeric)
- [Transformations](#transformations)
  - [Unary](#unary)
  - [NNF](#nnf)
  - [Tseitin](#tseitin)
- [Solvers](#solvers)
  - [BDD](#bdd)
  - [CDCL](#cdcl)
- [Examples](#examples)
  - [Tautology test for p or not p](#example-tautology-test-for-p-or-not-p)
  - [N-queens graphical BDD model](#example-n-queens-graphical-bdd-model)
  - [CDCL SAT solving](#example-cdcl-sat-solving)
  - [CDCL after applying the Tseitin transformation](#example-cdcl-after-applying-the-tseitin-transformation)
  - [Prime decomposition](#example-prime-decomposition) 

## Operators

### Boolean

| Operation             | Symbol | Function signature                |
|-----------------------|:------:|-----------------------------------|
| Negation              | ¬      | Not(Expression)                   |
| Conjunction           | ∧      | And(Expression...)                |
| Disjunction           | ∨      | Or(Expression...)                 |
| Exclusive disjunction | ⊗     | Xor(Expression...)                |
| Implication           | →      | Implies(Expression, Expression)   |
| Bi-implication        | ⟷     | Biimplies(Expression, Expression) |

### Numeric

| Operation             | Symbol* | Function signature                |
|-----------------------|:-------:|-----------------------------------|
| Equality              | =       | Equals(Number, Number)            |
| Shift                 | ≪      | Shift(Number, int)                |
| Addition              | +       | Add(Number, Number, Number)       |
| Multiplication        | ×       | Mult(Number, Number, Number)      |

## Transformations

### Unary

### NNF

### Tseitin

## Solvers

### BDD

### CDCL

## Examples

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
The image below shows the generated decision diagram for a 4x4 n-queens solution.
For any node, "true" edges are solid and "false" edges are dotted.
A variable p_{x}_{y} determines whether a queen should be placed on tile (x,y).
This model is found by solving the following expression:

```text
(((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((
(true ∧ ((((false ∨ p_0_0) ∨ p_0_1) ∨ p_0_2) ∨ p_0_3)) ∧ ((((false ∨ p_1_0) ∨ p_1_1) ∨ p_1_2) ∨ p_1_3)) ∧ 
((((false ∨ p_2_0) ∨ p_2_1) ∨ p_2_2) ∨ p_2_3)) ∧ ((((false ∨ p_3_0) ∨ p_3_1) ∨ p_3_2) ∨ p_3_3)) ∧ 
((((false ∨ p_0_0) ∨ p_1_0) ∨ p_2_0) ∨ p_3_0)) ∧ ((((false ∨ p_0_1) ∨ p_1_1) ∨ p_2_1) ∨ p_3_1)) ∧ 
((((false ∨ p_0_2) ∨ p_1_2) ∨ p_2_2) ∨ p_3_2)) ∧ ((((false ∨ p_0_3) ∨ p_1_3) ∨ p_2_3) ∨ p_3_3)) ∧ 
(( ¬ p_0_1) ∨ ( ¬ p_0_0))) ∧ (( ¬ p_0_2) ∨ ( ¬ p_0_0))) ∧ (( ¬ p_0_2) ∨ ( ¬ p_0_1))) ∧ (( ¬ p_0_3) ∨ 
( ¬ p_0_0))) ∧ (( ¬ p_0_3) ∨ ( ¬ p_0_1))) ∧ (( ¬ p_0_3) ∨ ( ¬ p_0_2))) ∧ (( ¬ p_1_1) ∨ ( ¬ p_1_0))) ∧ 
(( ¬ p_1_2) ∨ ( ¬ p_1_0))) ∧ (( ¬ p_1_2) ∨ ( ¬ p_1_1))) ∧ (( ¬ p_1_3) ∨ ( ¬ p_1_0))) ∧ (( ¬ p_1_3) ∨ 
( ¬ p_1_1))) ∧ (( ¬ p_1_3) ∨ ( ¬ p_1_2))) ∧ (( ¬ p_2_1) ∨ ( ¬ p_2_0))) ∧ (( ¬ p_2_2) ∨ ( ¬ p_2_0))) ∧ 
(( ¬ p_2_2) ∨ ( ¬ p_2_1))) ∧ (( ¬ p_2_3) ∨ ( ¬ p_2_0))) ∧ (( ¬ p_2_3) ∨ ( ¬ p_2_1))) ∧ (( ¬ p_2_3) ∨ 
( ¬ p_2_2))) ∧ (( ¬ p_3_1) ∨ ( ¬ p_3_0))) ∧ (( ¬ p_3_2) ∨ ( ¬ p_3_0))) ∧ (( ¬ p_3_2) ∨ ( ¬ p_3_1))) ∧ 
(( ¬ p_3_3) ∨ ( ¬ p_3_0))) ∧ (( ¬ p_3_3) ∨ ( ¬ p_3_1))) ∧ (( ¬ p_3_3) ∨ ( ¬ p_3_2))) ∧ (( ¬ p_1_0) ∨ 
( ¬ p_0_0))) ∧ (( ¬ p_2_0) ∨ ( ¬ p_0_0))) ∧ (( ¬ p_2_0) ∨ ( ¬ p_1_0))) ∧ (( ¬ p_3_0) ∨ ( ¬ p_0_0))) ∧ 
(( ¬ p_3_0) ∨ ( ¬ p_1_0))) ∧ (( ¬ p_3_0) ∨ ( ¬ p_2_0))) ∧ (( ¬ p_1_1) ∨ ( ¬ p_0_1))) ∧ (( ¬ p_2_1) ∨ 
( ¬ p_0_1))) ∧ (( ¬ p_2_1) ∨ ( ¬ p_1_1))) ∧ (( ¬ p_3_1) ∨ ( ¬ p_0_1))) ∧ (( ¬ p_3_1) ∨ ( ¬ p_1_1))) ∧ 
(( ¬ p_3_1) ∨ ( ¬ p_2_1))) ∧ (( ¬ p_1_2) ∨ ( ¬ p_0_2))) ∧ (( ¬ p_2_2) ∨ ( ¬ p_0_2))) ∧ (( ¬ p_2_2) ∨ 
( ¬ p_1_2))) ∧ (( ¬ p_3_2) ∨ ( ¬ p_0_2))) ∧ (( ¬ p_3_2) ∨ ( ¬ p_1_2))) ∧ (( ¬ p_3_2) ∨ ( ¬ p_2_2))) ∧ 
(( ¬ p_1_3) ∨ ( ¬ p_0_3))) ∧ (( ¬ p_2_3) ∨ ( ¬ p_0_3))) ∧ (( ¬ p_2_3) ∨ ( ¬ p_1_3))) ∧ (( ¬ p_3_3) ∨ 
( ¬ p_0_3))) ∧ (( ¬ p_3_3) ∨ ( ¬ p_1_3))) ∧ (( ¬ p_3_3) ∨ ( ¬ p_2_3))) ∧ (( ¬ p_1_0) ∨ ( ¬ p_0_1))) ∧ 
(( ¬ p_1_1) ∨ ( ¬ p_0_0))) ∧ (( ¬ p_1_1) ∨ ( ¬ p_0_2))) ∧ (( ¬ p_1_2) ∨ ( ¬ p_0_1))) ∧ (( ¬ p_1_2) ∨ 
( ¬ p_0_3))) ∧ (( ¬ p_1_3) ∨ ( ¬ p_0_2))) ∧ (( ¬ p_2_0) ∨ ( ¬ p_0_2))) ∧ (( ¬ p_2_0) ∨ ( ¬ p_1_1))) ∧ 
(( ¬ p_2_1) ∨ ( ¬ p_0_3))) ∧ (( ¬ p_2_1) ∨ ( ¬ p_1_0))) ∧ (( ¬ p_2_1) ∨ ( ¬ p_1_2))) ∧ (( ¬ p_2_2) ∨ 
( ¬ p_0_0))) ∧ (( ¬ p_2_2) ∨ ( ¬ p_1_1))) ∧ (( ¬ p_2_2) ∨ ( ¬ p_1_3))) ∧ (( ¬ p_2_3) ∨ ( ¬ p_0_1))) ∧ 
(( ¬ p_2_3) ∨ ( ¬ p_1_2))) ∧ (( ¬ p_3_0) ∨ ( ¬ p_0_3))) ∧ (( ¬ p_3_0) ∨ ( ¬ p_1_2))) ∧ (( ¬ p_3_0) ∨ 
( ¬ p_2_1))) ∧ (( ¬ p_3_1) ∨ ( ¬ p_1_3))) ∧ (( ¬ p_3_1) ∨ ( ¬ p_2_0))) ∧ (( ¬ p_3_1) ∨ ( ¬ p_2_2))) ∧ 
(( ¬ p_3_2) ∨ ( ¬ p_1_0))) ∧ (( ¬ p_3_2) ∨ ( ¬ p_2_1))) ∧ (( ¬ p_3_2) ∨ ( ¬ p_2_3))) ∧ (( ¬ p_3_3) ∨ 
( ¬ p_0_0))) ∧ (( ¬ p_3_3) ∨ ( ¬ p_1_1))) ∧ (( ¬ p_3_3) ∨ ( ¬ p_2_2)))
```

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

### Example: Prime decomposition

```go
bench := bdd_test.Bench{T: t}

// prime1 and prime2 are invisible to the solver
var prime1, prime2 uint = 13, 7

// feed the composite number
combined := prime1 * prime2

// estimate the number of bits needed
bits := int(math.Ceil(math.Log2(float64(combined))))
bits = 2 + (bits-(bits/2))*2

log.Printf("Using %d bits for prime computation", bits)

// construct two arbitrary numbers: a and b
a, b := Variable(bits/2), Variable(bits/2)

// construct a number c
c := Constant(combined, bits)

one := Constant(1, bits/2)

exprEq := operators.And(
    // a != 1
    operators.Not(Equals(a, one)),
    // b != 1
    operators.Not(Equals(b, one)),
)

// c == a * b
exprMul := Mult(a, b, c)

// a * b == c and c == number to test
expr := operators.And(exprEq, exprMul)

// prepare the expression tree
expr = algorithm.PruneUnary(expr)

// run bdd algorithm
tree := algorithm.FromExpression(expr)

bench.AssertSat(fmt.Sprintf("%d is a composed number", combined), tree)

if bdd.Sat(tree) {
    if model, ok := bdd.FindModel(tree); ok {
        aResolv, err := a.Resolve(model)
        if err != nil {
            t.Error(err)
        }
        bResolv, err := b.Resolve(model)
        if err != nil {
            t.Error(err)
        }
        t.Logf("Prime-decomposition of %d: %d x %d", combined, aResolv, bResolv)

        if aResolv * bResolv != combined {
            t.Error("a and b are not a valid decomposition of the original number")
        }
    } else {
        t.Fatal("Could not construct model of non-prime number")
    }
}
```

Output:

```verbose
> RUN TestPrimeDecomposition
> Using 10 bits for prime computation
> Prime-decomposition of 91: 13 x 7
> PASS: TestPrimeDecomposition (4.45s)
```