package simulation

import (
	"pathfinder/pkg/algo"
)

// Simulate runs the train movement simulation.
// It takes the disjoint paths and the number of trains, then prints the moves to stdout.
func Simulate(paths []algo.Path, numTrains int) error {
	// TODO: Calculate optimal distribution of trains to paths
	// TODO: Loop until all trains have reached the destination
	// TODO: In each turn, move trains and print output in format: T<id>-<station>
	return nil
}

// formatMove creates the string representation of a train's move.
func formatMove(trainID int, station string) string {
	// TODO: Format string
	return ""
}
