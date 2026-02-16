package parser

import (
	"io"
	"pathfinder/pkg/model"
)

// ParseNetwork reads the input from the provided reader and returns the constructed Graph.
// It validates the input format and returns detailed errors if parsing fails.
func ParseNetwork(r io.Reader) (*model.Graph, error) {
	// TODO: Initialize parser state (mode: stations/connections)
	// TODO: Loop through lines
	// TODO: Handle comments and whitespace
	// TODO: Parse station lines -> graph.AddStation
	// TODO: Parse connection lines -> graph.AddConnection
	return nil, nil // placeholder
}
