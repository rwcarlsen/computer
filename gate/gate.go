package gate

type Conn func(bool)

type Hub struct {
	outputs []Conn
}

func (h *Hub) Out(c Conn) {
	h.outputs = append(h.outputs, c)
}

func (h *Hub) In(v bool) {
	for _, in := range h.outputs {
		in(v)
	}
}

type And struct {
	in1 bool
	in2 bool
	outs []Conn
}

func (a *And) In1(v bool) {
	a.in1 = v
	for _, out := range a.outs {
		out(a.in1 && a.in2)
	}
}

func (a *And) In2(v bool) {
	a.in2 = v
	for _, out := range a.outs {
		out(a.in1 && a.in2)
	}
}

func (a *And) Val() bool {
	return a.in1 && a.in2
}

func (a *And) Out(c ...Conn) {
	a.outs = append(a.outs, c...)
}

type Or struct {
	in1 bool
	in2 bool
	outs []Conn
}

func (o *Or) In1(v bool) {
	o.in1 = v
	for _, out := range o.outs {
		out(o.in1 || o.in2)
	}
}

func (o *Or) In2(v bool) {
	o.in2 = v
	for _, out := range o.outs {
		out(o.in1 || o.in2)
	}
}

func (o *Or) Val() bool {
	return o.in1 || o.in2
}

func (o *Or) Out(c ...Conn) {
	o.outs = append(o.outs, c...)
}

type Not struct {
	in  bool
	outs []Conn
}

func (n *Not) In(v bool) {
	n.in = v
	for _, out := range n.outs {
		out(!n.in)
	}
}

func (n *Not) Val() bool {
	return !n.in
}

func (n *Not) Out(c ...Conn) {
	n.outs = append(n.outs, c...)
}
