package extra

import "testing"

func TestXor(t *testing.T) {
	inputs := [][]bool{
		{false, false, false},
		{false, true, true},
		{true, false, true},
		{true, true, false},
	}

	xor := Xor()
	for i, set := range inputs {
		x, y, out := set[0], set[1], set[2]

		xor.In(x, y)

		if xor.Output()[0] != out {
			t.Errorf("Set %v inputs %v, %v expected %v got %v", i+1, x, y, out, xor.Output()[0])
		}
	}
}

func TestMux(t *testing.T) {
	inputs := [][]bool{
		{true, true, false, true},
		{true, false, false, true},
		{false, true, false, false},
		{false, false, false, false},
		{true, true, true, true},
		{true, false, true, false},
		{false, true, true, true},
		{false, false, true, false},
	}

	mux := Mux()
	for i, set := range inputs {
		x, y, z, out := set[0], set[1], set[2], set[3]

		mux.In(x, y, z)

		if mux.Output()[0] != out {
			t.Errorf("Set %v inputs %v, %v, %v expected %v got %v", i+1, x, y, z, out, mux.Output()[0])
		}
	}
}

func TestDemux(t *testing.T) {
	inputs := [][]bool{
		{true, false, true, false},
		{true, false, true, false},
		{false, false, false, false},
		{false, false, false, false},
		{true, true, false, true},
		{true, true, false, true},
		{false, true, false, false},
		{false, true, false, false},
	}

	mux := Mux()
	for i, set := range inputs {
		x, y, z, out := set[0], set[1], set[2], set[3]

		mux.In(x, y, z)

		if mux.Output()[0] != out {
			t.Errorf("Set %v inputs %v, %v, %v expected %v got %v", i+1, x, y, z, out, mux.Output()[0])
		}
	}
}
