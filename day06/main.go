package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	X, Y int
}

type cell struct {
	Pos   point
	Value int
}

func main() {
	testPoints := loadTestCaseData()
	maxArea := calculateLargestArea(testPoints)
	fmt.Printf("TC01: largest area = %d\n", maxArea)

	points, _ := loadInputFromFile("input.txt")
	maxArea = calculateLargestArea(points)
	fmt.Printf("Largest non-infinite area: %d\n", maxArea)
}

func loadInputFromFile(filepath string) ([]point, error) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var points []point
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ", ")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])

		points = append(points, point{x, y})
	}

	return points, scanner.Err()
}

func loadTestCaseData() []point {
	return []point{
		point{1, 1},
		point{1, 6},
		point{8, 3},
		point{3, 4},
		point{5, 5},
		point{8, 9},
	}
}

func calculateLargestArea(points []point) int {
	// determine the size of the grid
	maxX, maxY := calculateMaximumDimensions(points)
	minX, minY := calculateMinimumDimensions(points)

	// populate the grid with cells referencing the nearest point
	grid := make([][]cell, maxX+1)
	for i := range grid {
		grid[i] = make([]cell, maxY+1)
	}

	areaMap := make(map[int]int)
	for i := 0; i <= maxX; i++ {
		for j := 0; j <= maxY; j++ {
			var p = point{i, j}
			index := findClosestPointIndex(p, points)
			grid[i][j] = cell{p, index}

			if index >= 0 {
				areaMap[index]++
			}
		}
	}

	// determine the largest area that isn't infinite
	maxArea := 0
	for k, v := range areaMap {
		if v > maxArea && !isInfinitePoint(points[k], k, grid, minX, minY, maxX, maxY) {
			maxArea = v
		}
	}

	return maxArea
}

func isInfinitePoint(p point, index int, grid [][]cell, minX int, minY int, maxX int, maxY int) bool {
	// if point is on the edge of point coordinates, it's infinite
	if p.X == maxX || p.Y == maxY || p.X == minX || p.Y == minY {
		return true // point is on an edge
	}

	// if point value is the same heading due north, east, south, or west at the edge, it's infinite
	if grid[p.X][maxY].Value == index || grid[p.X][minY].Value == index || grid[maxX][p.Y].Value == index || grid[minX][p.Y].Value == index {
		return true
	}

	return false
}
