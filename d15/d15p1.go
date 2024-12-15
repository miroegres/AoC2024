package main

import (
	"bufio"
	"fmt"
	"os"
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

	var mapLines []string
	var moves string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "#") || strings.Contains(line, ".") || strings.Contains(line, "O") || strings.Contains(line, "@") {
			mapLines = append(mapLines, line)
		} else {
			moves += line
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the initial map
	fmt.Println("Initial Map:")
	for _, line := range mapLines {
		fmt.Println(line)
	}

	// Simulate the movements and get the final map
	finalMap := simulateMovements(mapLines, moves)

	// Calculate the GPS coordinates of the boxes from the final map
	totalGPS := calculateGPSCoordinates(finalMap)
	fmt.Println("Total GPS coordinates sum:", totalGPS)

	endTime := time.Now()
	fmt.Println("Time taken:", endTime.Sub(startTime))
}

func simulateMovements(mapLines []string, moves string) []string {
	// Find the initial position of the robot
	var robotX, robotY int
	for y, line := range mapLines {
		if x := strings.Index(line, "@"); x != -1 {
			robotX, robotY = x, y
			break
		}
	}

	// Process each move
	for _, move := range moves {
		fmt.Printf("Move %c:\n", move)
		switch move {
		case '^':
			if canMove(mapLines, robotX, robotY-1) {
				robotY--
			} else if canPush(mapLines, robotX, robotY-1, 0, -1) {
				pushBoxes(mapLines, robotX, robotY-1, 0, -1)
				robotY--
			}
		case 'v':
			if canMove(mapLines, robotX, robotY+1) {
				robotY++
			} else if canPush(mapLines, robotX, robotY+1, 0, 1) {
				pushBoxes(mapLines, robotX, robotY+1, 0, 1)
				robotY++
			}
		case '<':
			if canMove(mapLines, robotX-1, robotY) {
				robotX--
			} else if canPush(mapLines, robotX-1, robotY, -1, 0) {
				pushBoxes(mapLines, robotX-1, robotY, -1, 0)
				robotX--
			}
		case '>':
			if canMove(mapLines, robotX+1, robotY) {
				robotX++
			} else if canPush(mapLines, robotX+1, robotY, 1, 0) {
				pushBoxes(mapLines, robotX+1, robotY, 1, 0)
				robotX++
			}
		}

		// Update the map with the new robot position
		mapLines = updateMap(mapLines, robotX, robotY)

		// Print the updated map
		for _, line := range mapLines {
			fmt.Println(line)
		}
	}

	return mapLines
}

func canMove(mapLines []string, x, y int) bool {
	// Check if the position is within bounds and not a wall or box
	if y < 0 || y >= len(mapLines) || x < 0 || x >= len(mapLines[y]) {
		return false
	}
	return mapLines[y][x] == '.'
}

func canPush(mapLines []string, x, y, dx, dy int) bool {
	// Check if the robot can push all boxes in the direction until a wall or another box is encountered
	if y < 0 || y >= len(mapLines) || x < 0 || x >= len(mapLines[y]) || mapLines[y][x] == '#' {
		return false
	}
	if mapLines[y][x] == '.' {
		return true
	}
	if mapLines[y][x] == 'O' {
		return canPush(mapLines, x+dx, y+dy, dx, dy)
	}
	return false
}

func pushBoxes(mapLines []string, x, y, dx, dy int) {
	// Recursively push all boxes in the direction until a wall or another box is encountered
	if mapLines[y][x] == 'O' {
		nextX, nextY := x+dx, y+dy
		if canPush(mapLines, nextX, nextY, dx, dy) {
			pushBoxes(mapLines, nextX, nextY, dx, dy)
			mapLines[nextY] = replaceAtIndex(mapLines[nextY], 'O', nextX)
			mapLines[y] = replaceAtIndex(mapLines[y], '.', x)
		}
	}
}

func updateMap(mapLines []string, robotX, robotY int) []string {
	// Create a copy of the map
	newMap := make([]string, len(mapLines))
	copy(newMap, mapLines)

	// Clear the old robot position
	for y, line := range newMap {
		newMap[y] = strings.Replace(line, "@", ".", 1)
	}

	// Set the new robot position
	newMap[robotY] = newMap[robotY][:robotX] + "@" + newMap[robotY][robotX+1:]

	return newMap
}

func replaceAtIndex(in string, r rune, i int) string {
	if i < 0 || i >= len(in) {
		return in
	}
	out := []rune(in)
	out[i] = r
	return string(out)
}

func calculateGPSCoordinates(mapLines []string) int {
	totalGPS := 0
	for y, line := range mapLines {
		for x, char := range line {
			if char == 'O' {
				gps := 100*y + x
				totalGPS += gps
			}
		}
	}
	return totalGPS
}
