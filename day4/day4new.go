package day4

import (
	"aoc24/helper"
	"fmt"
)

type point struct {
	x int
	y int
}

var direcs = []point{
	{x: -1, y: 0},  // left
	{x: 1, y: 0},   // right
	{x: 0, y: -1},  // up
	{x: 0, y: 1},   // down
	{x: -1, y: -1}, // upper left
	{x: 1, y: -1},  // upper right
	{x: -1, y: 1},  // lower left
	{x: 1, y: 1},   // lower right
}

var word = []rune{'X', 'M', 'A', 'S'}

func stringToRuneSlice(s string) []rune {
	var runes []rune
	for _, r := range s {
		runes = append(runes, r)
	}
	return runes
}

func walk(arr [][]rune, pos point, index int, direction []point) bool {
	if index == 4 {
		// base case
		fmt.Println("Match found starting at:", pos)
		return true
	}

	for _, direc := range direction {
		if pos.x+direc.x >= 0 && pos.x+direc.x < len(arr[0]) && pos.y+direc.y >= 0 && pos.y+direc.y < len(arr) {
			if arr[pos.y+direc.y][pos.x+direc.x] == word[index] {
				if walk(arr, point{x: pos.x + direc.x, y: pos.y + direc.y}, index+1, []point{direc}) {
					return true
				}
			}
		}

	}
	return false
}

func Day4New() {
	t := helper.ReadLines("day4/day4tst.txt")
	var runesArray [][]rune // Array of rune slices
	for _, s := range t {
		runes := stringToRuneSlice(s)
		runesArray = append(runesArray, runes)
	}

	var ret int
	for i, line := range runesArray {
		for j, r := range line {
			if r == 'X' {
				for _, dir := range direcs {
					p := point{x: j, y: i}
					if walk(runesArray, p, 1, []point{dir}) {
						ret++
					}
				}

			}
		}
	}

	println(ret)
}

var checkDirections = []point{
	{x: -1, y: 1},  //ul
	{x: 1, y: 1},   //ur
	{x: -1, y: -1}, //bl
	{x: 1, y: -1},  //br
}

func check(arr [][]rune, pos point) bool {

	for _, dirs := range checkDirections {
		if !(pos.x+dirs.x >= 0 && pos.x+dirs.x < len(arr[0]) && pos.y+dirs.y >= 0 && pos.y+dirs.y < len(arr)) {
			return false
		}
	}

	ul := arr[pos.y+1][pos.x-1]
	ur := arr[pos.y+1][pos.x+1]
	bl := arr[pos.y-1][pos.x-1]
	br := arr[pos.y-1][pos.x+1]

	if (ul == 'S' && ur == 'S' && bl == 'M' && br == 'M') || (ul == 'M' && ur == 'M' && bl == 'S' && br == 'S') || (ul == 'S' && ur == 'M' && bl == 'S' && br == 'M') || (ul == 'M' && ur == 'S' && bl == 'M' && br == 'S') {
		return true
	}
	return false

}

func Day4Two() {
	t := helper.ReadLines("day4/day4.txt")
	var runesArray [][]rune // Array of rune slices
	for _, s := range t {
		runes := stringToRuneSlice(s)
		runesArray = append(runesArray, runes)
	}
	var ret int
	for i, line := range runesArray {
		for j, r := range line {
			if r == 'A' {
				p := point{x: j, y: i}
				if check(runesArray, p) {
					ret++
				}
			}
		}
	}
	println(ret)
}
