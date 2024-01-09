package main

import (
	"time"

	"github.com/goserg/zg/animator"
	"github.com/goserg/zg/vector"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
)

func main() {
	run()
}

type Game struct {
	pos         vector.Vector[float64]
	posAnimator *animator.Animator[vector.Vector[float64]]

	img *ebiten.Image

	lastTick time.Time
}

func NewGame() *Game {
	g := Game{}
	g.lastTick = time.Now()
	g.pos = vector.Vector[float64]{
		X: 100,
		Y: 100,
	}

	g.posAnimator = animator.New[vector.Vector[float64]](
		func(start vector.Vector[float64],
			target vector.Vector[float64], dt float64) vector.Vector[float64] {
			return vector.Lerp(start, target, dt)
		},
		&animator.Options{
			Duration: time.Second / 4,
			EaseFunc: animator.CustomEase(-2, 0),
		},
	)

	g.img = ebiten.NewImage(50, 50)
	g.img.Fill(colornames.Antiquewhite)
	return &g
}

const speed = 100

func (g *Game) Update() error {
	dt := time.Since(g.lastTick)
	g.lastTick = time.Now()

	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyLeft):
		g.posAnimator.Start(g.pos, vector.Vector[float64]{
			X: g.pos.X - speed,
			Y: g.pos.Y,
		})
	case inpututil.IsKeyJustPressed(ebiten.KeyRight):
		g.posAnimator.Start(g.pos, vector.Vector[float64]{
			X: g.pos.X + speed,
			Y: g.pos.Y,
		})
	case inpututil.IsKeyJustPressed(ebiten.KeyUp):
		g.posAnimator.Start(g.pos, vector.Vector[float64]{
			X: g.pos.X,
			Y: g.pos.Y - speed,
		})
	case inpututil.IsKeyJustPressed(ebiten.KeyDown):
		g.posAnimator.Start(g.pos, vector.Vector[float64]{
			X: g.pos.X,
			Y: g.pos.Y + speed,
		})
	}

	if g.posAnimator.IsAnimating() {
		g.pos = g.posAnimator.Animate(dt)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	var opts ebiten.DrawImageOptions
	opts.GeoM.Translate(g.pos.X, g.pos.Y)
	screen.DrawImage(g.img, &opts)

	ebitenutil.DebugPrint(screen, "use arrow keys to move")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func run() error {
	return ebiten.RunGame(NewGame())
}
