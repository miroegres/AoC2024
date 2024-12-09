package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
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

	// Read the file into a 2D slice
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	for _, line := range lines {
		blockMap := generateBlockMap(line)
		//fmt.Println("Initial block map:", blockMap)
		finalBlockMap := moveBlocks(blockMap)
		//fmt.Println("Final block map:", finalBlockMap)
		checksum := calculateChecksum(finalBlockMap)
		fmt.Println("Checksum:", checksum)
	}

	endTime := time.Now()
	fmt.Println("Time taken:", endTime.Sub(startTime))
}

func generateBlockMap(diskMap string) []*big.Int {
	var blockMap []*big.Int
	fileID := big.NewInt(0)
	isFile := true

	for i := 0; i < len(diskMap); i++ {
		length, err := strconv.Atoi(string(diskMap[i]))
		if err != nil {
			fmt.Println("Error converting string to number:", err)
			return nil
		}

		if isFile {
			for j := 0; j < length; j++ {
				blockMap = append(blockMap, new(big.Int).Set(fileID))
			}
			fileID.Add(fileID, big.NewInt(1))
		} else {
			for j := 0; j < length; j++ {
				blockMap = append(blockMap, big.NewInt(-1))
			}
		}
		isFile = !isFile
	}

	return blockMap
}

func moveBlocks(blockMap []*big.Int) []*big.Int {
	for {
		moved := false
		// Find the rightmost used block
		for i := len(blockMap) - 1; i >= 0; i-- {
			if blockMap[i].Cmp(big.NewInt(-1)) != 0 {
				// Find the leftmost free block
				leftmostFree := -1
				for j, block := range blockMap {
					if block.Cmp(big.NewInt(-1)) == 0 {
						leftmostFree = j
						break
					}
				}
				if leftmostFree != -1 && leftmostFree < i {
					// Swap the rightmost used block with the leftmost free block
					blockMap[leftmostFree], blockMap[i] = blockMap[i], big.NewInt(-1)
					moved = true
					//fmt.Println("After moving:", blockMap)
					break
				}
			}
		}
		if !moved {
			break
		}
	}
	return blockMap
}

func calculateChecksum(blockMap []*big.Int) *big.Int {
	checksum := big.NewInt(0)
	for i, block := range blockMap {
		if block.Cmp(big.NewInt(-1)) != 0 {
			position := big.NewInt(int64(i))
			product := new(big.Int).Mul(position, block)
			checksum.Add(checksum, product)
		}
	}
	return checksum
}
