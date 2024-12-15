package day15

import (
	"fmt"
	"github.com/adriananderson/2024-advent-of-code/utils"
)

func Part2() int {
	defer utils.Timer("15-2")()
	var fileName = "day15/day15.txt"

	worldMap, movements, robotPos := readFile(fileName)
	//printWorld(worldMap)
	worldMap, robotPos = scaleWorld(worldMap)
	//printWorld(worldMap)

	for _, move := range movements {
		//fmt.Printf("Move %d %v\n", ii, move)
		robotPos = tryMove2(worldMap, robotPos, move)
		//printWorld(worldMap)
	}

	result := scoreWorld2(worldMap)

	return result
}

func scaleWorld(worldMap [][]int) (newWorldMap [][]int, robotPos Position) {
	newWorldMap = make([][]int, len(worldMap))
	for ii, row := range worldMap {
		newWorldMap[ii] = make([]int, len(row)*2)
		for jj := 0; jj < len(row); jj++ {
			switch worldMap[ii][jj] {
			case WALL:
				newWorldMap[ii][jj*2] = WALL
				newWorldMap[ii][jj*2+1] = WALL
			case EMPTY:
				newWorldMap[ii][jj*2] = EMPTY
				newWorldMap[ii][jj*2+1] = EMPTY
			case BOX:
				newWorldMap[ii][jj*2] = LEFTBOX
				newWorldMap[ii][jj*2+1] = RIGHTBOX
			case ROBOT:
				newWorldMap[ii][jj*2] = ROBOT
				newWorldMap[ii][jj*2+1] = EMPTY
				robotPos = Position{jj * 2, ii}
			}
		}
	}
	return newWorldMap, robotPos
}

func scoreWorld2(worldMap [][]int) int {
	result := 0
	for ii, row := range worldMap {
		for jj, cell := range row {
			if cell == LEFTBOX {
				result += 100*ii + jj
			}
		}
	}
	return result
}

func canMove(worldMap [][]int, objectPos Position, move int) (movePossible bool) {
	if move == NOMOVE { //newline
		fmt.Printf("attempted to move NOMOVE")
		return false
	}

	if worldMap[objectPos.yy][objectPos.xx] == EMPTY {
		return true
	} else if worldMap[objectPos.yy][objectPos.xx] == WALL {
		return false
	} else if worldMap[objectPos.yy][objectPos.xx] == ROBOT {
		switch move {
		case NORTH:
			return canMove(worldMap, Position{objectPos.xx, objectPos.yy - 1}, move)
		case EAST:
			return canMove(worldMap, Position{objectPos.xx + 1, objectPos.yy}, move)
		case SOUTH:
			return canMove(worldMap, Position{objectPos.xx, objectPos.yy + 1}, move)
		case WEST:
			return canMove(worldMap, Position{objectPos.xx - 1, objectPos.yy}, move)
		}
	} else if worldMap[objectPos.yy][objectPos.xx] == LEFTBOX {
		switch move {
		case NORTH:
			return canMove(worldMap, Position{objectPos.xx, objectPos.yy - 1}, move) &&
				canMove(worldMap, Position{objectPos.xx + 1, objectPos.yy - 1}, move)
		case EAST:
			return canMove(worldMap, Position{objectPos.xx + 2, objectPos.yy}, move)
		case SOUTH:
			return canMove(worldMap, Position{objectPos.xx, objectPos.yy + 1}, move) &&
				canMove(worldMap, Position{objectPos.xx + 1, objectPos.yy + 1}, move)
		case WEST:
			return canMove(worldMap, Position{objectPos.xx - 1, objectPos.yy}, move)
		}
	} else if worldMap[objectPos.yy][objectPos.xx] == RIGHTBOX {
		switch move {
		case NORTH:
			return canMove(worldMap, Position{objectPos.xx, objectPos.yy - 1}, move) &&
				canMove(worldMap, Position{objectPos.xx - 1, objectPos.yy - 1}, move)
		case EAST:
			return canMove(worldMap, Position{objectPos.xx + 1, objectPos.yy}, move)
		case SOUTH:
			return canMove(worldMap, Position{objectPos.xx, objectPos.yy + 1}, move) &&
				canMove(worldMap, Position{objectPos.xx - 1, objectPos.yy + 1}, move)
		case WEST:
			return canMove(worldMap, Position{objectPos.xx - 2, objectPos.yy}, move)
		}
	}

	fmt.Printf("Programmer failure")
	return false
}

