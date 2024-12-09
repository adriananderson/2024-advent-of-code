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

			answers := genAnswers(numbers)
			//check this line
		lineCheck:
			for _, ans := range answers {
				if ans == answer {
					result += answer
					break lineCheck
				}
			}

		}
	}

	return result
}

// takes a list of numbers and generates all possible solutions
func genAnswers(numbers []int) []int {
	return genAnswerSlice(numbers[0], numbers[1:])
}

func genAnswerSlice(front int, numbers []int) []int {

	sum := front + numbers[0]
	product := front * numbers[0]

	if len(numbers) > 1 {
		sumSlice := genAnswerSlice(sum, numbers[1:])
		productSlice := genAnswerSlice(product, numbers[1:])
		return append(sumSlice, productSlice...)
	}
	return []int{sum, product}
}
