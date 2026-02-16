package algo

import "pathfinder/pkg/model"

// Path represents a sequential list of station names.
type Path []string

// FindDisjointPaths finds multiple node-disjoint paths from start to end in the graph.
// It returns a slice of paths found.
func FindDisjointPaths(g *model.Graph, start, end string) ([]Path, error) {
	// TODO: Implement BFS to find shortest path
	// TODO: Iterate to find additional disjoint paths (Suurballe's or iterative blocking)
	// TODO: Return all found paths
	return nil, nil // placeholder
}

// bfs finds the shortest path in the graph prohibiting used nodes.
func bfs(g *model.Graph, start, end string, forbidden map[string]bool) Path {
	// TODO: Standard BFS implementation
	return nil
}
