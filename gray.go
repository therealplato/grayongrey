package grayongrey

import (
	"io"
	"math/rand"
	"strconv"
)

type World struct {
	turns  uint
	nodes  map[string]*node
	aliens []*alien
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
	aliens := make([]*alien, 0)
	nodeMap, nodeNames, err := parseInput(input)
	if err != nil {
		return nil, err
	}
	var i uint
	for ; i < attackers; i++ {
		a := &alien{
			name: "Alien " + strconv.Itoa(int(i)),
		}
		j := rand.Intn(len(nodeNames))
		target := nodeMap[nodeNames[j]]
		a.loc = target
		target.aliens[a] = struct{}{}
		aliens = append(aliens, a)
	}
	return &World{
		turns:  0,
		nodes:  nodeMap,
		aliens: aliens,
	}, nil
}

type node struct {
	name      string
	edges     []string
	destroyed bool
	aliens    map[*alien]struct{}
}

type alien struct {
	name      string
	loc       *node
	destroyed bool
}
