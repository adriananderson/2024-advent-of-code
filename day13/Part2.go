package day13

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
)

func Part2() int {
	defer utils.Timer("13-2")()
	var fileName = "day13/day13.txt"

	machines := readFile(fileName)
	result := 0

	for ii, machine := range machines {
		machines[ii].prize.row = 10000000000000 + machine.prize.row
		machines[ii].prize.column = 10000000000000 + machine.prize.column
	}

	for _, machine := range machines {
		//linear algebra, solve system of equations
		//see: https://en.wikipedia.org/wiki/Cramer%27s_rule
		// [ aa.row  bb.row | prize.row ]
		// [ aa.col  bb.col | prize.col ]
		partA1 := machine.bb.column*machine.prize.row - machine.bb.row*machine.prize.column
		partA2 := machine.aa.row*machine.bb.column - machine.aa.column*machine.bb.row
		mulA := partA1 / partA2

		partB1 := machine.aa.column*machine.prize.row - machine.aa.row*machine.prize.column
		partB2 := machine.aa.column*machine.bb.row - machine.aa.row*machine.bb.column
		mulB := partB1 / partB2

		if (machine.aa.row*mulA+machine.bb.row*mulB) == machine.prize.row &&
			(machine.aa.column*mulA+machine.bb.column*mulB) == machine.prize.column {
			//fmt.Printf("%d solution %v %v\n", zz, mulA, mulB)
			result += mulA*3 + mulB
		}

	}

	return result
}
