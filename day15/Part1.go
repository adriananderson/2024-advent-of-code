package day15

import (
	"fmt"
	"github.com/adriananderson/2024-advent-of-code/utils"
	"strings"
)

type Position struct {
	xx int
	yy int
}

const (
	EMPTY = iota
	WALL
	BOX
	ROBOT
	LEFTBOX
	RIGHTBOX
)

const (
	NOMOVE = iota
	NORTH
	EAST
	SOUTH
	WEST
)

func Part1() int {
	defer utils.Timer("15-1")()
	var fileName = "day15/day15.txt"

	worldMap, movements, robotPos := readFile(fileName)

	//printWorld(worldMap)
	for _, move := range movements {
		//fmt.Printf("Move %d %v\n", ii, move)
		robotPos = tryMove(worldMap, robotPos, move)
		//printWorld(worldMap)
	}

	result := scoreWorld(worldMap)

	return result
}

func scoreWorld(worldMap [][]int) int {
	result := 0
	for ii, row := range worldMap {
		for jj, cell := range row {
			if cell == BOX {
				result += 100*ii + jj
			}
		}
	}
	return result
}

func printWorld(worldMap [][]int) {
	for _, row := range worldMap {
		for _, cell := range row {
			switch cell {
			case WALL:
				fmt.Print("#")
			case BOX:
				fmt.Print("O")
			case ROBOT:
				fmt.Print("@")
			case EMPTY:
				fmt.Print(".")
			case LEFTBOX:
				fmt.Print("[")
			case RIGHTBOX:
				fmt.Print("]")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func tryMove(worldMap [][]int, robotPos Position, move int) (newRobotPos Position) {
	var intendedPosition Position
	var intendedObject int
	switch move {
	case NORTH:
		intendedPosition = Position{robotPos.xx, robotPos.yy - 1}
		intendedObject = worldMap[robotPos.yy-1][robotPos.xx]
	case EAST:
		intendedPosition = Position{robotPos.xx + 1, robotPos.yy}
		intendedObject = worldMap[robotPos.yy][robotPos.xx+1]
	case SOUTH:
		intendedPosition = Position{robotPos.xx, robotPos.yy + 1}
		intendedObject = worldMap[robotPos.yy+1][robotPos.xx]
	case WEST:
		intendedPosition = Position{robotPos.xx - 1, robotPos.yy}
		intendedObject = worldMap[robotPos.yy][robotPos.xx-1]
	case NOMOVE: // invalid
		return robotPos
	}

	if intendedObject == WALL { //nope
		newRobotPos = robotPos
	} else if intendedObject == EMPTY {
		worldMap[intendedPosition.yy][intendedPosition.xx] = worldMap[robotPos.yy][robotPos.xx]
		worldMap[robotPos.yy][robotPos.xx] = EMPTY
		newRobotPos = intendedPosition
	} else if intendedObject == BOX {
		tryMove(worldMap, intendedPosition, move)
		if worldMap[intendedPosition.yy][intendedPosition.xx] == EMPTY {
			worldMap[intendedPosition.yy][intendedPosition.xx] = worldMap[robotPos.yy][robotPos.xx]
			worldMap[robotPos.yy][robotPos.xx] = EMPTY
			newRobotPos = intendedPosition
		} else {
			newRobotPos = robotPos
		}
	}
	return newRobotPos
}

func readFile(fileName string) (worldMap [][]int, movements []int, robotPos Position) {
	fileContents, _ := utils.ReadFileAsText(fileName)

	fileParts := strings.Split(fileContents, "\n\n")
	mapLines := strings.Split(fileParts[0], "\n")
	movementLines := strings.Split(fileParts[1], "\n")

	worldMap = make([][]int, len(mapLines))
	for ii, mapLine := range mapLines {
		worldMap[ii] = make([]int, len(mapLine))
		squares := strings.Split(mapLine, "")
		for jj, square := range squares {
			switch square {
			case ".":
				worldMap[ii][jj] = EMPTY
			case "#":
				worldMap[ii][jj] = WALL
			case "O":
				worldMap[ii][jj] = BOX
			case "@":
				worldMap[ii][jj] = ROBOT
				robotPos.xx, robotPos.yy = jj, ii
			}
		}
	}

	movements = make([]int, len(fileParts[1]))
	index := 0
	for _, moveGroup := range movementLines {
		moveGroupParts := strings.Split(moveGroup, "")
		for _, move := range moveGroupParts {
			switch move {
			case "^":
				movements[index] = NORTH
			case ">":
				movements[index] = EAST
			case "v":
				movements[index] = SOUTH
			case "<":
				movements[index] = WEST
			}
			index++
		}
	}

	return worldMap, movements, robotPos
}
