package parser

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"pathfinder/pkg/model"
	"strconv"
	"strings"
	"unicode"
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
				return nil, fmt.Errorf("Line %d: Invalid station line format", lineNumber)
			}
			name := strings.TrimSpace(parts[0])
			xStr := strings.TrimSpace(parts[1])
			yStr := strings.TrimSpace(parts[2])

			x, err := strconv.Atoi(xStr)
			if err != nil || x < 0 {
				return nil, fmt.Errorf("Line %d: Any coordinates which are not a valid positive integer", lineNumber)
			}
			y, err := strconv.Atoi(yStr)
			if err != nil || y < 0 {
				return nil, fmt.Errorf("Line %d: Any coordinates which are not a valid positive integer", lineNumber)
			}

			// Validate name format
			for _, r := range name {
				if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_' {
					return nil, fmt.Errorf("Line %d: Invalid station names", lineNumber)
				}
			}

			// Validate spatial coordinate uniqueness across the station grid
			coordKey := xStr + "," + yStr
			if coordMap[coordKey] {
				return nil, fmt.Errorf("Line %d: Two stations exist at the exact same coordinate location", lineNumber)
			}
			coordMap[coordKey] = true

			err = graph.AddStation(name, x, y)
			if err != nil {
				return nil, fmt.Errorf("Line %d: %v", lineNumber, err) // Re-use the underlying exact struct error message from model.go
			}

			// Graph limitation constraints per project topology rules
			if len(graph.Stations) > 10000 {
				return nil, fmt.Errorf("Line %d: A map contains more than 10000 stations", lineNumber)
			}

		} else if mode == "connections" {
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				return nil, fmt.Errorf("Line %d: Invalid connection line", lineNumber)
			}
			from := strings.TrimSpace(parts[0])
			to := strings.TrimSpace(parts[1])

			err := graph.AddConnection(from, to)
			if err != nil {
				return nil, fmt.Errorf("Line %d: %v", lineNumber, err)
			}
		} else {
			return nil, fmt.Errorf("Line %d: Invalid data before sections", lineNumber)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if !hasStationsSection {
		return nil, errors.New("The map does not contain a 'stations:' section")
	}
	if !hasConnectionsSection {
		return nil, errors.New("The map does not contain a 'connections:' section")
	}

	return graph, nil
}
