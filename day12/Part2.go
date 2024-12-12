package day12

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
)

func Part2() int {
	defer utils.Timer("12-2")()
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
		//sides are too hard, let's just do corners
		numCorners := 0
		for _, plot := range neighborhood {
			hasSameAbove, hasSameLeft, hasSameRight, hasSameBelow := false, false, false, false
			hasSameUpperLeft, hasSameUpperRight, hasSameLowerLeft, hasSameLowerRight := false, false, false, false

			if plot.row > 0 && plot.column > 0 {
				hasSameUpperLeft = directory[plot.row][plot.column] == directory[plot.row-1][plot.column-1]
			}
			if plot.column > 0 {
				hasSameAbove = directory[plot.row][plot.column] == directory[plot.row][plot.column-1]
			}
			if plot.row < (numRows-1) && plot.column > 0 {
				hasSameUpperRight = directory[plot.row][plot.column] == directory[plot.row+1][plot.column-1]
			}
			if plot.row < (numRows - 1) {
				hasSameRight = directory[plot.row][plot.column] == directory[plot.row+1][plot.column]
			}
			if plot.row < (numRows-1) && plot.column < (numColumns-1) {
				hasSameLowerRight = directory[plot.row][plot.column] == directory[plot.row+1][plot.column+1]
			}
			if plot.column < (numColumns - 1) {
				hasSameBelow = directory[plot.row][plot.column] == directory[plot.row][plot.column+1]
			}
			if plot.row > 0 && plot.column < (numColumns-1) {
				hasSameLowerLeft = directory[plot.row][plot.column] == directory[plot.row-1][plot.column+1]
			}
			if plot.row > 0 {
				hasSameLeft = directory[plot.row][plot.column] == directory[plot.row-1][plot.column]
			}

			//upper left convex
			if !hasSameAbove && !hasSameLeft {
				numCorners++
			}
			//upper right convex
			if !hasSameAbove && !hasSameRight {
				numCorners++
			}
			//lower right convex
			if !hasSameBelow && !hasSameRight {
				numCorners++
			}
			//lower left convex
			if !hasSameBelow && !hasSameLeft {
				numCorners++
			}
			//upper left concave
			if hasSameAbove && hasSameLeft && !hasSameUpperLeft {
				numCorners++
			}
			//upper right concave
			if hasSameAbove && hasSameRight && !hasSameUpperRight {
				numCorners++
			}
			//lower right concave
			if hasSameBelow && hasSameRight && !hasSameLowerRight {
				numCorners++
			}
			//lower left concave
			if hasSameBelow && hasSameLeft && !hasSameLowerLeft {
				numCorners++
			}
		}

		result += area * numCorners
	}

	return result
}
