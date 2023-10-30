package internal

import (
	"math"

	r1 "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Bird  Bird
	Pipes []Pipe

	GameOver     bool
	Pause        bool
	FrameCounter float32
}

func (g *Game) Init() {
	centered_x := WIDTH / 2
	centered_y := 0

	g.Bird.Size = r1.Vector2{X: 25.0, Y: 25.0}
	g.Bird.Position = r1.Vector2{X: float32(centered_x), Y: float32(centered_y)}
	g.Bird.Speed = 0.5
	g.Bird.Color = r1.LightGray
	g.Bird.Flapping = -5
	g.Bird.Flapped = true
	g.FrameCounter = 1.0

	g.GameOver = false
	g.Pause = false
}

func (g *Game) Update() {
	if !g.GameOver {

		if g.Bird.Position.Y >= float32(HEIGHT)-g.Bird.Size.Y {
			g.Bird.Position.Y = float32(HEIGHT) - g.Bird.Size.Y
			g.FrameCounter = 0.0
		}

		t := math.Pow(float64(g.FrameCounter), 2)
		distance := float32(0.001 * 9.8 * t)
		originalSpeed := g.Bird.Speed
		g.Bird.Position.Y += originalSpeed + distance

		if r1.IsKeyPressed(r1.KeyQ) {
			g.GameOver = true
		}

		if r1.IsKeyPressed(r1.KeySpace) || r1.IsMouseButtonPressed(r1.MouseLeftButton) {
			g.Bird.Speed = float32(g.Bird.Flapping)
			g.Bird.Position.Y += originalSpeed
			g.FrameCounter = 0.0
		}

		g.FrameCounter++
	}
}

func (g *Game) Draw() {
	r1.BeginDrawing()

	r1.ClearBackground(r1.Color{R: 0, G: 0, B: 0, A: 1})
	r1.DrawRectangleV(g.Bird.Position, g.Bird.Size, g.Bird.Color)

	r1.EndDrawing()
}
