package day03

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
	"regexp"
	"strconv"
)

func Part1() int {
	defer utils.Timer("3-1")()
	var fileName = "day03/day03.txt"
	var result = 0

	if fileContent, err := utils.ReadFileAsText(fileName); err == nil {
		pattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
		parts := pattern.FindAllStringSubmatch(fileContent, -1)
		for _, part := range parts {
			first, _ := strconv.Atoi(part[1])
			second, _ := strconv.Atoi(part[2])
			result += first * second
		}
	}

	return result
}
