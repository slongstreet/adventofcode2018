package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	rawInput, _ := loadInputFromFile("input.txt")
	var polymer string

	// Part 1 - process all reactions until polymer is stable, then calculate length
	polymer = processPolymer(rawInput)

	fmt.Printf("Part one: polymer length: %d\n", len(polymer))

	// Part 2 - find shortest polymer if one unit type is removed
	shortestLength := len(rawInput)
	alpha := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < len(alpha); i++ {
		polymer = rawInput // reset polymer to default state
		polymer = strings.Replace(polymer, string(alpha[i]), "", -1)
		polymer = strings.Replace(polymer, string(unicode.ToUpper(rune(alpha[i]))), "", -1)

		polymer = processPolymer(polymer)

		if len(polymer) < shortestLength {
			shortestLength = len(polymer)
		}
	}

	fmt.Printf("Part two: polymer length: %d\n", shortestLength)
}

// Reads a string from the specified input file path.
func loadInputFromFile(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var text string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text = scanner.Text()
	}

	return text, scanner.Err()
}

// Process all reactions until polymer is stable.
// Return resulting polymer.
func processPolymer(polymer string) string {
	var reacted bool
	for {
		reacted, polymer = processPolymerReaction(polymer)
		if !reacted {
			break
		}
	}

	return polymer
}

// Remove the first reacting runes from the polymer string.
// If runes were removed, returns true; otherwise false.
// The string return value is the resulting polymer.
func processPolymerReaction(polymer string) (bool, string) {
	var r, neighbor rune
	for i := 1; i < len(polymer); i++ {
		r = rune(polymer[i])
		neighbor = rune(polymer[i-1])
		if willReact(r, neighbor) {
			separator := string(neighbor) + string(r)
			parts := strings.Split(polymer, separator)
			return true, strings.Join(parts, "")
		}
	}

	return false, polymer
}

// Returns true if two runes would react; false otherwise.
func willReact(r rune, neighbor rune) bool {
	sameRune := unicode.ToUpper(r) == unicode.ToUpper(neighbor)
	return sameRune && unicode.IsUpper(r) != unicode.IsUpper(neighbor)
}
