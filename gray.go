package grayongrey

import "io"

type World struct {
	turns uint
	nodes map[string]node
}

// Iterate performs one game loop
func (w *World) Iterate() {
	w.turns++
}

// Exists returns true if the game should continue from this state
func (w *World) Exists() bool {
	return w.turns < 10000
}

func New(input io.Reader, attackers uint) (*World, error) {
	nodes, err := parseInput(input)
	if err != nil {
		return nil, err
	}
	return &World{
		turns: 0,
		nodes: nodes,
	}, nil
}

type node struct {
	name  string
	edges []string
}
