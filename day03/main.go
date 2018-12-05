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

	var grid [1000][1000]int
	for _, sq := range squares {
		for i := sq.X; i < sq.X+sq.Width; i++ {
			for j := sq.Y; j < sq.Y+sq.Height; j++ {
				grid[i][j]++
				if grid[i][j] > 1 {
					multipleClaims[strconv.Itoa(i)+","+strconv.Itoa(j)] = true
				}
			}
		}
	}

	fmt.Printf("Square inches with multiple claims: %d.\n", len(multipleClaims))
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
