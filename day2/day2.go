package day2

import (
	"aoc24/helper"
	"strconv"
	"strings"
)

func splitReport(line string) []int {
	var s = strings.Split(line, " ")
	var n []int
	for i := range len(s) {
		ne, err := strconv.Atoi(s[i])
		if err != nil {
			return nil
		}
		n = append(n, ne)
	}
	return n
}

func checkReport(index int, levels []int, direction string) bool {
	// if end of list its pass
	if index == len(levels)-1 {
		return true
	}

	diff := 0
	if direction == "decreasing" {
		diff = levels[index] - levels[index+1]
	} else if direction == "increasing" {
		diff = levels[index+1] - levels[index]
	}

	// if distancer >2 and not minus
	if diff > 3 || diff <= 0 {
		return false
	}
	// idk but this is so smart
	return checkReport(index+1, levels, direction)
}

func getReportState(line string) bool {
	var levels = splitReport(line)
	if levels[0] > levels[1] {
		return checkReport(0, levels, "decreasing")
	} else if levels[0] < levels[1] {
		return checkReport(0, levels, "increasing")
	} else {
		return false
	}
}

func Day2() {
	var t = helper.ReadLines("day2/day2.txt")
	var cnt = 0
	for _, line := range t {
		var x = getReportState(line)
		if x {
			cnt++
		}
	}
	println(cnt)
}
