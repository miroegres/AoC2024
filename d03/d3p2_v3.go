package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	startTime := time.Now()
	filename := "input_.txt"
	sum2 := 0

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var substrings2 []string

	// Regular expressions to match "mul(x,y)", "do()", and "don't()"
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Find all substrings that match the patterns
		Matches := re.FindAllString(line, -1)

		// Add matches to substrings2 in the order they appear
		substrings2 = append(substrings2, Matches...)

	}

	//d3p2
	skipMultiplication := false
	for _, substr := range substrings2 {
		if substr == "don't()" {
			skipMultiplication = true
		} else if substr == "do()" {
			skipMultiplication = false
		} else if !skipMultiplication {
			// Extract numbers from the substring
			reNumbers := regexp.MustCompile(`\d{1,3}`)
			numbers := reNumbers.FindAllString(substr, -1)
			if len(numbers) == 2 {
				num1, err1 := strconv.Atoi(numbers[0])
				num2, err2 := strconv.Atoi(numbers[1])
				if err1 == nil && err2 == nil {
					sum2 += num1 * num2
					fmt.Printf("Substring: %s\n", substr)
				}
			}
		}
	}

	//fmt.Println(substrings)
	fmt.Println("Sum2: ", sum2)

	endTime := time.Now()
	fmt.Println("Cas: ", endTime.Sub(startTime))
}
