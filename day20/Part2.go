package day20

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
)

func Part2(fileName string, threshold int) int {
	defer utils.Timer("20-2")()

	mazeMap, startPosition, finishPosition := readFile(fileName)

	_, finishPositionMap := scoreMap(mazeMap, finishPosition, startPosition)
	startMap, _ := scoreMap(mazeMap, startPosition, finishPosition)

	noCheatScore := finishPositionMap[startPosition]
	cheatRange := 20
	result := 0

	//look at points closest to start up to finish line
	for startScore := 0; startScore <= (noCheatScore - threshold); startScore++ {
		for _, lowestPosition := range startMap[startScore] {
			for ii := -cheatRange; ii <= cheatRange; ii++ { // watch the off-by-one error
				for jj := -cheatRange; jj <= cheatRange; jj++ {
					deltaXX := 0
					if ii < 0 {
						deltaXX = -ii
					} else {
						deltaXX = ii
					}
					deltaYY := 0
					if jj < 0 {
						deltaYY = -jj
					} else {
						deltaYY = jj
					}
					if deltaXX+deltaYY > cheatRange || (deltaXX == 0 && deltaYY == 0) { //cheat range too high
						continue
					}

					newPos := Position{lowestPosition.xx + ii, lowestPosition.yy + jj}
					if finishScore, finishExists := finishPositionMap[newPos]; finishExists {
						if (deltaXX + deltaYY + finishScore + startScore) <= (noCheatScore - threshold) {
							result++
						}
					}
				}
			}
		}
	}

	return result
}
