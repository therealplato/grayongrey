package grayongrey

import (
	"bytes"
	"testing"
)

func TestNewWorld(t *testing.T) {
	t.Run("returns empty world with empty reader", func(t *testing.T) {
		w := New(&bytes.Buffer{}, 0)
		actual := len(w.Cities)
		if actual != 0 {
			t.Fatalf("expected 0 cities, had %v cities\n", actual)
		}
	})

	// Assumption: Cities can be disconnected and have no trailing space
	// Assumption: Input does not necessarily have trailing newline
	var fixtureOneCity = bytes.NewBufferString("Aberdeen")

	t.Run("returns populated world with city input", func(t *testing.T) {
		w := New(fixtureOneCity, 0)
		actual := len(w.Cities)
		if actual != 1 {
			t.Fatalf("expected 1 city, had %v cities\n", actual)
		}
	})
}
