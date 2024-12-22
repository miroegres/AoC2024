package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func nextSecretNumber(secret int) int {
	// Step 1: Multiply by 64, mix, and prune
	secret = (secret ^ (secret * 64)) % 16777216

	// Step 2: Divide by 32, mix, and prune
	secret = (secret ^ (secret / 32)) % 16777216

	// Step 3: Multiply by 2048, mix, and prune
	secret = (secret ^ (secret * 2048)) % 16777216

	return secret
}

func simulateSecretNumbers(initialSecrets []int, iterations int) [][]int {
	results := make([][]int, len(initialSecrets))

	for i, secret := range initialSecrets {
		prices := make([]int, iterations)
		for j := 0; j < iterations; j++ {
			secret = nextSecretNumber(secret)
			prices[j] = secret % 10 // Get the ones digit as the price
		}
		results[i] = prices
	}

	return results
}

func findBestSequence(prices [][]int) (bestSequence []int, maxBananas int) {
	sequences := generateAllSequences()
	for _, seq := range sequences {
		bananas := 0
		for _, buyerPrices := range prices {
			for i := 0; i <= len(buyerPrices)-len(seq)-1; i++ {
				if matchesSequence(buyerPrices[i:], seq) {
					bananas += buyerPrices[i+len(seq)]
					break
				}
			}
		}
		if bananas > maxBananas {
			maxBananas = bananas
			bestSequence = seq
		}
	}
	return
}

func generateAllSequences() [][]int {
	var sequences [][]int
	for a := -9; a <= 9; a++ {
		for b := -9; b <= 9; b++ {
			for c := -9; c <= 9; c++ {
				for d := -9; d <= 9; d++ {
					sequences = append(sequences, []int{a, b, c, d})
				}
			}
		}
	}
	return sequences
}

func matchesSequence(prices []int, sequence []int) bool {
	for i := 0; i < len(sequence); i++ {
		if i+1 >= len(prices) || prices[i+1]-prices[i] != sequence[i] {
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

	// Read the file into a slice of integers
	var initialSecrets []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return
		}
		initialSecrets = append(initialSecrets, num)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Simulate the creation of 2000 new secret numbers
	prices := simulateSecretNumbers(initialSecrets, 2000)

	// Find the best sequence of four price changes
	bestSequence, maxBananas := findBestSequence(prices)

	fmt.Println("Best sequence of price changes:", bestSequence)
	fmt.Println("Maximum bananas:", maxBananas)

	endTime := time.Now()
	fmt.Println("Time taken:", endTime.Sub(startTime))
}
