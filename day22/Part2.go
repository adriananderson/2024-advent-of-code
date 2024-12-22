package day22

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
)

type BananaPriceChange struct {
	price  int
	change int
}

type Sequence struct {
	a, b, c, d int
}

func Part2(fileName string, numIters int) int {
	defer utils.Timer("22-2")()

	secretNumbers := readFile(fileName)
	bananaPriceChanges := calcBananaPrices(secretNumbers, numIters)
	maxPrice := findMaxPrice(bananaPriceChanges)

	return maxPrice
}

func calcBananaPrices(secretNumbers []int, numIters int) (bananaPriceChange [][]BananaPriceChange) {
	bananaPriceChange = make([][]BananaPriceChange, len(secretNumbers))
	for i, line := range secretNumbers {
		bananaPriceChange[i] = make([]BananaPriceChange, numIters)
		num := line
		prev := num % 10
		for j := range numIters {
			num = nextNumber(num)
			bananaPriceChange[i][j] = getPriceChange(num, prev)
			prev = num % 10
		}
	}

	return bananaPriceChange
}

func getPriceChange(input int, previous int) BananaPriceChange {
	return BananaPriceChange{price: input % 10, change: input%10 - previous}
}

func findMaxPrice(bananaPriceChange [][]BananaPriceChange) int {
	seqMap := make(map[Sequence][]int)
	for i, _ := range bananaPriceChange {
		s := []int{bananaPriceChange[i][0].change, bananaPriceChange[i][1].change, bananaPriceChange[i][2].change}
		for j := 3; j < len(bananaPriceChange[i]); j++ {
			s = append(s, bananaPriceChange[i][j].change)

			if _, ok := seqMap[Sequence{s[0], s[1], s[2], s[3]}]; !ok {
				seqMap[Sequence{s[0], s[1], s[2], s[3]}] = make([]int, len(bananaPriceChange))
			}

			if seqMap[Sequence{s[0], s[1], s[2], s[3]}][i] == 0 {
				seqMap[Sequence{s[0], s[1], s[2], s[3]}][i] = bananaPriceChange[i][j].price
			}

			s = s[1:]
		}
	}
	maxPrice := 0
	for _, n := range seqMap {
		sum := 0
		for _, r := range n {
			sum += r
		}
		if sum > maxPrice {
			maxPrice = sum
		}
	}
	return maxPrice
}
