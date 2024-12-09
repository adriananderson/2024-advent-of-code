package day02

import (
	"2024/utils"
	"strconv"
	"strings"
)

func Part2() int {
	defer utils.Timer("2-2")()
	var filename = "day02/day02.txt"
	var result = 0

	if fileContent, err := utils.ReadFileAsLines(filename); err == nil {
		reports := make([][]int, len(fileContent)-1)
		for lineIdx := 0; lineIdx < len(fileContent)-1; lineIdx++ {
			parts := strings.Split(fileContent[lineIdx], " ")
			report := make([]int, len(parts))
			for cc := 0; cc < len(parts); cc++ {
				report[cc], _ = strconv.Atoi(parts[cc])
			}
			reports[lineIdx] = report
		}
		for idx := 0; idx < len(reports); idx++ {
			if IsSafeReportWithoutOne(reports[idx]) {
				result++
			}
		}
	}

	return result
}
