package ui

import "github.com/dennwc/dom/svg"

type Linkable interface {
	Positioner
	OnMove(fnc func(Pos))
}

func NewArrow(cont *svg.Container, src, dst Linkable) *Arrow {
	l := cont.NewLine()
	a := &Arrow{l: l, src: src, dst: dst}
	src.OnMove(func(_ Pos) {
		a.Update()
	})
	dst.OnMove(func(_ Pos) {
		a.Update()
	})
	a.Update()
	return a
}

type Arrow struct {
	l        *svg.Line
	src, dst Linkable
}

func (a *Arrow) Source() Linkable {
	return a.src
}

func (a *Arrow) Target() Linkable {
	return a.dst
}

func (a *Arrow) Reverse() {
	a.src, a.dst = a.dst, a.src
	a.Update()
}

func (a *Arrow) Update() {
	r1, r2 := a.src.Rect(), a.dst.Rect()
	c1, c2 := Center(r1), Center(r2)
	p1, p2 := c1, c2
	p1 = Border(r1, p1, c2)
	p2 = Border(r2, p2, c1)
	a.l.SetPos1(p1)
	a.l.SetPos2(p2)
}
