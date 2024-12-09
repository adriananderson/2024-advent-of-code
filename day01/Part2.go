package day01

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
	"sort"
	"strconv"
	"strings"
)

func calcCount(leftside, rightside []int) (result int) {
	var rightMap map[int]int
	rightMap = make(map[int]int)

	for idx := 0; idx < len(rightside); idx++ {
		rightMap[rightside[idx]] = 0
	}
	for idx := 0; idx < len(rightside); idx++ {
		rightMap[rightside[idx]] += 1
	}
	for idx := 0; idx < len(leftside); idx++ {
		result += rightMap[leftside[idx]] * leftside[idx]
	}

	return result
}

func Part2() int {
	defer utils.Timer("1-2")()
	var filename = "day01/day01.txt"
	var result = 0

	if fileContent, err := utils.ReadFileAsLines(filename); err == nil {
		left := make([]int, len(fileContent))
		right := make([]int, len(fileContent))
		for lineIdx := 0; lineIdx < len(fileContent); lineIdx++ {
			parts := strings.Split(fileContent[lineIdx], " ")
			left[lineIdx], _ = strconv.Atoi(parts[0])
			right[lineIdx], _ = strconv.Atoi(parts[(len(parts) - 1)])
		}
		sort.Ints(left)
		sort.Ints(right)
		result = calcCount(left, right)
	}

	return result
}
