package main

import (
	"math"
)

// Calculate the maximum X and Y coordinates for a supplied list of points.
func calculateMaximumDimensions(points []point) (int, int) {
	var maxX, maxY int
	for i := 0; i < len(points); i++ {
		if points[i].X > maxX {
			maxX = points[i].X
		}
		if points[i].Y > maxY {
			maxY = points[i].Y
		}
	}

	return maxX, maxY
}

// Calculate the minimum X and Y coordinates for a supplied list of points.
func calculateMinimumDimensions(points []point) (int, int) {
	minX := 4096
	minY := 4096

	for i := 0; i < len(points); i++ {
		if points[i].X < minX {
			minX = points[i].X
		}
		if points[i].Y < minY {
			minY = points[i].Y
		}
	}

	return minX, minY
}

// Calculate the Manhattan distance between two points.
func calculateDistanceBetweenPoints(a point, b point) int {
	return int(math.Abs(float64(a.X)-float64(b.X))) + int(math.Abs(float64(a.Y)-float64(b.Y)))
}

// Find the closest point index to the source point.  Returns -1 if
// src point matches one of the supplied points or there is a tie.
// Returns the index of the point closest to the src point otherwise.
func findClosestPointIndex(src point, points []point) int {
	closestDistance := 4096
	closestPointIndex := -1

	for i := 0; i < len(points); i++ {
		dist := calculateDistanceBetweenPoints(src, points[i])
		if dist == 0 {
			return i // source point equals one of the candidate points
		} else if dist == closestDistance {
			closestPointIndex = -1 // multiple points are equidistant from here
		} else if dist < closestDistance {
			closestDistance = dist
			closestPointIndex = i
		}
	}

	return closestPointIndex
}
