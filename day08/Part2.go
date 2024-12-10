package day08

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
	"strings"
)

func Part2() int {
	defer utils.Timer("8-2")()
	var fileName = "day08/day08.txt"

	antennas := make(map[string][]Position)
	antinodes := make(map[Position]string)

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

				deltaRow := (secondAntenna.row - firstAntenna.row)
				deltaColumn := (secondAntenna.column - firstAntenna.column)

				//I feel pleased that this works
				for antinodeRow, antinodeColumn := firstAntenna.row, firstAntenna.column; inBounds(antinodeRow, antinodeColumn, maxRow, maxColumn); antinodeRow, antinodeColumn = antinodeRow-deltaRow, antinodeColumn-deltaColumn {
					placeAntinode(antinodeRow, antinodeColumn, antinodes, label)
				}

				for antinodeRow, antinodeColumn := firstAntenna.row, firstAntenna.column; inBounds(antinodeRow, antinodeColumn, maxRow, maxColumn); antinodeRow, antinodeColumn = antinodeRow+deltaRow, antinodeColumn+deltaColumn {
					placeAntinode(antinodeRow, antinodeColumn, antinodes, label)
				}
			}
		}
	}

	return len(antinodes)
}

func inBounds(row, column, maxRow, maxColumn int) bool {
	return row >= 0 && row <= maxRow && column >= 0 && column <= maxColumn
}

func placeAntinode(antinodeRow int, antinodeColumn int, antinodes map[Position]string, label string) {
	antinodePosition := Position{row: antinodeRow, column: antinodeColumn}
	antinodes[antinodePosition] = label
}
