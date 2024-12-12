package day12

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
	"strings"
)

type Position struct {
	row    int
	column int
}

func Part1() int {
	defer utils.Timer("12-1")()
	//start := time.Now()
	var fileName = "day12/day12.txt"

	garden := readFile(fileName)

	result := 0

	numRows := len(garden)
	numColumns := len(garden[0])

	// neighborhoods of plots, key is the position of the upper left plot
	regions := make(map[Position][]Position)
	directory := make([][]Position, numRows)
	makeNeighborhoods(garden, directory, numColumns, regions)

	for _, neighborhood := range regions {
		area := len(neighborhood)
		perimeter := 0
		for _, plot := range neighborhood {
			if perimeter == 0 { //first plot
				perimeter = 4
				continue
			}

			sameAsLeft := false
			sameAsAbove := false
			if plot.row > 0 {
				sameAsLeft = directory[plot.row][plot.column] == directory[plot.row-1][plot.column]
			}
			if plot.column > 0 {
				sameAsAbove = directory[plot.row][plot.column] == directory[plot.row][plot.column-1]
			}
			if sameAsLeft && sameAsAbove {
			} else if sameAsLeft {
				perimeter += 2
			} else if sameAsAbove {
				perimeter += 2
			} else {
				perimeter += 4
			}
		}

		result += area * perimeter
	}

	return result
}

func makeNeighborhoods(garden [][]rune, directory [][]Position, numColumns int, regions map[Position][]Position) {
	for ii, row := range garden {
		directory[ii] = make([]Position, numColumns)
		for jj, _ := range row {
			//for each plot, find like neighbors
			sameAsLeft := false
			sameAsAbove := false
			if ii > 0 {
				sameAsLeft = garden[ii][jj] == garden[ii-1][jj]
			}
			if jj > 0 {
				sameAsAbove = garden[ii][jj] == garden[ii][jj-1]
			}

			if sameAsLeft && sameAsAbove {
				if directory[ii-1][jj] != directory[ii][jj-1] {
					// join!!
					oldNeighborhood := directory[ii][jj-1]
					newNeighborhood := directory[ii-1][jj]
					for _, plot := range regions[oldNeighborhood] {
						//move plots to new neighborhood
						regions[newNeighborhood] = append(regions[newNeighborhood], plot)
						//update directory
						directory[plot.row][plot.column] = newNeighborhood
					}
					//demolish old neighborhood
					delete(regions, oldNeighborhood)
				}
				directory[ii][jj] = directory[ii-1][jj]
				neighborhoodIdentifier := directory[ii][jj]
				regions[neighborhoodIdentifier] = append(regions[neighborhoodIdentifier], Position{row: ii, column: jj})
			} else if sameAsLeft {
				directory[ii][jj] = directory[ii-1][jj]
				neighborhoodIdentifier := directory[ii][jj]
				regions[neighborhoodIdentifier] = append(regions[neighborhoodIdentifier], Position{row: ii, column: jj})
			} else if sameAsAbove {
				directory[ii][jj] = directory[ii][jj-1]
				neighborhoodIdentifier := directory[ii][jj]
				regions[neighborhoodIdentifier] = append(regions[neighborhoodIdentifier], Position{row: ii, column: jj})
			} else { //new region
				thisPosition := Position{row: ii, column: jj}
				directory[ii][jj] = thisPosition
				regions[thisPosition] = append(regions[thisPosition], thisPosition)
			}
		}
	}
}

func readFile(fileName string) (garden [][]rune) {
	lines, _ := utils.ReadFileAsLines(fileName)

	garden = make([][]rune, len(lines)-1)
	for ii, line := range lines {
		words := strings.Split(line, "")
		if len(words) > 0 {
			garden[ii] = make([]rune, len(words))
			for jj, word := range words {
				garden[ii][jj] = []rune(word)[0]
			}
		}
	}

	return garden
}
