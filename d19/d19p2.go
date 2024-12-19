package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func countWaysToFormDesign(patterns []string, design string, memo map[string]int) int {
	if design == "" {
		return 1
	}
	if val, found := memo[design]; found {
		return val
	}
	totalWays := 0
	for _, pattern := range patterns {
		if strings.HasPrefix(design, pattern) {
			totalWays += countWaysToFormDesign(patterns, design[len(pattern):], memo)
		}
	}
	memo[design] = totalWays
	return totalWays
}

func main() {
	startTime := time.Now()
	filename := "input.txt"
	totalWays := 0

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read towel patterns
	var patterns []string
	if scanner.Scan() {
		line := scanner.Text()
		patterns = strings.Split(line, ", ")
	}

	fmt.Println("Patterns:", patterns) // Debug print

	// Read designs and count ways to form them
	for scanner.Scan() {
		design := scanner.Text()
		if design == "" {
			continue // Omit empty line as design
		}
		fmt.Println("Checking design:", design) // Debug print
		memo := make(map[string]int)
		ways := countWaysToFormDesign(patterns, design, memo)
		fmt.Printf("Ways to form '%s': %d\n", design, ways) // Debug print
		totalWays += ways
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Total ways to form all designs: ", totalWays)

	endTime := time.Now()
	fmt.Println("Time taken: ", endTime.Sub(startTime))
}
