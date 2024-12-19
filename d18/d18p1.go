package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"time"
)

const (
	numBytesToSimulate = 1024
	maxX               = 70
	maxY               = 70
)

var directions = [][2]int{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

type Point struct {
	x, y int
}

type Item struct {
	point    Point
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func isWithinGrid(x, y int) bool {
	return x >= 0 && x <= maxX && y >= 0 && y <= maxY
}

func simulateFallingBytes(bytePositions []Point) [][]rune {
	grid := make([][]rune, maxY+1)
	for i := range grid {
		grid[i] = make([]rune, maxX+1)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	for _, pos := range bytePositions {
		if isWithinGrid(pos.x, pos.y) {
			grid[pos.y][pos.x] = '#'
		}
	}
	return grid
}

func findMinSteps(grid [][]rune) (int, [][]rune) {
	start := Point{0, 0}
	goal := Point{maxX, maxY}
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{point: start, priority: 0})
	visited := make(map[Point]bool)
	gScore := make(map[Point]int)
	gScore[start] = 0
	cameFrom := make(map[Point]Point)

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item).point
		if current == goal {
			path := reconstructPath(cameFrom, current)
			for _, p := range path {
				if p != start && p != goal {
					grid[p.y][p.x] = 'O'
				}
			}
			return gScore[current], grid
		}
		visited[current] = true
		for _, dir := range directions {
			neighbor := Point{current.x + dir[0], current.y + dir[1]}
			if isWithinGrid(neighbor.x, neighbor.y) && !visited[neighbor] && grid[neighbor.y][neighbor.x] == '.' {
				tentativeGScore := gScore[current] + 1
				if score, exists := gScore[neighbor]; !exists || tentativeGScore < score {
					gScore[neighbor] = tentativeGScore
					fScore := tentativeGScore + abs(goal.x-neighbor.x) + abs(goal.y-neighbor.y)
					heap.Push(pq, &Item{point: neighbor, priority: fScore})
					cameFrom[neighbor] = current
				}
			}
		}
	}
	return -1, grid
}

func reconstructPath(cameFrom map[Point]Point, current Point) []Point {
	var path []Point
	for current != (Point{}) {
		path = append(path, current)
		current = cameFrom[current]
	}
	// Reverse the path
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	startTime := time.Now()
	filename := "input.txt"

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var bytePositions []Point
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var x, y int
		fmt.Sscanf(scanner.Text(), "%d,%d", &x, &y)
		bytePositions = append(bytePositions, Point{x, y})
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Simulate the specified number of bytes falling onto the memory space
	bytePositions = bytePositions[:numBytesToSimulate]

	// Update the grid with the falling bytes
	grid := simulateFallingBytes(bytePositions)

	// Find the minimum number of steps to reach the exit and draw the path
	minSteps, updatedGrid := findMinSteps(grid)

	fmt.Println("The minimum number of steps needed to reach the exit is", minSteps)
	fmt.Println("Grid after simulating the falling bytes and drawing the path:")
	for _, row := range updatedGrid {
		for _, cell := range row {
			fmt.Print(string(cell))
		}
		fmt.Println()
	}

	endTime := time.Now()
	fmt.Println("Time taken:", endTime.Sub(startTime))
}
