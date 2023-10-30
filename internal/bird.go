package internal

import rl "github.com/gen2brain/raylib-go/raylib"

type Bird struct {
	Size     rl.Vector2
	Position rl.Vector2
	Color    rl.Color
	Speed    float32
	Flapping float32
}
