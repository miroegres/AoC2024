package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	startTime := time.Now()
	filename := "input.txt"
	safe := 0

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var data [][]int

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Fields(line)
		var row []int
		for _, col := range columns {
			num, err := strconv.Atoi(col)
			if err == nil {
				row = append(row, num)
			}
		}
		data = append(data, row)
	}

	//d1p1
	// Iterate through each line of the data variable and through each value of the rows
	for i, row := range data {
		if len(row) < 2 {
			continue // Skip rows with less than 2 values
		}

		grad := 0  // not incr nor decr
		check := 1 // line seems OK
		fmt.Printf("Row %d: ", i)
		fmt.Println(row)
		for j := 1; j < len(row); j++ {
			value := row[j]
			diff := value - row[j-1]
			if diff > 3 || diff < -3 || diff == 0 {
				check = 0
				break
			}
			if j == 1 {
				grad = diff
			} else {
				if (grad > 0 && diff < 0) || (grad < 0 && diff > 0) {
					check = 0
					break
				}
			}
		}
		if check == 1 {
			safe++
		}
		fmt.Printf("line %d check=%d\n", i, check)
		fmt.Println("--------------------------")
	}

	//fmt.Println(len(column1))
	fmt.Println("Safe reports: ", safe)
	endTime := time.Now()
	fmt.Println("Cas: ", endTime.Sub(startTime))
}
