package day01

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
	"math"
	"sort"
	"strconv"
	"strings"
)

func calculateDistance(leftSide, rightSide []int) (result int) {
	for idx := 0; idx < len(leftSide); idx++ {
		result += int(math.Abs(float64(leftSide[idx]) - float64(rightSide[idx])))
	}

	return
}

func Part1() int {
	defer utils.Timer("1-1")()
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
		result = calculateDistance(left, right)
	}

	return result
}
