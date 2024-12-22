package day22

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
	"strconv"
)

func Part1(fileName string, numIters int) int {
	defer utils.Timer("22-1")()

	secretNumbers := readFile(fileName)

	for ii := 0; ii < numIters; ii++ {
		for jj, secretNumber := range secretNumbers {
			tmp := nextNumber(secretNumber)
			//fmt.Printf("%d     %d -> %d\n", ii, secretNumber, tmp)
			secretNumbers[jj] = tmp
		}
	}

	result := 0
	for _, secretNumber := range secretNumbers {
		result += secretNumber
	}

	return result
}

func nextNumber(secretNumber int) int {
	secretNumber = (secretNumber * 64) ^ secretNumber
	secretNumber = secretNumber % 16777216
	secretNumber = (secretNumber / 32) ^ secretNumber
	secretNumber = secretNumber % 16777216
	secretNumber = (secretNumber * 2048) ^ secretNumber
	secretNumber = secretNumber % 16777216
	return secretNumber
}

func readFile(fileName string) (secretNumbers []int) {
	fileLines, _ := utils.ReadFileAsLines(fileName)

	secretNumbers = make([]int, len(fileLines)-1)
	for ii, numberLine := range fileLines {
		if len(numberLine) > 0 {
			secretNumbers[ii], _ = strconv.Atoi(numberLine)
		}
	}

	return secretNumbers
}
