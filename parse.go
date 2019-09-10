package grayongrey

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

func parseInput(r io.Reader) ([]node, error) {
	var out []node
	bb, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	lines := bytes.Split(bb, []byte("\n"))
	for _, line := range lines {
		n, err := parseLine(line)
		if err != nil {
			return nil, err
		}
		out = append(out, n)
	}
	return out, nil
}

func parseLine(bb []byte) (node, error) {
	var n node
	fields := bytes.Fields(bb)
	if len(fields) > 5 {
		return n, fmt.Errorf("input had too many items: %q", string(bb))
	}
	return n, nil
}
