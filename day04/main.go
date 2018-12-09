package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	timeLayout = "2006-01-02 15:04"

	beginsShift = iota
	fallsAsleep
	wakesUp
)

type logEntry struct {
	timestamp time.Time
	guardID   int
	event     int
}

func main() {
	logEntries, _ := processInput("input.txt")
	fmt.Printf("log entry count: %d\n", len(logEntries))

	// Sort the log entries by timestamp
	sort.Slice(logEntries, func(i, j int) bool {
		return logEntries[i].timestamp.Before(logEntries[j].timestamp)
	})

	// Part 1 - determine who sleeps the most, and which minute they sleep most often
	sleepSum := make(map[int]float64)
	minuteMap := make(map[int]map[int]int)

	currentID := 0
	var sleepStart time.Time

	for _, entry := range logEntries {
		switch entry.event {
		case beginsShift:
			currentID = entry.guardID
		case fallsAsleep:
			sleepStart = entry.timestamp
		case wakesUp:
			sleepDuration := entry.timestamp.Sub(sleepStart).Minutes()
			sleepSum[currentID] += sleepDuration
			sleepEnd := float64(sleepStart.Minute()) + sleepDuration

			for i := float64(sleepStart.Minute()); i < sleepEnd; i++ {
				if minuteMap[currentID] == nil {
					minuteMap[currentID] = make(map[int]int)
				}
				minuteMap[currentID][int(math.Mod(i, 60))]++
			}
		}
	}

	// who slept the most?
	var biggestSleeper int
	var maxSleep float64
	for k, v := range sleepSum {
		if v > maxSleep {
			biggestSleeper = k
			maxSleep = v
		}
	}

	// which minute did they sleep the most in?
	sleepiestMinute := 0
	maxMinuteSleep := 0
	for k, v := range minuteMap[biggestSleeper] {
		if v > maxMinuteSleep {
			sleepiestMinute = k
			maxMinuteSleep = v
		}
	}

	fmt.Printf("Biggest sleeper: %d\n", biggestSleeper)
	fmt.Printf("Sleepiest minute: %d\n", sleepiestMinute)
	fmt.Printf("Part one answer: %d\n", biggestSleeper*sleepiestMinute)

	// Part 2 - which guard is most frequently asleep on the same minute?
	biggestSleeper = 0
	sleepiestMinute = 0
	maxTimesAsleep := 0
	for id, submap := range minuteMap {
		for k, v := range submap {
			if v > maxTimesAsleep {
				maxTimesAsleep = v
				sleepiestMinute = k
				biggestSleeper = id
			}
		}
	}

	fmt.Println()
	fmt.Printf("Most frequently asleep: %d\n", biggestSleeper)
	fmt.Printf("Sleepiest minute: %d\n", sleepiestMinute)
	fmt.Printf("Part two answer: %d\n", biggestSleeper*sleepiestMinute)
}

func processInput(filepath string) ([]logEntry, error) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var entries []logEntry
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entries = append(entries, parseInputLine(scanner.Text()))
	}

	return entries, scanner.Err()
}

func parseInputLine(input string) logEntry {
	// Sample formats:
	// [1518-11-03 00:04] Guard #3323 begins shift
	// [1518-03-01 00:10] falls asleep
	// [1518-06-14 00:45] wakes up
	parts := strings.Split(input, "] ")

	var entry logEntry
	entry.timestamp, _ = time.Parse(timeLayout, strings.TrimLeft(parts[0], "["))

	parts = strings.Split(parts[1], " ") // split RHS into parts
	if parts[0] == "Guard" {
		entry.event = beginsShift
		entry.guardID, _ = strconv.Atoi(strings.TrimLeft(parts[1], "#"))
	} else if parts[0] == "falls" {
		entry.event = fallsAsleep
	} else {
		entry.event = wakesUp
	}

	return entry
}
