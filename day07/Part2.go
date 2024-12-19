package day07

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
	"strconv"
	"strings"
)

func Part2() int {
	defer utils.Timer("7-2")()
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

			if isSolvable2(numbers, answer) {
				result += answer
			}
		}
	}

	return result
}

func isSolvable2(numbers []int, answer int) bool {
	if len(numbers) == 0 {
		return false
	}
	lastNumber := numbers[len(numbers)-1]
	if len(numbers) == 1 {
		return lastNumber == answer
	}
	if answer%lastNumber == 0 && isSolvable2(numbers[:len(numbers)-1], answer/lastNumber) {
		return true
	}
	if answer > lastNumber && isSolvable2(numbers[:len(numbers)-1], answer-lastNumber) {
		return true
	}

	answerString := strconv.Itoa(answer)
	lastNumberString := strconv.Itoa(lastNumber)
	if len(answerString) > len(lastNumberString) && strings.HasSuffix(answerString, lastNumberString) {
		newAnswer, _ := strconv.Atoi(answerString[:len(answerString)-len(lastNumberString)])
		return isSolvable2(numbers[:len(numbers)-1], newAnswer)
	}
	return false
}
