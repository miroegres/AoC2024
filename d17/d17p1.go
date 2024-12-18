package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()

	// Initialize registers
	registers := map[string]int{"A": 46337277, "B": 0, "C": 0}

	// Define the program
	program := []int{2, 4, 1, 1, 7, 5, 4, 4, 1, 4, 0, 3, 5, 5, 3, 0}

	// Run the program
	output := runProgram(registers, program)

	// Print the output
	fmt.Println("Output:", output)

	endTime := time.Now()
	fmt.Println("Time:", endTime.Sub(startTime))
}

func runProgram(registers map[string]int, program []int) string {
	ip := 0
	var output []string

	getComboValue := func(operand int) int {
		switch operand {
		case 0, 1, 2, 3:
			return operand
		case 4:
			return registers["A"]
		case 5:
			return registers["B"]
		case 6:
			return registers["C"]
		default:
			panic("Invalid combo operand")
		}
	}

	for ip < len(program) {
		opcode := program[ip]
		operand := program[ip+1]

		switch opcode {
		case 0: // adv
			denominator := 1 << getComboValue(operand)
			registers["A"] /= denominator
		case 1: // bxl
			registers["B"] ^= operand
		case 2: // bst
			registers["B"] = getComboValue(operand) % 8
		case 3: // jnz
			if registers["A"] != 0 {
				ip = operand
				continue
			}
		case 4: // bxc
			registers["B"] ^= registers["C"]
		case 5: // out
			output = append(output, strconv.Itoa(getComboValue(operand)%8))
		case 6: // bdv
			denominator := 1 << getComboValue(operand)
			registers["B"] = registers["A"] / denominator
		case 7: // cdv
			denominator := 1 << getComboValue(operand)
			registers["C"] = registers["A"] / denominator
		}

		ip += 2
	}

	return strings.Join(output, ",")
}
