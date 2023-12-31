package internal

import (
	"fmt"
	"math"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Bird     Bird
	Pipes    []Pipe
	PipesPos []rl.Vector2

	Camera       rl.Camera2D
	Score        int32
	GameOver     bool
	FrameCounter float32
}

func (g *Game) Init() {
	x, y := (WIDTH/2)-BIRD_SIZE, (HEIGHT / 2)

	g.Bird.Color = rl.LightGray
	g.Bird.Size = rl.NewVector2(float32(BIRD_SIZE), float32(BIRD_SIZE))
	g.Bird.Position = rl.NewVector2(float32(x), float32(y))
	g.Bird.Speed = 1.5
	g.Bird.Flapping = 0.7
	g.FrameCounter = 1.0

	g.Camera = rl.NewCamera2D(rl.NewVector2(float32(x), 0), rl.NewVector2(g.Bird.Position.X, g.Bird.Position.Y), 0.0, 1.0)

	g.PipesPos = make([]rl.Vector2, MAX_PIPES)
	for i := 0; i < int(MAX_PIPES); i++ {
		g.PipesPos[i].X = float32(480 + 360*i)
		g.PipesPos[i].Y = -float32(rl.GetRandomValue(0, 240))
	}

	pipesWidth := 60
	g.Pipes = make([]Pipe, MAX_PIPES*2)
	for i := 0; i < int(MAX_PIPES)*2; i += 2 {
		g.Pipes[i].Rec.X = g.PipesPos[i/2].X
		g.Pipes[i].Rec.Y = g.PipesPos[i/2].Y
		g.Pipes[i].Rec.Width = float32(pipesWidth)
		g.Pipes[i].Rec.Height = 420
		g.Pipes[i].Color = rl.Green

		g.Pipes[i+1].Rec.X = g.PipesPos[i/2].X
		g.Pipes[i+1].Rec.Y = 1200 + g.PipesPos[i/2].Y - 550
		g.Pipes[i+1].Rec.Width = float32(pipesWidth)
		g.Pipes[i+1].Rec.Height = 420

		g.Pipes[i/2].Active = true
	}

	g.Score = 0
	g.GameOver = false
}

func (g *Game) Update() {
	if !g.GameOver {

		isAtTop := g.Bird.Position.Y <= 0
		isAtBottom := g.Bird.Position.Y >= float32(HEIGHT)-g.Bird.Size.Y

		if isAtTop || isAtBottom {
			g.GameOver = true
		}

		t := float32(math.Pow(float64(g.FrameCounter), 2))
		distance := 0.5 * G * t
		adjustment := float32(0.0025)
		g.Bird.Position.Y += (g.Bird.Speed + distance) * adjustment

		if rl.IsKeyPressed(rl.KeyQ) {
			g.GameOver = true
		}

		if rl.IsKeyPressed(rl.KeySpace) || rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			g.FrameCounter = 0.0

			jumpHeight := 12.5

			for i := 1.2; i < jumpHeight; i += 1.2 {
				g.Bird.Position.Y += -(float32(i) * float32(g.Bird.Flapping))
				time.Sleep(10 * time.Millisecond)
			}
		}

		for i := 0; i < int(MAX_PIPES)*2; i++ {
			if rl.CheckCollisionRecs(rl.NewRectangle(g.Bird.Position.X, g.Bird.Position.Y, float32(BIRD_SIZE), float32(BIRD_SIZE)), g.Pipes[i].Rec) {
				g.GameOver = true
			} else if (g.PipesPos[i/2].X < g.Bird.Position.X-float32(BIRD_SIZE)) && g.Pipes[i/2].Active && !g.GameOver {
				g.Score += 1
				g.Pipes[i/2].Active = false
			}
		}

		if g.Score >= 0 {
			g.Bird.Position.X += 2.2
		}

		if g.Score >= 12 {
			g.Bird.Position.X += 2.6
		}

		if g.Score >= 24 {
			g.Bird.Position.X += 3.0
		}

		if g.Score >= 36 {
			g.Bird.Position.X += 3.2
		}

		g.Camera.Target = rl.NewVector2(g.Bird.Position.X, 0)

		g.FrameCounter++
	}
}

func (g *Game) Draw() {
	rl.BeginDrawing()

	rl.BeginMode2D(g.Camera)

	rl.ClearBackground(rl.Color{R: 0, G: 0, B: 0, A: 1})
	rl.DrawRectangleV(g.Bird.Position, g.Bird.Size, g.Bird.Color)

	for i := 0; i < int(MAX_PIPES); i++ {
		rl.DrawRectangle(int32(g.Pipes[i*2].Rec.X), int32(g.Pipes[i*2].Rec.Y), int32(g.Pipes[i*2].Rec.Width), int32(g.Pipes[i*2].Rec.Height), g.Pipes[i*2].Color)
		rl.DrawRectangle(int32(g.Pipes[i*2+1].Rec.X), int32(g.Pipes[i*2+1].Rec.Y), int32(g.Pipes[i*2+1].Rec.Width), int32(g.Pipes[i*2+1].Rec.Height), g.Pipes[i*2].Color)
	}

	rl.EndMode2D()

	rl.DrawText(fmt.Sprintf("%02d", g.Score), (WIDTH/2 - 1), 0, 32, rl.White)

	rl.EndDrawing()
}
