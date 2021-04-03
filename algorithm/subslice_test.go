package algorithm

import (
	"github.com/timbeurskens/gobdd/operators"
	"log"
	"testing"
)

func TestSubslice(t *testing.T) {
	s := []int{0, 1, 2, 3}
	t.Log(s[:len(s)-1])
}

func TestNewCDCLStack(t *testing.T) {
	s := NewCDCLStack(operators.CNF{})
	t.Log(s)
	s.Decide(operators.IVar(0))
	t.Log(s)
	s.Decide(operators.IVar(1))
	t.Log(s)
	s.Backtrack()
	t.Log(s)
	s.Decide(operators.IVar(2))
	t.Log(s)
}

func TestRange(t *testing.T) {
	a := []int{0}
	for i := 0; len(a[i:]) > 0; i++ {
		if a[i] < 10 {
			a = append(a, i+1)
		}
	}
	log.Println(a)
}
