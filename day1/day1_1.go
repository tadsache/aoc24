package day1

import (
	"aoc24/helper"
	"sort"
	"strconv"
	"strings"
)

func Day1() {

	var txt = helper.ReadLines("day1/day1.txt")

	var left []int
	var right []int
	for _, line := range txt {
		var split = strings.Split(line, "   ")
		leftValue, _ := strconv.Atoi(split[0])
		rightValue, _ := strconv.Atoi(split[1])
		left = append(left, leftValue)
		right = append(right, rightValue)
	}

	sort.Ints(left)
	sort.Ints(right)

	var res int
	for i := range len(left) {
		var r int
		if right[i] > left[i] {
			r = right[i] - left[i]
		} else {
			r = left[i] - right[i]
		}
		res = res + r
	}

	println(res)

}
