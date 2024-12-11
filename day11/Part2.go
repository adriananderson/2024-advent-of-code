package day11

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
	"strconv"
	"strings"
)

func Part2() int {
	defer utils.Timer("11-2")()
	//start := time.Now()
	var fileName = "day11/day11.txt"

	stones := readFileInt(fileName)

	stoneCheatSheet := make(map[int][]int, 5000)
	stoneCheatSheet[0] = []int{1}

	lengthCheatSheet := make(map[int]int, 5000)
	atoiCheatSheet := make(map[string]int, 5000)

	for ii := 0; ii < 75; ii++ {
		newStones := make(map[int]int, len(stones)*2)
		cacheHit := 0
		evenHit := 0
		otherHit := 0
		for stone, count := range stones {
			if cacheList, ok := stoneCheatSheet[stone]; ok {
				for _, cache := range cacheList {
					//add or increment newStone map
					if newStoneValue, ok := newStones[cache]; ok {
						newStones[cache] = newStoneValue + count
					} else {
						newStones[cache] = count
					}
				}
				cacheHit++
				continue
			}

			var stoneLength int
			var stoneLabel string
			if lengthCheat, ok := lengthCheatSheet[stone]; ok {
				stoneLength = lengthCheat
			} else {
				stoneLabel = strconv.Itoa(stone)
				stoneLength = len(stoneLabel)
				lengthCheatSheet[stone] = stoneLength
			}
			if stoneLength%2 == 0 {
				if stoneLabel == "" {
					stoneLabel = strconv.Itoa(stone)
				}
				headLabel := stoneLabel[:stoneLength/2]
				head := 0
				if headValue, ok := atoiCheatSheet[headLabel]; ok {
					head = headValue
				} else {
					atoiValue, _ := strconv.Atoi(headLabel)
					head = atoiValue
					atoiCheatSheet[headLabel] = head
				}
				//add or increment newStone map
				if newStoneValue, ok := newStones[head]; ok {
					newStones[head] = newStoneValue + count
				} else {
					newStones[head] = count
				}
				tailLabel := stoneLabel[stoneLength/2:]
				tail := 0
				if tailValue, ok := atoiCheatSheet[tailLabel]; ok {
					tail = tailValue
				} else {
					atoiValue, _ := strconv.Atoi(tailLabel)
					tail = atoiValue
					atoiCheatSheet[tailLabel] = tail
				}
				//add or increment newStone map
				if newStoneValue, ok := newStones[tail]; ok {
					newStones[tail] = newStoneValue + count
				} else {
					newStones[tail] = count
				}
				stoneCheatSheet[stone] = []int{head, tail}
				evenHit++
			} else {
				newStone := 2024 * stone
				//add or increment newStone map
				if newStoneValue, ok := newStones[newStone]; ok {
					newStones[newStone] = newStoneValue + count
				} else {
					newStones[newStone] = count
				}
				stoneCheatSheet[stone] = []int{newStone}
				otherHit++
			}
		}

		//fmt.Print(newStones)
		stones = newStones
		//result := 0
		//for _, stoneCount := range stones {
		//	result += stoneCount
		//}
		//fmt.Printf("progress %d len: %d     %v    cache: %d  even: %d  other:  %d    cache size: %d\n", ii, len(stones), time.Since(start), cacheHit, evenHit, otherHit, result)
	}

	result := 0
	for _, stoneCount := range stones {
		result += stoneCount
	}
	return result
}

func readFileInt(fileName string) (stones map[int]int) {
	line, _ := utils.ReadFileAsText(fileName)

	stoneLabels := strings.Split(line, " ")
	stones = make(map[int]int, len(stoneLabels))

	//order doesn't matter
	for _, stoneLabel := range stoneLabels {
		stone, _ := strconv.Atoi(stoneLabel)
		stones[stone] = 1
	}
	return stones
}
