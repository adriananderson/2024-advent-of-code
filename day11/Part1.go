package day11

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
	"strconv"
	"strings"
)

func Part1() int {
	defer utils.Timer("11-1")()
	//start := time.Now()
	var fileName = "day11/day11.txt"

	stones := readFile(fileName)

	for ii := 0; ii < 25; ii++ {
		newStones := make([]string, 0, len(stones)*2)
		for _, stone := range stones {
			stoneLength := len(stone)

			if stone == "0" {
				newStones = append(newStones, "1")
			} else if stoneLength%2 == 0 {
				newStones = append(newStones, stone[:stoneLength/2])
				tail, _ := strconv.Atoi(stone[stoneLength/2:])
				tailLabel := strconv.Itoa(tail)
				newStones = append(newStones, tailLabel)
			} else {
				stoneValue, _ := strconv.Atoi(stone)
				newStoneValue := 2024 * stoneValue
				newStone := strconv.Itoa(newStoneValue)
				newStones = append(newStones, newStone)
			}
		}
		stones = newStones
		//sort.Strings(stones)
		//fmt.Print(stones)
		//fmt.Printf("progress %d len: %d     %v\n", ii, len(stones), time.Since(start))

	}

	return len(stones)
}

func readFile(fileName string) (stones []string) {
	line, _ := utils.ReadFileAsText(fileName)

	stones = strings.Split(line, " ")
	//stones = make([]string, len(stoneLabels))
	//
	//for ii, stoneLabel := range stoneLabels {
	//	stones[ii], _ = strconv.Atoi(stoneLabel)
	//}
	return stones
}
