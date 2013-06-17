package gate

import "testing"

func TestAnd(t *testing.T) {
	inputs := [][]bool{
		{false, false, false},
		{false, true, false},
		{true, false, false},
		{true, true, true},
	}

	and := And()
	for i, set := range inputs {
		x, y, out := set[0], set[1], set[2]
		and.In1(x)
		and.In2(y)

		if and.Output() != out {
			t.Errorf("Set %v inputs %v, %v expected %v got %v", i+1, x, y, out, and.Output())
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

	or := Or()
	for i, set := range inputs {
		x, y, out := set[0], set[1], set[2]
		or.In1(x)
		or.In2(y)

		if or.Output() != out {
			t.Errorf("Set %v inputs %v, %v expected %v got %v", i+1, x, y, out, or.Output())
		}
	}
}

func TestNot(t *testing.T) {
	inputs := [][]bool{
		{false, true},
		{true, false},
	}

	not := Not()
	for i, set := range inputs {
		x, out := set[0], set[1]
		not.In1(x)

		if not.Output() != out {
			t.Errorf("Set %v input %v expected %v got %v", i+1, x, out, not.Output())
		}
	}
}

func TestMulti(t *testing.T) {
	n1 := Not()
	n2 := Not()
	n3 := Not()
	n4 := Not()

	n1.Out(n2.In1)
	n1.Out(n3.In1)
	n1.Out(n4.In1)

	n1.In1(true)

	if !n2.Output() {
		t.Errorf("Expected %v, got %v", true, false)
	} else if !n3.Output() {
		t.Errorf("Expected %v, got %v", true, false)
	} else if !n4.Output() {
		t.Errorf("Expected %v, got %v", true, false)
	}

	n1.In1(false)

	if n2.Output() {
		t.Errorf("Expected %v, got %v", false, true)
	} else if n3.Output() {
		t.Errorf("Expected %v, got %v", false, true)
	} else if n4.Output() {
		t.Errorf("Expected %v, got %v", false, true)
	}
}

func TestChain(t *testing.T) {
	and := And()
	or := Or()
	not := Not()

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
		not.In1(z)

		if or.Output() != out {
			t.Errorf("Set %v inputs %v, %v, %v expected %v got %v", i+1, x, y, z, out, or.Output())
		}
	}
}
