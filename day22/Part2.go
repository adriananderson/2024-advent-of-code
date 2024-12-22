package day22

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
	"strings"
)

type Delta struct {
	price  int
	change int
}

type Sequence struct {
	a, b, c, d int
}

func Part2(fileName string, numIters int) int {
	defer utils.Timer("22-2")()

	secretNumbers := readFile(fileName)

	sellersBestPrices, bestPricePerSeller := generateChanges(secretNumbers, numIters)

	highestNumBananas := 0

	//assume best sequence is price 9 for someone ...
	for ii := 0; ii < len(secretNumbers); ii++ {
		listOfNineLabels := sellersBestPrices[ii][9]
		for jj := 0; jj < len(listOfNineLabels); jj++ {
			numBananas := SellAtSeqence(listOfNineLabels[jj], bestPricePerSeller)
			if numBananas > highestNumBananas {
				highestNumBananas = numBananas
			}
		}
	}

	return highestNumBananas
}

func SellAtSeqence(label string, bestPricePerSeller [] /*sellerindex*/ map[string]int) int {
	result := 0
	for ii := 0; ii < len(bestPricePerSeller); ii++ {
		result += bestPricePerSeller[ii][label]
	}
	return result
}

func generateChanges(secretNumbers []int, numIters int) (sellersBestPrices [] /*sellerindex*/ [] /*price*/ [] /*list*/ string, bestPricePerSeller [] /*sellerindex*/ map[string]int) {

	sellersBestPrices = make([][][]string, len(secretNumbers))
	bestPricePerSeller = make([]map[string]int, len(secretNumbers))
	deltas := make([][]Delta, len(secretNumbers))

	for jj, secretNumber := range secretNumbers {
		sellersBestPrices[jj] = make([][]string, 10)
		bestPricePerSeller[jj] = make(map[string]int)
		deltas[jj] = make([]Delta, numIters)

		thisSecretNumber := secretNumber
		for ii := 0; ii < numIters; ii++ {
			newSecretNumber := nextNumber(thisSecretNumber)
			//fmt.Printf("%d     %d -> %d\n", ii, thisSecretNumber, newSecretNumber)
			price := newSecretNumber % 10
			prevPrice := thisSecretNumber % 10

			deltas[jj][ii].price = price
			deltas[jj][ii].change = price - prevPrice
			//fmt.Printf(" %d   %d\n", deltas[jj][ii].price, deltas[jj][ii].change)

			if ii > 3 {
				label := CreateLabelForPriceChanges(deltas[jj][ii-3].change, deltas[jj][ii-2].change, deltas[jj][ii-1].change, deltas[jj][ii].change)
				if _, exists := bestPricePerSeller[jj][label]; !exists {
					bestPricePerSeller[jj][label] = price
					sellersBestPrices[jj][price] = append(sellersBestPrices[jj][price], label)
				}
			}

			thisSecretNumber = newSecretNumber
		}
	}

	return sellersBestPrices, bestPricePerSeller
}

/*
*
create label for 4 day price changes
*/
func CreateLabelForPriceChanges(aa, bb, cc, dd int) string {
	var sb strings.Builder
	sb.WriteRune(rune('J' + aa))
	sb.WriteRune(rune('J' + bb))
	sb.WriteRune(rune('J' + cc))
	sb.WriteRune(rune('J' + dd))
	return sb.String()
}

func PriceChangesFromLabel(label string) (priceChanges [4]int) {
	priceChanges = [4]int{}
	for i := 0; i < len(label); i++ {
		priceChanges[i] = int(label[i] - 'J')
		if priceChanges[i] > 10 {
			priceChanges[i] -= 256
		}
	}
	return priceChanges
}
