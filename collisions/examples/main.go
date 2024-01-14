package main

import (
	"math"
	"time"

	"github.com/goserg/zg/collisions"
	vectorF "github.com/goserg/zg/vector"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"
)

func main() {
	g := Game{}
	g.lastDraw = time.Now()
	p := collisions.NewRect(vectorF.Vector{X: 10, Y: 10}, vectorF.Vector{X: 30, Y: 20})
	g.rects = append(g.rects, &p)
	r := collisions.NewRect(vectorF.Vector{X: 100, Y: 100}, vectorF.Vector{X: 80, Y: 50})
	g.rects = append(g.rects, &r)
	r1 := collisions.NewRect(vectorF.Vector{X: 70, Y: 70}, vectorF.Vector{X: 80, Y: 50})
	g.rects = append(g.rects, &r1)
	r3 := collisions.NewRect(vectorF.Vector{X: 200, Y: 270}, vectorF.Vector{X: 80, Y: 50})
	g.rects = append(g.rects, &r3)
	r4 := collisions.NewRect(vectorF.Vector{X: 400, Y: 300}, vectorF.Vector{X: 80, Y: 50})
	g.rects = append(g.rects, &r4)
	ebiten.RunGame(&g)
}

type Game struct {
	rects    []*collisions.Rect
	speed    vectorF.Vector
	lastDraw time.Time
}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		x, y := ebiten.CursorPosition()
		g.speed = vectorF.Vector{
			X: float64(x),
			Y: float64(y),
		}.Sub(g.rects[0].Pos).Normalize().Mul(500)
	} else {
		g.speed = vectorF.Vector{}
	}
	dt := time.Since(g.lastDraw)
	g.lastDraw = time.Now()
	for i := 1; i < len(g.rects); i++ {
		r := g.rects[i]
		col, _, contactNormal, contactTime := collisions.DynamicRectVsRect(*g.rects[0], g.speed, *r, dt)
		if col && contactTime < 1 {
			g.speed.X += contactNormal.X * math.Abs(g.speed.X) * (1 - contactTime)
			g.speed.Y += contactNormal.Y * math.Abs(g.speed.Y) * (1 - contactTime)
		}
	}

	g.rects[0].Pos = g.rects[0].Pos.Add(g.speed.Mul(dt.Seconds()))

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, r := range g.rects {
		vector.StrokeRect(
			screen,
			float32(r.Pos.X),
			float32(r.Pos.Y),
			float32(r.Size.X),
			float32(r.Size.Y),
			1,
			colornames.Yellow,
			true,
		)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
