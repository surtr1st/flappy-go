package internal

type Game struct {
	Bird  Bird
	Pipes []Pipe

	GameOver bool
	Pause    bool
}
