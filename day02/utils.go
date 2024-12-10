package day02

import "github.com/adriananderson/2024-advent-of-code/utils"

func IsSafeUpReport(report []int) bool {
	isSafe := true
	for idx := 0; idx < len(report)-1; idx++ {
		if 1 > (report[idx+1]-report[idx]) || (report[idx+1]-report[idx]) > 3 {
			isSafe = false
			break
		}
	}
	return isSafe
}

func IsSafeDownReport(report []int) bool {
	isSafe := true
	for idx := 0; idx < len(report)-1; idx++ {
		if 1 > (report[idx]-report[idx+1]) || (report[idx]-report[idx+1]) > 3 {
			isSafe = false
			break
		}
	}
	return isSafe
}

func IsSafeReport(report []int) bool {
	return IsSafeUpReport(report) || IsSafeDownReport(report)
}

func IsSafeReportWithoutOne(report []int) bool {
	isSafe := false
	for idx := 0; idx < len(report); idx++ {
		if IsSafeReport(utils.RemoveElementAt(idx, report)) {
			isSafe = true
			break
		}
	}
	return isSafe
}
