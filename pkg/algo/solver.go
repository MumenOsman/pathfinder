package algo

import "pathfinder/pkg/model"

// Path represents a sequential list of station names.
type Path []string

// FindDisjointPaths finds multiple node-disjoint paths from start to end in the graph.
// It uses an iterative BFS to find greedy disjoint paths.
// It returns a slice of paths found.
func FindDisjointPaths(g *model.Graph, start, end string) ([]Path, error) {
	var paths []Path
	forbiddenNodes := make(map[string]bool)
	forbiddenEdges := make(map[string]bool)

	for {
		p := bfs(g, start, end, forbiddenNodes, forbiddenEdges)
		if p == nil {
			break
		}
		paths = append(paths, p)

		// Mark nodes as forbidden
		for i := 1; i < len(p)-1; i++ {
			forbiddenNodes[p[i]] = true
		}

		// Mark all edges used in this path as forbidden so we don't reuse tracks
		for i := 0; i < len(p)-1; i++ {
			edge1 := p[i] + "-" + p[i+1]
			edge2 := p[i+1] + "-" + p[i]
			forbiddenEdges[edge1] = true
			forbiddenEdges[edge2] = true
		}
	}

	return paths, nil
}

// bfs finds the shortest path in the graph prohibiting used nodes and edges.
func bfs(g *model.Graph, start, end string, forbiddenNodes map[string]bool, forbiddenEdges map[string]bool) Path {
	queue := [][]string{{start}}

	// visited map MUST be newly created for each BFS run
	visited := make(map[string]bool)
	visited[start] = true

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		curr := path[len(path)-1]

		if curr == end {
			return path
		}

		for _, neighbor := range g.Connections[curr] {
			edge1 := curr + "-" + neighbor
			if !visited[neighbor] && !forbiddenNodes[neighbor] && !forbiddenEdges[edge1] {
				visited[neighbor] = true

				newPath := make([]string, len(path))
				copy(newPath, path)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}
	return nil
}
