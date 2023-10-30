package internal

import r1 "github.com/gen2brain/raylib-go/raylib"

type Game struct {
	Bird  Bird
	Pipes []Pipe

	GameOver bool
	Pause    bool
}

func (g *Game) Init() {
	g.Bird.Size = r1.Vector2{X: 25, Y: 25}
	g.Bird.Position = r1.Vector2{X: 0, Y: 0}
	g.Bird.Speed = r1.Vector2{X: 0.5, Y: 0.5}
	g.Bird.Color = r1.Color{R: 255, G: 255, B: 255, A: 1}

	g.GameOver = false
	g.Pause = false
}

func (g *Game) Update() {
	if !g.GameOver {
		if r1.IsKeyPressed(r1.KeyQ) {
			g.GameOver = true
		}
	}
}

func (g *Game) Draw() {
	r1.BeginDrawing()

	r1.ClearBackground(r1.RayWhite)

	if !g.GameOver {
		r1.DrawRectangleV(g.Bird.Position, g.Bird.Size, g.Bird.Color)
	}

	r1.EndDrawing()
}
