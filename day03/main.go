package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type square struct {
	ID            int
	X, Y          int
	Width, Height int
}

func main() {
	squares, _ := processInput("input.txt")
	fmt.Printf("Loaded %d squares...\n", len(squares))

	// Part 1 - calculate number of square inches with multiple claims
	multipleClaims := make(map[string]bool) // store multi claims as we go to avoid multiple enumerations
	overlapCheck := make([]int, len(squares)+1)

	var grid [1000][1000]int
	for _, sq := range squares {
		for i := sq.X; i < sq.X+sq.Width; i++ {
			for j := sq.Y; j < sq.Y+sq.Height; j++ {
				if grid[i][j] == 0 {
					grid[i][j] = sq.ID // if grid location is unused, store our ID here.
				} else {
					// grid location is shared, so mark it as having multiple claims
					multipleClaims[strconv.Itoa(i)+","+strconv.Itoa(j)] = true
					overlapCheck[max(grid[i][j], 0)] = 1 // store that this square has an overlap...
					overlapCheck[sq.ID] = 1              // ...and the current square too.
					grid[i][j] = -1
				}
			}
		}
	}

	fmt.Printf("Square inches with multiple claims: %d.\n", len(multipleClaims))

	// Part 2 - find the only claim which doesn't overlap
	var nonoverlappingID int
	for id, val := range overlapCheck {
		if val == 0 {
			nonoverlappingID = id
			break
		}
	}

	fmt.Printf("Non-overlapping claim ID: %d\n", nonoverlappingID)
}

// Reads and parses square definitions from a file.
func processInput(filepath string) ([]square, error) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var squares []square
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		squares = append(squares, parseInputLine(scanner.Text()))
	}

	return squares, scanner.Err()
}

func parseInputLine(input string) square {
	// Sample input format:
	// #17 @ 878,268: 13x11
	parts := strings.Split(input, " ")

	var sq square
	sq.ID, _ = strconv.Atoi(strings.Split(parts[0], "#")[1])

	coordinates := strings.Split(parts[2], ",")
	sq.X, _ = strconv.Atoi(coordinates[0])
	sq.Y, _ = strconv.Atoi(strings.Split(coordinates[1], ":")[0])

	dimensions := strings.Split(parts[3], "x")
	sq.Width, _ = strconv.Atoi(dimensions[0])
	sq.Height, _ = strconv.Atoi(dimensions[1])

	return sq
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
