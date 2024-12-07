package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalSum := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			continue
		}

		expectedResult, err := strconv.Atoi(parts[0])
		if err != nil {
			continue
		}

		operands := strings.Fields(parts[1])
		if isValidEquation(expectedResult, operands) {
			totalSum += expectedResult
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Total sum of valid equations:", totalSum)
}

func isValidEquation(expectedResult int, operands []string) bool {
	n := len(operands)
	if n == 0 {
		return false
	}

	// Convert operands to integers
	nums := make([]int, n)
	for i, operand := range operands {
		num, err := strconv.Atoi(operand)
		if err != nil {
			return false
		}
		nums[i] = num
	}

	// Generate all possible equations with +, *, and ||
	return evaluate(nums, expectedResult, 0, nums[0])
}

func evaluate(nums []int, expectedResult, index, currentResult int) bool {
	if index == len(nums)-1 {
		return currentResult == expectedResult
	}

	// Try addition
	if evaluate(nums, expectedResult, index+1, currentResult+nums[index+1]) {
		return true
	}

	// Try multiplication
	if evaluate(nums, expectedResult, index+1, currentResult*nums[index+1]) {
		return true
	}

	// Try concatenation
	concatenated := concatenate(currentResult, nums[index+1])
	if evaluate(nums, expectedResult, index+1, concatenated) {
		return true
	}

	return false
}

func concatenate(a, b int) int {
	concatenated, _ := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	return concatenated
}
