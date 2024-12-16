package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"time"
)

// Direction represents a movement direction with its cost
type Direction struct {
	dx, dy, cost int
}

// PriorityQueue implements heap.Interface and holds Items
type PriorityQueue []*Item

// Item represents a state in the priority queue
type Item struct {
	cost      int
	x, y      int
	direction string
	index     int
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
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
	var maze [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		maze = append(maze, []rune(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Directions: (dx, dy, cost)
	directions := map[string]Direction{
		"E": {0, 1, 1},
		"W": {0, -1, 1},
		"N": {-1, 0, 1},
		"S": {1, 0, 1},
	}

	// Rotation costs
	rotationCost := 1000

	// Find the start and end positions
	var start, end [2]int
	for i := range maze {
		for j := range maze[i] {
			if maze[i][j] == 'S' {
				start = [2]int{i, j}
			} else if maze[i][j] == 'E' {
				end = [2]int{i, j}
			}
		}
	}

	// Priority queue for Dijkstra's algorithm
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{cost: 0, x: start[0], y: start[1], direction: "E"})
	visited := make(map[[3]interface{}]bool)

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		cost, x, y, direction := item.cost, item.x, item.y, item.direction

		if [2]int{x, y} == end {
			fmt.Printf("The lowest score a Reindeer could possibly get is %d.\n", cost)
			break
		}

		if visited[[3]interface{}{x, y, direction}] {
			continue
		}

		visited[[3]interface{}{x, y, direction}] = true

		// Move forward in the current direction
		dir := directions[direction]
		nx, ny := x+dir.dx, y+dir.dy
		if nx >= 0 && nx < len(maze) && ny >= 0 && ny < len(maze[0]) && maze[nx][ny] != '#' {
			heap.Push(pq, &Item{cost: cost + dir.cost, x: nx, y: ny, direction: direction})
		}

		// Rotate clockwise and counterclockwise
		for newDirection := range directions {
			if newDirection != direction {
				heap.Push(pq, &Item{cost: cost + rotationCost, x: x, y: y, direction: newDirection})
			}
		}
	}

	endTime := time.Now()
	fmt.Println("Time taken: ", endTime.Sub(startTime))
}
