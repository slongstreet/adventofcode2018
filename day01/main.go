package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	frequencies, err := processInput("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %s", err)
	}

	// Part 1 - calculate final frequency
	total := 0
	for _, f := range frequencies {
		total += f
	}

	fmt.Printf("Final frequency value = %d\n", total)

	// Part 2 - calculate first repeat frequency value
	// Test Cases 01-04:
	fmt.Printf("TC01: %d\n", findFirstRepeatFrequency([]int{1, -1}))
	fmt.Printf("TC02: %d\n", findFirstRepeatFrequency([]int{3, 3, 4, -2, -4}))
	fmt.Printf("TC03: %d\n", findFirstRepeatFrequency([]int{-6, 3, 8, 5, -6}))
	fmt.Printf("TC04: %d\n", findFirstRepeatFrequency([]int{7, 7, -2, -7, -4}))

	firstRepeat := findFirstRepeatFrequency(frequencies)
	fmt.Printf("First repeat value = %d\n", firstRepeat)
}

// Reads whitespace-separated integers from a file.
func processInput(filepath string) ([]int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var frequencies []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return frequencies, err
		}

		frequencies = append(frequencies, f)
	}

	return frequencies, scanner.Err()
}

// Finds the first repeated frequency total from looping through the frequency inputs.
func findFirstRepeatFrequency(frequencies []int) int {
	dict := map[int]int{0: 1} // map to store count of values
	total := 0
	firstRepeat := 0
	found := false

	for {
		for _, f := range frequencies {
			total += f
			dict[total]++
			if dict[total] == 2 {
				firstRepeat = total
				found = true
				break
			}
		}

		if found {
			break // we found our first repeat, so we can stop looping
		}
	}

	return firstRepeat
}
