package grayongrey

import (
	"bytes"
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
		w, err := New(bytes.NewBufferString("Aberdeen"), 1)
		if err != nil {
			t.Fatalf("expected nil error, got %v\n", err)
		}
		actual := len(w.nodes["Aberdeen"].aliens)
		if actual != 1 {
			t.Fatalf("expected 1 alien at Aberdeen, had %v\n", actual)
		}
	})
}
