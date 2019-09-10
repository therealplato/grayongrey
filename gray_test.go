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
