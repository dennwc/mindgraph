package ui

import (
	"github.com/dennwc/dom"
	"math"
)

func Center(r dom.Rect) Pos {
	return r.Min.Add(r.Size().Div(2))
}

func Border(r dom.Rect, p1, p2 Pos) Pos {
	if p := intersectX(p1, p2, r.Min.Y, r.Min.X, r.Max.X); p != nil {
		return *p
	}
	if p := intersectY(p1, p2, r.Min.X, r.Min.Y, r.Max.Y); p != nil {
		return *p
	}
	if p := intersectX(p1, p2, r.Max.Y, r.Min.X, r.Max.X); p != nil {
		return *p
	}
	if p := intersectY(p1, p2, r.Max.X, r.Min.Y, r.Max.Y); p != nil {
		return *p
	}
	return Center(r)
}

func intersectX(ps1, pe1 Pos, y, xs2, xe2 int) *Pos {
	v1 := pe1.Sub(ps1)

	ax1, ax2 := float64(v1.X), float64(xe2-xs2)
	ay1 := float64(v1.Y)

	bx1, bx2 := float64(ps1.X), float64(xs2)
	by1, by2 := float64(ps1.Y), float64(y)

	// ax1 * t1 + bx1 = ax2 * t2 + bx2
	// ay1 * t1 + by1 = ay2 * t2 + by2

	// ax1 * t1 + bx1 = ax2 * t2 + bx2
	// ay1 * t1 + by1 = by2

	t1 := (by2 - by1) / ay1
	t2 := (ax1*t1 + bx1 - bx2) / ax2
	if t1 < 0 || t1 > 1 || math.IsNaN(t1) {
		return nil
	} else if t2 < 0 || t2 > 1 || math.IsNaN(t2) {
		return nil
	}
	return &Pos{int(ax1*t1 + bx1), int(ay1*t1 + by1)}
}

func intersectY(ps1, pe1 Pos, x, ys2, ye2 int) *Pos {
	v1 := pe1.Sub(ps1)

	ax1 := float64(v1.X)
	ay1, ay2 := float64(v1.Y), float64(ye2-ys2)

	bx1, bx2 := float64(ps1.X), float64(x)
	by1, by2 := float64(ps1.Y), float64(ys2)

	// ax1 * t1 + bx1 = ax2 * t2 + bx2
	// ay1 * t1 + by1 = ay2 * t2 + by2

	// ax1 * t1 + bx1 = bx2
	// ay1 * t1 + by1 = ay2 * t2 + by2

	t1 := (bx2 - bx1) / ax1
	t2 := (ay1*t1 + by1 - by2) / ay2
	if t1 < 0 || t1 > 1 || math.IsNaN(t1) {
		return nil
	} else if t2 < 0 || t2 > 1 || math.IsNaN(t2) {
		return nil
	}
	return &Pos{int(ax1*t1 + bx1), int(ay1*t1 + by1)}
}
