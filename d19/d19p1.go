package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func canFormDesign(patterns []string, design string) bool {
	if design == "" {
		return true
	}
	for _, pattern := range patterns {
		if strings.HasPrefix(design, pattern) {
			if canFormDesign(patterns, design[len(pattern):]) {
				return true
			}
		}
	}
	return false
}

func main() {
	startTime := time.Now()
	filename := "input.txt"
	sum := 0

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

	// Read designs and check if they can be formed
	for scanner.Scan() {
		design := scanner.Text()
		fmt.Println("Checking design:", design) // Debug print
		if canFormDesign(patterns, design) {
			sum++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("count: ", sum)

	endTime := time.Now()
	fmt.Println("Time taken: ", endTime.Sub(startTime))
}
