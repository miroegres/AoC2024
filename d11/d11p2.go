package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func hashStones(stones []int) string {
	hash := sha256.New()
	for _, stone := range stones {
		hash.Write([]byte(strconv.Itoa(stone)))
	}
	return hex.EncodeToString(hash.Sum(nil))
}

func transformStone(stone int, memo map[int][]int) []int {
	if val, found := memo[stone]; found {
		return val
	}

	var transformed []int
	if stone == 0 {
		transformed = []int{1}
	} else if len(strconv.Itoa(stone))%2 == 0 {
		mid := len(strconv.Itoa(stone)) / 2
		left, _ := strconv.Atoi(strconv.Itoa(stone)[:mid])
		right, _ := strconv.Atoi(strconv.Itoa(stone)[mid:])
		transformed = []int{left, right}
	} else {
		transformed = []int{stone * 2024}
	}

	memo[stone] = transformed
	return transformed
}

func processChunks(stones []int, memo map[int][]int, cache map[string][]int) []int {
	chunkSize := 1000
	var newStones []int

	for i := 0; i < len(stones); i += chunkSize {
		end := i + chunkSize
		if end > len(stones) {
			end = len(stones)
		}
		chunk := stones[i:end]
		hash := hashStones(chunk)
		if cachedResult, found := cache[hash]; found {
			newStones = append(newStones, cachedResult...)
		} else {
			var transformedChunk []int
			for _, stone := range chunk {
				transformedChunk = append(transformedChunk, transformStone(stone, memo)...)
			}
			cache[hash] = transformedChunk
			newStones = append(newStones, transformedChunk...)
		}
	}

	return newStones
}

func main() {
	startTime := time.Now()
	filename := "inputT.txt"

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

	memo := make(map[int][]int)
	cache := make(map[string][]int)

	for iteration := 1; iteration <= 75; iteration++ {
		stones = processChunks(stones, memo, cache)

		if iteration <= 8 {
			fmt.Printf("Iteration %d: %v\n", iteration, stones)
		} else {
			fmt.Printf("Iteration %d: Number of stones = %d\n", iteration, len(stones))
		}
	}

	endTime := time.Now()
	fmt.Println("Time taken: ", endTime.Sub(startTime))
}
