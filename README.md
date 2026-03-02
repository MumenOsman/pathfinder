# Pathfinder

Pathfinder is a highly robust and mathematically optimized train routing simulation built in Go. It determines the most efficient disjoint routes combining multiple tracks across a complex network topography and simulates iterative train dispatches across those routes to reach final destinations in the strictly minimal elapsed time possible.

## Features

- Parse custom network map topographies mapped to unique Cartesian coordinates.
- Fully fault-tolerant architecture built to trap format inconsistencies, topological constraints, and route logic mapping rules.
- Robust iterative Breadth-First Searches evaluating both Node-Disjoint and strictly Edge-Disjoint pathways, guaranteeing mathematically perfect routes across heavily congested areas.
- Automated Load-Balancers distributing traffic intelligently between short high-traffic parallel lanes and alternative longer sweeping routes based on train volume vs transit delays.

## Algorithms Used

1. **Iterative Breadth-First Search (BFS)**: Central to the `algo` package, it iteratively discovers routes from start to end. With each successful discovery, its specific nodes and edges are added to a "forbidden" map, allowing subsequent BFS iterations to strictly discover node-disjoint and edge-disjoint paths, natively avoiding network collisions and infinite loops.
2. **Greedy Load-Balancing (Simulation)**: Once all independent paths are mapped, the simulation calculates the optimal distribution of $N$ trains. It greedily assigns each train to the path that yields the earliest absolute arrival time (offsetting path length by current queue congestion), guaranteeing the fewest possible turns.

## Build and Run

You'll need an active Go runtime environment on your machine.

**1. Clone the repository:**
```bash
git clone https://github.com/MumenOsman/pathfinder.git
cd pathfinder
```

**2. Compile the executable:**
```bash
go build -o pathfinder.exe .
```

*Note: On Linux/macOS omit `.exe` from the output switch depending on preferences.*

**3. Run the application:**
To start a simulation, you must provide exactly four arguments: 
- `network map file`
- `start station`
- `end station`
- `number of trains`

```bash
./pathfinder.exe network.map waterloo st_pancras 4
```

## E2E Test Suite

The repository contains an automated top-to-bottom testing suite covering failure injections, node looping complexities, and performance timechecks. Execute the testing suite using:
```bash
go test -v .
```