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
		finalBlockMap := moveFiles(blockMap)
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

func moveFiles(blockMap []*big.Int) []*big.Int {
	fileLengths := make(map[int]int)
	for _, fileID := range blockMap {
		if fileID.Cmp(big.NewInt(-1)) != 0 {
			id := int(fileID.Int64())
			fileLengths[id]++
		}
	}

	maxFileID := len(fileLengths) - 1

	for fileID := maxFileID; fileID >= 0; fileID-- {
		fileLength := fileLengths[fileID]
		leftmostFreeSpanStart := -1
		leftmostFreeSpanLength := 0

		for i := 0; i < len(blockMap); i++ {
			if blockMap[i].Cmp(big.NewInt(-1)) == 0 {
				if leftmostFreeSpanStart == -1 {
					leftmostFreeSpanStart = i
				}
				leftmostFreeSpanLength++
			} else {
				leftmostFreeSpanStart = -1
				leftmostFreeSpanLength = 0
			}

			if leftmostFreeSpanLength == fileLength {
				break
			}
		}

		// Ensure the free span is to the left of the current file
		currentFileStart := -1
		for i, block := range blockMap {
			if block.Cmp(big.NewInt(int64(fileID))) == 0 {
				currentFileStart = i
				break
			}
		}

		if leftmostFreeSpanStart != -1 && leftmostFreeSpanStart < currentFileStart && leftmostFreeSpanLength >= fileLength {
			// Move the file to the leftmost free span
			for i := len(blockMap) - 1; i >= 0; i-- {
				if blockMap[i].Cmp(big.NewInt(int64(fileID))) == 0 {
					blockMap[i] = big.NewInt(-1)
				}
			}

			for i := leftmostFreeSpanStart; i < leftmostFreeSpanStart+fileLength; i++ {
				blockMap[i] = big.NewInt(int64(fileID))
			}

			//fmt.Println("After moving file", fileID, ":", blockMap)
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
