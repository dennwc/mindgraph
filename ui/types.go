package ui

import "github.com/dennwc/dom"

type Pos = dom.Point
type Size = dom.Point

type Positioner interface {
	Rect() dom.Rect
}

type Movable interface {
	Positioner
	RelMovable
	MoveTo(p Pos)
}

type RelMovable interface {
	Move(dp Pos)
}

type Resizable interface {
	Positioner
	Resize(sz Size)
}

type Linkable interface {
	Positioner
	OnMove(fnc func(Pos))
}
