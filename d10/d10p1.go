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

	// Convert lines to a 2D slice of integers
	mapData := make([][]int, len(lines))
	for i, line := range lines {
		mapData[i] = make([]int, len(line))
		for j, char := range line {
			mapData[i][j] = int(char - '0')
		}
	}

	// Find all valid trails and sum their scores
	totalScore, trailCount := findTrails(mapData)
	fmt.Println("Total score of all trailheads:", totalScore)
	fmt.Println("Number of trails found:", trailCount)

	endTime := time.Now()
	fmt.Println("Time taken:", endTime.Sub(startTime))
}

func findTrails(mapData [][]int) (int, int) {
	rows := len(mapData)
	cols := len(mapData[0])
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // Up, Down, Left, Right

	isValid := func(x, y int) bool {
		return x >= 0 && x < rows && y >= 0 && y < cols
	}

	var dfs func(x, y int, visited map[[2]int]bool) int
	dfs = func(x, y int, visited map[[2]int]bool) int {
		if mapData[x][y] == 9 {
			visited[[2]int{x, y}] = true
			return 1
		}
		count := 0
		for _, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]
			if isValid(nx, ny) && mapData[nx][ny] == mapData[x][y]+1 && !visited[[2]int{nx, ny}] {
				visited[[2]int{x, y}] = true
				count += dfs(nx, ny, visited)
			}
		}
		return count
	}

	totalScore := 0
	trailCount := 0
	for i := range mapData {
		for j := range mapData[i] {
			if mapData[i][j] == 0 {
				trailCount++
				visited := make(map[[2]int]bool)
				score := dfs(i, j, visited)
				totalScore += score
				//fmt.Printf("Trailhead at (%d, %d) with score %d:\n", i, j, score)
				printTrailMap(mapData, visited)
			}
		}
	}

	return totalScore, trailCount
}

func printTrailMap(mapData [][]int, visited map[[2]int]bool) {
	for i := range mapData {
		for j := range mapData[i] {
			if visited[[2]int{i, j}] {
				//fmt.Printf("%d ", mapData[i][j])
			} else {
				//fmt.Print(". ")
			}
		}
		//fmt.Println()
	}
	//fmt.Println()
}
