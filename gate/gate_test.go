package gate

import (
	"testing"
)

func TestAnd(t *testing.T) {
	inputs := [][]bool{
		{false, false, false},
		{false, true, false},
		{true, false, false},
		{true, true, true},
	}

	and := new(And)
	for i, set := range inputs {
		x, y, out := set[0], set[1], set[2]
		and.In1(x)
		and.In2(y)

		if and.Val() != out {
			t.Errorf("Set %v inputs %v, %v expected %v got %v", i+1, x, y, out, and.Val())
		}
	}
}

func TestOr(t *testing.T) {
	inputs := [][]bool{
		{false, false, false},
		{false, true, true},
		{true, false, true},
		{true, true, true},
	}

	or := new(Or)
	for i, set := range inputs {
		x, y, out := set[0], set[1], set[2]
		or.In1(x)
		or.In2(y)

		if or.Val() != out {
			t.Errorf("Set %v inputs %v, %v expected %v got %v", i+1, x, y, out, or.Val())
		}
	}
}

func TestNot(t *testing.T) {
	inputs := [][]bool{
		{false, true},
		{true, false},
	}

	not := new(Not)
	for i, set := range inputs {
		x, out := set[0], set[1]
		not.In(x)

		if not.Val() != out {
			t.Errorf("Set %v input %v expected %v got %v", i+1, x, out, not.Val())
		}
	}
}

func TestMulti(t *testing.T) {
	n1 := new(Not)
	n2 := new(Not)
	n3 := new(Not)
	n4 := new(Not)

	n1.Out(n2.In)
	n1.Out(n3.In)
	n1.Out(n4.In)

	n1.In(true)

	if !n2.Val() {
		t.Errorf("Expected %v, got %v", true, false)
	} else if !n3.Val() {
		t.Errorf("Expected %v, got %v", true, false)
	} else if !n4.Val() {
		t.Errorf("Expected %v, got %v", true, false)
	}

	n1.In(false)

	if n2.Val() {
		t.Errorf("Expected %v, got %v", false, true)
	} else if n3.Val() {
		t.Errorf("Expected %v, got %v", false, true)
	} else if n4.Val() {
		t.Errorf("Expected %v, got %v", false, true)
	}
}

func TestChain(t *testing.T) {
	and := new(And)
	or := new(Or)
	not := new(Not)

	and.Out(or.In1)
	not.Out(or.In2)

	inputs := [][]bool{
		{false, false, false, true},
		{false, false, true, false},
		{false, true, false, true},
		{false, true, true, false},
		{true, false, false, true},
		{true, false, true, false},
		{true, true, false, true},
		{true, true, true, true},
	}

	for i, set := range inputs {
		x, y, z, out := set[0], set[1], set[2], set[3]
		and.In1(x)
		and.In2(y)
		not.In(z)

		if or.Val() != out {
			t.Errorf("Set %v inputs %v, %v, %v expected %v got %v", i+1, x, y, z, out, or.Val())
		}
	}
}
