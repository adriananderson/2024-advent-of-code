package day05

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
	"strconv"
	"strings"
)

type Rule struct {
	first  string
	second string
}

func followsRules(first string, second string, rules []Rule) (followsRules bool) {
	for _, rule := range rules {
		if rule.first == second && rule.second == first {
			return false
		}
	}
	return true
}

func Part1() int {
	defer utils.Timer("5-1")()
	var fileName = "day05/day05.txt"
	var result = 0
	var rules []Rule

	if fileContent, err := utils.ReadFileAsLines(fileName); err == nil {
		for _, line := range fileContent {
			//rules
			if strings.Contains(line, "|") {
				pages := strings.Split(line, "|")
				dep := Rule{pages[0], pages[1]}
				rules = append(rules, dep)
			}
			//pages
			if strings.Contains(line, ",") {
				pages := strings.Split(line, ",")
				valid := true
				for ii, page := range pages {
					for _, nextPage := range pages[ii:] {
						if !followsRules(page, nextPage, rules) {
							valid = false
						}
					}
				}
				if valid {
					mid, _ := strconv.Atoi(pages[len(pages)/2])
					result += mid
				}
			}
		}
	}

	return result
}
