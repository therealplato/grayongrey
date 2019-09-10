package grayongrey

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
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

// Assumption: city names may not have equals signs
// Assumption: directions are lowercase
var regexpDirection = regexp.MustCompile(`^(north|east|south|west)=([^=]+)$`)

func parseLine(bb []byte) (node, error) {
	var n node
	fields := bytes.Fields(bb)
	if len(fields) > 5 {
		return n, fmt.Errorf("input had too many items: %q", string(bb))
	}
	for i := 1; i < len(fields); i++ {
		groups := regexpDirection.FindSubmatch(fields[i])
		if len(groups) != 2 {
			return n, fmt.Errorf("input direction did not match north=Beirut: %q", string(fields[i]))
		}
	}
	return n, nil
}
