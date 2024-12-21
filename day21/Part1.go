package day21

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
	"strconv"
	"strings"
)

func Part1(fileName string) int {
	defer utils.Timer("21-1")()

	numericKeys := readFile(fileName)
	numericPaths := PreCalcNumericPaths()
	directionPaths := PreCalcDirectionPaths()

	firstRobotDirectionKeys := make([]string, len(numericKeys))
	secondRobotDirectionKeys := make([]string, len(firstRobotDirectionKeys))
	thirdRobotDirectionKeys := make([]string, len(secondRobotDirectionKeys))

	for ii := 0; ii < len(numericKeys); ii++ {
		firstRobotDirectionKeys[ii] = ProcessKeyMap(numericKeys[ii], numericPaths)
		//fmt.Printf(" %s -> %s\n", numericKeys[ii], firstRobotDirectionKeys[ii])
		secondRobotDirectionKeys[ii] = ProcessKeyMap(firstRobotDirectionKeys[ii], directionPaths)
		//fmt.Printf(" %s -> %s\n", firstRobotDirectionKeys[ii], secondRobotDirectionKeys[ii])
		thirdRobotDirectionKeys[ii] = ProcessKeyMap(secondRobotDirectionKeys[ii], directionPaths)
		//fmt.Printf(" %s -> %s\n", secondRobotDirectionKeys[ii], thirdRobotDirectionKeys[ii])
	}

	result := 0
	for ii := 0; ii < len(numericKeys); ii++ {
		numericKeyLine := numericKeys[ii]
		numericValue, _ := strconv.Atoi(string(numericKeyLine[:len(numericKeyLine)-1]))
		length := len(thirdRobotDirectionKeys[ii])
		result += numericValue * length
		//fmt.Printf(" %d   %d\n", numericValue, length)
	}

	return result
}

func ProcessKeyMap(directionKeyLine string, directionPaths map[rune]map[rune]string) (directionString string) {
	firstPosition := 'A'
	var sb strings.Builder
	for _, nextPosition := range directionKeyLine {
		segment := directionPaths[firstPosition][nextPosition]
		sb.WriteString(segment)
		firstPosition = nextPosition
	}
	directionString = sb.String()
	//fmt.Printf("%s  %d\n", sb.String(), len(sb.String()))
	return directionString
}

func PreCalcDirectionPaths() (directionPaths map[rune]map[rune]string) {
	directionPaths = make(map[rune]map[rune]string)
	directionPaths['A'] = make(map[rune]string)
	directionPaths['A']['A'] = "A"
	directionPaths['A']['^'] = "<A"
	directionPaths['A']['<'] = "v<<A"
	directionPaths['A']['v'] = "<vA"
	directionPaths['A']['>'] = "vA"
	directionPaths['^'] = make(map[rune]string)
	directionPaths['^']['A'] = ">A"
	directionPaths['^']['^'] = "A"
	directionPaths['^']['<'] = "v<A"
	directionPaths['^']['v'] = "vA"
	directionPaths['^']['>'] = "v>A"
	directionPaths['<'] = make(map[rune]string)
	directionPaths['<']['A'] = ">>^A"
	directionPaths['<']['^'] = ">^A"
	directionPaths['<']['<'] = "A"
	directionPaths['<']['v'] = ">A"
	directionPaths['<']['>'] = ">>A"
	directionPaths['v'] = make(map[rune]string)
	directionPaths['v']['A'] = ">^A"
	directionPaths['v']['^'] = "^A"
	directionPaths['v']['<'] = "<A"
	directionPaths['v']['v'] = "A"
	directionPaths['v']['>'] = ">A"
	directionPaths['>'] = make(map[rune]string)
	directionPaths['>']['A'] = "^A"
	directionPaths['>']['^'] = "<^A"
	directionPaths['>']['<'] = "<<A"
	directionPaths['>']['v'] = "<A"
	directionPaths['>']['>'] = "A"

	return directionPaths
}

