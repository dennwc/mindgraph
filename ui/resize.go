package ui

import "github.com/dennwc/dom"

func IsOnBorder(el *dom.Element, e *dom.MouseEvent, border int) bool {
	rect := el.GetBoundingClientRect()
	rect.Min = rect.Min.Add(Pos{border, border})
	rect.Max = rect.Max.Sub(Pos{border, border})
	pos := e.ClientPos()
	if pos.X > rect.Max.X && pos.Y > rect.Max.Y {
		// bottom-right corner
		return true
	}
	// TODO: handle other borders
	return false
}

func Resize(el Resizable) MoveHandler {
	return MoveHandlerFunc(func(dp Pos) {
		el.Resize(el.Rect().Size().Add(dp))
	})
}
