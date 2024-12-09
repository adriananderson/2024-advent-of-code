package day05

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
	"slices"
	"strconv"
	"strings"
)

func Part2() int {
	defer utils.Timer("5-2")()
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
			outerLoop:
				for ii, page := range pages {
					for _, nextPage := range pages[ii:] {
						if !followsRules(page, nextPage, rules) {
							valid = false
							break outerLoop
						}
					}
				}
				if !valid {
					slices.SortFunc(pages, func(a, b string) (ret int) {
						ret = 1
						if followsRules(a, b, rules) {
							ret = -1
						}
						return ret
					})

					mid, _ := strconv.Atoi(pages[len(pages)/2])
					result += mid
				}
			}
		}
	}

	return result
}
