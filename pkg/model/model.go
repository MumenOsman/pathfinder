package model

import "errors"

// Station represents a node in the graph.
type Station struct {
	Name string
	X    int
	Y    int
}

// Graph represents the network of stations and connections.
type Graph struct {
	Stations    map[string]*Station
	Connections map[string][]string // Adjacency list
}

// NewGraph initializes and returns a new Graph.
func NewGraph() *Graph {
	return &Graph{
		Stations:    make(map[string]*Station),
		Connections: make(map[string][]string),
	}
}

// AddStation adds a station to the graph.
// It should return an error if the station already exists.
func (g *Graph) AddStation(name string, x, y int) error {
	if _, exists := g.Stations[name]; exists {
		return errors.New("duplicate station names: " + name)
	}
	g.Stations[name] = &Station{
		Name: name,
		X:    x,
		Y:    y,
	}
	return nil
}

// AddConnection adds a connection between two stations.
// It should return an error if either station does not exist or if the connection already exists.
func (g *Graph) AddConnection(from, to string) error {
	if _, exists := g.Stations[from]; !exists {
		return errors.New("a connection is made with a station which does not exist: " + from)
	}
	if _, exists := g.Stations[to]; !exists {
		return errors.New("a connection is made with a station which does not exist: " + to)
	}
	if from == to {
		return errors.New("duplicate connection between " + from + " and " + to)
	}

	for _, conn := range g.Connections[from] {
		if conn == to {
			return errors.New("duplicate routes exist between two stations, including in reverse: " + from + "-" + to)
		}
	}

	g.Connections[from] = append(g.Connections[from], to)
	g.Connections[to] = append(g.Connections[to], from)
	return nil
}
