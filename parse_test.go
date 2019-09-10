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
