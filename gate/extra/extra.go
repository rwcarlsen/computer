package extra

import (
	"github.com/rwcarlsen/computer/gate"
)

func xorMap(vals []bool, gs []*gate.Gate) {
	and1 := gs[1]
	or := gs[3]
	x, y := vals[0], vals[1]

	and1.In1(x)
	and1.In2(y)
	or.In1(x)
	or.In2(y)
}

// Xor takes 2 inputs and produces 1 output.
func Xor() gate.Chip {
	and1 := gate.And()
	and2 := gate.And()
	not := gate.Not()
	or := gate.Or()

	and1.Out(not.In1)
	not.Out(and2.In1)
	or.Out(and2.In2)

	return gate.NewChip(xorMap, 2, 1, and2, and1, not, or)
}

func muxMap(vals []bool, gs []*gate.Gate) {
	and1 := gs[1]
	and2 := gs[2]
	not := gs[3]
	a, b, sel := vals[0], vals[1], vals[2]

	and1.In1(a)
	not.In1(sel)
	and2.In1(b)
	and2.In2(sel)
}

// Mux takes 3 inputs a, b, and sel (in that order) and produces 1 output. a is
// selected when sel=false. b is selected when sel=true.
func Mux() gate.Chip {
	and1 := gate.And()
	and2 := gate.And()
	or := gate.Or()
	not := gate.Not()

	not.Out(and1.In2)
	and1.Out(or.In1)
	and2.Out(or.In2)

	return gate.NewChip(muxMap, 3, 1, or, and1, and2, not)
}

func demuxMap(vals []bool, gs []*gate.Gate) {
	and1 := gs[0]
	and2 := gs[1]
	not := gs[2]
	x, sel := vals[0], vals[1]

	and1.In1(x)
	not.In1(sel)
	and2.In1(x)
	and2.In2(sel)
}

// Demux takes 2 inputs x, sel (in that order) and produces 2 outputs a, b. a
// is selected if sel=false. b is selected if sel=true
func Dmux() gate.Chip {
	and1 := gate.And()
	and2 := gate.And()
	not := gate.Not()

	not.Out(and1.In2)

	return gate.NewChip(demuxMap, 2, 2, and1, and2, not)
}

func arrayMap(vals []bool, gs []*gate.Gate) {
	for i := range gs {
		gs[i].In1(vals[i*2])
		gs[i].In2(vals[i*2+1])
	}
}

func arrayNotMap(vals []bool, gs []*gate.Gate) {
	for i := range gs {
		gs[i].In1(vals[i])
	}
}

// And16 takes 32 inputs (a, b; a, b; a, b...) and produces 16 outputs.
func And16() gate.Chip {
	gates := make([]*gate.Gate, 16)
	for i := range gates {
		gates[i] = gate.And()
	}
	return gate.NewChip(arrayMap, 32, 16, gates...)
}

// Or16 takes 32 inputs (a, b; a, b; a, b...) and produces 16 outputs.
func Or16() gate.Chip {
	gates := make([]*gate.Gate, 16)
	for i := range gates {
		gates[i] = gate.Or()
	}
	return gate.NewChip(arrayMap, 32, 16, gates...)
}

// Not16 takes 16 inputs and produces 16 outputs.
func Not16() gate.Chip {
	gates := make([]*gate.Gate, 16)
	for i := range gates {
		gates[i] = gate.Not()
	}
	return gate.NewChip(arrayNotMap, 16, 16, gates...)
}

// Mux16 takes 33 inputs (a1, a2, a3...; b1, b2, b3..., sel) and produces 16
// outputs.
func Mux16() gate.Chip {
	return nil
}
