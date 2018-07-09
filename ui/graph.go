package ui

import (
	"github.com/dennwc/dom"
	"github.com/dennwc/dom/svg"
)

func NewGraph(root *svg.Container) *Graph {
	slink := root.NewG()
	snode := root.NewG()
	g := &Graph{sroot: root, slink: &slink.Container, snode: &snode.Container}
	g.m = NewMouseArea(root.DOMElement())
	g.root = g.newBox(nil, &snode.Container, "Graph")
	re := root.DOMElement()
	re.OnClick(func(_ *dom.MouseEvent) {
		g.Deselect()
	})
	return g
}

type Graph struct {
	sroot        *svg.Container
	slink, snode *svg.Container
	m            *MouseArea
	root         *Box
	sel          *Object
	links        []*Arrow
}

func (g *Graph) Root() *Box {
	return g.root
}
func (g *Graph) Links() []*Arrow {
	return append([]*Arrow{}, g.links...)
}
func (g *Graph) Link(n1, n2 Linkable) *Arrow {
	a := NewArrow(g.slink, n1, n2)
	g.links = append(g.links, a)
	return a
}
func (g *Graph) Deselect() {
	if g.sel != nil {
		g.sel.Selected(false)
	}
	g.sel = nil
}
func (g *Graph) Select(n *Object) {
	g.Deselect()
	if n != nil {
		n.Selected(true)
	}
	g.sel = n
}
func (g *Graph) RegisterNode(n *Object) {
	n.OnClick(func(_ *dom.MouseEvent) {
		g.Select(n)
	})
	n.OnMouseDown(func(e *dom.MouseEvent) {
		if e.Button() != dom.MouseLeft {
			return
		}
		g.Select(n)
	})
}
func (g *Graph) NewNode(name string, pos Pos) *Object {
	return g.root.NewNode(name, pos)
}

func (g *Graph) newBox(parent *Box, cont *svg.Container, name string) *Box {
	return &Box{g: g, parent: parent, cont: cont, name: name}
}

func (g *Graph) NewBox(name string, pos Pos) *Box {
	return g.root.NewBox(name, pos)
}
