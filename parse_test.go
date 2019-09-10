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
	m, err := parseInput(input)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if len(m) != 1 {
		t.Fatalf("expected one node, got %v", len(m))
	}
	_, ok := m["Athens"]
	if !ok {
		t.Fatal("expected node map keyed on Athens, but was missing")
	}
}

func TestParseInputIgnoresEmptyLines(t *testing.T) {
	input := bytes.NewBufferString(`Athens

Beirut
`)
	nodes, err := parseInput(input)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if len(nodes) != 2 {
		t.Fatalf("expected two nodes, got %v", len(nodes))
	}
}

func TestProcessLinePopulatesMissingNodes(t *testing.T) {
	m := make(map[string]*node)
	err := processLine([]byte("Athens north=Beirut"), m)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if len(m) != 2 {
		t.Fatalf("expected two nodes, got %v", len(m))
	}
	_, ok := m["Athens"]
	if !ok {
		t.Fatal("expected node map keyed on Athens, but was missing")
	}
	_, ok = m["Beirut"]
	if !ok {
		t.Fatal("expected node map keyed on Athens, but was missing")
	}
}
