package grayongrey

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
)

func parseInput(r io.Reader) (map[string]node, error) {
	var out = make(map[string]node)
	bb, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	lines := bytes.Split(bb, []byte("\n"))
	for _, line := range lines {
		err := processLine(line, out)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

// Assumption: city names may not have equals signs
// Assumption: directions are lowercase
var regexpDirection = regexp.MustCompile(`^(north|east|south|west)=([^=]+)$`)

func processLine(bb []byte, nodes map[string]node) error {
	var n node
	fields := bytes.Fields(bb)
	if len(fields) < 1 {
		return fmt.Errorf("input line had zero items: %q", string(bb))
	}
	if len(fields) > 5 {
		return fmt.Errorf("input had too many items: %q", string(bb))
	}
	n.name = string(fields[0])
	for i := 1; i < len(fields); i++ {
		groups := regexpDirection.FindSubmatch(fields[i])
		// groups[0] is full match
		// groups[1] is direction
		// groups[2] is destination
		if len(groups) != 3 {
			return fmt.Errorf("input direction did not match north=Beirut: %q", string(fields[i]))
		}
		// dir := string(groups[1])
		dest := string(groups[2])
		n.edges = append(n.edges, dest)
	}
	nodes[n.name] = n
	return nil
}
