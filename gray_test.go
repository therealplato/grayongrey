package grayongrey

import (
	"bytes"
	"math/rand"
	"testing"
)

func TestNewWorld(t *testing.T) {
	// Assumption: Cities can be disconnected and have no trailing space
	// Assumption: Input does not necessarily have trailing newline

	t.Run("returns populated world with node input", func(t *testing.T) {
		w, err := New(bytes.NewBufferString("Aberdeen"), 0)
		if err != nil {
			t.Fatalf("expected nil error, got %v\n", err)
		}
		actual := len(w.nodes)
		if actual != 1 {
			t.Fatalf("expected 1 node, had %v cities\n", actual)
		}
	})

	t.Run("places aliens on nodes", func(t *testing.T) {
		w, err := New(bytes.NewBufferString("Aberdeen"), 2)
		if err != nil {
			t.Fatalf("expected nil error, got %v\n", err)
		}
		actual := len(w.nodes["Aberdeen"].aliens)
		if actual != 2 {
			t.Fatalf("expected 2 aliens at Aberdeen, had %v\n", actual)
		}
		if len(w.aliens) != 2 {
			t.Fatalf("expected 2 aliens in world slice, had %v\n", len(w.aliens))
		}
	})
}

func TestTerminationConditions(t *testing.T) {
	t.Run("with no aliens left", func(t *testing.T) {
		w, err := New(bytes.NewBufferString("Aberdeen"), 0)
		if err != nil {
			t.Fatalf("expected nil error, got %v\n", err)
		}
		actual := w.GameOver()
		if actual != true {
			t.Fatal("expected termination with no aliens but game isnt over")
		}
	})
	t.Run("with 10001 moves", func(t *testing.T) {
		w, err := New(bytes.NewBufferString("Aberdeen"), 1)
		if err != nil {
			t.Fatalf("expected nil error, got %v\n", err)
		}
		w.turns = 10001
		actual := w.GameOver()
		if actual != true {
			t.Fatal("expected termination from turn count but game isnt over", err)
		}
	})
}

func TestIterateMovesAlien(t *testing.T) {
	// make deterministic:
	rand.Seed(1)
	a1 := &alien{}
	n1 := &node{
		name:  "Arvada",
		edges: make(map[string]*node),
		aliens: map[*alien]struct{}{
			a1: struct{}{},
		},
	}
	n2 := &node{
		name:   "Boulder",
		edges:  make(map[string]*node),
		aliens: map[*alien]struct{}{},
	}
	a1.loc = n1
	n1.edges["west"] = n2
	n2.edges["east"] = n1
	w := World{
		nodes: map[string]*node{
			"Arvada":  n1,
			"Boulder": n2,
		},
		aliens: []*alien{a1},
	}
	w.Iterate()
	if len(n1.aliens) != 0 {
		t.Fatal("alien did not leave Arvada")
	}
	if len(n2.aliens) != 1 {
		t.Fatal("alien did not arrive in Boulder")
	}
	if w.aliens[0].loc != n2 {
		t.Fatal("alien did not update own loc to Boulder")
	}
}

func TestDestroyedAlienDoesNotMove(t *testing.T) {
	// make deterministic:
	rand.Seed(1)
	a1 := &alien{destroyed: true}
	n1 := &node{
		name:  "Arvada",
		edges: make(map[string]*node),
		aliens: map[*alien]struct{}{
			a1: struct{}{},
		},
	}
	n2 := &node{
		name:   "Boulder",
		edges:  make(map[string]*node),
		aliens: map[*alien]struct{}{},
	}
	a1.loc = n1
	n1.edges["west"] = n2
	n2.edges["east"] = n1
	w := World{
		nodes: map[string]*node{
			"Arvada":  n1,
			"Boulder": n2,
		},
		aliens: []*alien{a1},
	}
	w.Iterate()
	if len(n1.aliens) != 1 {
		t.Fatal("alien should not leave Arvada")
	}
	if len(n2.aliens) != 0 {
		t.Fatal("alien should not arrive in Boulder")
	}
	if w.aliens[0].loc != n1 {
		t.Fatal("alien should not update own loc")
	}
}

func TestBrawlKillsAliens(t *testing.T) {
	rand.Seed(1)
	buf := bytes.NewBufferString(`Athens south=Beirut
Beirut north=Athens south=Charleston
Charleston north=Beirut`)
	w, err := New(buf, 0)
	if err != nil {
		t.Fatal(err)
	}
	a1 := &alien{name: "alien 1"}
	a2 := &alien{name: "alien 2"}
	a1.loc = w.nodes["Beirut"]
	a2.loc = w.nodes["Beirut"]
	w.nodes["Beirut"].aliens = map[*alien]struct{}{
		a1: struct{}{},
		a2: struct{}{},
	}
	w.aliens = []*alien{a1, a2}
	w.Brawl()
	w.Log()
	if a1.destroyed != true {
		t.Fatal("alien 1 should have been destroyed")
	}
	if a2.destroyed != true {
		t.Fatal("alien 2 should have been destroyed")
	}

	if w.nodes["Athens"].destroyed == true {
		t.Fatal("Athens should not have been destroyed")
	}
	if w.nodes["Beirut"].destroyed != true {
		t.Fatal("Beirut should have been destroyed")
	}
	if w.nodes["Charleston"].destroyed == true {
		t.Fatal("Charleston should not have been destroyed")
	}

	if len(w.nodes["Athens"].edges) != 0 {
		t.Fatal("Athens should have had all roads destroyed")
	}
	if len(w.nodes["Beirut"].edges) != 0 {
		t.Fatal("Beirut should have had all roads destroyed")
	}
	if len(w.nodes["Charleston"].edges) != 0 {
		t.Fatal("Charleston should have had all roads destroyed")
	}
}

func TestDestroyedCitiesAreImpassable(t *testing.T) {
	rand.Seed(1)
	buf := bytes.NewBufferString(`Athens south=Beirut
Beirut north=Athens south=Charleston
Charleston north=Beirut`)
	w, err := New(buf, 0)
	if err != nil {
		t.Fatal(err)
	}
	a1 := &alien{name: "alien 1"}
	w.nodes["Athens"].aliens = map[*alien]struct{}{
		a1: struct{}{},
	}
	a1.loc = w.nodes["Athens"]
	w.aliens = []*alien{a1}
	w.nodes["Beirut"].destroyed = true
	w.Iterate()
	if a1.loc != w.nodes["Athens"] {
		t.Fatal("alien 1 should have been stuck in Athens")
	}
}
