package grayongrey

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
)

func parseInput(r io.Reader) (map[string]node, []string, error) {
	var outMap = make(map[string]node)
	var outNames []string
	bb, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, nil, err
	}
	lines := bytes.Split(bb, []byte("\n"))
	for _, line := range lines {
		node, err := processLine(line)
		if err != nil {
			return nil, nil, err
		}
		outMap[node.name] = node
		outNames = append(outNames, node.name)
	}
	return outMap, outNames, nil
}

// Assumption: city names may not have equals signs
// Assumption: directions are lowercase
var regexpDirection = regexp.MustCompile(`^(north|east|south|west)=([^=]+)$`)

var errEmptyLine = errors.New("input line was empty")

func processLine(bb []byte) (node, error) {
	var n node
	fields := bytes.Fields(bb)
	if len(fields) < 1 {
		return n, errEmptyLine
	}
	if len(fields) > 5 {
		return n, fmt.Errorf("input had too many items: %q", string(bb))
	}
	n.name = string(fields[0])
	for i := 1; i < len(fields); i++ {
		groups := regexpDirection.FindSubmatch(fields[i])
		// groups[0] is full match
		// groups[1] is direction
		// groups[2] is destination
		if len(groups) != 3 {
			return n, fmt.Errorf("input direction did not match north=Beirut: %q", string(fields[i]))
		}
		// dir := string(groups[1])
		dest := string(groups[2])
		n.edges = append(n.edges, dest)
	}
	return n, nil
}
