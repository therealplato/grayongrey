package grayongrey

import (
	"bytes"
	"testing"
)

func TestParseInputErrors(t *testing.T) {
	t.Run("invalid direction", func(t *testing.T) {
		input := bytes.NewBufferString("a up=ISS")
		_, err := parseInput(input)
		if err == nil {
			t.Fatal("expected parse error, got nil")
		}
	})
	t.Run("extra directions", func(t *testing.T) {
		input := bytes.NewBufferString("a north=b east=c south=d west=e up=ISS")
		_, err := parseInput(input)
		if err == nil {
			t.Fatal("expected parse error, got nil")
		}
	})
}

func TestParseInputAcceptsIsolatedCity(t *testing.T) {
	input := bytes.NewBufferString("Athens")
	nodes, err := parseInput(input)
	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
	if len(nodes) != 1 {
		t.Fatalf("expected one node, got %v", len(nodes))
	}
	_, ok := nodes["Athens"]
	if !ok {
		t.Fatal("expected node map keyed on Athens, but was missing")
	}
}