func doMove(worldMap [][]int, objectPos Position, move int) (newPos Position) {
	newPos = objectPos
	if worldMap[objectPos.yy][objectPos.xx] == EMPTY {
		//fmt.Printf("attempted to move EMPTY\n")
	} else if worldMap[objectPos.yy][objectPos.xx] == WALL {
		fmt.Printf("attempted to move WALL\n")
	} else if worldMap[objectPos.yy][objectPos.xx] == ROBOT {
		switch move {
		case NORTH:
			newPos = Position{objectPos.xx, objectPos.yy - 1}
			doMove(worldMap, newPos, move)
			worldMap[objectPos.yy-1][objectPos.xx] = ROBOT
			worldMap[objectPos.yy][objectPos.xx] = EMPTY
		case EAST:
			newPos = Position{objectPos.xx + 1, objectPos.yy}
			doMove(worldMap, newPos, move)
			worldMap[objectPos.yy][objectPos.xx+1] = ROBOT
			worldMap[objectPos.yy][objectPos.xx] = EMPTY
		case SOUTH:
			newPos = Position{objectPos.xx, objectPos.yy + 1}
			doMove(worldMap, newPos, move)
			worldMap[objectPos.yy+1][objectPos.xx] = ROBOT
			worldMap[objectPos.yy][objectPos.xx] = EMPTY
		case WEST:
			newPos = Position{objectPos.xx - 1, objectPos.yy}
			doMove(worldMap, newPos, move)
			worldMap[objectPos.yy][objectPos.xx-1] = ROBOT
			worldMap[objectPos.yy][objectPos.xx] = EMPTY
		}
	} else if worldMap[objectPos.yy][objectPos.xx] == LEFTBOX {
		switch move {
		case NORTH:
			doMove(worldMap, Position{objectPos.xx, objectPos.yy - 1}, move)
			worldMap[objectPos.yy-1][objectPos.xx] = LEFTBOX
			worldMap[objectPos.yy][objectPos.xx] = EMPTY
			doMove(worldMap, Position{objectPos.xx + 1, objectPos.yy - 1}, move)
			worldMap[objectPos.yy-1][objectPos.xx+1] = RIGHTBOX
			worldMap[objectPos.yy][objectPos.xx+1] = EMPTY
		case EAST:
			doMove(worldMap, Position{objectPos.xx + 2, objectPos.yy}, move)
			worldMap[objectPos.yy][objectPos.xx+2] = RIGHTBOX
			worldMap[objectPos.yy][objectPos.xx+1] = LEFTBOX
			worldMap[objectPos.yy][objectPos.xx] = EMPTY
		case SOUTH:
			doMove(worldMap, Position{objectPos.xx, objectPos.yy + 1}, move)
			worldMap[objectPos.yy+1][objectPos.xx] = LEFTBOX
			worldMap[objectPos.yy][objectPos.xx] = EMPTY
			doMove(worldMap, Position{objectPos.xx + 1, objectPos.yy + 1}, move)
			worldMap[objectPos.yy+1][objectPos.xx+1] = RIGHTBOX
			worldMap[objectPos.yy][objectPos.xx+1] = EMPTY
		case WEST:
			doMove(worldMap, Position{objectPos.xx - 1, objectPos.yy}, move)
			worldMap[objectPos.yy][objectPos.xx-1] = LEFTBOX
			worldMap[objectPos.yy][objectPos.xx] = RIGHTBOX
			worldMap[objectPos.yy][objectPos.xx+1] = EMPTY
		}
	} else if worldMap[objectPos.yy][objectPos.xx] == RIGHTBOX {
		switch move {
		case NORTH:
			doMove(worldMap, Position{objectPos.xx - 1, objectPos.yy - 1}, move)
			worldMap[objectPos.yy-1][objectPos.xx-1] = LEFTBOX
			worldMap[objectPos.yy][objectPos.xx-1] = EMPTY
			doMove(worldMap, Position{objectPos.xx, objectPos.yy - 1}, move)
			worldMap[objectPos.yy-1][objectPos.xx] = RIGHTBOX
			worldMap[objectPos.yy][objectPos.xx] = EMPTY
		case EAST:
			doMove(worldMap, Position{objectPos.xx + 1, objectPos.yy}, move)
			worldMap[objectPos.yy][objectPos.xx+1] = RIGHTBOX
			worldMap[objectPos.yy][objectPos.xx] = LEFTBOX
			worldMap[objectPos.yy][objectPos.xx-1] = EMPTY
		case SOUTH:
			doMove(worldMap, Position{objectPos.xx - 1, objectPos.yy + 1}, move)
			worldMap[objectPos.yy+1][objectPos.xx-1] = LEFTBOX
			worldMap[objectPos.yy][objectPos.xx-1] = EMPTY
			doMove(worldMap, Position{objectPos.xx, objectPos.yy + 1}, move)
			worldMap[objectPos.yy+1][objectPos.xx] = RIGHTBOX
			worldMap[objectPos.yy][objectPos.xx] = EMPTY
		case WEST:
			doMove(worldMap, Position{objectPos.xx - 2, objectPos.yy}, move)
			worldMap[objectPos.yy][objectPos.xx-2] = LEFTBOX
			worldMap[objectPos.yy][objectPos.xx-1] = RIGHTBOX
			worldMap[objectPos.yy][objectPos.xx] = EMPTY
		}
	}
	return newPos
}

func tryMove2(worldMap [][]int, objectPos Position, move int) (robotPos Position) {
	if move == NOMOVE { //newline
		return
	}

	if canMove(worldMap, objectPos, move) {
		robotPos = doMove(worldMap, objectPos, move)
	} else {
		robotPos = objectPos
	}
	return robotPos
}
