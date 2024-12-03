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

func isSafe(row []int) bool {
	if len(row) < 2 {
		return false
	}

	grad := 0
	for j := 1; j < len(row); j++ {
		value := row[j]
		diff := value - row[j-1]
		if diff > 3 || diff < -3 || diff == 0 {
			return false
		}
		if j == 1 {
			grad = diff
		} else {
			if (grad > 0 && diff < 0) || (grad < 0 && diff > 0) {
				return false
			}
		}
	}
	return true
}

func canBeMadeSafe(row []int) bool {
	for j := 0; j < len(row); j++ {
		newRow := append([]int{}, row[:j]...)
		newRow = append(newRow, row[j+1:]...)
		if isSafe(newRow) {
			return true
		}
	}
	return false
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
			if err == nil {
				row = append(row, num)
			}
		}
		data = append(data, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Iterate through each line of the data variable and through each value of the rows
	for i, row := range data {
		if isSafe(row) || canBeMadeSafe(row) {
			safe++
		}
		fmt.Printf("Row %d: %v\n", i, row)
	}

	fmt.Println("Safe reports: ", safe)
	endTime := time.Now()
	fmt.Println("Time taken: ", endTime.Sub(startTime))
}
