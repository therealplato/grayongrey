package grayongrey

import "io"

type World struct {
	turns uint
}

// Iterate performs one game loop
func (w *World) Iterate() {
	w.turns++
}

// Exists returns true if the game should continue from this state
func (w *World) Exists() bool {
	return w.turns < 10000
}

func New(input io.Reader, attackers uint) *World {
	return &World{}
}
