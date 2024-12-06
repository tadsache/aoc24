package day3

import (
	"aoc24/helper"
	"regexp"
	"strconv"
	"strings"
)

func sum(str string) int {
	trim := strings.TrimSuffix(strings.TrimPrefix(str, "mul("), ")")

	numbers := strings.Split(trim, ",")
	one, _ := strconv.Atoi(numbers[0])
	two, _ := strconv.Atoi(numbers[1])
	return one * two
}

func Day3() {
	var text = helper.Read("day3/day3.txt")

	r := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)

	matches := r.FindAllString(text, -1)

	var retval int
	idk := true
	for _, match := range matches {
		// level one without the if
		if match == "do()" {
			idk = true
		} else if match == "don't()" {
			idk = false
		}
		if idk && strings.HasPrefix(match, "mul") {
			retval += sum(match)
		}
	}
	println(retval)
}
