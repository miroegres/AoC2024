package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Point struct {
	x, y int
}

var directions = []Point{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

func main() {
	startTime := time.Now()
	filename := "input.txt"
	cheatTime := 100 // Variable to specify the amount of picoseconds to save by cheating

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var racetrack [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		racetrack = append(racetrack, []rune(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	start, end := findStartEnd(racetrack)
	shortestPath := bfs(racetrack, start, end)
	fmt.Println("Shortest path length without cheating:", shortestPath)

	bestCheatsCount, cheatsMap := findBestCheats(racetrack, start, end, cheatTime)
	fmt.Println("\nRacetrack with Cheats Marked:")
	printRacetrack(cheatsMap)

	fmt.Printf("\nNumber of possible cheats that save at least %d picoseconds: %d\n", cheatTime, bestCheatsCount)

	endTime := time.Now()
	fmt.Println("Execution time:", endTime.Sub(startTime))
}

func findStartEnd(racetrack [][]rune) (Point, Point) {
	var start, end Point
	for y, row := range racetrack {
		for x, cell := range row {
			if cell == 'S' {
				start = Point{x, y}
			} else if cell == 'E' {
				end = Point{x, y}
			}
		}
	}
	return start, end
}

func bfs(racetrack [][]rune, start, end Point) int {
	queue := []Point{start}
	visited := make(map[Point]bool)
	visited[start] = true
	distance := make(map[Point]int)
	distance[start] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			return distance[current]
		}

		for _, dir := range directions {
			next := Point{current.x + dir.x, current.y + dir.y}
			if isValidMove(racetrack, next) && !visited[next] {
				queue = append(queue, next)
				visited[next] = true
				distance[next] = distance[current] + 1
			}
		}
	}
	return -1 // No path found
}

func isValidMove(racetrack [][]rune, point Point) bool {
	if point.y < 0 || point.y >= len(racetrack) || point.x < 0 || point.x >= len(racetrack[0]) {
		return false
	}
	return racetrack[point.y][point.x] != '#'
}

func bfsWithCheat(racetrack [][]rune, start, end, cheatStart, cheatEnd Point) int {
	queue := []Point{start}
	visited := make(map[Point]bool)
	visited[start] = true
	distance := make(map[Point]int)
	distance[start] = 0
	cheated := false

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			return distance[current]
		}

		for _, dir := range directions {
			next := Point{current.x + dir.x, current.y + dir.y}
			if isValidMove(racetrack, next) && !visited[next] {
				queue = append(queue, next)
				visited[next] = true
				distance[next] = distance[current] + 1
			} else if !cheated && current == cheatStart && next == cheatEnd {
				queue = append(queue, next)
				visited[next] = true
				distance[next] = distance[current] + 1
				cheated = true
			}
		}
	}
	return -1 // No path found
}

func findBestCheats(racetrack [][]rune, start, end Point, cheatTime int) (int, [][]rune) {
	originalPathLength := bfs(racetrack, start, end)
	cheatsCount := 0
	cheatsMap := make([][]rune, len(racetrack))
	for i := range racetrack {
		cheatsMap[i] = make([]rune, len(racetrack[i]))
		copy(cheatsMap[i], racetrack[i])
	}

	for y := range racetrack {
		for x := range racetrack[0] {
			if racetrack[y][x] == '#' {
				continue
			}

			for _, dir := range directions {
				cheatStart := Point{x, y}
				cheatEnd1 := Point{x + dir.x, y + dir.y}
				cheatEnd2 := Point{x + 2*dir.x, y + 2*dir.y}

				if isValidMove(racetrack, cheatEnd2) {
					pathLengthWithCheat1 := bfsWithCheat(racetrack, start, end, cheatStart, cheatEnd1)
					pathLengthWithCheat2 := bfsWithCheat(racetrack, start, end, cheatStart, cheatEnd2)

					if pathLengthWithCheat1 > 0 && originalPathLength-pathLengthWithCheat1 >= cheatTime {
						cheatsCount++
						cheatsMap[cheatStart.y][cheatStart.x] = 'C'
						cheatsMap[cheatEnd1.y][cheatEnd1.x] = 'C'
					}

					if pathLengthWithCheat2 > 0 && originalPathLength-pathLengthWithCheat2 >= cheatTime {
						cheatsCount++
						cheatsMap[cheatStart.y][cheatStart.x] = 'C'
						cheatsMap[cheatEnd2.y][cheatEnd2.x] = 'C'
					}
				}
			}
		}
	}

	return cheatsCount, cheatsMap
}

func printRacetrack(racetrack [][]rune) {
	for _, row := range racetrack {
		fmt.Println(string(row))
	}
}
