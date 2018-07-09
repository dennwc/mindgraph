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

	s := svg.NewFullscreen()
	g := ui.NewGraph(&s.Container)

	bnew.OnClick(func(_ dom.Event) {
		name := fname.Value()
		p := ui.Pos{10, 10}
		if sel := g.Selected(); sel != nil {
			sel.NewBox(name, p)
		} else {
			g.NewBox(name, p)
		}
	})

	n1 := g.NewBox("Foo", ui.Pos{170, 10})
	n2 := g.NewBox("Bar", ui.Pos{50, 10})

	b := g.NewBox("Sub", ui.Pos{10, 70})
	b.NewBox("A", ui.Pos{10, 40})

	b2 := b.NewBox("Sub 2", ui.Pos{140, 20})
	b2.Resize(ui.Size{140, 70})
	n3 := b2.NewBox("B", ui.Pos{10, 20})

	g.Link(n1, n3)
	g.Link(b, n2)

	dom.Loop()
}
