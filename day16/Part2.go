package day16

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
)

func Part2(fileName string) int {
	defer utils.Timer("16-2")()

	mazeMap, startPosition, endPosition := readFile(fileName)

	scoreMap := make(map[int][]PositionFacing)
	scoreMap[0] = []PositionFacing{PositionFacing{endPosition.xx, endPosition.yy, NORTH},
		PositionFacing{endPosition.xx, endPosition.yy, EAST},
		PositionFacing{endPosition.xx, endPosition.yy, SOUTH},
		PositionFacing{endPosition.xx, endPosition.yy, WEST}}

	positionScoreMap := make(map[PositionFacing]int)
	positionScoreMap[PositionFacing{endPosition.xx, endPosition.yy, NORTH}] = 0
	positionScoreMap[PositionFacing{endPosition.xx, endPosition.yy, EAST}] = 0
	positionScoreMap[PositionFacing{endPosition.xx, endPosition.yy, SOUTH}] = 0
	positionScoreMap[PositionFacing{endPosition.xx, endPosition.yy, WEST}] = 0

	positionOriginMap := make(map[PositionFacing]map[Position]bool)

	makeMapScore2(scoreMap, mazeMap, positionScoreMap, positionOriginMap, startPosition)

	return len(positionOriginMap[startPosition])
}

