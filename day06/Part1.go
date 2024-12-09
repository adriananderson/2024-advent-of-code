package day06

import (
	"2024/utils"
)

type Position struct {
	row    int
	column int
}

const (
	NORTH = iota
	EAST
	SOUTH
	WEST
)

func Part1() int {
	defer utils.Timer("6-1")()
	var fileName = "day06/day06.txt"

	var guard Position
	var facing = NORTH
	var obstructions [][]bool
	var visited [][]bool
	
	//read in initial state
	obstructions, guard = readFile(fileName)

	visited = findVisited(obstructions, guard, facing)

	return countVisited(visited)
}

func findVisited(obstructions [][]bool, guard Position, facing int) [][]bool {
	visited := make([][]bool, len(obstructions)+1)
	for row := range obstructions {
		visited[row] = make([]bool, len(obstructions[row])+1)
	}

whileLoop:
	for {
		visited[guard.row][guard.column] = true
		switch facing {
		case NORTH:
			if guard.row == 0 {
				break whileLoop
			}
			if obstructions[guard.row-1][guard.column] == false {
				guard = Position{guard.row - 1, guard.column}
			} else {
				facing = EAST
			}
		case EAST:
			if guard.column == len(obstructions)-1 {
				break whileLoop
			}
			if obstructions[guard.row][guard.column+1] == false {
				guard = Position{guard.row, guard.column + 1}
			} else {
				facing = SOUTH
			}
		case SOUTH:
			if guard.row == len(obstructions)-1 {
				break whileLoop
			}
			if obstructions[guard.row+1][guard.column] == false {
				guard = Position{guard.row + 1, guard.column}
			} else {
				facing = WEST
			}
		case WEST:
			if guard.column == 0 {
				break whileLoop
			}
			if obstructions[guard.row][guard.column-1] == false {
				guard = Position{guard.row, guard.column - 1}
			} else {
				facing = NORTH
			}
		}
	}
	return visited
}

func readFile(fileName string) ([][]bool, Position) {
	var obstructions [][]bool
	var guard Position
	if fileContent, err := utils.ReadFileAsLines(fileName); err == nil {
		obstructions = make([][]bool, len(fileContent)-1)
		for row, line := range fileContent {
			if len(line) > 1 {
				obstructions[row] = make([]bool, len(line))
				for column, char := range line {
					switch char {
					case '#':
						obstructions[row][column] = true
					case '^':
						guard = Position{row, column}
					}
				}
			}
		}
	}
	return obstructions, guard
}

func countVisited(visited [][]bool) (result int) {
	result = 0
	for _, visitedRow := range visited {
		for _, visitedColumn := range visitedRow {
			if visitedColumn {
				result += 1
			}
		}
	}
	return result
}
