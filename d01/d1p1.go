package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	sum := 0
	sim := 0

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	//inputString := "four82sev7en"
	//reFirst := regexp.MustCompile(`\d`)
	//reLast := regexp.MustCompile(`\d\D*$`)

	/*firstDigit := reFirst.FindString(inputString)
	lastDigit := reLast.FindString(inputString)
	fmt.Println("First Digit:", firstDigit)
	fmt.Println("Last Digit:", lastDigit)*/

	// Create slices to hold the numbers
	var column1 []int     //input
	var column2 []int     //input
	var column3 [1000]int //distance, must be set to specific length
	var column4 [1000]int //occurences
	var column5 [1000]int //similarity

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Fields(line)
		if len(columns) >= 2 {
			num1, err1 := strconv.Atoi(columns[0])
			num2, err2 := strconv.Atoi(columns[1])
			if err1 == nil && err2 == nil {
				column1 = append(column1, num1)
				column2 = append(column2, num2)
			}
		}
	}

	// Sort column1
	sort.Ints(column1)
	sort.Ints(column2)

	//d1p1
	for i := 0; i < 1000; i++ {
		column3[i] = absInt(column1[i] - column2[i])
		sum += column3[i]
	}

	//d1p2
	for i, value1 := range column1 {
		count := 0
		for _, value2 := range column2 {
			if value1 == value2 {
				count++
			}
		}
		column4[i] = count
		column5[i] = column1[i] * column4[i]
		sim += column5[i]
	}

	//fmt.Println(len(column1))
	fmt.Println("Sum: ", sum)
	fmt.Println("Similarity: ", sim)
	endTime := time.Now()
	fmt.Println("Cas: ", endTime.Sub(startTime))
}
