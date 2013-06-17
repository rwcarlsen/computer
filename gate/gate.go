package gate

type Conn func(bool)

type Gate struct {
	in1  bool
	in2  bool
	outs []Conn
	fn   func(bool, bool) bool
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

func (d *Gate) Val() bool {
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
