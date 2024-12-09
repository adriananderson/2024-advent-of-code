package day04

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
	"strings"
)

func countCrosses(wordMap [][]string, row int, column int) (count int) {
	if wordMap[row][column] == "A" {
		if ((wordMap[row-1][column-1] == "M" && wordMap[row+1][column+1] == "S") ||
			(wordMap[row-1][column-1] == "S" && wordMap[row+1][column+1] == "M")) &&
			((wordMap[row-1][column+1] == "M" && wordMap[row+1][column-1] == "S") ||
				(wordMap[row-1][column+1] == "S" && wordMap[row+1][column-1] == "M")) {
			return 1
		}
	}

	return 0
}

func Part2() int {
	defer utils.Timer("4-2")()
	var fileName = "day04/day04.txt"
	var result = 0

	if fileContent, err := utils.ReadFileAsLines(fileName); err == nil {
		wordMap := make([][]string, len(fileContent[0]))
		for row := 0; row < len(fileContent)-1; row++ {
			letters := strings.Split(fileContent[row], "")
			wordMap[row] = letters
		}

		for row := 1; row < len(wordMap)-1; row++ {
			for col := 1; col < len(wordMap)-1; col++ {
				result += countCrosses(wordMap, row, col)
			}
		}
	}

	return result
}
