package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func transformStones(stones []int) []int {
	var newStones []int
	for _, stone := range stones {
		if stone == 0 {
			newStones = append(newStones, 1)
		} else if len(strconv.Itoa(stone))%2 == 0 {
			mid := len(strconv.Itoa(stone)) / 2
			left, _ := strconv.Atoi(strconv.Itoa(stone)[:mid])
			right, _ := strconv.Atoi(strconv.Itoa(stone)[mid:])
			newStones = append(newStones, left, right)
		} else {
			newStones = append(newStones, stone*2024)
		}
	}
	return newStones
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

	var stones []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		for _, num := range numbers {
			stone, _ := strconv.Atoi(num)
			stones = append(stones, stone)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	for iteration := 1; iteration <= 75; iteration++ {
		stones = transformStones(stones)
		if iteration <= 5 {
			fmt.Printf("Iteration %d: %v\n", iteration, stones)
		} else {
			fmt.Printf("Iteration %d: Number of stones = %d\n", iteration, len(stones))
		}
	}

	endTime := time.Now()
	fmt.Println("Time taken: ", endTime.Sub(startTime))
}
