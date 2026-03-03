package main_test

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func TestEndToEndErrors(t *testing.T) {
	// Build the binary once
	cmd := exec.Command("go", "build", "-o", "pathfinder.exe", ".")
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}

	tests := []struct {
		name           string
		args           []string
		expectedErr    string
		expectExitCode int
	}{
		{
			name:           "Too few arguments",
			args:           []string{"arg1", "arg2"},
			expectedErr:    "Error: Too few command line arguments",
			expectExitCode: 1,
		},
		{
			name:           "Invalid number of trains",
			args:           []string{"network.map", "start", "end", "abc"},
			expectedErr:    "Error: Invalid number of trains",
			expectExitCode: 1,
		},
		{
			name:           "Same start and end station",
			args:           []string{"network.map", "stationA", "stationA", "5"},
			expectedErr:    "Error: Start and end station are the same",
			expectExitCode: 1,
		},
		{
			name:           "Map file does not exist",
			args:           []string{"nonexistent.map", "start", "end", "5"},
			expectedErr:    "Error: cannot open file",
			expectExitCode: 1,
		},
		{
			name:           "Start station does not exist",
			args:           []string{"network_beethoven.map", "nowhere", "part", "5"},
			expectedErr:    "Error: Start station does not exist",
			expectExitCode: 1,
		},
		{
			name:           "End station does not exist",
			args:           []string{"network_small_large.map", "small", "nowhere", "5"},
			expectedErr:    "Error: End station does not exist",
			expectExitCode: 1,
		},
		{
			name:           "Trigger parsing error (duplicate)",
			args:           []string{"network_error_test.map", "apple", "banana", "5"},
			expectedErr:    "Error: Line 4: Duplicate station names: apple",
			expectExitCode: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cmd := exec.Command("./pathfinder.exe", tc.args...)
			var stderr bytes.Buffer
			cmd.Stderr = &stderr

			err := cmd.Run()

			if err == nil {
				t.Errorf("Expected an error (exit code), but command succeeded")
			}

			output := stderr.String()
			if !strings.Contains(output, tc.expectedErr) {
				t.Errorf("Expected stderr to contain %q, but got %q", tc.expectedErr, output)
			}
		})
	}
}

func TestPerformanceAndLogic(t *testing.T) {
	// These tests ensure it doesn't hang excessively and finishes correctly.
	tests := []struct {
		name          string
		args          []string
		expectedTurns int
	}{
		{
			name:          "Jungle to Desert (10 trains)",
			args:          []string{"network_jungle.map", "jungle", "desert", "10"},
			expectedTurns: 8,
		},
		{
			name:          "Beginning to Terminus (20 trains)",
			args:          []string{"network_terminus.map", "beginning", "terminus", "20"},
			expectedTurns: 11,
		},
		{
			name:          "Small to Large (9 trains)",
			args:          []string{"network_small_large.map", "small", "large", "9"},
			expectedTurns: 8,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cmd := exec.Command("./pathfinder.exe", tc.args...)
			var stdout bytes.Buffer
			cmd.Stdout = &stdout

			err := cmd.Run()
			if err != nil {
				t.Fatalf("Command failed unexpectedly: %v", err)
			}

			output := strings.TrimSpace(stdout.String())
			lines := strings.Split(output, "\n")
			if len(lines) != tc.expectedTurns {
				t.Errorf("Expected %d turns, got %d turns", tc.expectedTurns, len(lines))
			}
		})
	}
}
