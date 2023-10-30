package main

import (
	i "flappy/internal"

	r1 "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	r1.InitWindow(i.WIDTH, i.HEIGHT, i.TITLE)

	r1.SetTargetFPS(i.TARGET_FPS)

	game := i.Game{}
	game.Init()

	for !r1.WindowShouldClose() {
		if game.GameOver {
			break
		}
		game.Update()
		game.Draw()
	}

	r1.CloseWindow()
}
