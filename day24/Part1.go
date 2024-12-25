package day24

import (
	"fmt"
	"github.com/adriananderson/2024-advent-of-code/utils"
	"strings"
)

const (
	AND = iota
	OR
	XOR
)

type Gate struct {
	left, right string
	logic       int
	output      string
}

func Part1(fileName string) int {
	defer utils.Timer("24-1")()

	valueMap, wireGateMap := readFile(fileName)

	result := 0
	for ii := 99; ii >= 0; ii-- {
		key := fmt.Sprintf("z%02d", ii)
		if _, exists := wireGateMap[key]; exists {
			value := evaluateGate(valueMap, wireGateMap, key)
			result = result << 1
			if value {
				result = result | 1
			}
		}
	}

	return result
}

func evaluateGate(valueMap map[string]bool, wireGateMap map[string]Gate, gateLabel string) bool {
	if len(gateLabel) == 0 {
		fmt.Printf("gateLabel is empty\n")
		return false
	}

	if value, exists := valueMap[gateLabel]; exists {
		return value
	}
	gate := wireGateMap[gateLabel]
	var result bool
	left := evaluateGate(valueMap, wireGateMap, gate.left)
	right := evaluateGate(valueMap, wireGateMap, gate.right)
	switch gate.logic {
	case AND:
		result = left && right
	case OR:
		result = left || right
	case XOR:
		result = left != right
	}
	valueMap[gateLabel] = result
	return result
}

func readFile(fileName string) (valueMap map[string]bool, wireGateMap map[string]Gate) {
	fileLines, _ := utils.ReadFileAsLines(fileName)

	valueMap = make(map[string]bool)
	wireGateMap = make(map[string]Gate)

	for _, line := range fileLines {
		if len(line) > 0 {
			if strings.Contains(line, ": ") {
				parts := strings.Split(line, ": ")
				valueMap[parts[0]] = (parts[1] == "1")
			} else if strings.Contains(line, " -> ") {
				parts := strings.Split(line, " ")
				var logic int
				switch parts[1] {
				case "AND":
					logic = AND
				case "OR":
					logic = OR
				case "XOR":
					logic = XOR
				}
				gate := Gate{parts[0], parts[2], logic, parts[4]}
				wireGateMap[parts[4]] = gate
			}
		}
	}

	return valueMap, wireGateMap
}
