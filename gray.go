package grayongrey

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"strings"
)

type World struct {
	turns  uint
	nodes  map[string]*node
	aliens []*alien
}

// Iterate performs one game loop
func (w *World) Iterate() {
	w.turns++
	for _, a := range w.aliens {
		a.move()
	}
	w.Brawl()
}

// GameOver returns true if 10000 moves have passed or all aliens are terminated
func (w *World) GameOver() bool {
	living := 0
	for _, a := range w.aliens {
		if a.destroyed == false {
			living++
		}
	}
	return (w.turns >= 10000 || living == 0)
}

func (w *World) Brawl() {
	for _, node := range w.nodes {
		participants := make([]string, 0)
		if len(node.aliens) > 1 {
			// destroy aliens
			for a, _ := range node.aliens {
				a.destroyed = true
				participants = append(participants, a.name)
			}
			// destroy city
			node.destroyed = true
			// destroy roads
			for k, v := range node.edges {
				switch k {
				case "north":
					delete(v.edges, "south")
					delete(node.edges, "north")
				case "east":
					delete(v.edges, "west")
					delete(node.edges, "east")
				case "south":
					delete(v.edges, "north")
					delete(node.edges, "south")
				case "west":
					delete(v.edges, "east")
					delete(node.edges, "west")
				}
			}
			msg := strings.Join(participants, ", ")
			msg = node.name + " destroyed by " + msg
			fmt.Println(msg)
		}
	}
}

// Log formats and outputs the world state
func (w *World) Log() {
	// // ignoring sorting
	msg := ""
	for _, node := range w.nodes {
		msg = msg + node.name + " "
		for dir, node2 := range node.edges {
			msg = msg + dir + "=" + node2.name + " "
		}
		msg = strings.TrimSpace(msg)
		msg += "\n"
	}
	fmt.Println(msg)
}

// New takes input world data and number of attackers and creates a *World state
func New(input io.Reader, n uint) (*World, error) {
	aliens := make([]*alien, 0)
	nodeNames := make([]string, 0)
	nodeMap, err := parseInput(input)
	if err != nil {
		return nil, err
	}
	if len(nodeMap) == 0 {
		return nil, errors.New("New called with no input")
	}
	for k, _ := range nodeMap {
		nodeNames = append(nodeNames, k)
	}
	var i uint
	for ; i < n; i++ {
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
	edges     map[string]*node
	destroyed bool
	aliens    map[*alien]struct{}
}

type alien struct {
	name      string
	loc       *node
	destroyed bool
}

// move randomly picks an available edge, or no edge, then updates the alien's location and the locations' aliens
func (a *alien) move() {
	if a.destroyed {
		return
	}
	available := make([]*node, 0)
	for _, v := range a.loc.edges {
		if !v.destroyed {
			available = append(available, v)
		}
	}
	n := len(available)
	i := rand.Intn(n + 1)
	if i == 0 {
		// Don't move
		return
	}
	dest := available[i-1]

	// leave this place
	delete(a.loc.aliens, a)

	// go to that place
	dest.aliens[a] = struct{}{}

	// update self
	a.loc = dest
}
