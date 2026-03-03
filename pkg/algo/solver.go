package algo

import (
	"fmt"
	"pathfinder/pkg/model"
)

// Path represents a sequential list of station names.
type Path []string

type edge struct {
	to   int
	cap  int
	flow int
	cost int
	rev  *edge
}

func addEdge(adj [][]*edge, u, v, cap, cost int) {
	e1 := &edge{to: v, cap: cap, flow: 0, cost: cost}
	e2 := &edge{to: u, cap: 0, flow: 0, cost: -cost}
	e1.rev = e2
	e2.rev = e1
	adj[u] = append(adj[u], e1)
	adj[v] = append(adj[v], e2)
}

func spfa(adj [][]*edge, start, end int) (bool, []*edge) {
	dist := make([]int, len(adj))
	for i := range dist {
		dist[i] = 1e9 // Infinity
	}
	parentEdge := make([]*edge, len(adj))
	inQueue := make([]bool, len(adj))

	queue := []int{start}
	dist[start] = 0
	inQueue[start] = true

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		inQueue[u] = false

		for _, e := range adj[u] {
			if e.cap-e.flow > 0 && dist[u]+e.cost < dist[e.to] {
				dist[e.to] = dist[u] + e.cost
				parentEdge[e.to] = e
				if !inQueue[e.to] {
					queue = append(queue, e.to)
					inQueue[e.to] = true
				}
			}
		}
	}

	if dist[end] == 1e9 {
		return false, nil
	}

	var path []*edge
	curr := end
	for curr != start {
		e := parentEdge[curr]
		path = append([]*edge{e}, path...)
		curr = e.rev.to
	}
	return true, path
}

func extractPaths(adj [][]*edge, start, end int, idToName map[int]string) []Path {
	var paths []Path

	for _, startEdge := range adj[start] {
		if startEdge.flow == 1 {
			var path Path
			path = append(path, idToName[start])

			curr := startEdge.to
			for curr != end {
				stationID := curr
				path = append(path, idToName[stationID])

				outNode := curr + 1 // by our node mapping, out node is in node + 1
				var nextNode int
				for _, outE := range adj[outNode] {
					if outE.flow == 1 {
						nextNode = outE.to
						break
					}
				}
				curr = nextNode
			}
			path = append(path, idToName[end])
			paths = append(paths, path)
		}
	}
	return paths
}

func calcTurns(paths []Path, numTrains int) int {
	if len(paths) == 0 {
		return 0
	}
	pathUsage := make([]int, len(paths))
	for i := 0; i < numTrains; i++ {
		bestIdx := 0
		earliestArrival := (len(paths[0]) - 1) + pathUsage[0]

		for j := 1; j < len(paths); j++ {
			arrival := (len(paths[j]) - 1) + pathUsage[j]
			if arrival < earliestArrival {
				earliestArrival = arrival
				bestIdx = j
			}
		}

		pathUsage[bestIdx]++
	}

	maxTurns := 0
	for j := 0; j < len(paths); j++ {
		if pathUsage[j] > 0 {
			turns := (len(paths[j]) - 1) + pathUsage[j] - 1
			if turns > maxTurns {
				maxTurns = turns
			}
		}
	}
	return maxTurns
}

// FindDisjointPaths finds the optimal combination of node-disjoint paths from start to end in the graph
// to minimize the total number of turns required to move numTrains trains.
// It uses a min-cost max-flow algorithm with node splitting (Suurballe's variant).
func FindDisjointPaths(g *model.Graph, start, end string, numTrains int) ([]Path, error) {
	// Node mapping:
	// 0: start_out
	// 1: end_in
	// other stations: 2i = in, 2i+1 = out
	stationIDs := make(map[string]int)
	idToName := make(map[int]string)

	startNode := 0
	endNode := 1
	idToName[startNode] = start
	idToName[endNode] = end

	idCounter := 2
	for name := range g.Stations {
		if name == start || name == end {
			continue
		}
		stationIDs[name] = idCounter
		idToName[idCounter] = name
		idCounter += 2
	}

	adj := make([][]*edge, idCounter)

	// Add internal node capacities
	for name, id := range stationIDs {
		// in -> out
		idToName[id+1] = name
		addEdge(adj, id, id+1, 1, 0)
	}

	addedEdges := make(map[string]bool)

	for fromName, connections := range g.Connections {
		for _, toName := range connections {
			edgeKey := fmt.Sprintf("%s-%s", fromName, toName)
			if addedEdges[edgeKey] {
				continue
			}
			addedEdges[edgeKey] = true

			var u, v int
			if fromName == start {
				u = startNode
			} else if fromName == end {
				continue // No flow leaves end
			} else {
				u = stationIDs[fromName] + 1 // out node
			}

			if toName == start {
				continue // No flow enters start
			} else if toName == end {
				v = endNode
			} else {
				v = stationIDs[toName] // in node
			}

			addEdge(adj, u, v, 1, 1) // capacity 1, cost 1 (edges represent actual travel distance)
		}
	}

	var bestPaths []Path
	minTurns := int(1e9)

	for {
		found, path := spfa(adj, startNode, endNode)
		if !found {
			break
		}

		// Augment flow
		for _, e := range path {
			e.flow++
			e.rev.flow--
		}

		currentPaths := extractPaths(adj, startNode, endNode, idToName)
		turns := calcTurns(currentPaths, numTrains)

		if turns < minTurns {
			minTurns = turns
			// Deep copy
			bestPaths = make([]Path, len(currentPaths))
			for i, p := range currentPaths {
				bestPaths[i] = make(Path, len(p))
				copy(bestPaths[i], p)
			}
		}
	}

	if len(bestPaths) == 0 {
		return nil, nil // Error wrapper handles len == 0
	}

	return bestPaths, nil
}
