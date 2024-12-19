package day07

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
	"strconv"
	"strings"
)

func Part1() int {
	defer utils.Timer("7-1")()
	var fileName = "day07/day07.txt"

	result := 0

	if fileLines, err := utils.ReadFileAsLines(fileName); err == nil {
		for _, line := range fileLines {
			parts := strings.Split(line, ":")
			if len(parts) != 2 {
				break
			}
			answer, _ := strconv.Atoi(parts[0])
			//fmt.Println(answer)
			stringNums := strings.Split(strings.TrimSpace(parts[1]), " ")
			numbers := make([]int, len(stringNums))
			for ii, num := range stringNums {
				numbers[ii], _ = strconv.Atoi(num)
			}

			if isSolvable(numbers, answer) {
				result += answer
			}
		}
	}

	return result
}

func isSolvable(numbers []int, answer int) bool {
	if len(numbers) == 0 {
		return false
	}
	lastNumber := numbers[len(numbers)-1]
	if len(numbers) == 1 {
		return lastNumber == answer
	}
	if answer%lastNumber == 0 && isSolvable(numbers[0:len(numbers)-1], answer/lastNumber) {
		return true
	}
	if answer > lastNumber && isSolvable(numbers[0:len(numbers)-1], answer-lastNumber) {
		return true
	}
	return false
}
