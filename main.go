package main

import (
	"fmt"
	"os"
	"pathfinder/pkg/algo"
	"pathfinder/pkg/parser"
	"pathfinder/pkg/simulation"
	"strconv"
)

func main() {
	if len(os.Args) < 5 {
		fmt.Fprintln(os.Stderr, "Error: Too few command line arguments")
		os.Exit(1)
	}

	// Strictly limit to 5 arguments as bonuses are not implemented
	if len(os.Args) > 5 {
		fmt.Fprintln(os.Stderr, "Error: Too many command line arguments")
		os.Exit(1)
	}
	networkMapFile := os.Args[1]
	startStation := os.Args[2]
	endStation := os.Args[3]
	numTrainsStr := os.Args[4]

	numTrains, err := strconv.Atoi(numTrainsStr)
	if err != nil || numTrains <= 0 {
		fmt.Fprintln(os.Stderr, "Error: Invalid number of trains")
		os.Exit(1)
	}

	if startStation == endStation {
		fmt.Fprintln(os.Stderr, "Error: Start and end station are the same")
		os.Exit(1)
	}

	// Read and parse the target network graph map
	file, err := os.Open(networkMapFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: cannot open file %s\n", networkMapFile)
		os.Exit(1)
	}
	defer file.Close()

	graph, err := parser.ParseNetwork(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if _, exists := graph.Stations[startStation]; !exists {
		fmt.Fprintln(os.Stderr, "Error: Start station does not exist")
		os.Exit(1)
	}
	if _, exists := graph.Stations[endStation]; !exists {
		fmt.Fprintln(os.Stderr, "Error: End station does not exist")
		os.Exit(1)
	}

	// Discover node and edge disjoint paths ensuring absolute efficiency
	paths, err := algo.FindDisjointPaths(graph, startStation, endStation, numTrains)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if len(paths) == 0 {
		fmt.Fprintln(os.Stderr, "Error: No path exists between the start and end stations")
		os.Exit(1)
	}

	// Calculate iterative train dispatches and output movements turn by turn
	err = simulation.Simulate(paths, numTrains)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
