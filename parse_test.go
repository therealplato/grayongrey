package grayongrey

import (
	"bytes"
	"testing"
)

func TestParseInputErrors(t *testing.T) {
	input := bytes.NewBufferString("a up=ISS")
	_, err := parseInput(input)
	if err == nil {
		t.Fatal("expected parse error, got nil")
	}
}
