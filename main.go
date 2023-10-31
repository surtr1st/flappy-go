package main

import (
	i "flappy/internal"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(i.WIDTH, i.HEIGHT, i.TITLE)

	rl.SetTargetFPS(i.TARGET_FPS)

	game := i.Game{}
	game.Init()

	for !rl.WindowShouldClose() {
		if game.GameOver {
			break
		}
		go game.Update()
		game.Draw()
	}

	rl.CloseWindow()
}
