package parser

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"pathfinder/pkg/model"
	"strconv"
	"strings"
)

// ParseNetwork reads the input from the provided reader and returns the constructed Graph.
// It validates the input format and returns detailed errors if parsing fails.
func ParseNetwork(r io.Reader) (*model.Graph, error) {
	scanner := bufio.NewScanner(r)
	graph := model.NewGraph()
	mode := "none"

	coordMap := make(map[string]bool)
	hasStationsSection := false
	hasConnectionsSection := false
	var lineNumber int

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		// Handle comments
		commentIdx := strings.Index(line, "#")
		if commentIdx != -1 {
			line = line[:commentIdx]
		}
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if line == "stations:" {
			mode = "stations"
			hasStationsSection = true
			continue
		} else if line == "connections:" {
			mode = "connections"
			hasConnectionsSection = true
			continue
		}

		if mode == "stations" {
			parts := strings.Split(line, ",")
			if len(parts) != 3 {
				// Station formats are strictly predefined as name,x,y
				return nil, fmt.Errorf("line %d: invalid station line: %s", lineNumber, line)
			}
			name := strings.TrimSpace(parts[0])
			xStr := strings.TrimSpace(parts[1])
			yStr := strings.TrimSpace(parts[2])

			x, err := strconv.Atoi(xStr)
			if err != nil || x < 0 {
				return nil, fmt.Errorf("line %d: any coordinates which are not a valid positive integer: %s", lineNumber, xStr)
			}
			y, err := strconv.Atoi(yStr)
			if err != nil || y < 0 {
				return nil, fmt.Errorf("line %d: any coordinates which are not a valid positive integer: %s", lineNumber, yStr)
			}

			// Validate name format
			for _, r := range name {
				if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_') {
					return nil, fmt.Errorf("line %d: invalid station names: %s", lineNumber, name)
				}
			}

			// Validate spatial coordinate uniqueness across the station grid
			coordKey := xStr + "," + yStr
			if coordMap[coordKey] {
				return nil, fmt.Errorf("line %d: two stations exist at the exact same coordinate location: %s", lineNumber, coordKey)
			}
			coordMap[coordKey] = true

			err = graph.AddStation(name, x, y)
			if err != nil {
				return nil, fmt.Errorf("line %d: duplicate station names: %s", lineNumber, name)
			}

			// Graph limitation constraints per project topology rules
			if len(graph.Stations) > 10000 {
				return nil, fmt.Errorf("line %d: a map contains more than 10000 stations", lineNumber)
			}

		} else if mode == "connections" {
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				return nil, fmt.Errorf("line %d: invalid connection line: %s", lineNumber, line)
			}
			from := strings.TrimSpace(parts[0])
			to := strings.TrimSpace(parts[1])

			err := graph.AddConnection(from, to)
			if err != nil {
				return nil, fmt.Errorf("line %d: %v", lineNumber, err)
			}
		} else {
			return nil, fmt.Errorf("line %d: invalid data before sections: %s", lineNumber, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if !hasStationsSection {
		return nil, errors.New("the map does not contain a 'stations:' section")
	}
	if !hasConnectionsSection {
		return nil, errors.New("the map does not contain a 'connections:' section")
	}

	return graph, nil
}
