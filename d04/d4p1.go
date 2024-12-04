package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	startTime := time.Now()
	filename := "input.txt"
	outputFilename := "output.txt"
	sum := 0

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	outputFile, err := os.Create(outputFilename)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)

	// Read the file into a 2D slice
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Define the patterns to search for
	patterns := []string{"XMAS"}

	// Function to check for patterns in all directions
	checkPattern := func(x, y, dx, dy int, pattern string) bool {
		for i := 0; i < len(pattern); i++ {
			nx, ny := x+dx*i, y+dy*i
			if nx < 0 || ny < 0 || nx >= len(lines) || ny >= len(lines[nx]) || lines[nx][ny] != pattern[i] {
				return false
			}
		}
		return true
	}

	// Search for patterns in all directions
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			for _, pattern := range patterns {
				directions := [][2]int{{0, 1}, {1, 0}, {1, 1}, {1, -1}, {0, -1}, {-1, 0}, {-1, -1}, {-1, 1}}
				for _, dir := range directions {
					if checkPattern(i, j, dir[0], dir[1], pattern) {
						sum++
						writer.WriteString(fmt.Sprintf("%s found at [%d,%d]\n", pattern, i, j))
					}
				}
			}
		}
	}

	writer.Flush()

	fmt.Println("count: ", sum)

	endTime := time.Now()
	fmt.Println("Cas: ", endTime.Sub(startTime))
}
