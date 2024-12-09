package day04

import (
	"2024/utils"
	"strings"
)

func isWord(wordMap [][]string, word string, row, rowDir, column, columnDir int) bool {
	for c := 0; c < len(word); c++ {
		if row < 0 || row >= len(wordMap) || column < 0 || column >= len(wordMap[row]) {
			return false
		}
		value := wordMap[row][column]
		if value != string(word[c]) {
			return false
		}
		row += rowDir
		column += columnDir
	}

	return true
}

func countWords(wordMap [][]string, word string, row, column int) (count int) {
	directions := []int{-1, 0, 1}
	for _, rowDir := range directions {
		for _, colDir := range directions {
			if isWord(wordMap, word, row, rowDir, column, colDir) {
				//fmt.Printf("Found word %d %d %d %d\n", row, rowDir, column, colDir)
				count++
			}
		}
	}
	return
}

func Part1() int {
	defer utils.Timer("4-1")()
	var fileName = "day04/day04.txt"
	var result = 0

	if fileContent, err := utils.ReadFileAsLines(fileName); err == nil {
		wordMap := make([][]string, len(fileContent[0]))
		for row := 0; row < len(fileContent)-1; row++ {
			letters := strings.Split(fileContent[row], "")
			wordMap[row] = letters
		}

		for row := 0; row < len(wordMap); row++ {
			for col := 0; col < len(wordMap); col++ {
				result += countWords(wordMap, "XMAS", row, col)
			}
		}
	}

	return result
}
