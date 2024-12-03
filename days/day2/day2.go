package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkForSafeReport(report []int) bool {
	if len(report) <= 1 {
		return true
	}
	/*
	  1. get the diff to next
	  2. check if same "direction" as previous diff
	  3. check for diff smaller than 4
	*/
	prevDiffs := []int{}
	for idx := 0; idx < len(report)-1; idx++ {
		diffToNext := report[idx] - report[idx+1]
		if diffToNext > 3 || diffToNext < -3 {
			return false
		}
		// check if same diff "direction" as previous
		if len(prevDiffs) > 0 {
			if prevDiffs[idx-1] > 0 && diffToNext > 0 {
				prevDiffs = append(prevDiffs, diffToNext)
			} else if prevDiffs[idx-1] < 0 && diffToNext < 0 {
				prevDiffs = append(prevDiffs, diffToNext)
			} else {
				return false
			}
		} else {
			prevDiffs = append(prevDiffs, diffToNext)
		}
	}

	return true
}

func firstLevelUnsafe(report []int) bool {
	if len(report) < 3 {
		return false
	}
	// first > second < third
	if report[0] > report[1] && report[1] < report[2] {
		return true
	} else if report[0] < report[1] && report[1] > report[2] {
		return true
	}

	return false
}

func checkForSafeReportPart2(report []int, corrected bool) bool {
	if len(report) < 2 {
		return true
	}
	/*
	  1. get the diff to next
	  2. check if same "direction" as previous diff
	  3. check for diff smaller than 4
	*/

	prevDiffs := []int{}
	for idx := 0; idx < len(report)-1; idx++ {
		diffToNext := report[idx] - report[idx+1]
		if diffToNext > 3 || diffToNext < -3 {
			// skip and make level safe
			if corrected == false {
				newReport := append([]int{}, report[:idx]...)
				newReport = append(newReport, report[idx+1:]...)
				fmt.Printf("Trying with new report %v, previous report %v\n", newReport, report)
				return checkForSafeReportPart2(newReport, true)
			}
			return false
		}
		// check if same diff "direction" as previous
		if len(prevDiffs) > 0 {
			if prevDiffs[idx-1] > 0 && diffToNext > 0 {
				prevDiffs = append(prevDiffs, diffToNext)
			} else if prevDiffs[idx-1] < 0 && diffToNext < 0 {
				prevDiffs = append(prevDiffs, diffToNext)
			} else {
				if corrected == false {
					newReport := append([]int{}, report[:idx]...)
					newReport = append(newReport, report[idx+1:]...)
					fmt.Printf("Trying with new report %v, previous report %v\n", newReport, report)
					return checkForSafeReportPart2(newReport, true)
				}
				return false
			}
		} else {
			if firstLevelUnsafe(report) {
				newReport := append([]int{}, report[:idx]...)
				newReport = append(newReport, report[idx+1:]...)
				fmt.Printf("Trying with new report %v, previous report %v\n", newReport, report)
				return checkForSafeReportPart2(newReport, true)
			}
			prevDiffs = append(prevDiffs, diffToNext)
		}
	}

	fmt.Printf("%v is safe\n", report)
	return true
}

func part1() int {
	file, _ := os.Open("input.txt")
	defer file.Close()

	levels := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, " ")

		level := []int{}
		for _, numStr := range nums {
			num, _ := strconv.Atoi(numStr)
			level = append(level, num)
		}

		levels = append(levels, level)
	}

	safeLevelsCount := 0
	for _, level := range levels {
		if checkForSafeReport(level) {
			safeLevelsCount++
		}
	}

	fmt.Println(safeLevelsCount)

	return safeLevelsCount
}

func part2() int {
	file, _ := os.Open("input.txt")
	defer file.Close()

	levels := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, " ")

		level := []int{}
		for _, numStr := range nums {
			num, _ := strconv.Atoi(numStr)
			level = append(level, num)
		}

		levels = append(levels, level)
	}

	safeLevelsCount := 0
	for _, level := range levels {
		if checkForSafeReportPart2(level, false) {
			safeLevelsCount++
		}
	}

	fmt.Println(safeLevelsCount)

	return safeLevelsCount
}
func main() {
	part1()
	part2()
}
