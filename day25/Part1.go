package day25

import (
	"fmt"
	"github.com/adriananderson/2024-advent-of-code/utils"
	"strings"
)

func Part1(fileName string) int {
	defer utils.Timer("25-1")()

	keys, locks := readFile(fileName)

	result := 0
	for _, key := range keys {
		for _, lock := range locks {
			if keyFitsLock(key, lock) {
				fmt.Printf("key %v fits lock %v\n", key, lock)
				result++
			} else {
				fmt.Printf("key %v DOES NOT FIT lock %v\n", key, lock)
			}
		}
	}

	return result
}

func keyFitsLock(key [5]int, lock [5]int) bool {
	for ii := 0; ii < 5; ii++ {
		if key[ii]+lock[ii] > 5 {
			return false
		}
	}
	return true
}

func readFile(fileName string) (keys [][5]int, locks [][5]int) {
	fileText, _ := utils.ReadFileAsText(fileName)

	keys = make([][5]int, 0)
	locks = make([][5]int, 0)

	parts := strings.Split(fileText, "\n\n")
	for _, unit := range parts {
		lines := strings.Split(unit, "\n")
		thing := [5]int{}
		if lines[0] == "#####" { //lock
			for ii := 1; ii < 7; ii++ {
				for jj := 0; jj < 5; jj++ {
					if lines[ii][jj] == '#' {
						thing[jj]++
					}
				}
			}
			locks = append(locks, thing)
		} else { //key
			for ii := 0; ii < 6; ii++ {
				for jj := 0; jj < 5; jj++ {
					if lines[ii][jj] == '#' {
						thing[jj]++
					}
				}
			}
			keys = append(keys, thing)
		}
	}

	return keys, locks
}
