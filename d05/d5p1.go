package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()
	filename := "input.txt"
	sum := 0
	fixedSum := 0

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var Rules [][]string
	var updates [][]int
	var validUpdates [][]int
	var invalidUpdates [][]int
	var fixedUpdates [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			rule := strings.Split(line, "|")
			Rules = append(Rules, rule)
		} else if strings.Contains(line, ",") {
			numStrings := strings.Split(line, ",")
			var nums []int
			for _, numStr := range numStrings {
				num, err := strconv.Atoi(strings.TrimSpace(numStr))
				if err != nil {
					fmt.Println("Error converting string ", num, " to int:", err)
					return
				}
				nums = append(nums, num)
			}
			updates = append(updates, nums)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Loop A: Go through all the updates
	for update_ind, update := range updates {
		valid := true

		// Loop B: Check each item in the update
		for i := 0; i < len(update); i++ {
			for _, rule := range Rules {
				first, err1 := strconv.Atoi(rule[0])
				second, err2 := strconv.Atoi(rule[1])
				if err1 != nil || err2 != nil {
					fmt.Println("Error converting rule to int:", err1, err2)
					return
				}
				firstIndex := indexOf(update, first)
				secondIndex := indexOf(update, second)
				if firstIndex != -1 && secondIndex != -1 && firstIndex >= secondIndex {
					valid = false
					fmt.Println("update #", update_ind, " pages:", update, " - rule:", rule, " failed")
					break
				}
			}
			if !valid {
				break
			}
		}

		if valid {
			validUpdates = append(validUpdates, update)
		} else {
			invalidUpdates = append(invalidUpdates, update)
		}
	}

	validCount := len(validUpdates)

	// Calculate the sum of the middle values of each valid update
	for _, update := range validUpdates {
		middleIndex := len(update) / 2
		sum += update[middleIndex]
	}

	// Fix invalid updates
	for _, update := range invalidUpdates {
		fixedUpdate := fixUpdate(update, Rules)
		fixedUpdates = append(fixedUpdates, fixedUpdate)
	}

	// Calculate the sum of the middle values of each fixed update
	for _, update := range fixedUpdates {
		middleIndex := len(update) / 2
		fixedSum += update[middleIndex]
	}

	fmt.Println("updates: ", len(updates))
	fmt.Println("valid updates: ", validCount)
	fmt.Println("sum of middle values: ", sum)
	fmt.Println(fixedUpdates)
	fmt.Println("sum of middle values of fixed updates: ", fixedSum)

	endTime := time.Now()
	fmt.Println("Cas: ", endTime.Sub(startTime))
}

// Helper function to find the index of an element in a slice
func indexOf(slice []int, element int) int {
	for i, v := range slice {
		if v == element {
			return i
		}
	}
	return -1
}

// Helper function to fix an update according to the rules
func fixUpdate(update []int, rules [][]string) []int {
	for _, rule := range rules {
		first, err1 := strconv.Atoi(rule[0])
		second, err2 := strconv.Atoi(rule[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Error converting rule to int:", err1, err2)
			return update
		}
		firstIndex := indexOf(update, first)
		secondIndex := indexOf(update, second)
		if firstIndex != -1 && secondIndex != -1 && firstIndex >= secondIndex {
			// Swap the elements to satisfy the rule
			update[firstIndex], update[secondIndex] = update[secondIndex], update[firstIndex]
		}
	}
	return update
}
