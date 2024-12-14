package day14

import (
	"fmt"
	"github.com/adriananderson/2024-advent-of-code/utils"
	"strconv"
	"strings"
)

type Position struct {
	xx int
	yy int
}
type Bound struct {
	min int
	max int
}

func Part1() int {
	defer utils.Timer("14-1")()
	var fileName = "day14/day14.txt"

	//maxXX, maxYY := 11, 7
	maxXX, maxYY := 101, 103
	robotPos, robotVel := readFile(fileName)

	//printRobots(robotPos, maxXX, maxYY)
	for ii := 0; ii < 100; ii++ {
		advanceRobots(robotPos, robotVel, maxXX, maxYY)
		//printRobots(robotPos, maxXX, maxYY)
	}

	q1Count, q2Count, q3Count, q4Count := countQuadrants(maxXX, maxYY, robotPos)

	score := q1Count * q2Count * q3Count * q4Count

	return score
}

func printRobots(robotPositions []Position, maxXX int, maxYY int) {
	grid := make([][]int, maxXX)
	for xx := 0; xx < maxXX; xx++ {
		grid[xx] = make([]int, maxYY)
	}
	for _, pos := range robotPositions {
		grid[pos.xx][pos.yy]++
	}
	fmt.Println()
	for yy := 0; yy < maxYY; yy++ {
		for xx := 0; xx < maxXX; xx++ {
			if grid[xx][yy] > 0 {
				fmt.Printf("%d", grid[xx][yy])
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func countQuadrants(maxXX int, maxYY int, robotPos []Position) (int, int, int, int) {
	q1xBounds := Bound{0, maxXX/2 - 1}
	q1yBounds := Bound{0, maxYY/2 - 1}

	q2xBounds := Bound{maxXX/2 + 1, maxXX - 1}
	q2yBounds := Bound{0, maxYY/2 - 1}

	q3xBounds := Bound{0, maxXX/2 - 1}
	q3yBounds := Bound{maxYY/2 + 1, maxYY - 1}

	q4xBounds := Bound{maxXX/2 + 1, maxXX - 1}
	q4yBounds := Bound{maxYY/2 + 1, maxYY - 1}

	q1Count, q2Count, q3Count, q4Count := 0, 0, 0, 0

	for ii := 0; ii < len(robotPos); ii++ {
		if robotPos[ii].xx >= q1xBounds.min && robotPos[ii].xx <= q1xBounds.max && robotPos[ii].yy >= q1yBounds.min && robotPos[ii].yy <= q1yBounds.max {
			q1Count++
		}
		if robotPos[ii].xx >= q2xBounds.min && robotPos[ii].xx <= q2xBounds.max && robotPos[ii].yy >= q2yBounds.min && robotPos[ii].yy <= q2yBounds.max {
			q2Count++
		}
		if robotPos[ii].xx >= q3xBounds.min && robotPos[ii].xx <= q3xBounds.max && robotPos[ii].yy >= q3yBounds.min && robotPos[ii].yy <= q3yBounds.max {
			q3Count++
		}
		if robotPos[ii].xx >= q4xBounds.min && robotPos[ii].xx <= q4xBounds.max && robotPos[ii].yy >= q4yBounds.min && robotPos[ii].yy <= q4yBounds.max {
			q4Count++
		}
	}
	return q1Count, q2Count, q3Count, q4Count
}

func advanceRobots(robotPos []Position, robotVel []Position, maxXX int, maxYY int) {
	for jj := 0; jj < len(robotPos); jj++ {
		xx := robotPos[jj].xx + robotVel[jj].xx
		if xx >= maxXX {
			xx -= maxXX
		} else if xx < 0 {
			xx += maxXX
		}
		robotPos[jj].xx = xx
		yy := robotPos[jj].yy + robotVel[jj].yy
		if yy >= maxYY {
			yy -= maxYY
		} else if yy < 0 {
			yy += maxYY
		}
		robotPos[jj].yy = yy
	}
}

func readFile(fileName string) (robotPositions []Position, robotVelocities []Position) {
	lines, _ := utils.ReadFileAsLines(fileName)

	robotPositions = make([]Position, len(lines))
	robotVelocities = make([]Position, len(lines))
	for ii, line := range lines {
		parts := strings.Split(line, " ")

		locParts := strings.Split(parts[0], "=")
		locBits := strings.Split(locParts[1], ",")
		xx, _ := strconv.Atoi(locBits[0])
		yy, _ := strconv.Atoi(locBits[1])
		robotPositions[ii] = Position{xx: xx, yy: yy}

		velParts := strings.Split(parts[1], "=")
		velBits := strings.Split(velParts[1], ",")
		xx, _ = strconv.Atoi(velBits[0])
		yy, _ = strconv.Atoi(velBits[1])
		robotVelocities[ii] = Position{xx: xx, yy: yy}
	}

	return robotPositions, robotVelocities
}
