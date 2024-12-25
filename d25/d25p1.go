package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// parseHeights converts a schematic into a list of heights.
func parseHeights(schematic []string, isLock bool) []int {
	heights := make([]int, len(schematic[0]))
	for col := 0; col < len(schematic[0]); col++ {
		height := 0
		if isLock {
			// Count '#' downward for locks
			for row := 0; row < len(schematic); row++ {
				if schematic[row][col] == '#' {
					height++
				} else {
					break
				}
			}
			heights[col] = height - 1 // Adjust to start from 0
		} else {
			// Count '#' upward for keys
			for row := len(schematic) - 1; row >= 0; row-- {
				if schematic[row][col] == '#' {
					height++
				} else {
					break
				}
			}
			heights[col] = height - 1 // Adjust to start from 0
		}
	}
	return heights
}

// isLock determines if a schematic represents a lock.
func isLock(schematic []string) bool {
	for _, char := range schematic[0] {
		if char != '#' {
			return false
		}
	}
	return true
}

// fitsTogether checks if a lock and key can fit together without overlapping.
func fitsTogether(lock, key []int) bool {
	for i := 0; i < len(lock); i++ {
		if lock[i]+key[i] >= 6 { // Sum exceeding 5 means overlap
			return false
		}
	}
	return true
}

func main() {
	startTime := time.Now()

	filename := "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file into lines
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Parse locks and keys from the input
	var locks, keys [][]string
	var current []string

	for _, line := range lines {
		if line == "" {
			if len(current) > 0 {
				if isLock(current) {
					locks = append(locks, current)
				} else {
					keys = append(keys, current)
				}
				current = nil
			}
			continue
		}
		current = append(current, line)
	}
	if len(current) > 0 {
		if isLock(current) {
			locks = append(locks, current)
		} else {
			keys = append(keys, current)
		}
	}

	// Convert schematics to heights
	lockHeights := make([][]int, len(locks))
	keyHeights := make([][]int, len(keys))

	for i, lock := range locks {
		lockHeights[i] = parseHeights(lock, true)
		fmt.Printf("Lock %d heights: %v\n", i, lockHeights[i])
	}
	for i, key := range keys {
		keyHeights[i] = parseHeights(key, false)
		fmt.Printf("Key %d heights: %v\n", i, keyHeights[i])
	}

	// Check all lock/key pairs
	sum := 0
	for i, lock := range lockHeights {
		for j, key := range keyHeights {
			if fitsTogether(lock, key) {
				fmt.Printf("Lock %d and Key %d fit together\n", i, j)
				sum++
			}
		}
	}

	fmt.Println("Number of fitting lock/key pairs:", sum)
	endTime := time.Now()
	fmt.Println("Execution time:", endTime.Sub(startTime))
}
