package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		blockMap := generateBlockMap(line)
		finalMap := moveBlocks(blockMap)
		checksum := calculateChecksum(finalMap)
		fmt.Println("Checksum:", checksum)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	endTime := time.Now()
	fmt.Println("Time taken:", endTime.Sub(startTime))
}

func generateBlockMap(diskMap string) string {
	var result strings.Builder
	fileID := 0
	isFile := true

	for i := 0; i < len(diskMap); i++ {
		length, err := strconv.Atoi(string(diskMap[i]))
		if err != nil {
			fmt.Println("Error converting string to number:", err)
			return ""
		}

		if isFile {
			result.WriteString(strings.Repeat(strconv.Itoa(fileID), length))
			fileID++
		} else {
			result.WriteString(strings.Repeat(".", length))
		}
		isFile = !isFile
	}

	return result.String()
}

func moveBlocks(blockMap string) string {
	runes := []rune(blockMap)
	for {
		moved := false
		// Find the rightmost used block
		for i := len(runes) - 1; i >= 0; i-- {
			if runes[i] != '.' {
				// Find the leftmost free block
				leftmostFree := strings.IndexRune(string(runes), '.')
				if leftmostFree != -1 && leftmostFree < i {
					// Swap the rightmost used block with the leftmost free block
					runes[leftmostFree], runes[i] = runes[i], '.'
					moved = true
					break
				}
			}
		}
		if !moved {
			break
		}
	}
	return string(runes)
}

func calculateChecksum(blockMap string) *big.Int {
	checksum := big.NewInt(0)
	for i, block := range blockMap {
		if block != '.' {
			fileID, err := strconv.Atoi(string(block))
			if err != nil {
				fmt.Println("Error converting string to number:", err)
				return big.NewInt(0)
			}
			position := big.NewInt(int64(i))
			id := big.NewInt(int64(fileID))
			product := new(big.Int).Mul(position, id)
			checksum.Add(checksum, product)
		}
	}
	return checksum
}
