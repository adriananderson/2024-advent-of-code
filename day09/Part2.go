package day09

import (
	"github.com/adriananderson/2024-advent-of-code/utils"
	"strings"
)

type File struct {
	id   int
	size int
}

type DiskBlock struct {
	size      int
	files     []File
	available int
	offset    int
}

func Part2() int {
	defer utils.Timer("9-2")()
	var fileName = "day09/day09.txt"

	diskBlocks := readFile2(fileName)

	//move blocks
	for index := len(diskBlocks) - 1; index >= 0; index-- {
		for ii := 0; ii < index; ii++ {
			if len(diskBlocks[index].files) > 0 {
				if diskBlocks[ii].available >= diskBlocks[index].size {
					diskBlocks[ii].available -= diskBlocks[index].size
					diskBlocks[index].available = diskBlocks[index].size //not really needed since we only do a single pass
					diskBlocks[ii].files = append(diskBlocks[ii].files, diskBlocks[index].files[0])
					clear(diskBlocks[index].files)
					break
				}
			}
		}
	}
	checksum := checksum2(diskBlocks)

	return checksum
}

func checksum2(diskBlocks []DiskBlock) (checksum int) {
	checksum = 0
	for _, block := range diskBlocks {
		ii := block.offset
		for _, file := range block.files {
			for jj := 0; jj < file.size; jj++ {
				checksum += file.id * ii
				ii++
			}
		}
	}
	return checksum
}

func readFile2(fileName string) (diskBlocks []DiskBlock) {
	initialDiskDescription, _ := utils.ReadFileAsText(fileName)
	blockDescriptions := strings.Split(initialDiskDescription, "")

	diskBlocks = make([]DiskBlock, len(initialDiskDescription)-1)
	offset := 0
	for ii, blockDescription := range blockDescriptions {
		length := (int)(blockDescription[0] - '0')
		if length >= 0 && length <= 9 {
			diskBlocks[ii].size = length
			diskBlocks[ii].offset = offset
			if ii%2 == 0 { //file
				diskBlocks[ii].files = make([]File, 1)
				diskBlocks[ii].files[0] = File{id: ii / 2, size: length}
			} else { //empty
				diskBlocks[ii].available = length
				diskBlocks[ii].files = make([]File, 0)
			}
		}
		offset += length
	}
	return diskBlocks
}
