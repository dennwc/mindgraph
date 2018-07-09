package ui

import (
	"github.com/dennwc/dom"
	"github.com/dennwc/dom/svg"
)

var (
	_ Movable  = (*Box)(nil)
	_ Linkable = (*Box)(nil)
)

type Box struct {
	g        *Graph
	parent   *Box
	name     string
	sub      []*Box
	selected bool
	size     Size

	cont *svg.Container
	*svg.G
	r *svg.Rect
	t *svg.Text

	onMove []func(Pos)
}

func (b *Box) IsSubgraph() bool {
	return len(b.sub) != 0
}
func (b *Box) attached() bool {
	return b.r != nil
}
func (b *Box) padding() int {
	if b.IsSubgraph() {
		return 4
	}
	return 10
}
func (b *Box) updateRect() {
	pad := b.padding()
	b.r.SetRound(pad, pad)
	if b.IsSubgraph() {
		if b.selected {
			b.r.Fill("#00eeff55")
		} else {
			b.r.Fill("#0000ff33")
		}
	} else {
		if b.selected {
			b.r.Fill("#ffee00")
		} else {
			b.r.Fill("#ffcc00")
		}
	}
}

func (b *Box) updateSize() {
	sz := b.size
	b.r.SetSize(sz.X, sz.Y)
	if b.IsSubgraph() {
		b.t.Translate(0, 0)
	} else {
		b.t.Translate(0, float64(sz.Y)/2)
	}
}

func (b *Box) updateText() {
	sz := b.size
	var dy dom.Unit
	if b.IsSubgraph() {
		dy = dom.Em(1)
	} else {
		b.t.Translate(0, float64(sz.Y/2))
		dy = dom.Em(0.3)
	}
	pad := b.padding()
	b.t.SetDPos(dom.Px(pad), dy)
}

func (b *Box) Update() {
	if !b.attached() {
		return
	}
	if b.size == (Size{}) {
		b.size = Size{100, 40}
	}
	b.updateSize()
	b.updateRect()
	b.updateText()
}

func (b *Box) Selected(v bool) {
	b.selected = v
	b.updateRect()
}
func (b *Box) Resize(sz Size) {
	b.size = sz
	b.updateSize()
}

func (b *Box) Rect() dom.Rect {
	if b.G == nil {
		return dom.Rect{}
	}
	r := b.G.DOMElement().GetBoundingClientRect()
	zero := b.g.sroot.DOMElement().GetBoundingClientRect().Min
	r.Min = r.Min.Sub(zero)
	r.Max = r.Max.Sub(zero)
	return r
}

func (b *Box) AbsPosition() Pos {
	return b.Rect().Min
}
func (b *Box) zeroPos() Pos {
	var zero Pos
	if b.parent != nil {
		zero = b.parent.AbsPosition()
	}
	return zero
}
func (b *Box) Position() (Pos, Positioner) {
	zero := b.zeroPos()
	var rel Positioner
	if b.parent != nil {
		rel = b.parent
	}
	return b.AbsPosition().Sub(zero), rel
}

func (b *Box) MoveTo(p Pos) {
	b.G.Translate(float64(p.X), float64(p.Y))
	if len(b.onMove) == 0 {
		return
	}
	p = b.zeroPos().Add(p)
	for _, fnc := range b.onMove {
		fnc(p)
	}
}

func (b *Box) Move(dp Pos) {
	p, _ := b.Position()
	b.MoveTo(p.Add(dp))
}

func (b *Box) newBox(name string) *Box {
	return b.g.newBox(b, nil, name)
}

func (b *Box) NewBox(name string, pos Pos) *Box {
	if b.cont == nil {
		panic("box not attached")
	}
	toBox := !b.IsSubgraph()

	b2 := b.newBox(name)
	b2.attachSVG(b.cont)
	b2.Update()
	b2.MoveTo(pos)

	b.g.Selectable(b2)
	b.sub = append(b.sub, b2)
	if toBox {
		b.size = Size{300, 200}
		b.Update()
	}
	return b2
}
func (b *Box) attachSVG(parent *svg.Container) {
	if b.G == nil {
		b.G = parent.NewG()
		b.cont = &b.G.Container
	}
	b.cont.OnMouseDown(func(e *dom.MouseEvent) {
		if e.Button() != dom.MouseLeft {
			return
		}
		const border = 10
		if IsOnBorder(b.r.DOMElement(), e, border) {
			// on the border
			b.g.m.TrackMove(e, Resize(b))
		} else {
			// inside client area
			b.g.m.TrackMove(e, Drag(b))
		}
	})

	if b.r == nil {
		r := b.G.NewRect(0, 0)
		r.Stroke("#000")
		b.r = r
	}

	if b.t == nil {
		t := b.G.NewText(b.name)
		t.Selectable(false)
		b.t = t
	}
}
func (b *Box) OnMove(fnc func(Pos)) {
	b.onMove = append(b.onMove, fnc)
	if b.parent != nil {
		b.parent.OnMove(fnc)
	}
}
