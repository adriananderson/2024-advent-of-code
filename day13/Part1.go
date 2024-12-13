package day13

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
	"strconv"
	"strings"
)

type Position struct {
	row    int
	column int
}

type Machine struct {
	aa    Position
	bb    Position
	prize Position
}

func Part1() int {
	defer utils.Timer("13-1")()
	var fileName = "day13/day13-test.txt"

	machines := readFile(fileName)
	result := 0

	for _, machine := range machines {
		bestToSpend := 1000
		for ii := 0; ii <= 100; ii++ {
			for jj := 0; jj <= 100; jj++ {
				if (machine.aa.row*ii+machine.bb.row*jj) == machine.prize.row &&
					(machine.aa.column*ii+machine.bb.column*jj) == machine.prize.column {
					cost := ii*3 + jj
					//fmt.Printf("%d  solution: aa %d, bb %d   cost %d\n", zz, ii, jj, cost)
					if cost < bestToSpend {
						bestToSpend = cost
					}
				}
			}
		}
		if bestToSpend < 1000 {
			result += bestToSpend
		}
	}

	return result
}

func readFile(fileName string) (machines []Machine) {
	lines, _ := utils.ReadFileAsLines(fileName)

	machines = make([]Machine, len(lines)/4)
	for ii, line := range lines {
		switch ii % 4 {
		case 0:
			words := strings.Split(line, " ")
			xx := strings.Split(words[2], "+")[1]
			row, _ := strconv.Atoi(xx[:len(xx)-1])
			column, _ := strconv.Atoi(strings.Split(words[3], "+")[1])
			machines[ii/4].aa = Position{row: row, column: column}
		case 1:
			words := strings.Split(line, " ")
			xx := strings.Split(words[2], "+")[1]
			row, _ := strconv.Atoi(xx[:len(xx)-1])
			column, _ := strconv.Atoi(strings.Split(words[3], "+")[1])
			machines[ii/4].bb = Position{row: row, column: column}
		case 2:
			words := strings.Split(line, " ")
			xx := strings.Split(words[1], "=")[1]
			row, _ := strconv.Atoi(xx[:len(xx)-1])
			column, _ := strconv.Atoi(strings.Split(words[2], "=")[1])
			machines[ii/4].prize = Position{row: row, column: column}
		}

	}

	return machines
}
