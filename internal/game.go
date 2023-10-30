package internal

import (
	"math"
	"time"

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
	centered_x := (WIDTH / 2) - BIRD_SIZE
	centered_y := (HEIGHT / 2) - BIRD_SIZE

	g.Bird.Size = r1.Vector2{X: float32(BIRD_SIZE), Y: float32(BIRD_SIZE)}
	g.Bird.Position = r1.Vector2{X: float32(centered_x), Y: float32(centered_y)}
	g.Bird.Speed = 1.5
	g.Bird.Color = r1.LightGray
	g.Bird.Flapping = 0.7
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

		t := float32(math.Pow(float64(g.FrameCounter), 2))
		distance := 0.5 * G * t
		g.Bird.Position.Y += (g.Bird.Speed + distance) * 0.0010

		if r1.IsKeyPressed(r1.KeyQ) {
			g.GameOver = true
		}

		if r1.IsKeyPressed(r1.KeySpace) || r1.IsMouseButtonPressed(r1.MouseLeftButton) {
			g.FrameCounter = 0.0

			for i := 1.2; i < 12.0; i += 1.2 {
				g.Bird.Position.Y += -(float32(i) * float32(g.Bird.Flapping))
				time.Sleep(10 * time.Millisecond)
			}
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
