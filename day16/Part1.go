package day16

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
	"strings"
)

type Position struct {
	xx int
	yy int
}

type PositionFacing struct {
	xx  int
	yy  int
	dir int
}

const (
	ANY = iota
	NORTH
	EAST
	SOUTH
	WEST
)

func Part1(fileName string) int {
	defer utils.Timer("16-1")()

	mazeMap, startPosition, endPosition := readFile(fileName)

	scoreMap := make(map[int][]PositionFacing)
	scoreMap[0] = []PositionFacing{{endPosition.xx, endPosition.yy, NORTH},
		{endPosition.xx, endPosition.yy, EAST},
		{endPosition.xx, endPosition.yy, SOUTH},
		{endPosition.xx, endPosition.yy, WEST}}

	positionScoreMap := make(map[PositionFacing]int)
	positionScoreMap[PositionFacing{endPosition.xx, endPosition.yy, NORTH}] = 0
	positionScoreMap[PositionFacing{endPosition.xx, endPosition.yy, EAST}] = 0
	positionScoreMap[PositionFacing{endPosition.xx, endPosition.yy, SOUTH}] = 0
	positionScoreMap[PositionFacing{endPosition.xx, endPosition.yy, WEST}] = 0

	makeMapScore(scoreMap, mazeMap, positionScoreMap, startPosition)

	return positionScoreMap[startPosition]
}

func makeMapScore(scoreMap map[int][]PositionFacing, mazeMap [][]bool, positionScoreMap map[PositionFacing]int, startPosition PositionFacing) {
	lowestScore := 0
	for {
		for _, lowestSquare := range scoreMap[lowestScore] {
			//movement
			switch lowestSquare.dir {
			case NORTH:
				if !mazeMap[lowestSquare.yy+1][lowestSquare.xx] {
					positionSouthOfMe := PositionFacing{lowestSquare.xx, lowestSquare.yy + 1, NORTH}
					if _, exists := positionScoreMap[positionSouthOfMe]; !exists {
						scoreMap[lowestScore+1] = append(scoreMap[lowestScore+1], positionSouthOfMe)
						positionScoreMap[positionSouthOfMe] = lowestScore + 1
					}
				}
			case EAST:
				if !mazeMap[lowestSquare.yy][lowestSquare.xx-1] {
					positionWestOfMe := PositionFacing{lowestSquare.xx - 1, lowestSquare.yy, EAST}
					if _, exists := positionScoreMap[positionWestOfMe]; !exists {
						scoreMap[lowestScore+1] = append(scoreMap[lowestScore+1], positionWestOfMe)
						positionScoreMap[positionWestOfMe] = lowestScore + 1
					}
				}
			case SOUTH:
				if !mazeMap[lowestSquare.yy-1][lowestSquare.xx] {
					positionNorthOfMe := PositionFacing{lowestSquare.xx, lowestSquare.yy - 1, SOUTH}
					if _, exists := positionScoreMap[positionNorthOfMe]; !exists {
						scoreMap[lowestScore+1] = append(scoreMap[lowestScore+1], positionNorthOfMe)
						positionScoreMap[positionNorthOfMe] = lowestScore + 1
					}
				}
			case WEST:
				if !mazeMap[lowestSquare.yy][lowestSquare.xx+1] {
					positionEastOfMe := PositionFacing{lowestSquare.xx + 1, lowestSquare.yy, WEST}
					if _, exists := positionScoreMap[positionEastOfMe]; !exists {
						scoreMap[lowestScore+1] = append(scoreMap[lowestScore+1], positionEastOfMe)
						positionScoreMap[positionEastOfMe] = lowestScore + 1
					}
				}
			}
			//turns ... a bit of duplication here, maybe tidy up
			switch lowestSquare.dir {
			case NORTH:
				eastTurn := PositionFacing{lowestSquare.xx, lowestSquare.yy, EAST}
				if _, exists := positionScoreMap[eastTurn]; !exists {
					scoreMap[lowestScore+1000] = append(scoreMap[lowestScore+1000], eastTurn)
					positionScoreMap[eastTurn] = lowestScore + 1000
				}
				westTurn := PositionFacing{lowestSquare.xx, lowestSquare.yy, WEST}
				if _, exists := positionScoreMap[westTurn]; !exists {
					scoreMap[lowestScore+1000] = append(scoreMap[lowestScore+1000], westTurn)
					positionScoreMap[westTurn] = lowestScore + 1000
				}
			case EAST:
				southTurn := PositionFacing{lowestSquare.xx, lowestSquare.yy, SOUTH}
				if _, exists := positionScoreMap[southTurn]; !exists {
					scoreMap[lowestScore+1000] = append(scoreMap[lowestScore+1000], southTurn)
					positionScoreMap[southTurn] = lowestScore + 1000
				}
				northTurn := PositionFacing{lowestSquare.xx, lowestSquare.yy, NORTH}
				if _, exists := positionScoreMap[northTurn]; !exists {
					scoreMap[lowestScore+1000] = append(scoreMap[lowestScore+1000], northTurn)
					positionScoreMap[northTurn] = lowestScore + 1000
				}
			case SOUTH:
				westTurn := PositionFacing{lowestSquare.xx, lowestSquare.yy, WEST}
				if _, exists := positionScoreMap[westTurn]; !exists {
					scoreMap[lowestScore+1000] = append(scoreMap[lowestScore+1000], westTurn)
					positionScoreMap[westTurn] = lowestScore + 1000
				}
				eastTurn := PositionFacing{lowestSquare.xx, lowestSquare.yy, EAST}
				if _, exists := positionScoreMap[eastTurn]; !exists {
					scoreMap[lowestScore+1000] = append(scoreMap[lowestScore+1000], eastTurn)
					positionScoreMap[eastTurn] = lowestScore + 1000
				}
			case WEST:
				northTurn := PositionFacing{lowestSquare.xx, lowestSquare.yy, NORTH}
				if _, exists := positionScoreMap[northTurn]; !exists {
					scoreMap[lowestScore+1000] = append(scoreMap[lowestScore+1000], northTurn)
					positionScoreMap[northTurn] = lowestScore + 1000
				}
				southTurn := PositionFacing{lowestSquare.xx, lowestSquare.yy, SOUTH}
				if _, exists := positionScoreMap[southTurn]; !exists {
					scoreMap[lowestScore+1000] = append(scoreMap[lowestScore+1000], southTurn)
					positionScoreMap[southTurn] = lowestScore + 1000
				}
			}

		}

		if _, ok := positionScoreMap[startPosition]; ok {
			break
		}

		lowestScore++
	}
}

func readFile(fileName string) (mazeMap [][]bool, startPosition PositionFacing, endPosition Position) {
	mazeLines, _ := utils.ReadFileAsLines(fileName)

	mazeMap = make([][]bool, len(mazeLines))
	for ii, mazeLine := range mazeLines {
		mazeMap[ii] = make([]bool, len(mazeLine))
		squares := strings.Split(mazeLine, "")
		for jj, square := range squares {
			switch square {
			case "#":
				mazeMap[ii][jj] = true
			case "E":
				endPosition.xx, endPosition.yy = jj, ii
			case "S":
				startPosition.xx, startPosition.yy, startPosition.dir = jj, ii, EAST
			}
		}
	}

	return mazeMap, startPosition, endPosition
}
