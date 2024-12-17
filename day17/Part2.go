package day17

import (
	"fmt"
	"github.com/adriananderson/2024-advent-of-code/utils"
	"slices"
)

func Part2(fileName string) int {
	defer utils.Timer("17-2")()

	registers, program := readFile(fileName)

	result := 0

	programLength := len(program)
	minA := 1 << (3 * (programLength - 1))
	maxA := 1 << (3 * (programLength))

	//fmt.Printf(" %d %d     %o %o\n", minA, maxA, minA, maxA)

	increments := make([]int, programLength)
	offset := 3
	for ii := 0; ii < programLength; ii++ {
		increments[ii] = 1 << (3 * (programLength - ii - 1))
	}

	for ii := minA; ii <= maxA; ii += increments[offset] {
		registers[0], registers[1], registers[2] = ii, 0, 0

		output := runProgram(program, registers)

		if slices.Equal(output, program) {
			result = ii
			//fmt.Printf("%d  %o out   %v        prog %v\n", ii, ii, output, program)
			break
		}

		//fmt.Printf("%d  %o out   %v        prog %v\n", ii, ii, output, program)

		if slices.Equal(output[(programLength-1-offset):], program[(programLength-1-offset):]) {
			//fmt.Printf("      offset: %d  %o\n", offset, increments[offset])
			offset++
		}
	}

	return result
}

func Part3(fileName string) int {
	defer utils.Timer("17-2")()

	registers, program := readFile(fileName)

	result := 0

	//printProgram(program)
	//oldOutput := runProgram(instructionPointer, program, registers)

	programLength := len(program)
	minA := 1 << (3 * (programLength - 1))
	maxA := 1 << (3 * (programLength))

	fmt.Printf(" %d %d     %o %o\n", minA, maxA, minA, maxA)

	increments := make([]int, programLength)
	offset := 0
	for ii := 0; ii < programLength; ii++ {
		increments[ii] = 1 << (3 * (programLength - ii - 1))
	}

	for ii := minA; ii <= maxA; ii += increments[offset] {
		registers[0], registers[1], registers[2] = ii, 0, 0

		output := runProgram(program, registers)

		if slices.Equal(output, program) {
			result = ii
			fmt.Printf("%d  %o out   %v        prog %v\n", ii, ii, output, program)
			break
		}

		fmt.Printf("%d  %o out   %v        prog %v\n", ii, ii, output, program)

		if slices.Equal(output[(programLength-2-offset):], program[(programLength-2-offset):]) {
			fmt.Printf("      offset: %d  %o\n", offset, increments[offset])
			offset++
		}
		//for program[programLength-2-offset] == output[programLength-2-offset] {
		//	fmt.Printf("      offset: %d  %o\n", offset, increments[offset])
		//	offset++
		//}
	}

	return result
}

func printProgram(program []int) {
	for ii := 0; ii < len(program); ii += 2 {

		switch program[ii] {
		case 0: //adv
			fmt.Printf("%d %d \tADV combo(%d)\n", program[ii], program[ii+1], program[ii+1])
		case 1: //bxl
			fmt.Printf("%d %d \tBXL (%d)\n", program[ii], program[ii+1], program[ii+1])
		case 2: //bst
			fmt.Printf("%d %d \tBST combo(%d)\n", program[ii], program[ii+1], program[ii+1])
		case 3: //jnz
			fmt.Printf("%d %d \tJNZ (%d)\n", program[ii], program[ii+1], program[ii+1])
		case 4: //bxc
			fmt.Printf("%d %d \tBXC (%d)\n", program[ii], program[ii+1], program[ii+1])
		case 5: //out
			fmt.Printf("%d %d \tOUT combo(%d)\n", program[ii], program[ii+1], program[ii+1])
		case 6: //bdv
			fmt.Printf("%d %d \tBDV (%d)\n", program[ii], program[ii+1], program[ii+1])
		case 7: //cdv
			fmt.Printf("%d %d \tCDV (%d)\n", program[ii], program[ii+1], program[ii+1])
		}
	}
}
