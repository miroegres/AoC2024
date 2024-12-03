package main

import (
	"bufio"
	"fmt"
	"os"
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

// Function to check if a row is strictly increasing or decreasing
func checkRow(row []int) bool {
	isIncreasing := true
	isDecreasing := true

	for j, value := range row {
		if j > 0 {
			diff := absInt(value - row[j-1])
			if diff == 0 || diff > 3 {
				return false
			}

			if !isIncreasing && !isDecreasing {
				return false
			}

			if value > row[j-1] {
				isDecreasing = false
			} else if value < row[j-1] {
				isIncreasing = false
			}
		}
	}

	return isIncreasing || isDecreasing
}

func main() {
	startTime := time.Now()

	filename := "input.txt"
	safe := 0

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var data [][]int

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Fields(line)
		var row []int
		for _, col := range columns {
			num, err := strconv.Atoi(col)
			if err != nil {
				fmt.Printf("Error parsing number in line: %s\n", line)
				continue
			}
			row = append(row, num)
		}
		data = append(data, row)
	}

	// Iterate through each line of the data variable and through each value of the rows
	for i, row := range data {
		isSafe := false

		fmt.Printf("Row %d: ", i)
		fmt.Println(row)

		// Check if removing any single value makes the row safe
		for j := 0; !isSafe && j < len(row); j++ {
			modifiedRow := append(row[:j], row[j+1:]...)
			if checkRow(modifiedRow) {
				isSafe = true
			}
		}

		var checkResult int
		if isSafe {
			checkResult = 1
		} else {
			checkResult = 0
		}

		fmt.Printf("Line %d check=%d\n", i, checkResult)
		fmt.Println("--------------------------")
	}

	fmt.Println("Safe reports:", safe)

	endTime := time.Now()
	fmt.Println("Elapsed time:", endTime.Sub(startTime))
}
