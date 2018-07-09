package ui

import "github.com/dennwc/dom"

type Selectable interface {
	OnClick(h dom.MouseEventHandler)
	Selected(v bool)
}

func (g *Graph) Deselect() {
	if g.sel != nil {
		g.sel.Selected(false)
	}
	g.sel = nil
}

func (g *Graph) Select(n *Box) {
	g.Deselect()
	if n != nil {
		n.Selected(true)
	}
	g.sel = n
}

func (g *Graph) Selectable(n *Box) {
	n.OnClick(func(e *dom.MouseEvent) {
		if e.Button() != dom.MouseLeft {
			return
		}
		g.Select(n)
	})
}
