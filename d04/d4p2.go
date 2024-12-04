package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	count := 0
	rows := len(grid)
	cols := len(grid[0])

	// Check for patterns
	for i := 1; i < rows-1; i++ { // Start from the second line and stop before the last line
		for j := 1; j < cols-1; j++ { // Skip the first and last characters in each line
			if grid[i][j] == 'A' {
				condition1 := true

				condition2 := i-1 >= 0 && j-1 >= 0 && i+1 < rows && j+1 < cols &&
					grid[i-1][j-1] == 'M' && grid[i+1][j+1] == 'S'
				condition3 := i-1 >= 0 && j-1 >= 0 && i+1 < rows && j+1 < cols &&
					grid[i-1][j-1] == 'S' && grid[i+1][j+1] == 'M'
				condition4 := i-1 >= 0 && j+1 < cols && i+1 < rows && j-1 >= 0 &&
					grid[i-1][j+1] == 'M' && grid[i+1][j-1] == 'S'
				condition5 := i-1 >= 0 && j+1 < cols && i+1 < rows && j-1 >= 0 &&
					grid[i-1][j+1] == 'S' && grid[i+1][j-1] == 'M'

				if condition1 && (condition2 || condition3) && (condition4 || condition5) {
					count++
				}
			}
		}
	}

	fmt.Println("Total occurrences:", count)
}
