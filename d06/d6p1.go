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

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

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

	// Convert lines to a 2D slice (matrix)
	matrix := make([][]rune, len(lines))
	visited := make([][]rune, len(lines))
	for i, line := range lines {
		matrix[i] = []rune(line)
		visited[i] = make([]rune, len(line))
		copy(visited[i], matrix[i])
	}

	// Find the initial position and direction of the guard
	var x, y int
	var direction rune
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == '^' {
				x, y = i, j
				direction = '^'
				break
			}
		}
	}

	// Directions: N, E, S, W
	directions := []rune{'^', '>', 'v', '<'}
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}
	dirIndex := 0 // Start facing North

	// Function to turn right
	turnRight := func() {
		dirIndex = (dirIndex + 1) % 4
		direction = directions[dirIndex]
		matrix[x][y] = direction
	}

	// Simulate the guard's movement
	steps := 0
	visited[x][y] = 'X' // Mark the starting position as visited
	for {
		nx, ny := x+dx[dirIndex], y+dy[dirIndex]

		// Check if the guard has reached the edge of the map
		if nx < 0 || ny < 0 || nx >= len(matrix) || ny >= len(matrix[0]) {
			break
		}

		// Check if the next cell is an obstacle
		if matrix[nx][ny] == '#' {
			turnRight()
		} else {
			// Move forward
			x, y = nx, ny
			if visited[x][y] != 'X' {
				steps++
				visited[x][y] = 'X'
			}
		}
	}

	// Count the number of 'X' in the visited map
	xCount := 0
	for _, row := range visited {
		for _, cell := range row {
			if cell == 'X' {
				xCount++
			}
		}
	}

	fmt.Println("Spaces stepped on:", steps)
	fmt.Println("Visited cells map:")
	for _, line := range visited {
		fmt.Println(string(line))
	}
	fmt.Println("Number of 'X' in the map:", xCount)
	endTime := time.Now()
	fmt.Println("Time taken:", endTime.Sub(startTime))
}
