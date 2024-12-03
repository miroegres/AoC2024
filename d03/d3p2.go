package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()
	filename := "input.txt"

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var data [][]int
	var substrings []string
	var substrings2 []string
	addToSubstrings2 := true

	// Regular expressions to match "mul(x,y)", "do()", and "don't()"
	reMul := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	reDo := regexp.MustCompile(`do\(\)`)
	reDont := regexp.MustCompile(`don't\(\)`)

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Fields(line)
		var row []int
		for _, col := range columns {
			num, err := strconv.Atoi(col)
			if err == nil {
				row = append(row, num)
			}
		}
		data = append(data, row)

		// Find all substrings that match the patterns
		mulMatches := reMul.FindAllString(line, -1)
		doMatches := reDo.FindAllString(line, -1)
		dontMatches := reDont.FindAllString(line, -1)

		substrings = append(substrings, mulMatches...)

		// Handle "do()" and "don't()"
		if len(doMatches) > 0 {
			addToSubstrings2 = true
		}
		if len(dontMatches) > 0 {
			addToSubstrings2 = false
		}

		// Add "mul(x,y)" to substrings2 based on the current state
		if addToSubstrings2 {
			substrings2 = append(substrings2, mulMatches...)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Substrings found: ", substrings)
	fmt.Println("Substrings2 found: ", substrings2)

	// Parse and multiply numbers in substrings
	for _, substr := range substrings2 {
		// Extract numbers from the substring
		reNumbers := regexp.MustCompile(`\d{1,3}`)
		numbers := reNumbers.FindAllString(substr, -1)
		if len(numbers) == 2 {
			num1, err1 := strconv.Atoi(numbers[0])
			num2, err2 := strconv.Atoi(numbers[1])
			if err1 == nil && err2 == nil {
				product := num1 * num2
				fmt.Printf("Substring: %s, Product: %d\n", substr, product)
			}
		}
	}

	endTime := time.Now()
	fmt.Println("Time taken: ", endTime.Sub(startTime))
}
