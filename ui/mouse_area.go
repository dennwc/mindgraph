package ui

import "github.com/dennwc/dom"

func NewMouseArea(root *dom.Element) *MouseArea {
	m := &MouseArea{root: root}
	root.OnMouseMove(m.onMove)
	root.OnMouseUp(m.onUp)
	return m
}

type MouseArea struct {
	root *dom.Element
	move struct {
		h    MoveHandler
		from Pos
	}
}

func (m *MouseArea) Rect() dom.Rect {
	return m.root.GetBoundingClientRect()
}
func (m *MouseArea) onMove(e *dom.MouseEvent) {
	if m.move.h != nil {
		pos := e.ClientPos()
		dp := pos.Sub(m.move.from)
		if dp == (Pos{}) {
			// prevent occasional move events when position stays the same
			return
		}
		m.move.from = pos
		m.move.h.Move(dp)
	}
}
func (m *MouseArea) onUp(e *dom.MouseEvent) {
	m.onMove(e)
	if m.move.h != nil {
		m.move.h.EndMove()
		m.move.h = nil
	}
}
func (m *MouseArea) TrackMove(e *dom.MouseEvent, h MoveHandler) {
	m.move.h = h
	m.move.from = e.ClientPos()
}

type MoveHandler interface {
	RelMovable
	EndMove()
}

type MoveHandlerFunc func(dp Pos)

func (fnc MoveHandlerFunc) Move(dp Pos) {
	fnc(dp)
}

func (MoveHandlerFunc) EndMove() {}
