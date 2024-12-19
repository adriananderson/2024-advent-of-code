package day19

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
	"strings"
)

func Part1(fileName string) int {
	defer utils.Timer("19-1")()

	availablePatterns, desiredPatterns, maxAvailableLength := readFile(fileName)
	result := 0
	isMatchCache := make(map[string]bool)
	for _, desired := range desiredPatterns {
		//fmt.Printf("::: %s\n", desired)
		if isMatch(desired, availablePatterns, maxAvailableLength, isMatchCache) {
			//fmt.Printf("match %s\n", desired)
			result++
			//} else {
			//	fmt.Printf("no match %s\n", desired)
		}
	}

	return result
}

func isMatch(desired string, availablePatterns map[string]bool, maxAvailableLength int, isMatchCache map[string]bool) (match bool) {

	if result, exists := isMatchCache[desired]; exists {
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

	if len(lengths) == 0 {
		isMatchCache[desired] = false
		return false
	}
	if len(desired) == lengths[len(lengths)-1] {
		isMatchCache[desired] = true
		return true
	}
	match = false
	for _, length := range lengths {
		//fmt.Printf("        :: %s\n", desired[length:])
		match = isMatch(desired[length:], availablePatterns, maxAvailableLength, isMatchCache)
		if match {
			break
		}
	}
	isMatchCache[desired] = match
	return match
}

func readFile(fileName string) (availablePatterns map[string]bool, desiredPatterns []string, maxAvailableLength int) {
	fileLines, _ := utils.ReadFileAsLines(fileName)

	availablePatterns = make(map[string]bool)
	maxAvailableLength = 0
	for ii := 0; ii < len(fileLines); ii++ {
		fileLine := fileLines[ii]
		if len(fileLine) < 2 {
			continue
		}
		if ii == 0 {
			availablePatternSplice := strings.Split(fileLine, ", ")
			for _, availablePattern := range availablePatternSplice {
				availablePatterns[availablePattern] = true
				if len(availablePattern) > maxAvailableLength {
					maxAvailableLength = len(availablePattern)
				}
			}
		} else {
			desiredPatterns = append(desiredPatterns, fileLine)
		}
	}

	return availablePatterns, desiredPatterns, maxAvailableLength
}
