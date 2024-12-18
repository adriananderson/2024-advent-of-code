package day18

import (
	"fmt"
	"github.com/adriananderson/2024-advent-of-code/utils"
	"math"
	"strconv"
	"strings"
)

func Part2(fileName string, maxSize int, start int) string {
	defer utils.Timer("18-2")()

	wallCoordinates := readWallCoordinates(fileName)
	wallMap := makeMaze(maxSize)

	low := start
	high := len(wallCoordinates) - 1

	for low <= high {
		mid := (low + high) / 2

		//fmt.Printf("low: %d   mid: %d   high: %d\n", low, mid, high)

		loadWallMap(wallCoordinates, wallMap, mid)
		pathLengthMid := solveMaze(wallMap, maxSize)

		loadWallMap(wallCoordinates, wallMap, mid-1)
		pathLengthPrevMid := solveMaze(wallMap, maxSize)

		if pathLengthPrevMid < math.MaxInt && pathLengthMid == math.MaxInt {
			//fmt.Printf("solved %d\n", mid)
			return fmt.Sprintf("%d,%d", wallCoordinates[mid-1][0], wallCoordinates[mid-1][1])
		} else if pathLengthPrevMid == math.MaxInt {
			//fmt.Printf("blocked %d\n", mid)
			high = mid - 2
		} else if pathLengthMid < math.MaxInt {
			//fmt.Printf("safe %d\n", mid)
			low = mid + 1
		}
	}

	return "no blockage found"
}

func readWallCoordinates(fileName string) (wallCoordinates [][]int) {
	coordLines, _ := utils.ReadFileAsLines(fileName)

	wallCoordinates = make([][]int, len(coordLines)-1)
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
	for ii := lineLimit; ii < len(coordLines); ii++ {
		wallMap[coordLines[ii][0]][coordLines[ii][1]] = false
	}
}
