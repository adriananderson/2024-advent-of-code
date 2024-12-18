package day18

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
	"math"
	"strconv"
	"strings"
)

type Position struct {
	xx int
	yy int
}

func Part1(fileName string, maxSize int, limit int) int {
	defer utils.Timer("18-1")()

	wallMap := readFile(fileName, maxSize, limit)
	pathLength := solveMaze(wallMap, maxSize)

	return pathLength
}

func solveMaze(wallMap [][]bool, maxSize int) (pathLength int) {

	solutionMap := make([][]int, maxSize+1)
	for xx := range solutionMap {
		solutionMap[xx] = make([]int, maxSize+1)
		for yy := range solutionMap[xx] {
			solutionMap[xx][yy] = math.MaxInt
		}
	}
	solutionMap[maxSize][maxSize] = 0
	advance(wallMap, solutionMap, maxSize, maxSize, maxSize)

	return solutionMap[0][0]
}

func advance(wallMap [][]bool, solutionMap [][]int, xx int, yy int, maxSize int) {
	nextLength := solutionMap[xx][yy] + 1
	//north
	if yy > 0 {
		if solutionMap[xx][yy-1] > nextLength && !wallMap[xx][yy-1] {
			solutionMap[xx][yy-1] = nextLength
			advance(wallMap, solutionMap, xx, yy-1, maxSize)
		}
	}
	//east
	if xx < maxSize {
		if solutionMap[xx+1][yy] > nextLength && !wallMap[xx+1][yy] {
			solutionMap[xx+1][yy] = nextLength
			advance(wallMap, solutionMap, xx+1, yy, maxSize)
		}
	}
	//south
	if yy < maxSize {
		if solutionMap[xx][yy+1] > nextLength && !wallMap[xx][yy+1] {
			solutionMap[xx][yy+1] = nextLength
			advance(wallMap, solutionMap, xx, yy+1, maxSize)
		}
	}
	//west
	if xx > 0 {
		if solutionMap[xx-1][yy] > nextLength && !wallMap[xx-1][yy] {
			solutionMap[xx-1][yy] = nextLength
			advance(wallMap, solutionMap, xx-1, yy, maxSize)
		}
	}
}

func readFile(fileName string, maxSize int, lineLimit int) (mazeMap [][]bool) {
	coordLines, _ := utils.ReadFileAsLines(fileName)

	mazeMap = make([][]bool, maxSize+1)
	for ii := 0; ii <= maxSize; ii++ {
		mazeMap[ii] = make([]bool, maxSize+1)
	}

	for ii := 0; ii < lineLimit; ii++ {
		coords := strings.Split(coordLines[ii], ",")
		xx, _ := strconv.Atoi(coords[0])
		yy, _ := strconv.Atoi(coords[1])
		mazeMap[xx][yy] = true
	}

	return mazeMap
}
