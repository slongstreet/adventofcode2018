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
	polymer, _ := loadInputFromFile("input.txt")

	// Part 1 - process all reactions until polymer is stable, then calculate length
	var reacted bool
	for {
		reacted, polymer = processPolymer(polymer)
		if !reacted {
			break
		}
	}

	fmt.Printf("Polymer length: %d\n", len(polymer))
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

// Remove the first reacting runes from the polymer string.
// If runes were removed, returns true; otherwise false.
// The string return value is the resulting polymer.
func processPolymer(polymer string) (bool, string) {
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
