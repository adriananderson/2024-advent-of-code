package day17

import (
	"fmt"
	"github.com/adriananderson/2024-advent-of-code/utils"
	"math"
	"strconv"
	"strings"
)

func Part1(fileName string) string {
	defer utils.Timer("17-1")()

	registers, program := readFile(fileName)

	output := runProgram(program, registers)

	result := ""
	for i := 0; i < len(output); i++ {
		result += strconv.Itoa(output[i])
		if i < len(output)-1 {
			result += ","
		}
	}

	return result
}

func runProgram(program []int, registers [3]int) []int {
	instructionPointer := 0
	output := make([]int, 0)

	for {
		if instructionPointer > len(program)-1 {
			break
		}
		switch program[instructionPointer] {
		case 0: //adv
			num := registers[0]
			denom := math.Pow(2, (float64)(combo(program[instructionPointer+1], registers)))
			registers[0] = num / (int)(denom)
		case 1: //bxl
			registers[1] = registers[1] ^ program[instructionPointer+1]
		case 2: //bst
			registers[1] = combo(program[instructionPointer+1], registers) % 8
		case 3: //jnz
			if registers[0] != 0 {
				instructionPointer = program[instructionPointer+1] - 2
			}
		case 4: //bxc
			registers[1] = registers[1] ^ registers[2]
		case 5: //out
			output = append(output, combo(program[instructionPointer+1], registers)%8)
		case 6: //bdv
			num := registers[0]
			denom := math.Pow(2, (float64)(combo(program[instructionPointer+1], registers)))
			registers[1] = num / (int)(denom)
		case 7: //cdv
			num := registers[0]
			denom := math.Pow(2, (float64)(combo(program[instructionPointer+1], registers)))
			registers[2] = num / (int)(denom)
		}
		instructionPointer += 2
	}
	return output
}

func combo(value int, registers [3]int) int {
	switch value {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	case 4:
		return registers[0]
	case 5:
		return registers[1]
	case 6:
		return registers[2]
	default:
		fmt.Printf("invalid value %d\n", value)
		return -1
	}
}

func readFile(fileName string) (registers [3]int, program []int) {
	programLines, _ := utils.ReadFileAsLines(fileName)

	for _, line := range programLines {
		lineParts := strings.Split(line, ": ")
		switch lineParts[0] {
		case "Register A":
			registers[0], _ = strconv.Atoi(lineParts[1])
		case "Register B":
			registers[1], _ = strconv.Atoi(lineParts[1])
		case "Register C":
			registers[2], _ = strconv.Atoi(lineParts[1])
		case "Program":
			opcodes := strings.Split(lineParts[1], ",")
			program = make([]int, len(opcodes))
			for ii, opcode := range opcodes {
				program[ii], _ = strconv.Atoi(opcode)
			}
		}
	}
	return registers, program
}
