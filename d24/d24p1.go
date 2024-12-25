package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Gate struct {
	input1, input2, output string
	op                     string
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

	// Read the file into a slice of strings
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Parse initial wire values and gate connections
	wireValues := make(map[string]int)
	var gates []Gate
	for _, line := range lines {
		if strings.Contains(line, ":") {
			parts := strings.Split(line, ": ")
			wireValues[parts[0]] = int(parts[1][0] - '0')
		} else if strings.Contains(line, "->") {
			parts := strings.Split(line, " -> ")
			inputs := strings.Fields(parts[0])
			gates = append(gates, Gate{inputs[0], inputs[2], parts[1], inputs[1]})
		}
	}

	// Simulate the gates
	processed := make(map[string]bool)
	for len(processed) < len(gates) {
		for _, gate := range gates {
			if processed[gate.output] {
				continue
			}
			val1, ok1 := wireValues[gate.input1]
			val2, ok2 := wireValues[gate.input2]
			if ok1 && ok2 {
				var result int
				switch gate.op {
				case "AND":
					result = val1 & val2
				case "OR":
					result = val1 | val2
				case "XOR":
					result = val1 ^ val2
				}
				wireValues[gate.output] = result
				processed[gate.output] = true
			}
		}
	}

	// Combine the bits from all wires starting with 'z'
	var binaryResult string
	for i := 0; ; i++ {
		wire := fmt.Sprintf("z%02d", i)
		if val, exists := wireValues[wire]; exists {
			binaryResult = fmt.Sprintf("%d%s", val, binaryResult)
		} else {
			break
		}
	}

	// Convert binary result to decimal
	var decimalResult int
	fmt.Sscanf(binaryResult, "%b", &decimalResult)

	fmt.Println("Decimal result:", decimalResult)

	endTime := time.Now()
	fmt.Println("Execution time:", endTime.Sub(startTime))
}
