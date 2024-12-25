package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	// Read the file into a slice of connections
	var connections [][2]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "-")
		connections = append(connections, [2]string{parts[0], parts[1]})
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Create a map to store the network
	network := make(map[string]map[string]bool)
	for _, conn := range connections {
		a, b := conn[0], conn[1]
		if network[a] == nil {
			network[a] = make(map[string]bool)
		}
		if network[b] == nil {
			network[b] = make(map[string]bool)
		}
		network[a][b] = true
		network[b][a] = true
	}

	// Find all sets of three inter-connected computers, considering order doesn't matter
	setsOfThree := make(map[[3]string]bool)
	for a := range network {
		for b := range network[a] {
			for c := range network[b] {
				if c != a && network[a][c] {
					set := [3]string{a, b, c}
					sort.Strings(set[:])
					setsOfThree[set] = true
				}
			}
		}
	}

	// Print all sets of three inter-connected computers
	fmt.Println("Sets of three inter-connected computers:")
	for set := range setsOfThree {
		fmt.Println(set)
	}

	// Count how many sets contain at least one computer with a name that starts with 't'
	count := 0
	for set := range setsOfThree {
		for _, computer := range set {
			if strings.HasPrefix(computer, "t") {
				count++
				break
			}
		}
	}

	fmt.Println("Count:", count)

	endTime := time.Now()
	fmt.Println("Time:", endTime.Sub(startTime))
}
