package internal

import r1 "github.com/gen2brain/raylib-go/raylib"

type Bird struct {
	Size     r1.Vector2
	Position r1.Vector2
	Color    r1.Color
	Speed    float32
	Flapping int32
	Flapped  bool
}
