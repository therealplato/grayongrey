package grayongrey

import (
	"io"
	"math/rand"
)

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

// New takes input world data and number of attackers and creates a *World state
func New(input io.Reader, attackers uint) (*World, error) {
	nodeMap, nodeNames, err := parseInput(input)
	if err != nil {
		return nil, err
	}
	var i uint
	for ; i < attackers; i++ {
		j := rand.Intn(len(nodeNames))
		target := nodeMap[nodeNames[j]]
	}
	return &World{
		turns: 0,
		nodes: nodeMap,
	}, nil
}

type node struct {
	name      string
	edges     []string
	destroyed bool
	aliens    []alien
}

type alien struct {
	name string
}