func makeMapScore2(scoreMap map[int][]PositionFacing, mazeMap [][]bool, positionScoreMap map[PositionFacing]int, positionOriginMap map[PositionFacing]map[Position]bool, startPosition PositionFacing) {
	lowestScore := 0
	for {
		for _, lowestPositionFacing := range scoreMap[lowestScore] {
			if _, exists := positionOriginMap[lowestPositionFacing]; !exists {
				positionOriginMap[lowestPositionFacing] = make(map[Position]bool)
				thisPosition := Position{lowestPositionFacing.xx, lowestPositionFacing.yy}
				positionOriginMap[lowestPositionFacing][thisPosition] = true
			}

			//movement
			switch lowestPositionFacing.dir {
			case NORTH:
				if !mazeMap[lowestPositionFacing.yy+1][lowestPositionFacing.xx] {
					positionSouthOfMe := PositionFacing{lowestPositionFacing.xx, lowestPositionFacing.yy + 1, NORTH}
					score, exists := positionScoreMap[positionSouthOfMe]
					if !exists { //first
						scoreMap[lowestScore+1] = append(scoreMap[lowestScore+1], positionSouthOfMe)
						positionScoreMap[positionSouthOfMe] = lowestScore + 1
						positionOriginMap[positionSouthOfMe] = make(map[Position]bool)
						thatPosition := Position{positionSouthOfMe.xx, positionSouthOfMe.yy}
						positionOriginMap[positionSouthOfMe][thatPosition] = true
						for key, _ := range positionOriginMap[lowestPositionFacing] {
							positionOriginMap[positionSouthOfMe][key] = true
						}
					} else if score == lowestScore+1 { //tie
						for key, _ := range positionOriginMap[lowestPositionFacing] {
							positionOriginMap[positionSouthOfMe][key] = true
						}
					}
				}
			case EAST:
				if !mazeMap[lowestPositionFacing.yy][lowestPositionFacing.xx-1] {
					positionWestOfMe := PositionFacing{lowestPositionFacing.xx - 1, lowestPositionFacing.yy, EAST}
					score, exists := positionScoreMap[positionWestOfMe]
					if !exists { //first
						scoreMap[lowestScore+1] = append(scoreMap[lowestScore+1], positionWestOfMe)
						positionScoreMap[positionWestOfMe] = lowestScore + 1
						positionOriginMap[positionWestOfMe] = make(map[Position]bool)
						thatPosition := Position{positionWestOfMe.xx, positionWestOfMe.yy}
						positionOriginMap[positionWestOfMe][thatPosition] = true
						for key, _ := range positionOriginMap[lowestPositionFacing] {
							positionOriginMap[positionWestOfMe][key] = true
						}
					} else if score == lowestScore+1 { //tie
						for key, _ := range positionOriginMap[lowestPositionFacing] {
							positionOriginMap[positionWestOfMe][key] = true
						}
					}
				}
			case SOUTH:
				if !mazeMap[lowestPositionFacing.yy-1][lowestPositionFacing.xx] {
					positionNorthOfMe := PositionFacing{lowestPositionFacing.xx, lowestPositionFacing.yy - 1, SOUTH}
					score, exists := positionScoreMap[positionNorthOfMe]
					if !exists { //first
						scoreMap[lowestScore+1] = append(scoreMap[lowestScore+1], positionNorthOfMe)
						positionScoreMap[positionNorthOfMe] = lowestScore + 1
						positionOriginMap[positionNorthOfMe] = make(map[Position]bool)
						thatPosition := Position{positionNorthOfMe.xx, positionNorthOfMe.yy}
						positionOriginMap[positionNorthOfMe][thatPosition] = true
						for key, _ := range positionOriginMap[lowestPositionFacing] {
							positionOriginMap[positionNorthOfMe][key] = true
						}
					} else if score == lowestScore+1 { //tie
						for key, _ := range positionOriginMap[lowestPositionFacing] {
							positionOriginMap[positionNorthOfMe][key] = true
						}
					}
				}
			case WEST:
				if !mazeMap[lowestPositionFacing.yy][lowestPositionFacing.xx+1] {
					positionEastOfMe := PositionFacing{lowestPositionFacing.xx + 1, lowestPositionFacing.yy, WEST}
					score, exists := positionScoreMap[positionEastOfMe]
					if !exists { //first
						scoreMap[lowestScore+1] = append(scoreMap[lowestScore+1], positionEastOfMe)
						positionScoreMap[positionEastOfMe] = lowestScore + 1
						positionOriginMap[positionEastOfMe] = make(map[Position]bool)
						thatPosition := Position{positionEastOfMe.xx, positionEastOfMe.yy}
						positionOriginMap[positionEastOfMe][thatPosition] = true
						for key, _ := range positionOriginMap[lowestPositionFacing] {
							positionOriginMap[positionEastOfMe][key] = true
						}
					} else if score == lowestScore+1 { //tie
						for key, _ := range positionOriginMap[lowestPositionFacing] {
							positionOriginMap[positionEastOfMe][key] = true
						}
					}
				}
			}
			//turns ... a bit of duplication here, maybe tidy up
			switch lowestPositionFacing.dir {
			case NORTH:
				eastTurn := PositionFacing{lowestPositionFacing.xx, lowestPositionFacing.yy, EAST}
				if _, exists := positionScoreMap[eastTurn]; !exists {
					scoreMap[lowestScore+1000] = append(scoreMap[lowestScore+1000], eastTurn)
					positionScoreMap[eastTurn] = lowestScore + 1000
					positionOriginMap[eastTurn] = make(map[Position]bool)
					thatPosition := Position{eastTurn.xx, eastTurn.yy}
					positionOriginMap[eastTurn][thatPosition] = true
					for key, _ := range positionOriginMap[lowestPositionFacing] {
						positionOriginMap[eastTurn][key] = true
					}
				}
				westTurn := PositionFacing{lowestPositionFacing.xx, lowestPositionFacing.yy, WEST}
				if _, exists := positionScoreMap[westTurn]; !exists {
					scoreMap[lowestScore+1000] = append(scoreMap[lowestScore+1000], westTurn)
					positionScoreMap[westTurn] = lowestScore + 1000
					positionOriginMap[westTurn] = make(map[Position]bool)
					thatPosition := Position{westTurn.xx, westTurn.yy}
					positionOriginMap[westTurn][thatPosition] = true
					for key, _ := range positionOriginMap[lowestPositionFacing] {
						positionOriginMap[westTurn][key] = true
					}
				}
			case EAST:
				southTurn := PositionFacing{lowestPositionFacing.xx, lowestPositionFacing.yy, SOUTH}
				if _, exists := positionScoreMap[southTurn]; !exists {
					scoreMap[lowestScore+1000] = append(scoreMap[lowestScore+1000], southTurn)
					positionScoreMap[southTurn] = lowestScore + 1000
					positionOriginMap[southTurn] = make(map[Position]bool)
					thatPosition := Position{southTurn.xx, southTurn.yy}
					positionOriginMap[southTurn][thatPosition] = true
					for key, _ := range positionOriginMap[lowestPositionFacing] {
						positionOriginMap[southTurn][key] = true
					}
				}
				northTurn := PositionFacing{lowestPositionFacing.xx, lowestPositionFacing.yy, NORTH}
				if _, exists := positionScoreMap[northTurn]; !exists {
					scoreMap[lowestScore+1000] = append(scoreMap[lowestScore+1000], northTurn)
					positionScoreMap[northTurn] = lowestScore + 1000
					positionOriginMap[northTurn] = make(map[Position]bool)
					thatPosition := Position{northTurn.xx, northTurn.yy}
					positionOriginMap[northTurn][thatPosition] = true
					for key, _ := range positionOriginMap[lowestPositionFacing] {
						positionOriginMap[northTurn][key] = true
					}
				}
			case SOUTH:
				westTurn := PositionFacing{lowestPositionFacing.xx, lowestPositionFacing.yy, WEST}
				if _, exists := positionScoreMap[westTurn]; !exists {
					scoreMap[lowestScore+1000] = append(scoreMap[lowestScore+1000], westTurn)
					positionScoreMap[westTurn] = lowestScore + 1000
					positionOriginMap[westTurn] = make(map[Position]bool)
					thatPosition := Position{westTurn.xx, westTurn.yy}
					positionOriginMap[westTurn][thatPosition] = true
					for key, _ := range positionOriginMap[lowestPositionFacing] {
						positionOriginMap[westTurn][key] = true
					}
				}
				eastTurn := PositionFacing{lowestPositionFacing.xx, lowestPositionFacing.yy, EAST}
				if _, exists := positionScoreMap[eastTurn]; !exists {
					scoreMap[lowestScore+1000] = append(scoreMap[lowestScore+1000], eastTurn)
					positionScoreMap[eastTurn] = lowestScore + 1000
					positionOriginMap[eastTurn] = make(map[Position]bool)
					thatPosition := Position{eastTurn.xx, eastTurn.yy}
					positionOriginMap[eastTurn][thatPosition] = true
					for key, _ := range positionOriginMap[lowestPositionFacing] {
						positionOriginMap[eastTurn][key] = true
					}
				}
			case WEST:
				northTurn := PositionFacing{lowestPositionFacing.xx, lowestPositionFacing.yy, NORTH}
				if _, exists := positionScoreMap[northTurn]; !exists {
					scoreMap[lowestScore+1000] = append(scoreMap[lowestScore+1000], northTurn)
					positionScoreMap[northTurn] = lowestScore + 1000
					positionOriginMap[northTurn] = make(map[Position]bool)
					thatPosition := Position{northTurn.xx, northTurn.yy}
					positionOriginMap[northTurn][thatPosition] = true
					for key, _ := range positionOriginMap[lowestPositionFacing] {
						positionOriginMap[northTurn][key] = true
					}
				}
				southTurn := PositionFacing{lowestPositionFacing.xx, lowestPositionFacing.yy, SOUTH}
				if _, exists := positionScoreMap[southTurn]; !exists {
					scoreMap[lowestScore+1000] = append(scoreMap[lowestScore+1000], southTurn)
					positionScoreMap[southTurn] = lowestScore + 1000
					positionOriginMap[southTurn] = make(map[Position]bool)
					thatPosition := Position{southTurn.xx, southTurn.yy}
					positionOriginMap[southTurn][thatPosition] = true
					for key, _ := range positionOriginMap[lowestPositionFacing] {
						positionOriginMap[southTurn][key] = true
					}
				}
			}

		}

		if _, ok := positionScoreMap[startPosition]; ok {
			break
		}

		lowestScore++
	}
}
