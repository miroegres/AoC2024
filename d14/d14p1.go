package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Robot struct {
	px, py, vx, vy int
}

func parseLine(line string) Robot {
	parts := strings.Split(line, " ")
	pos := strings.Split(parts[0][2:], ",")
	vel := strings.Split(parts[1][2:], ",")
	px, _ := strconv.Atoi(pos[0])
	py, _ := strconv.Atoi(pos[1])
	vx, _ := strconv.Atoi(vel[0])
	vy, _ := strconv.Atoi(vel[1])
	return Robot{px, py, vx, vy}
}

func main() {
	startTime := time.Now()
	filename := "input.txt"
	seconds := 100
	width := 101  //101
	height := 103 //103

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var robots []Robot
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		robots = append(robots, parseLine(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Initialize the grid
	grid := make([][]int, height)
	for i := range grid {
		grid[i] = make([]int, width)
	}

	// Simulate the motion of the robots
	for _, robot := range robots {
		x, y := robot.px, robot.py
		vx, vy := robot.vx, robot.vy

		for t := 0; t < seconds; t++ {
			x = (x + vx + width) % width
			y = (y + vy + height) % height
		}

		grid[y][x]++
	}

	// Count the number of robots in each quadrant
	q1, q2, q3, q4 := 0, 0, 0, 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if y == height/2 || x == width/2 {
				continue
			}
			if y < height/2 && x < width/2 {
				q1 += grid[y][x]
			} else if y < height/2 && x > width/2 {
				q2 += grid[y][x]
			} else if y > height/2 && x < width/2 {
				q3 += grid[y][x]
			} else if y > height/2 && x > width/2 {
				q4 += grid[y][x]
			}
		}
	}

	// Calculate the safety factor
	safetyFactor := q1 * q2 * q3 * q4

	fmt.Printf("Safety factor: %d\n", safetyFactor)
	fmt.Printf("robots in quadrants: %d, %d, %d, %d\n", q1, q2, q3, q4)

	endTime := time.Now()
	fmt.Println("Time taken:", endTime.Sub(startTime))
}
