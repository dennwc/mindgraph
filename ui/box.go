package ui

import (
	"github.com/dennwc/dom"
	"github.com/dennwc/dom/svg"
)

var _ Movable = (*Box)(nil)

type Box struct {
	g      *Graph
	parent *Box
	*svg.G
	r      *svg.Rect
	cont   *svg.Container
	name   string
	sub    []*Object
	subb   []*Box
	onMove []func(Pos)
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

func (b *Box) NewNode(name string, pos Pos) *Object {
	if b.cont == nil {
		panic("box not attached")
	}
	n := NewObject(b, b.cont, name)
	b.sub = append(b.sub, n)
	n.OnMouseDown(func(e *dom.MouseEvent) {
		if e.Button() != dom.MouseLeft {
			return
		}
		if IsOnBorder(n.r.DOMElement(), e, 10) {
			// on the border
			b.g.m.TrackMove(e, Resize(n))
		} else {
			// inside client area
			b.g.m.TrackMove(e, Drag(n))
		}
	})
	b.g.RegisterNode(n)
	n.MoveTo(pos)
	return n
}

func (b *Box) newBox(name string) *Box {
	return b.g.newBox(b, nil, name)
}

func (b *Box) NewBox(name string, pos Pos) *Box {
	b2 := b.newBox(name)
	b2.attachSVG(b.cont)
	b.subb = append(b.subb, b2)
	b2.MoveTo(pos)
	return b2
}
func (b *Box) Resize(sz Size) {
	b.r.SetSize(sz.X, sz.Y)
}
func (b *Box) attachSVG(parent *svg.Container) {
	b.G = parent.NewG()
	b.cont = &b.G.Container
	const (
		w, h   = 300, 200
		pad    = 4
		border = 10
	)
	b.cont.OnMouseDown(func(e *dom.MouseEvent) {
		if e.Button() != dom.MouseLeft {
			return
		}
		if IsOnBorder(b.r.DOMElement(), e, border) {
			// on the border
			b.g.m.TrackMove(e, Resize(b))
		} else {
			// inside client area
			b.g.m.TrackMove(e, Drag(b))
		}
	})

	r := b.G.NewRect(w, h)
	r.Stroke("#000")
	r.Fill("#0000ff33")
	r.SetRound(pad, pad)
	b.r = r

	t := b.G.NewText(b.name)
	t.SetDPos(dom.Px(pad), dom.Em(1))
	t.Selectable(false)
}
func (b *Box) OnMove(fnc func(Pos)) {
	b.onMove = append(b.onMove, fnc)
	if b.parent != nil {
		b.parent.OnMove(fnc)
	}
}
