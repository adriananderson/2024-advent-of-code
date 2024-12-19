package day19

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
)

func Part2(fileName string) int {
	defer utils.Timer("19-1")()

	availablePatterns, desiredPatterns, maxAvailableLength := readFile(fileName)
	result := 0
	countMatchCache := make(map[string]int)
	for _, desired := range desiredPatterns {
		//fmt.Printf("::: %s\n", desired)
		result += countMatch(desired, availablePatterns, maxAvailableLength, countMatchCache)
	}

	return result
}

func countMatch(desired string, availablePatterns map[string]bool, maxAvailableLength int, countMatchCache map[string]int) (matchCount int) {

	if result, exists := countMatchCache[desired]; exists {
		return result
	}

	//find matches
	lengths := make([]int, 0)
	maxSearch := min(maxAvailableLength, len(desired))
	for ii := 1; ii <= maxSearch; ii++ {
		substring := desired[0:ii]
		if availablePatterns[substring] {
			lengths = append(lengths, ii)
		}
	}

	matchCount = 0
	for _, length := range lengths {
		if length == len(desired) {
			matchCount++
		} else {
			//fmt.Printf("        :: %s\n", desired[length:])
			matchCount += countMatch(desired[length:], availablePatterns, maxAvailableLength, countMatchCache)
		}
	}
	countMatchCache[desired] = matchCount
	return matchCount
}
