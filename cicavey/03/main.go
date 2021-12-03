package main

import (
	"aoc2021/aoc"
	"fmt"
	"strconv"
)

type histbin struct {
	zeros int64
	ones  int64
}

func hist(lines []string) []histbin {
	count := make([]histbin, len(lines[0]))

	for _, line := range lines {
		for i, c := range line {
			if c == '0' {
				count[i].zeros++
			} else {
				count[i].ones++
			}
		}
	}

	return count
}

func main() {

	input := aoc.ReadLines("input.txt")

	// Part 1
	count := hist(input)
	gamma := int64(0)
	epsilon := int64(0)
	for _, b := range count {
		gamma <<= 1
		epsilon <<= 1

		if b.ones > b.zeros {
			gamma |= 1
		} else {
			epsilon |= 1
		}
	}
	fmt.Println(gamma * epsilon)

	currentInput := input
	pos := 0
	for {
		count = hist(currentInput)
		filterVal := '0'
		if count[pos].ones >= count[pos].zeros {
			filterVal = '1'
		}

		currentInput = aoc.Filter(currentInput, func(v string) bool {
			return []rune(v)[pos] == filterVal
		})

		pos++

		if len(currentInput) == 1 {
			break
		}
	}
	oxygenRating, _ := strconv.ParseInt(currentInput[0], 2, 64)

	currentInput = input
	pos = 0
	for {
		count = hist(currentInput)
		filterVal := '1'
		if count[pos].zeros <= count[pos].ones {
			filterVal = '0'
		}

		currentInput = aoc.Filter(currentInput, func(v string) bool {
			return []rune(v)[pos] == filterVal
		})

		pos++

		if len(currentInput) == 1 {
			break
		}
	}
	co2ScrubberRating, _ := strconv.ParseInt(currentInput[0], 2, 64)

	fmt.Println(oxygenRating * co2ScrubberRating)

}
