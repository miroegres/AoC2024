package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Machine struct {
	ax, ay, bx, by, px, py int
}

func extractCoordinates(line string) (int, int) {
	parts := strings.Split(line, ", ")
	xPart := strings.Split(parts[0], "X+")[1]
	yPart := strings.Split(parts[1], "Y+")[1]
	x, _ := strconv.Atoi(xPart)
	y, _ := strconv.Atoi(yPart)
	return x, y
}

func extractPrize(line string) (int, int) {
	parts := strings.Split(line, ", ")
	xPart := strings.Split(parts[0], "X=")[1]
	yPart := strings.Split(parts[1], "Y=")[1]
	x, _ := strconv.Atoi(xPart)
	y, _ := strconv.Atoi(yPart)
	return x, y
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

	var machines []Machine
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Button A:") {
			ax, ay := extractCoordinates(line)
			scanner.Scan()
			bx, by := extractCoordinates(scanner.Text())
			scanner.Scan()
			px, py := extractPrize(scanner.Text())
			machines = append(machines, Machine{ax, ay, bx, by, px, py})
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Summarize the input data
	/*
		for i, machine := range machines {
			fmt.Printf("Machine %d:\n", i+1)
			fmt.Printf("  Button A increments: X+%d, Y+%d\n", machine.ax, machine.ay)
			fmt.Printf("  Button B increments: X+%d, Y+%d\n", machine.bx, machine.by)
			fmt.Printf("  Prize location: X=%d, Y=%d\n", machine.px, machine.py)
		}
	*/

	totalTokens := 0
	prizesWon := 0

	// Calculate the minimum tokens needed for each machine
	for _, machine := range machines {
		minTokens := -1
		for a := 0; a <= 100; a++ {
			for b := 0; b <= 100; b++ {
				if a*machine.ax+b*machine.bx == machine.px && a*machine.ay+b*machine.by == machine.py {
					tokens := a*3 + b
					if minTokens == -1 || tokens < minTokens {
						minTokens = tokens
					}
				}
			}
		}
		if minTokens != -1 {
			// fmt.Printf("Machine %d: Minimum tokens needed = %d\n", i+1, minTokens)
			totalTokens += minTokens
			prizesWon++
		} else {
			// fmt.Printf("Machine %d: No solution found\n", i+1)
		}
	}

	fmt.Printf("Total tokens needed to win all possible prizes: %d\n", totalTokens)
	fmt.Printf("Total prizes won: %d\n", prizesWon)

	endTime := time.Now()
	fmt.Println("Time taken:", endTime.Sub(startTime))
}
