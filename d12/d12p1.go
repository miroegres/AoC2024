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

	// Convert lines to a 2D slice of runes
	garden := make([][]rune, len(lines))
	for i, line := range lines {
		garden[i] = []rune(line)
	}

	// Calculate areas and perimeters
	totalPrice := 0
	visited := make([][]bool, len(garden))
	for i := range visited {
		visited[i] = make([]bool, len(garden[i]))
	}

	for i := range garden {
		for j := range garden[i] {
			if !visited[i][j] {
				area, perimeter := exploreRegion(garden, visited, i, j, garden[i][j])
				price := area * perimeter
				totalPrice += price
				fmt.Printf("Plant %c: Area = %d, Perimeter = %d, Price = %d\n", garden[i][j], area, perimeter, price)
			}
		}
	}

	endTime := time.Now()
	fmt.Println("Time: ", endTime.Sub(startTime))
	fmt.Println("Total Price: ", totalPrice)
}

func exploreRegion(garden [][]rune, visited [][]bool, x, y int, plant rune) (int, int) {
	if x < 0 || x >= len(garden) || y < 0 || y >= len(garden[x]) || visited[x][y] || garden[x][y] != plant {
		return 0, 0
	}

	visited[x][y] = true
	area := 1
	perimeter := 0

	// Check all four directions
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, d := range directions {
		nx, ny := x+d[0], y+d[1]
		if nx < 0 || nx >= len(garden) || ny < 0 || ny >= len(garden[nx]) || garden[nx][ny] != plant {
			perimeter++
		} else {
			a, p := exploreRegion(garden, visited, nx, ny, plant)
			area += a
			perimeter += p
		}
	}

	return area, perimeter
}
