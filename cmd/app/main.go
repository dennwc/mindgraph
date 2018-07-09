// +build wasm

package main

import (
	"github.com/dennwc/dom"
	"github.com/dennwc/dom/svg"
	"github.com/dennwc/mindgraph/ui"
)

func init() {
	dom.Body.Style().SetMarginsRaw("0")
}

func main() {
	println("running")

	fname := dom.Doc.NewInput("text")
	dom.Body.AppendChild(fname)

	bnew := dom.Doc.NewButton("New")
	dom.Body.AppendChild(bnew)

	bnewb := dom.Doc.NewButton("New Box")
	dom.Body.AppendChild(bnewb)

	s := svg.NewFullscreen()
	g := ui.NewGraph(&s.Container)

	bnew.OnClick(func(_ dom.Event) {
		name := fname.Value()
		g.NewNode(name, ui.Pos{10, 10})
	})

	bnewb.OnClick(func(_ dom.Event) {
		name := fname.Value()
		g.NewBox(name, ui.Pos{10, 10})
	})

	n1 := g.NewNode("Foo", ui.Pos{170, 10})
	n2 := g.NewNode("Bar", ui.Pos{50, 10})

	b := g.NewBox("Sub", ui.Pos{10, 70})
	b.NewNode("A", ui.Pos{10, 40})

	b2 := b.NewBox("Sub 2", ui.Pos{140, 20})
	b2.Resize(ui.Size{140, 70})
	n3 := b2.NewNode("B", ui.Pos{10, 20})

	g.Link(n1, n3)
	g.Link(b, n2)

	dom.Loop()
}
