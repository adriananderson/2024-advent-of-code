package day06

import "github.com/adriananderson/2024-advent-of-code/utils"

func Part2() int {
	defer utils.Timer("6-2")()

	var fileName = "day06/day06.txt"

	var guard Position
	var facing = NORTH
	var obstructions [][]bool

	var result = 0

	obstructions, guard = readFile(fileName) //read in initial state
	visited := findVisited(obstructions, guard, facing)

	for row := range obstructions {
		for column := range obstructions[row] {
			if visited[row][column] == true { //skip if not part of path
				obstructions[row][column] = true
				if isLoop(guard, facing, obstructions) {
					result += 1
				}
				obstructions[row][column] = false
			}
		}
	}

	return result
}

// the path will either end in going off the board or repeating a visited square facing the same direction
func isLoop(guard Position, facing int, obstructions [][]bool) (loop bool) {

	visited := make([][][]bool, len(obstructions)+1)
	for row := range obstructions {
		visited[row] = make([][]bool, len(obstructions[row])+1)
		for column := range obstructions[row] {
			visited[row][column] = make([]bool, WEST+1)
		}
	}

whileLoop:
	for {
		if visited[guard.row][guard.column][facing] {
			return true //found loop
		}

		visited[guard.row][guard.column][facing] = true
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

	return false //off board
}
