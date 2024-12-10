package day10

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
	"strings"
)

type Position struct {
	row    int
	column int
	height int
}

func Part1() int {
	defer utils.Timer("10-1")()
	var fileName = "day10/day10.txt"

	topographicMap, zeros := readFile(fileName)

	result := 0
	for _, zero := range zeros {
		result += len(findTrailHeadSummits(zero, topographicMap))
	}

	return result
}

func readFile(fileName string) ([][]int, []Position) {
	fileLines, _ := utils.ReadFileAsLines(fileName)
	topographicMap := make([][]int, len(fileLines))
	zeros := make([]Position, 0)

	for row, line := range fileLines {
		topographicMap[row] = make([]int, len(line))
		elevations := strings.Split(line, "")
		for column, elevation := range elevations {
			elevationValue := (int)(elevation[0] - '0')
			topographicMap[row][column] = elevationValue

			if elevationValue == 0 {
				zeros = append(zeros, Position{row: row, column: column, height: 0})
			}
		}
	}
	return topographicMap, zeros
}

func findTrailHeadSummits(location Position, topographicMap [][]int) (summits map[Position]int) {
	summits = make(map[Position]int)
	if location.height == 9 {
		summits[location] = 1
		return summits
	}

	if location.row > 0 && topographicMap[location.row-1][location.column] == (location.height+1) { //north
		northSummits := findTrailHeadSummits(Position{row: location.row - 1, column: location.column, height: location.height + 1}, topographicMap)
		for position, value := range northSummits {
			if rating, ok := summits[position]; ok {
				summits[position] = rating + value
			} else {
				summits[position] = value
			}
		}
	}
	if location.column < len(topographicMap)-1 && topographicMap[location.row][location.column+1] == (location.height+1) { //east
		eastSummits := findTrailHeadSummits(Position{row: location.row, column: location.column + 1, height: location.height + 1}, topographicMap)
		for position, value := range eastSummits {
			if rating, ok := summits[position]; ok {
				summits[position] = rating + value
			} else {
				summits[position] = value
			}
		}
	}
	if location.row < len(topographicMap[0])-1 && topographicMap[location.row+1][location.column] == (location.height+1) { //south
		southSummits := findTrailHeadSummits(Position{row: location.row + 1, column: location.column, height: location.height + 1}, topographicMap)
		for position, value := range southSummits {
			if rating, ok := summits[position]; ok {
				summits[position] = rating + value
			} else {
				summits[position] = value
			}
		}
	}
	if location.column > 0 && topographicMap[location.row][location.column-1] == (location.height+1) { //west
		westSummits := findTrailHeadSummits(Position{row: location.row, column: location.column - 1, height: location.height + 1}, topographicMap)
		for position, value := range westSummits {
			if rating, ok := summits[position]; ok {
				summits[position] = rating + value
			} else {
				summits[position] = value
			}
		}
	}

	return summits
}
