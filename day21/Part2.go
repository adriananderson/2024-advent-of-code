package day21

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
	"strconv"
	"strings"
)

type Coordinates struct {
	X, Y int
}

func Part2(fileName string, numRobots int) int {
	defer utils.Timer("21-2")()

	numericKeys := readFile(fileName)

	numericalMap := make(map[string]Coordinates)
	numericalMap["A"] = Coordinates{2, 0}
	numericalMap["0"] = Coordinates{1, 0}
	numericalMap["1"] = Coordinates{0, 1}
	numericalMap["2"] = Coordinates{1, 1}
	numericalMap["3"] = Coordinates{2, 1}
	numericalMap["4"] = Coordinates{0, 2}
	numericalMap["5"] = Coordinates{1, 2}
	numericalMap["6"] = Coordinates{2, 2}
	numericalMap["7"] = Coordinates{0, 3}
	numericalMap["8"] = Coordinates{1, 3}
	numericalMap["9"] = Coordinates{2, 3}

	directionalMap := make(map[string]Coordinates)
	directionalMap["A"] = Coordinates{2, 1}
	directionalMap["^"] = Coordinates{1, 1}
	directionalMap["<"] = Coordinates{0, 0}
	directionalMap["v"] = Coordinates{1, 0}
	directionalMap[">"] = Coordinates{2, 0}

	return getSequence(numericKeys, numericalMap, directionalMap, numRobots)
}

/*
// +---+---+---+
// | 7 | 8 | 9 |
// +---+---+---+
// | 4 | 5 | 6 |
// +---+---+---+
// | 1 | 2 | 3 |
// +---+---+---+
//	   | 0 | A |
//	   +---+---+
*/
func getPressesForNumericPad(input []string, numericalMap map[string]Coordinates) []string {
	current := numericalMap["A"]
	output := []string{}

	for _, char := range input {
		dest := numericalMap[char]
		diffX, diffY := dest.X-current.X, dest.Y-current.Y

		horizontal, vertical := []string{}, []string{}

		for i := 0; i < utils.Abs(diffX); i++ {
			if diffX >= 0 {
				horizontal = append(horizontal, ">")
			} else {
				horizontal = append(horizontal, "<")
			}
		}

		for i := 0; i < utils.Abs(diffY); i++ {
			if diffY >= 0 {
				vertical = append(vertical, "^")
			} else {
				vertical = append(vertical, "v")
			}
		}

		if current.Y == 0 && dest.X == 0 { //if on lower level and going to left-most
			output = append(output, vertical...)
			output = append(output, horizontal...)

		} else if current.X == 0 && dest.Y == 0 { //if left-most going to lowest level
			output = append(output, horizontal...)
			output = append(output, vertical...)
		} else if diffX < 0 {
			output = append(output, horizontal...)
			output = append(output, vertical...)
		} else {
			output = append(output, vertical...)
			output = append(output, horizontal...)
		}

		current = dest
		output = append(output, "A")
	}
	return output
}

/*
//     +---+---+
//     | ^ | A |
// +---+---+---+
// | < | v | > |
// +---+---+---+
*/
func getPressesForDirectionalPad(input []string, start string, directionlMap map[string]Coordinates) []string {
	current := directionlMap[start]
	output := []string{}

	for _, char := range input {
		dest := directionlMap[char]
		diffX, diffY := dest.X-current.X, dest.Y-current.Y

		horizontal, vertical := []string{}, []string{}

		for i := 0; i < utils.Abs(diffX); i++ {
			if diffX >= 0 {
				horizontal = append(horizontal, ">")
			} else {
				horizontal = append(horizontal, "<")
			}
		}

		for i := 0; i < utils.Abs(diffY); i++ {
			if diffY >= 0 {
				vertical = append(vertical, "^")
			} else {
				vertical = append(vertical, "v")
			}
		}

		if current.X == 0 && dest.Y == 1 { //if on left arrow going to upper row
			output = append(output, horizontal...)
			output = append(output, vertical...)
		} else if current.Y == 1 && dest.X == 0 { //if upper row going to left arrow
			output = append(output, vertical...)
			output = append(output, horizontal...)
		} else if diffX < 0 {
			output = append(output, horizontal...)
			output = append(output, vertical...)
		} else {
			output = append(output, vertical...)
			output = append(output, horizontal...)
		}
		current = dest
		output = append(output, "A")
	}
	return output
}

func getSequence(input []string, numericalMap, directionalMap map[string]Coordinates, maxNumRobots int) int {
	count := 0
	cache := make(map[string][]int)
	for _, line := range input {
		row := strings.Split(line, "")
		seq1 := getPressesForNumericPad(row, numericalMap)
		num := getCountAfterRobots(seq1, 1, maxNumRobots, cache, directionalMap)
		numericValue, _ := strconv.Atoi(line[:len(line)-1])
		count += numericValue * num
	}
	return count
}

func getCountAfterRobots(input []string, robot int, maxNumRobots int, cache map[string][]int, directionalMap map[string]Coordinates) int {
	if val, ok := cache[strings.Join(input, "")]; ok {
		if val[robot-1] != 0 {
			return val[robot-1]
		}
	} else {
		cache[strings.Join(input, "")] = make([]int, maxNumRobots)
	}

	seq := getPressesForDirectionalPad(input, "A", directionalMap)
	cache[strings.Join(input, "")][0] = len(seq)

	if robot == maxNumRobots {
		return len(seq)
	}

	splitSeq := getIndividualSteps(seq)

	count := 0
	for _, s := range splitSeq {
		c := getCountAfterRobots(s, maxNumRobots, robot+1, cache, directionalMap)
		if _, ok := cache[strings.Join(s, "")]; !ok {
			cache[strings.Join(s, "")] = make([]int, maxNumRobots)
		}
		cache[strings.Join(s, "")][0] = c
		count += c
	}

	cache[strings.Join(input, "")][robot-1] = count
	return count
}

func getIndividualSteps(input []string) [][]string {
	output := [][]string{}
	current := []string{}
	for _, char := range input {
		current = append(current, char)

		if char == "A" {
			output = append(output, current)
			current = []string{}
		}
	}
	return output
}
