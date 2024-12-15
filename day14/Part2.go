package day14

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
)

func Part2() int {
	defer utils.Timer("14-2")()
	var fileName = "day14/day14.txt"

	//maxXX, maxYY := 11, 7
	maxXX, maxYY := 101, 103
	robotPos, robotVel := readFile(fileName)

	//printRobots(robotPos, maxXX, maxYY)
	//looking for minimum safety factor because that's what reddit says
	bestSafetyFactor := 99999999999999
	lowestTime := 0
	for ii := 0; ii < 10403; ii++ {
		advanceRobots(robotPos, robotVel, maxXX, maxYY)
		//printRobots(robotPos, maxXX, maxYY)
		q1Count, q2Count, q3Count, q4Count := countQuadrants(maxXX, maxYY, robotPos)

		score := q1Count * q2Count * q3Count * q4Count
		if score < bestSafetyFactor {
			bestSafetyFactor = score
			lowestTime = ii + 1
			//fmt.Printf("better safety %d %d\n", ii, score)
			//printRobots(robotPos, maxXX, maxYY)
		}
	}

	return lowestTime
}
