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

// processLine is noop if empty line
// If city does not exist, add it
// Update city's edges
// If connected city does not exist, add it
// Update connected city's edges
func processLine(bb []byte, m map[string]*node) error {
	fields := bytes.Fields(bb)
	if len(fields) < 1 {
		return errEmptyLine
	}
	if len(fields) > 5 {
		return fmt.Errorf("input had too many items: %q", string(bb))
	}
	name := string(fields[0])
	rawEdges := make(map[string]string)
	for i := 1; i < len(fields); i++ {
		groups := regexpDirection.FindSubmatch(fields[i])
		// groups[0] is full match
		// groups[1] is direction
		// groups[2] is destination
		if len(groups) != 3 {
			return fmt.Errorf("input edge was not shaped like north=Beirut: %q", string(fields[i]))
		}
		dir := string(groups[1])
		dest := string(groups[2])
		rawEdges[dir] = dest
	}
	updateNodeMap(m, name, rawEdges)
	return nil
}

func updateNodeMap(m map[string]*node, name string, rawEdges map[string]string) {
	center, ok := m[name]
	if !ok {
		center = &node{
			name:      name,
			edges:     make(map[string]*node),
			destroyed: false,
			aliens:    make(map[*alien]struct{}),
		}
	}
	for dir, label := range rawEdges {
		dest, ok := m[label]
		if !ok {
			// implicitly create this newly seen node name
			dest = &node{
				name:      label,
				edges:     make(map[string]*node),
				destroyed: false,
				aliens:    make(map[*alien]struct{}),
			}
		}
		// update edges
		switch dir {
		case "north":
			center.edges["north"] = dest
			dest.edges["south"] = center
		case "east":
			center.edges["east"] = dest
			dest.edges["west"] = center
		case "south":
			center.edges["south"] = dest
			dest.edges["north"] = center
		case "west":
			center.edges["west"] = dest
			dest.edges["east"] = center
		}
		m[label] = dest
	}
	m[name] = center
}
