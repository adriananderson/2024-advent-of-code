package day08

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
	"strconv"
	"strings"
)

type Position struct {
	row    int
	column int
}

func Part1() int {
	defer utils.Timer("8-1")()
	var fileName = "day08/day08.txt"

	antennas := make(map[string][]Position)
	antinodes := make(map[string]string)

	maxColumn, maxRow := 0, 0

	//read in antennas
	if fileLines, err := utils.ReadFileAsLines(fileName); err == nil {
		for row, line := range fileLines {
			if len(line) < maxRow {
				break
			}
			maxRow = row
			letters := strings.Split(line, "")
			maxColumn = len(line) - 1
			for column, letter := range letters {
				if letter != "." {
					if antennas[letter] == nil {
						antennas[letter] = make([]Position, 0)
					}
					antennas[letter] = append(antennas[letter], Position{row: row, column: column})
				}
			}
		}
	}

	for label, antennaPositions := range antennas {
		for ii := 0; ii < len(antennaPositions); ii++ {
			for jj := ii + 1; jj < len(antennaPositions); jj++ {
				//calc antinodes
				firstAntenna := antennaPositions[ii]
				secondAntenna := antennaPositions[jj]

				antinodeRow := firstAntenna.row - (secondAntenna.row - firstAntenna.row)
				antinodeColumn := firstAntenna.column - (secondAntenna.column - firstAntenna.column)

				if antinodeRow >= 0 && antinodeColumn >= 0 && antinodeRow <= maxRow && antinodeColumn <= maxColumn {
					key := strconv.Itoa(antinodeRow) + "-" + strconv.Itoa(antinodeColumn)
					antinodes[key] = label
				}

				antinodeRow = secondAntenna.row + (secondAntenna.row - firstAntenna.row)
				antinodeColumn = secondAntenna.column + (secondAntenna.column - firstAntenna.column)

				if antinodeRow >= 0 && antinodeColumn >= 0 && antinodeRow <= maxRow && antinodeColumn <= maxColumn {
					key := strconv.Itoa(antinodeRow) + "-" + strconv.Itoa(antinodeColumn)
					antinodes[key] = label
				}
			}
		}
	}

	return len(antinodes)
}
