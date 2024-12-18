package day18

import (
	"fmt"
	"github.com/adriananderson/2024-advent-of-code/utils"
	"math"
	"strconv"
	"strings"
)

func Part2(fileName string, maxSize int, limit int) string {
	defer utils.Timer("18-2")()

	wallCoordinates := readWallCoordinates(fileName)
	wallMap := makeMaze(maxSize)

	for ii := limit; ii < len(wallCoordinates); ii++ {
		loadWallMap(wallCoordinates, wallMap, ii)
		pathLength := solveMaze(wallMap, maxSize)
		if pathLength == math.MaxInt {
			return fmt.Sprintf("%d,%d", wallCoordinates[ii-1][0], wallCoordinates[ii-1][1])
		}
	}

	return "no blockage found"
}

func readWallCoordinates(fileName string) (wallCoordinates [][]int) {
	coordLines, _ := utils.ReadFileAsLines(fileName)

	wallCoordinates = make([][]int, len(coordLines))
	for ii, coordinate := range coordLines {
		if len(coordinate) > 0 {
			wallCoordinates[ii] = make([]int, 2)
			xxyy := strings.Split(coordinate, ",")
			xx, _ := strconv.Atoi(xxyy[0])
			yy, _ := strconv.Atoi(xxyy[1])
			wallCoordinates[ii][0] = xx
			wallCoordinates[ii][1] = yy
		}
	}

	return wallCoordinates
}

func makeMaze(maxSize int) (mazeMap [][]bool) {
	mazeMap = make([][]bool, maxSize+1)
	for ii := 0; ii <= maxSize; ii++ {
		mazeMap[ii] = make([]bool, maxSize+1)
	}

	return mazeMap
}

func loadWallMap(coordLines [][]int, wallMap [][]bool, lineLimit int) {

	for ii := 0; ii < lineLimit; ii++ {
		wallMap[coordLines[ii][0]][coordLines[ii][1]] = true
	}
}
