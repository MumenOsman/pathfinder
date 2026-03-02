package simulation

import (
	"fmt"
	"pathfinder/pkg/algo"
	"strings"
)

type Train struct {
	ID      int
	PathIdx int
	Pos     int
}

// Simulate runs the train movement simulation.
// It takes the disjoint paths and the number of trains, then prints the moves to stdout.
func Simulate(paths []algo.Path, numTrains int) error {
	if len(paths) == 0 {
		return fmt.Errorf("no paths available")
	}

	// Calculate optimal distribution of trains to paths
	pathUsage := make([]int, len(paths))
	for i := 0; i < numTrains; i++ {
		bestPathIdx := 0
		earliestArrival := len(paths[0]) - 1 + pathUsage[0]

		for j := 1; j < len(paths); j++ {
			arrival := len(paths[j]) - 1 + pathUsage[j]
			if arrival < earliestArrival {
				earliestArrival = arrival
				bestPathIdx = j
			}
		}

		pathUsage[bestPathIdx]++
	}

	var activeTrains []*Train
	pathUsageRemaining := make([]int, len(paths))
	copy(pathUsageRemaining, pathUsage)

	nextTrainID := 1
	finishedTrains := 0

	for finishedTrains < numTrains {
		var moves []string

		// 1. Move active trains
		for _, t := range activeTrains {
			if t.Pos < len(paths[t.PathIdx])-1 {
				t.Pos++
				moves = append(moves, fmt.Sprintf("T%d-%s", t.ID, paths[t.PathIdx][t.Pos]))
				if t.Pos == len(paths[t.PathIdx])-1 {
					finishedTrains++
				}
			}
		}

		// 2. Spawn new trains
		for pIdx := 0; pIdx < len(paths); pIdx++ {
			if pathUsageRemaining[pIdx] > 0 {
				t := &Train{
					ID:      nextTrainID,
					PathIdx: pIdx,
					Pos:     1,
				}
				nextTrainID++
				activeTrains = append(activeTrains, t)
				moves = append(moves, fmt.Sprintf("T%d-%s", t.ID, paths[pIdx][1]))
				if t.Pos == len(paths[pIdx])-1 {
					finishedTrains++
				}
				pathUsageRemaining[pIdx]--
			}
		}

		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}
	}

	return nil
}
