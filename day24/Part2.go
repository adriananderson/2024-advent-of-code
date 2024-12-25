package day24

import (
	"fmt"
	"github.com/adriananderson/2024-advent-of-code/utils"
	"sort"
	"strconv"
	"strings"
)

type FriendlyGateLabel struct {
	label string
	num   int
}

func Part2(fileName string, maxBit int) string {
	defer utils.Timer("24-2")()

	_, wireGateMap := readFile(fileName)

	//for ii := 2; ii <= maxBit; ii++ {
	//	zLabel := fmt.Sprintf("z%02d", ii)
	//	printGate(zLabel, wireGateMap, 0)
	//}

	labeledGates := make(map[FriendlyGateLabel]Gate)
	gateDescriptions := make(map[string]FriendlyGateLabel)
	for _, gate := range wireGateMap {
		gateLabel := getGateLabel(gate, wireGateMap)
		gateDescriptions[gate.output] = gateLabel
		labeledGates[gateLabel] = gate
	}

	badGates := checkLabeledGates(labeledGates, gateDescriptions, wireGateMap, maxBit)
	badGateList := make([]string, 0)
	for gateLabel := range badGates {
		badGateList = append(badGateList, gateLabel)
	}
	sort.Strings(badGateList)
	return strings.Join(badGateList, ",")
}

func checkLabeledGates(labeledGates map[FriendlyGateLabel]Gate, gateDescriptions map[string]FriendlyGateLabel, wireGateMap map[string]Gate, maxBit int) (naughtyGates map[string]bool) {
	naughtyGates = make(map[string]bool)
	//check z-lines
	for ii := 3; ii < maxBit; ii++ {
		zLabel := fmt.Sprintf("z%02d", ii)
		actualDescription := gateDescriptions[zLabel]
		expectedDescription := FriendlyGateLabel{label: "XOR+", num: ii}
		actualZGate := labeledGates[expectedDescription]

		if actualDescription != expectedDescription {
			naughtyGates[actualZGate.output] = true
			naughtyGates[zLabel] = true
		}

		//only have to check heads
		//xorrLabel := FriendlyGateLabel{label: "XOR+", num: ii}
		orr1Label := FriendlyGateLabel{label: "OR+", num: ii - 1}
		andd1Label := FriendlyGateLabel{label: "AND+", num: ii - 1}
		xorLabel := FriendlyGateLabel{label: "XOR", num: ii}
		xor1Label := FriendlyGateLabel{label: "XOR", num: ii - 1}
		orr2Label := FriendlyGateLabel{label: "OR+", num: ii - 2}
		and1Label := FriendlyGateLabel{label: "AND", num: ii - 1}

		//XOR+
		leftDesc := gateDescriptions[actualZGate.left]
		rightDesc := gateDescriptions[actualZGate.right]

		if leftDesc != xorLabel && leftDesc != orr1Label {
			naughtyGates[actualZGate.left] = true
			if rightDesc == xorLabel {
				naughtyGates[labeledGates[orr1Label].output] = true
			} else {
				naughtyGates[labeledGates[xorLabel].output] = true
			}
		}
		if rightDesc != xorLabel && rightDesc != orr1Label {
			naughtyGates[actualZGate.right] = true
			if leftDesc == xorLabel {
				naughtyGates[labeledGates[orr1Label].output] = true
			} else {
				naughtyGates[labeledGates[xorLabel].output] = true
			}
		}

		//OR+
		orrGate := labeledGates[orr1Label]
		leftDesc = gateDescriptions[orrGate.left]
		rightDesc = gateDescriptions[orrGate.right]
		if leftDesc != and1Label && leftDesc != andd1Label {
			naughtyGates[orrGate.left] = true
			if rightDesc == and1Label {
				naughtyGates[labeledGates[andd1Label].output] = true
			} else {
				naughtyGates[labeledGates[and1Label].output] = true
			}
		}
		if rightDesc != and1Label && rightDesc != andd1Label {
			naughtyGates[orrGate.right] = true
			if leftDesc == and1Label {
				naughtyGates[labeledGates[andd1Label].output] = true
			} else {
				naughtyGates[labeledGates[and1Label].output] = true
			}
		}

		//AND+
		anddGate := labeledGates[andd1Label]
		leftDesc = gateDescriptions[anddGate.left]
		rightDesc = gateDescriptions[anddGate.right]

		if leftDesc != xor1Label && leftDesc != orr2Label {
			naughtyGates[anddGate.left] = true
			if rightDesc == xor1Label {
				naughtyGates[labeledGates[orr2Label].output] = true
			} else {
				naughtyGates[labeledGates[xor1Label].output] = true
			}
		}
		if rightDesc != xor1Label && rightDesc != orr2Label {
			naughtyGates[actualZGate.right] = true
			if leftDesc == xor1Label {
				naughtyGates[labeledGates[orr2Label].output] = true
			} else {
				naughtyGates[labeledGates[xor1Label].output] = true
			}
		}

	}

	return naughtyGates
}

func getGateNumber(gate Gate, wireGateMap map[string]Gate) (result int) {

	if (gate.left[0] == 'x' || gate.right[0] == 'x') && (gate.left[0] == 'y' || gate.right[0] == 'y') {
		numLabel := gate.left[1:]
		result, _ = strconv.Atoi(numLabel)
	} else {
		leftGate := wireGateMap[gate.left]
		leftNum := getGateNumber(leftGate, wireGateMap)
		rightGate := wireGateMap[gate.right]
		rightNum := getGateNumber(rightGate, wireGateMap)
		result = max(leftNum, rightNum)
	}

	return result
}

func getGateLabel(gate Gate, wireGateMap map[string]Gate) (gateLabel FriendlyGateLabel) {

	if (gate.left[0] == 'x' || gate.right[0] == 'x') && (gate.left[0] == 'y' || gate.right[0] == 'y') {
		if gate.logic == XOR {
			gateLabel.label = "XOR"
		} else if gate.logic == AND {
			gateLabel.label = "AND"
		}
	} else {
		if gate.logic == XOR {
			gateLabel.label = "XOR+"
		} else if gate.logic == AND {
			gateLabel.label = "AND+"
		} else if gate.logic == OR {
			gateLabel.label = "OR+"
		}
	}

	gateLabel.num = getGateNumber(gate, wireGateMap)

	return gateLabel
}

func printGate(key string, wireGateMap map[string]Gate, depth int) {
	if depth > 35 {
		return
	}
	gate := wireGateMap[key]

	logic := ""
	switch gate.logic {
	case 0:
		logic = "AND"
	case 1:
		logic = "OR "
	case 2:
		logic = "XOR"
	}

	fmt.Printf("%*s%s %s %s\n", depth, " ", key, logic, gate.left)
	if _, exists := wireGateMap[gate.left]; exists {
		printGate(gate.left, wireGateMap, depth+9)
	}
	fmt.Printf("%*s      - %s\n", depth, " ", gate.right)
	if _, exists := wireGateMap[gate.right]; exists {
		printGate(gate.right, wireGateMap, depth+9)
	}
}
