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

func simulateSecretNumbers(initialSecrets []int, iterations int) []int {
	results := make([]int, len(initialSecrets))

	for i, secret := range initialSecrets {
		for j := 0; j < iterations; j++ {
			secret = nextSecretNumber(secret)
		}
		results[i] = secret
	}

	return results
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
	finalSecrets := simulateSecretNumbers(initialSecrets, 2000)

	// Calculate the sum of the final secret numbers
	for _, secret := range finalSecrets {
		sum += secret
	}

	fmt.Println("Sum of the 2000th secret number generated by each buyer:", sum)

	endTime := time.Now()
	fmt.Println("Time taken:", endTime.Sub(startTime))
}