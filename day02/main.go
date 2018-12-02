package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Test Cases 01-07:
	evaluateTestCase("TC01", "abcdef", 0, 0)
	evaluateTestCase("TC02", "bababc", 1, 1)
	evaluateTestCase("TC03", "abbcde", 1, 0)
	evaluateTestCase("TC04", "abcccd", 0, 1)
	evaluateTestCase("TC05", "aabcdd", 1, 0)
	evaluateTestCase("TC06", "abcdee", 1, 0)
	evaluateTestCase("TC07", "ababab", 0, 1)

	// Part 1 - Calculate checksum of input file
	lines, err := processInput("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %s", err)
	}

	doublesCount := 0
	triplesCount := 0
	for _, str := range lines {
		d, t := analyzeString(str)
		doublesCount += d
		triplesCount += t
	}

	checksum := doublesCount * triplesCount
	fmt.Printf("Checksum = %d\n", checksum)
}

// Reads whitespace-separated strings from a file.
func processInput(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var inputs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	return inputs, scanner.Err()
}

// Analyzes input string for doubles and triples.
// First return arg is count of doubles; second is count of triples.
func analyzeString(input string) (int, int) {
	doubles := 0
	triples := 0

	// Count the frequency of runes in the input string.
	dict := make(map[rune]int)
	for _, r := range input {
		dict[r]++
	}

	// Tally up the doubles and triples.
	for _, v := range dict {
		if v == 2 {
			doubles = 1 // Per the instructions, count only once
		} else if v == 3 {
			triples = 1 // Per the instructions, count only once
		}
	}

	return doubles, triples
}

// Evaluate and print test case results.
func evaluateTestCase(testCaseID string, input string, expectedDoubles int, expectedTriples int) {
	d, t := analyzeString(input)
	var pass string
	if d == expectedDoubles && t == expectedTriples {
		pass = "PASS"
	} else {
		pass = "FAIL"
	}

	fmt.Printf("%s: 2x: %d, 3x: %d => %s\n", testCaseID, d, t, pass)
}
