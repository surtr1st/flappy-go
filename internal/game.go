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
	x, y := (WIDTH/2)-BIRD_SIZE, (HEIGHT/2)-BIRD_SIZE

	g.Bird.Color = rl.LightGray
	g.Bird.Size = rl.NewVector2(float32(BIRD_SIZE), float32(BIRD_SIZE))
	g.Bird.Position = rl.NewVector2(float32(x), float32(y))
	g.Bird.Speed = 1.5
	g.Bird.Flapping = 0.7
	g.FrameCounter = 1.0

	pipesPos := make([]rl.Vector2, MAX_PIPES)
	for i := 0; i < int(MAX_PIPES); i++ {
		pipesPos[i].X = float32(480 + 360*i)
		pipesPos[i].Y = -float32(rl.GetRandomValue(0, 240))
	}

	pipesWidth := 60
	g.Pipes = make([]Pipe, MAX_PIPES*2)
	for i := 0; i < int(MAX_PIPES)*2; i += 2 {
		g.Pipes[i].Rec.X = pipesPos[i/2].X
		g.Pipes[i].Rec.Y = pipesPos[i/2].Y
		g.Pipes[i].Rec.Width = float32(pipesWidth)
		g.Pipes[i].Rec.Height = 400
		g.Pipes[i].Color = rl.Green

		g.Pipes[i+1].Rec.X = pipesPos[i/2].X
		g.Pipes[i+1].Rec.Y = 1200 + pipesPos[i/2].Y - 550
		g.Pipes[i+1].Rec.Width = float32(pipesWidth)
		g.Pipes[i+1].Rec.Height = 400
	}

	g.GameOver = false
	g.Pause = false
}

func (g *Game) Update() {
	if !g.GameOver {

		isAtBottom := g.Bird.Position.Y >= float32(HEIGHT)-g.Bird.Size.Y

		if isAtBottom {
			g.Bird.Position.Y = float32(HEIGHT) - g.Bird.Size.Y
			g.FrameCounter = 0.0
		}

		t := float32(math.Pow(float64(g.FrameCounter), 2))
		distance := 0.5 * G * t
		adjustment := float32(0.0010)
		g.Bird.Position.Y += (g.Bird.Speed + distance) * adjustment

		if rl.IsKeyPressed(rl.KeyQ) {
			g.GameOver = true
		}

		if rl.IsKeyPressed(rl.KeySpace) || rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			g.FrameCounter = 0.0

			jumpHeight := 12.0

			for i := 1.2; i < jumpHeight; i += 1.2 {
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

	for i := 0; i < int(MAX_PIPES); i++ {
		rl.DrawRectangle(int32(g.Pipes[i*2].Rec.X), int32(g.Pipes[i*2].Rec.Y), int32(g.Pipes[i*2].Rec.Width), int32(g.Pipes[i*2].Rec.Height), g.Pipes[i*2].Color)
		rl.DrawRectangle(int32(g.Pipes[i*2+1].Rec.X), int32(g.Pipes[i*2+1].Rec.Y), int32(g.Pipes[i*2+1].Rec.Width), int32(g.Pipes[i*2+1].Rec.Height), g.Pipes[i*2].Color)

		rl.DrawRectangleLines(int32(g.Pipes[i*2].Rec.X), int32(g.Pipes[i*2].Rec.Y), int32(g.Pipes[i*2].Rec.Width), int32(g.Pipes[i*2].Rec.Height), rl.Black)
		rl.DrawRectangleLines(int32(g.Pipes[i*2+1].Rec.X), int32(g.Pipes[i*2+1].Rec.Y), int32(g.Pipes[i*2+1].Rec.Width), int32(g.Pipes[i*2+1].Rec.Height), rl.Black)
	}

	rl.EndDrawing()
}
