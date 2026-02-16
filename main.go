package main

import (
	"fmt"
	"os"
)

func main() {
	// Validate command line arguments
	if len(os.Args) != 5 {
		fmt.Fprintln(os.Stderr, "Error: Incorrect number of command line arguments")
		fmt.Fprintln(os.Stderr, "Usage: go run . [network_map] [start_station] [end_station] [valid_number_of_trains]")
		os.Exit(1)
	}

	// Parse arguments
	// networkMapFile := os.Args[1]
	// startStation := os.Args[2]
	// endStation := os.Args[3]
	// numTrainsStr := os.Args[4]

	// TODO: Parse numTrains to int and validate > 0

	// TODO: Open network map file

	// TODO: Parse network map using parser.ParseNetwork

	// TODO: Validate start and end stations exist in the graph

	// TODO: Run algo.FindDisjointPaths

	// TODO: Run simulation.Simulate
}
