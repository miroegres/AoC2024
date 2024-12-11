package main

import (
	"fmt"
	"strconv"
	"time"
)

// Memoization cache
var memo = make(map[string]int64)

// Recursive function to calculate the number of stones after a given number of iterations
func functionRec(stone int64, iterations int) int64 {
	// Create a unique key for the memoization cache
	key := fmt.Sprintf("%d_%d", stone, iterations)
	if val, found := memo[key]; found {
		return val
	}

	if iterations == 0 {
		return 1
	}

	var result int64
	if stone == 0 {
		result = functionRec(1, iterations-1)
	} else if len(strconv.FormatInt(stone, 10))%2 == 0 {
		mid := len(strconv.FormatInt(stone, 10)) / 2
		left, _ := strconv.ParseInt(strconv.FormatInt(stone, 10)[:mid], 10, 64)
		right, _ := strconv.ParseInt(strconv.FormatInt(stone, 10)[mid:], 10, 64)
		result = functionRec(left, iterations-1) + functionRec(right, iterations-1)
	} else {
		result = functionRec(stone*2024, iterations-1)
	}

	memo[key] = result
	return result
}

func main() {
	startTime := time.Now()

	// Initial stones
	stones := []int64{814, 1183689, 0, 1, 766231, 4091, 93836, 46}

	// Number of iterations
	iterations := 75

	// Calculate the total number of stones after the given number of iterations
	var totalStones int64
	for _, stone := range stones {
		totalStones += functionRec(stone, iterations)
	}

	fmt.Printf("Number of stones after %d iterations: %d\n", iterations, totalStones)

	endTime := time.Now()
	fmt.Println("Time taken: ", endTime.Sub(startTime))
}
