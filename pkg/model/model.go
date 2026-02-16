package model

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
	// TODO: Check if station exists
	// TODO: logic to add station
	return nil
}

// AddConnection adds a connection between two stations.
// It should return an error if either station does not exist or if the connection already exists.
func (g *Graph) AddConnection(from, to string) error {
	// TODO: Validate stations exist
	// TODO: Add connection to both from and to lists (undirected)
	return nil
}
