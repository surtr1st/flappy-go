package main

import r1 "github.com/gen2brain/raylib-go/raylib"

var (
	TITLE      = "Flappy"
	WIDTH      = int32(800)
	HEIGHT     = int32(600)
	TARGET_FPS = int32(60)
)

func main() {
	r1.InitWindow(WIDTH, HEIGHT, TITLE)

	r1.SetTargetFPS(TARGET_FPS)

	for !r1.WindowShouldClose() {
	}

	r1.CloseWindow()
}
