package day5

import (
	"aoc24/helper"
	"fmt"
	"strings"
)

func splitInput(lines []string) ([]string, []string) {
	var sortLines []string
	var resLines []string
	var x = 0
	for _, line := range lines {
		if line == "" {
			x = 1
		} else if x == 0 {
			sortLines = append(sortLines, line)
		} else if x == 1 {
			resLines = append(resLines, line)
		}
	}
	return sortLines, resLines
}

func checkInList(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func removeElement(slice []string, index int) []string {
	// Check if the index is within the bounds of the slice
	if index < 0 || index >= len(slice) {
		fmt.Println("Index out of range")
		return slice // Return the unchanged slice
	}
	// Remove the element at the index
	return append(slice[:index], slice[index+1:]...)
}

func sort(lines []string) []string {

	var srtarr []string
	for _, line := range lines {
		split := strings.Split(line, "|")
		lft := split[0]
		rgt := split[1]

		if len(srtarr) == 0 {
			// start point
			srtarr = append(srtarr, lft)
			srtarr = append(srtarr, rgt)
		} else {

			// case
			//1. both are not in the list
			//2. left is in the list
			//3. right is the list
			lftb := checkInList(srtarr, lft)
			rgtb := checkInList(srtarr, rgt)

			if !lftb && !rgtb {
				srtarr = append(srtarr, lft)
				srtarr = append(srtarr, rgt)
			} else if lftb && !rgtb {
				srtarr = appendAfter(srtarr, lft, rgt)
			} else if !lftb && rgtb {
				srtarr = appendBefore(srtarr, lft, rgt)
			} else {
				// if both are there check if left is beeing legt
				var lv int
				var rv int
				for i, v := range srtarr {
					if v == lft {
						lv = i
					}
					if v == rgt {
						rv = i
					}
				}

				if lv > rv {
					// do stuff
					switchSlices(srtarr, lv, rv)
				}
			}
		}
	}
	return srtarr
}

func switchSlices(arr []string, lftv int, rgtv int) []string {
	subSlice := arr[rgtv:lftv] // get the slice

	// Loop from 2 to 5 inclusive
	for i := lftv; i < rgtv; i++ {
		removeElement(arr, i)
	}

	arr = append(arr, subSlice...)
	return arr
}

func appendBefore(arr []string, lft string, rgt string) []string {
	for i, val := range arr {
		if val == rgt {
			arr = append(arr[:i], append([]string{lft}, arr[i:]...)...)
			break
		}
	}
	return arr // fixme do return inside the loop
}

func appendAfter(arr []string, lft string, rgt string) []string {
	for i, val := range arr {
		if val == lft {
			arr = append(arr[:i+1], append([]string{rgt}, arr[i+1:]...)...)
			break
		}
	}
	return arr
}

func Day5() {

	lines := helper.ReadLines("day5/day5tst.txt")
	sortLines, _ := splitInput(lines)

	var x = sort(sortLines)
	println(x)

}
