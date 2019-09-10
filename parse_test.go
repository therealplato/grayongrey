package grayongrey

import (
	"bytes"
	"testing"
)

func TestParseInputErrors(t *testing.T) {
	t.Run("invalid direction", func(t *testing.T) {
		input := bytes.NewBufferString("a up=ISS")
		_, _, err := parseInput(input)
		if err == nil {
			t.Fatal("expected parse error, got nil")
		}
	})
	t.Run("extra directions", func(t *testing.T) {
		input := bytes.NewBufferString("a north=b east=c south=d west=e up=ISS")
		_, _, err := parseInput(input)
		if err == nil {
			t.Fatal("expected parse error, got nil")
		}
	})
}

func TestParseInputAcceptsIsolatedCity(t *testing.T) {
	input := bytes.NewBufferString("Athens")
	nodes, names, err := parseInput(input)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if len(nodes) != 1 {
		t.Fatalf("expected one node, got %v", len(nodes))
	}
	_, ok := nodes["Athens"]
	if !ok {
		t.Fatal("expected node map keyed on Athens, but was missing")
	}
	if len(names) != 1 {
		t.Fatalf("expected one name, got %v", len(names))
	}
	if names[0] != "Athens" {
		t.Fatalf("expected name Athens, got %v", names[0])
	}

}

func TestParseInputIgnoresEmptyLines(t *testing.T) {
	input := bytes.NewBufferString(`Athens

Beirut
`)
	nodes, names, err := parseInput(input)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if len(nodes) != 2 {
		t.Fatalf("expected two nodes, got %v", len(nodes))
	}
	if len(names) != 2 {
		t.Fatalf("expected two names, got %v", len(names))
	}
}

func TestParseInputPopulatesMissingNodes(t *testing.T) {
	t.Skip("Assumption: Every city has a line present in the input")
}
