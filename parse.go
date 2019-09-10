package grayongrey

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
)

func parseInput(r io.Reader) (map[string]*node, error) {
	var outMap = make(map[string]*node)
	bb, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	lines := bytes.Split(bb, []byte("\n"))
	for _, line := range lines {
		err := processLine(line, outMap)
		if err != nil {
			if err == errEmptyLine {
				continue
			}
			return nil, err
		}
	}
	return outMap, nil
}

// Assumption: city names may not have equals signs
// Assumption: directions are lowercase
var regexpDirection = regexp.MustCompile(`^(north|east|south|west)=([^=]+)$`)

var errEmptyLine = errors.New("input line was empty")

func processLine(bb []byte, m map[string]*node) error {
	fields := bytes.Fields(bb)
	if len(fields) < 1 {
		return errEmptyLine
	}
	if len(fields) > 5 {
		return fmt.Errorf("input had too many items: %q", string(bb))
	}
	name := string(fields[0])
	n, _ := m[name]
	if n == nil {
		n = &node{
			name:   name,
			aliens: make(map[*alien]struct{}),
		}
	}

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
	m[n.name] = n
	return nil
}
