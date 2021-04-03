package bdd_test

import (
	"fmt"
	"github.com/timbeurskens/gobdd/operators"
	"github.com/timbeurskens/gobdd/operators/bdd"
	"testing"
)

type Bench struct {
	T testing.TB
}

func (B Bench) Assert(descr string, expr bool) {
	if !expr {
		B.T.Errorf("Assertion failed for {%s}", descr)
	}
}

func (B Bench) AssertInfo(descr string, expr bool, info ...interface{}) {
	if !expr {
		B.T.Error(fmt.Sprintf("Assertion failed for {%s}", descr), info)
	} else {
		B.T.Log("ok:", descr, info)
	}
}

func (B Bench) AssertSize(descr string, size int, n operators.Node) {
	realSize := operators.Size(n)
	B.AssertInfo(descr, size == realSize, realSize)
}

func (B Bench) AssertNotEquivalent(descr string, a, b operators.Node) {
	B.Assert(descr, !bdd.Equivalent(a, b))
}

func (B Bench) AssertEquivalent(descr string, a, b operators.Node) {
	B.AssertInfo(descr, bdd.Equivalent(a, b), a, b)
}

func (B Bench) AssertSat(descr string, a operators.Node) {
	B.Assert(descr, bdd.Sat(a))
	B.T.Log(bdd.FindModel(a))
}

func (B Bench) AssertUnsat(descr string, a operators.Node) {
	B.Assert(descr, bdd.Unsat(a))
}

func (B Bench) AssertTautology(descr string, a operators.Node) {
	B.Assert(descr, bdd.Tautology(a))
}
