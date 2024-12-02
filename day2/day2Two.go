package day2

import (
	"aoc24/helper"
	"strconv"
	"strings"
)

func splitReport2(line string) []int {
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

func checkReport2(index int, levels []int, direction string) bool {

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
	return checkReport2(index+1, levels, direction)
}

func removeFromSlice(arr []int, pos int) []int {
	return append(arr[:pos], arr[pos+1:]...)
}

func recusreReportState(levels []int) bool {
	if levels[0] > levels[1] {
		return checkReport2(0, levels, "decreasing")
	} else if levels[0] < levels[1] {
		return checkReport2(0, levels, "increasing")
	} else {
		return false
	}
}

func withforloop(levels []int) bool {
	var x = recusreReportState(levels)
	if x {
		return true
	}

	for i := 0; i < len(levels); i++ {
		lvls := make([]int, len(levels))
		// wow... fails me again
		copy(lvls, levels)
		lvls = removeFromSlice(lvls, i)
		var xi = recusreReportState(lvls)
		if xi {
			return true
		}
	}
	return false
}

func getReportState2(line string) bool {
	var levels = splitReport2(line)
	return withforloop(levels)

}

func Day2Two() {
	var t = helper.ReadLines("day2/day2.txt")
	var cnt = 0
	for _, line := range t {
		var x = getReportState2(line)
		if x {
			cnt++
		}
	}
	println(cnt)
}
