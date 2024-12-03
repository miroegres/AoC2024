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

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	startTime := time.Now()
	filename := "input.txt"
	sum := 0
	sum2 := 0

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

	reMul := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	reDo := regexp.MustCompile(`do\(\)`)
	reDont := regexp.MustCompile(`don't\(\)`)

	// Regular expression to match "mul(x,y)" where x and y are numbers with 1 to 3 digits
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

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

		// Find all substrings that match the pattern
		matches := re.FindAllString(line, -1)
		substrings = append(substrings, matches...)

		mulMatches := reMul.FindAllString(line, -1)
		doMatches := reDo.FindAllString(line, -1)
		dontMatches := reDont.FindAllString(line, -1)

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

	//d3p1
	// Iterate through each line of the data and find the mul(x,y) strings
	for _, substr := range substrings {
		// Extract numbers from the substring
		reNumbers := regexp.MustCompile(`\d{1,3}`)
		numbers := reNumbers.FindAllString(substr, -1)
		if len(numbers) == 2 {
			num1, err1 := strconv.Atoi(numbers[0])
			num2, err2 := strconv.Atoi(numbers[1])
			if err1 == nil && err2 == nil {
				sum += num1 * num2
				//fmt.Printf("Substring: %s, Product: %d\n", substr, product)
			}
		}
	}

	//d3p2
	for _, substr := range substrings2 {
		// Extract numbers from the substring
		reNumbers := regexp.MustCompile(`\d{1,3}`)
		numbers := reNumbers.FindAllString(substr, -1)
		if len(numbers) == 2 {
			num1, err1 := strconv.Atoi(numbers[0])
			num2, err2 := strconv.Atoi(numbers[1])
			if err1 == nil && err2 == nil {
				sum2 += num1 * num2
				//fmt.Printf("Substring: %s, Product: %d\n", substr, product)
			}
		}
	}

	//fmt.Println(substrings)
	fmt.Println("Sum: ", sum)
	fmt.Println("Sum2: ", sum2)

	endTime := time.Now()
	fmt.Println("Cas: ", endTime.Sub(startTime))
}
