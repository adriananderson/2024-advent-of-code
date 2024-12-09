package day09

import (
	"2024/utils"
	"strings"
)

func Part1() int {
	defer utils.Timer("9-1")()
	var fileName = "day09/day09.txt"

	totalSize, diskBlocks := readFile(fileName)

	//move blocks
	for nextOpening, lastFilled := findNextOpening(0, totalSize, diskBlocks), findLastFilled(totalSize-1, 0, diskBlocks); nextOpening != -1 && lastFilled != -1; nextOpening, lastFilled = findNextOpening(nextOpening, lastFilled, diskBlocks), findLastFilled(lastFilled, nextOpening, diskBlocks) {
		diskBlocks[nextOpening] = diskBlocks[lastFilled]
		diskBlocks[lastFilled] = -1
	}

	checkSum := checksum(diskBlocks)

	return checkSum
}

func checksum(diskBlocks []int) int {
	checkSum := 0
	for ii, block := range diskBlocks {
		if block != -1 {
			checkSum += ii * block
		}
	}
	return checkSum
}

func readFile(fileName string) (int, []int) {
	initialDiskDescription, _ := utils.ReadFileAsText(fileName)
	blockDescriptions := strings.Split(initialDiskDescription, "")
	totalSize := 0
	for _, blockDescription := range blockDescriptions {
		length := (int)(blockDescription[0] - '0')
		if length >= 0 && length <= 9 {
			totalSize += length
		}
	}
	diskBlocks := make([]int, totalSize)
	index := 0
	for ii, blockDescription := range blockDescriptions {
		length := (int)(blockDescription[0] - '0')
		if length >= 0 && length <= 9 {
			if ii%2 == 0 { //file
				for jj := 0; jj < length; jj++ {
					diskBlocks[index] = (ii / 2)
					index++
				}
			} else { //empty
				for jj := 0; jj < length; jj++ {
					diskBlocks[index] = -1
					index++
				}
			}
		}
	}
	return totalSize, diskBlocks
}

func findNextOpening(nextOpening int, lastFilled int, diskBlocks []int) int {
	for ii := nextOpening; ii < lastFilled; ii++ {
		if diskBlocks[ii] == -1 {
			return ii
		}
	}
	return -1
}

func findLastFilled(lastFilled int, nextOpening int, diskBlocks []int) int {
	for ii := lastFilled; ii > nextOpening; ii-- {
		if diskBlocks[ii] != -1 {
			return ii
		}
	}
	return -1
}
