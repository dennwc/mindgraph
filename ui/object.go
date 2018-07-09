package ui

import (
	"github.com/dennwc/dom"
	"github.com/dennwc/dom/svg"
)

func NewObject(parent *Box, p *svg.Container, text string) *Object {
	o := &Object{parent: parent, text: text}
	o.attachSVG(p)
	return o
}

var _ Movable = (*Object)(nil)

type Object struct {
	text string
	*svg.G
	parent *Box
	t      *svg.Text
	r      *svg.Rect
	onMove []func(Pos)
}

func (o *Object) Rect() dom.Rect {
	if o.G == nil {
		return dom.Rect{}
	}
	r := o.G.DOMElement().GetBoundingClientRect()
	zero := o.parent.g.sroot.DOMElement().GetBoundingClientRect().Min
	r.Min = r.Min.Sub(zero)
	r.Max = r.Max.Sub(zero)
	return r
}

func (o *Object) Resize(sz Size) {
	o.r.SetSize(sz.X, sz.Y)
	o.t.Translate(0, float64(sz.Y)/2)
}

func (o *Object) Move(dp Pos) {
	p, _ := o.Position()
	o.MoveTo(p.Add(dp))
}

func (o *Object) AbsPosition() Pos {
	return o.Rect().Min
}

func (o *Object) zeroPos() Pos {
	var zero Pos
	if o.parent != nil {
		zero = o.parent.AbsPosition()
	}
	return zero
}
func (o *Object) Position() (Pos, Positioner) {
	zero := o.zeroPos()
	var rel Positioner
	if o.parent != nil {
		rel = o.parent
	}
	return o.AbsPosition().Sub(zero), rel
}

func (o *Object) Selected(v bool) {
	if v {
		o.r.Fill("#ffee00")
	} else {
		o.r.Fill("#ffcc00")
	}
}
func (o *Object) attachSVG(parent *svg.Container) {
	o.G = parent.NewG()

	const (
		w, h = 100, 40
		pad  = 10
	)
	o.r = o.G.NewRect(w, h)
	o.r.Stroke("#000")
	o.r.SetRound(pad, pad)

	o.t = o.G.NewText(o.text)
	o.t.Translate(0, h/2)
	o.t.SetDPos(dom.Px(pad), dom.Em(0.3))
	o.t.Selectable(false)

	o.Selected(false)
}
func (o *Object) MoveTo(pos Pos) {
	o.Translate(float64(pos.X), float64(pos.Y))
	if len(o.onMove) == 0 {
		return
	}
	p := o.zeroPos().Add(pos)
	for _, fnc := range o.onMove {
		fnc(p)
	}
}
func (o *Object) OnMove(fnc func(Pos)) {
	o.onMove = append(o.onMove, fnc)
	if o.parent != nil {
		o.parent.OnMove(fnc)
	}
}
