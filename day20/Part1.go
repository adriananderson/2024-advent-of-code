package day20

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
	"strings"
)

type Offset struct {
	position Position
	distance int
}

type Position struct {
	xx int
	yy int
}

func Part1(fileName string, threshold int) int {
	defer utils.Timer("20-1")()

	mazeMap, startPosition, finishPosition := readFile(fileName)

	_, finishPositionMap := scoreMap(mazeMap, finishPosition, startPosition)
	startMap, _ := scoreMap(mazeMap, startPosition, finishPosition)

	noCheatScore := finishPositionMap[startPosition]
	//fmt.Printf("noCheatScore: %d\n", noCheatScore)

	result := 0
	//look at points closest to start up to finish line
	//for startScore := 0; startScore < (noCheatScore - threshold); startScore++ {
	for startScore := 0; startScore < noCheatScore; startScore++ {
		//check if one short cut is shorter
		//           1
		//          8#2
		//         7#0#3
		//          6#4
		//           5
		//going from point 0 -> {1-8}
		for _, lowestPosition := range startMap[startScore] {
			positionOne := Position{lowestPosition.xx, lowestPosition.yy - 2}
			if finishScore, finishExists := finishPositionMap[positionOne]; finishExists {
				if (2 + finishScore + startScore) <= (noCheatScore - threshold) {
					result++
				}
			}
			positionTwo := Position{lowestPosition.xx + 1, lowestPosition.yy - 1}
			if finishScore, finishExists := finishPositionMap[positionTwo]; finishExists {
				if (2 + finishScore + startScore) <= (noCheatScore - threshold) {
					result++
				}
			}
			positionThree := Position{lowestPosition.xx + 2, lowestPosition.yy}
			if finishScore, finishExists := finishPositionMap[positionThree]; finishExists {
				if (2 + finishScore + startScore) <= (noCheatScore - threshold) {
					result++
				}
			}
			positionFour := Position{lowestPosition.xx + 1, lowestPosition.yy + 1}
			if finishScore, finishExists := finishPositionMap[positionFour]; finishExists {
				if (2 + finishScore + startScore) <= (noCheatScore - threshold) {
					result++
				}
			}
			positionFive := Position{lowestPosition.xx, lowestPosition.yy + 2}
			if finishScore, finishExists := finishPositionMap[positionFive]; finishExists {
				if (2 + finishScore + startScore) <= (noCheatScore - threshold) {
					result++
				}
			}
			positionSix := Position{lowestPosition.xx - 1, lowestPosition.yy + 1}
			if finishScore, finishExists := finishPositionMap[positionSix]; finishExists {
				if (2 + finishScore + startScore) <= (noCheatScore - threshold) {
					result++
				}
			}
			positionSeven := Position{lowestPosition.xx - 2, lowestPosition.yy}
			if finishScore, finishExists := finishPositionMap[positionSeven]; finishExists {
				if (2 + finishScore + startScore) <= (noCheatScore - threshold) {
					result++
				}
			}
			positionEight := Position{lowestPosition.xx - 1, lowestPosition.yy - 1}
			if finishScore, finishExists := finishPositionMap[positionEight]; finishExists {
				if (2 + finishScore + startScore) <= (noCheatScore - threshold) {
					result++
				}
			}
		}
	}

	return result
}

func scoreMap(mazeMap [][]bool, startPosition Position, endPosition Position) (scoreMap map[int][]Position, positionMap map[Position]int) {
	scoreMap = make(map[int][]Position)
	scoreMap[0] = []Position{startPosition}
	positionMap = make(map[Position]int)
	positionMap[startPosition] = 0
	lowestScore := 0
	for {
		for _, lowestSquare := range scoreMap[lowestScore] {
			if !mazeMap[lowestSquare.yy-1][lowestSquare.xx] { //NORTH
				northPosition := Position{lowestSquare.xx, lowestSquare.yy - 1}
				if _, exists := positionMap[northPosition]; !exists {
					scoreMap[lowestScore+1] = append(scoreMap[lowestScore+1], northPosition)
					positionMap[northPosition] = lowestScore + 1
				}
			}
			if !mazeMap[lowestSquare.yy][lowestSquare.xx+1] { //EAST
				eastPosition := Position{lowestSquare.xx + 1, lowestSquare.yy}
				if _, exists := positionMap[eastPosition]; !exists {
					scoreMap[lowestScore+1] = append(scoreMap[lowestScore+1], eastPosition)
					positionMap[eastPosition] = lowestScore + 1
				}
			}
			if !mazeMap[lowestSquare.yy+1][lowestSquare.xx] { //SOUTH
				southPosition := Position{lowestSquare.xx, lowestSquare.yy + 1}
				if _, exists := positionMap[southPosition]; !exists {
					scoreMap[lowestScore+1] = append(scoreMap[lowestScore+1], southPosition)
					positionMap[southPosition] = lowestScore + 1
				}
			}
			if !mazeMap[lowestSquare.yy][lowestSquare.xx-1] { //WEST
				westPosition := Position{lowestSquare.xx - 1, lowestSquare.yy}
				if _, exists := positionMap[westPosition]; !exists {
					scoreMap[lowestScore+1] = append(scoreMap[lowestScore+1], westPosition)
					positionMap[westPosition] = lowestScore + 1
				}
			}
		}

		if _, exists := positionMap[endPosition]; exists {
			break
		}
		lowestScore++
	}

	return scoreMap, positionMap
}

func readFile(fileName string) (mazeMap [][]bool, startPosition Position, endPosition Position) {
	mazeLines, _ := utils.ReadFileAsLines(fileName)

	mazeMap = make([][]bool, len(mazeLines))
	for yy, mazeLine := range mazeLines {
		mazeMap[yy] = make([]bool, len(mazeLine))
		squares := strings.Split(mazeLine, "")
		for xx, square := range squares {
			switch square {
			case "#":
				mazeMap[yy][xx] = true
			case "E":
				endPosition.xx, endPosition.yy = xx, yy
			case "S":
				startPosition.xx, startPosition.yy = xx, yy
			}
		}
	}

	return mazeMap, startPosition, endPosition
}