func PreCalcNumericPaths() (numericPaths map[rune]map[rune]string) {
	numericPaths = make(map[rune]map[rune]string)
	numericPaths['A'] = make(map[rune]string)
	numericPaths['A']['A'] = "A"
	numericPaths['A']['0'] = "<A"
	numericPaths['A']['1'] = "^<<A"
	numericPaths['A']['2'] = "<^A"
	numericPaths['A']['3'] = "^A"
	numericPaths['A']['4'] = "^^<<A"
	numericPaths['A']['5'] = "<^^A"
	numericPaths['A']['6'] = "^^A"
	numericPaths['A']['7'] = "^^^<<A"
	numericPaths['A']['8'] = "<^^^A"
	numericPaths['A']['9'] = "^^^A"
	numericPaths['0'] = make(map[rune]string)
	numericPaths['0']['A'] = ">A"
	numericPaths['0']['0'] = "A"
	numericPaths['0']['1'] = "^<A"
	numericPaths['0']['2'] = "^A"
	numericPaths['0']['3'] = "^>A"
	numericPaths['0']['4'] = "^^<A"
	numericPaths['0']['5'] = "^^A"
	numericPaths['0']['6'] = "^^>A"
	numericPaths['0']['7'] = "^^^<A"
	numericPaths['0']['8'] = "^^^A"
	numericPaths['0']['9'] = "^^^>A"

	numericPaths['1'] = make(map[rune]string)
	numericPaths['1']['A'] = ">>vA"
	numericPaths['1']['0'] = ">vA"
	numericPaths['1']['1'] = "A"
	numericPaths['1']['2'] = ">A"
	numericPaths['1']['3'] = ">>A"
	numericPaths['1']['4'] = "^A"
	numericPaths['1']['5'] = "^>A"
	numericPaths['1']['6'] = "^>>A"
	numericPaths['1']['7'] = "^^A"
	numericPaths['1']['8'] = "^^>A"
	numericPaths['1']['9'] = "^^>>A"
	numericPaths['2'] = make(map[rune]string)
	numericPaths['2']['A'] = "v>A"
	numericPaths['2']['0'] = "vA"
	numericPaths['2']['1'] = "<A"
	numericPaths['2']['2'] = "A"
	numericPaths['2']['3'] = ">A"
	numericPaths['2']['4'] = "<^A"
	numericPaths['2']['5'] = "^A"
	numericPaths['2']['6'] = "^>A"
	numericPaths['2']['7'] = "<^^A"
	numericPaths['2']['8'] = "^^A"
	numericPaths['2']['9'] = ">^^A"
	numericPaths['3'] = make(map[rune]string)
	numericPaths['3']['A'] = "vA"
	numericPaths['3']['0'] = "<vA"
	numericPaths['3']['1'] = "<<A"
	numericPaths['3']['2'] = "<A"
	numericPaths['3']['3'] = "A"
	numericPaths['3']['4'] = "<<^A"
	numericPaths['3']['5'] = "<^A"
	numericPaths['3']['6'] = "^A"
	numericPaths['3']['7'] = "<<^^A"
	numericPaths['3']['8'] = "<^^A"
	numericPaths['3']['9'] = "^^A"

	numericPaths['4'] = make(map[rune]string)
	numericPaths['4']['A'] = ">>vvA"
	numericPaths['4']['0'] = ">vvA"
	numericPaths['4']['1'] = "vA"
	numericPaths['4']['2'] = ">vA"
	numericPaths['4']['3'] = ">>vA"
	numericPaths['4']['4'] = "A"
	numericPaths['4']['5'] = ">A"
	numericPaths['4']['6'] = ">>A"
	numericPaths['4']['7'] = "^A"
	numericPaths['4']['8'] = "^>A"
	numericPaths['4']['9'] = "^>>A"
	numericPaths['5'] = make(map[rune]string)
	numericPaths['5']['A'] = "vv>A"
	numericPaths['5']['0'] = "vvA"
	numericPaths['5']['1'] = "<vA"
	numericPaths['5']['2'] = "vA"
	numericPaths['5']['3'] = ">vA"
	numericPaths['5']['4'] = "<A"
	numericPaths['5']['5'] = "A"
	numericPaths['5']['6'] = ">A"
	numericPaths['5']['7'] = "<^A"
	numericPaths['5']['8'] = "^A"
	numericPaths['5']['9'] = "^>A"
	numericPaths['6'] = make(map[rune]string)
	numericPaths['6']['A'] = "vvA"
	numericPaths['6']['0'] = "<vvA"
	numericPaths['6']['1'] = "<<vA"
	numericPaths['6']['2'] = "<vA"
	numericPaths['6']['3'] = "vA"
	numericPaths['6']['4'] = "<<A"
	numericPaths['6']['5'] = "<A"
	numericPaths['6']['6'] = "A"
	numericPaths['6']['7'] = "<<^A"
	numericPaths['6']['8'] = "<^A"
	numericPaths['6']['9'] = "^A"

	numericPaths['7'] = make(map[rune]string)
	numericPaths['7']['A'] = ">>vvvA"
	numericPaths['7']['0'] = ">vvvA"
	numericPaths['7']['1'] = "vvA"
	numericPaths['7']['2'] = "vv>A"
	numericPaths['7']['3'] = "vv>>A"
	numericPaths['7']['4'] = "vA"
	numericPaths['7']['5'] = "v>A"
	numericPaths['7']['6'] = "v>>A"
	numericPaths['7']['7'] = "A"
	numericPaths['7']['8'] = ">A"
	numericPaths['7']['9'] = ">>A"
	numericPaths['8'] = make(map[rune]string)
	numericPaths['8']['A'] = "vvv>A"
	numericPaths['8']['0'] = "vvvA"
	numericPaths['8']['1'] = "vv<A"
	numericPaths['8']['2'] = "vvA"
	numericPaths['8']['3'] = "vv>A"
	numericPaths['8']['4'] = "<vA"
	numericPaths['8']['5'] = "vA"
	numericPaths['8']['6'] = "v>A"
	numericPaths['8']['7'] = "<A"
	numericPaths['8']['8'] = "A"
	numericPaths['8']['9'] = ">A"
	numericPaths['9'] = make(map[rune]string)
	numericPaths['9']['A'] = "vvvA"
	numericPaths['9']['0'] = "<vvvA"
	numericPaths['9']['1'] = "<<vvA"
	numericPaths['9']['2'] = "<vvA"
	numericPaths['9']['3'] = "vvA"
	numericPaths['9']['4'] = "<<vA"
	numericPaths['9']['5'] = "<vA"
	numericPaths['9']['6'] = "vA"
	numericPaths['9']['7'] = "<<A"
	numericPaths['9']['8'] = "<A"
	numericPaths['9']['9'] = "A"

	return numericPaths
}

func readFile(fileName string) (numericKeys []string) {
	keyLines, _ := utils.ReadFileAsLines(fileName)

	numericKeys = make([]string, len(keyLines)-1)
	for ii, keyLine := range keyLines {
		if len(keyLine) > 0 {
			numericKeys[ii] = keyLine
		}
	}

	return numericKeys
}
