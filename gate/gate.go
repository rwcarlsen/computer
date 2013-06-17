package gate

type Conn func(bool)

type Gate struct {
	in1  bool
	in2  bool
	outs []Conn
	fn   func(bool, bool) bool
}

func NewGate(fn func(x, y bool) bool) *Gate {
	return &Gate{fn: fn}
}

func (d *Gate) In1(v bool) {
	d.in1 = v
	for _, out := range d.outs {
		out(d.fn(d.in1, d.in2))
	}
}

func (d *Gate) In2(v bool) {
	d.in2 = v
	for _, out := range d.outs {
		out(d.fn(d.in1, d.in2))
	}
}

func (d *Gate) Output() bool {
	return d.fn(d.in1, d.in2)
}

func (d *Gate) Out(c ...Conn) {
	d.outs = append(d.outs, c...)
}

func And() *Gate {
	return &Gate{fn: func(x, y bool) bool { return x && y }}
}

func Or() *Gate {
	return &Gate{fn: func(x, y bool) bool { return x || y }}
}

func Not() *Gate {
	return &Gate{fn: func(x, y bool) bool { return !x }}
}

type Chip interface {
	In(...bool)
	InPin(int) Conn
	Out(int, ...Conn)
	Output() []bool
}

type chip struct {
	mapFn   func([]bool, []*Gate)
	gateSet []*Gate
	nOuts   int
	outs    [][]Conn
	ins     []bool
}

// NewChip creates a new compositie chip composed of one or more connected
// gates. The output gates for the chip must be first and in desired order.
// nOuts must be less than or equal to the number of gates.
func NewChip(mapFn func([]bool, []*Gate), nIns, nOuts int, gates ...*Gate) Chip {
	return &chip{
		mapFn:   mapFn,
		gateSet: gates,
		nOuts:   nOuts,
		outs:    make([][]Conn, nOuts),
		ins:     make([]bool, nIns),
	}
}

func (c *chip) In(vals ...bool) {
	c.ins = vals
	c.mapFn(vals, c.gateSet)
	for i, outs := range c.outs {
		for _, out := range outs {
			out(c.gateSet[i].Output())
		}
	}
}

func (c *chip) InPin(i int) Conn {
	return func(v bool) {
		c.ins[i] = v
		c.In(c.ins...)
	}
}

func (c *chip) Out(i int, cs ...Conn) {
	c.outs[i] = append(c.outs[i], cs...)
}

func (c *chip) Output() []bool {
	vals := make([]bool, c.nOuts)
	for i := range vals {
		vals[i] = c.gateSet[i].Output()
	}
	return vals
}
