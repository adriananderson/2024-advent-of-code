package day23

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
	"sort"
	"strings"
)

func Part1(fileName string) int {
	defer utils.Timer("23-1")()

	nodes, connections := readFile(fileName)

	parties := make(map[string]bool)
	for node := range nodes {
		for destNode := range connections[node] {
			for otherNode := range connections[destNode] {
				if connections[node][otherNode] {
					connectedNodes := []string{node, destNode, otherNode}
					sort.Strings(connectedNodes)
					//fmt.Printf("%s,%s,%s\n", connectedNodes[0], connectedNodes[1], connectedNodes[2])
					if node[0] == 't' || destNode[0] == 't' || otherNode[0] == 't' {
						key := strings.Join(connectedNodes, ",")
						parties[key] = true
					}
				}
			}
		}
	}

	return len(parties)
}

func readFile(fileName string) (nodeMap map[string]bool, connections map[string]map[string]bool) {
	fileLines, _ := utils.ReadFileAsLines(fileName)

	nodeMap = make(map[string]bool)
	connections = make(map[string]map[string]bool)
	for _, connectionLine := range fileLines {
		if len(connectionLine) > 0 {
			nodes := strings.Split(connectionLine, "-")
			nodeMap[nodes[0]] = true
			nodeMap[nodes[1]] = true

			if _, exists := connections[nodes[0]]; !exists {
				connections[nodes[0]] = make(map[string]bool)
				connections[nodes[0]][nodes[1]] = true
			} else {
				connections[nodes[0]][nodes[1]] = true
			}
			if _, exists := connections[nodes[1]]; !exists {
				connections[nodes[1]] = make(map[string]bool)
				connections[nodes[1]][nodes[0]] = true
			} else {
				connections[nodes[1]][nodes[0]] = true
			}
		}
	}

	return nodeMap, connections
}
