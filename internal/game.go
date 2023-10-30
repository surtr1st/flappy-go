package internal

import (
	"math"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
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

	g.Bird.Size = rl.NewVector2(float32(BIRD_SIZE), float32(BIRD_SIZE))
	g.Bird.Position = rl.NewVector2(float32(centered_x), float32(centered_y))
	g.Bird.Speed = 1.5
	g.Bird.Color = rl.LightGray
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

		if rl.IsKeyPressed(rl.KeyQ) {
			g.GameOver = true
		}

		if rl.IsKeyPressed(rl.KeySpace) || rl.IsMouseButtonPressed(rl.MouseLeftButton) {
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
	rl.BeginDrawing()

	rl.ClearBackground(rl.Color{R: 0, G: 0, B: 0, A: 1})
	rl.DrawRectangleV(g.Bird.Position, g.Bird.Size, g.Bird.Color)

	rl.EndDrawing()
}
