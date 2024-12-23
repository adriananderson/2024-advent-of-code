package day23

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
	"sort"
	"strings"
)

func Part2(fileName string) string {
	defer utils.Timer("23-2")()

	//nodes, connections := readFile(fileName)
	_, connections := readFile(fileName)

	password := findMaxCliques(connections)
	//var longestParty []string
	//for node := range nodes {
	//	members := make(map[string]bool)
	//	members[node] = true
	//	nonMembers := make(map[string]bool)
	//	newLongestParty := findLongestParty(node, members, nonMembers, connections)
	//	if len(newLongestParty) > len(longestParty) {
	//		sort.Strings(newLongestParty)
	//		longestParty = newLongestParty
	//	}
	//}
	//
	//return strings.Join(longestParty, ",")
	return password
}

func findMaxCliques(lanMap map[string]map[string]bool) string {
	maxClique := []string{}
	allComputers := []string{}
	for key, _ := range lanMap {
		allComputers = append(allComputers, key)
	}
	cliques := BronKerbosch([]string{}, allComputers, []string{}, lanMap, [][]string{})
	for _, c := range cliques {
		if len(c) > len(maxClique) {
			maxClique = c
		}
	}
	sort.Strings(maxClique)
	return strings.Join(maxClique, ",")
}

func BronKerbosch(currentClique []string, yetToConsider []string, alreadyConsidered []string, lanMap map[string]map[string]bool, cliques [][]string) [][]string {
	if len(yetToConsider) == 0 && len(alreadyConsidered) == 0 {
		cliques = append(cliques, append([]string{}, currentClique...))
		return cliques
	}

	for index := 0; index < len(yetToConsider); {
		node := yetToConsider[index]
		newYetToConsider := []string{}
		newAlreadyConsidered := []string{}

		for _, n := range yetToConsider {
			if _, ok := lanMap[node][n]; ok {
				newYetToConsider = append(newYetToConsider, n)
			}
		}

		for _, n := range alreadyConsidered {
			if _, ok := lanMap[node][n]; ok {
				newAlreadyConsidered = append(newAlreadyConsidered, n)
			}
		}

		cliques = BronKerbosch(append(currentClique, node), newYetToConsider, newAlreadyConsidered, lanMap, cliques)

		yetToConsider = append(yetToConsider[:index], yetToConsider[index+1:]...)
		alreadyConsidered = append(alreadyConsidered, node)
	}
	return cliques
}

func findLongestParty(head string, members map[string]bool, nonMembers map[string]bool, connections map[string]map[string]bool) []string {
	longestLength := len(members)
	longestParty := make([]string, 0)
	for member := range members {
		longestParty = append(longestParty, member)
	}
	for possibleNode := range connections[head] {
		if _, exists := members[possibleNode]; exists {
			continue
		}
		if _, exists := nonMembers[possibleNode]; exists {
			continue
		}
		hasConnectionsToExisting := true
		for existingMember := range members {
			if !connections[possibleNode][existingMember] {
				hasConnectionsToExisting = false
			}
		}
		if !hasConnectionsToExisting {
			continue
		}
		newMembers := make(map[string]bool)
		for key, value := range members {
			newMembers[key] = value
		}
		newMembers[possibleNode] = true
		newNonMembers := make(map[string]bool)
		for key, value := range nonMembers {
			newNonMembers[key] = value
		}
		newLongestParty := findLongestParty(head, newMembers, newNonMembers, connections)
		newLength := len(newLongestParty)
		if newLength > longestLength {
			longestLength = newLength
			longestParty = newLongestParty
			//fmt.Printf("nlp: %v\n", longestParty)
		}
		nonMembers[possibleNode] = true
	}

	return longestParty
}
