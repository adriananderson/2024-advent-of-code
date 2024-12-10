package day10

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
)

func Part2() int {
	defer utils.Timer("10-1")()
	var fileName = "day10/day10.txt"

	topographicMap, zeros := readFile(fileName)

	result := 0
	for _, zero := range zeros {
		for _, rating := range findTrailHeadSummits(zero, topographicMap) {
			result += rating
		}
	}

	return result
}
