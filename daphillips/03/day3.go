package main

import (
	"aoc2021/utils"
	"fmt"
	"log"
	"strconv"
)

func part1(lines []string) {
	// because these are zero and our input is sufficiently small, we can safely ignore two's complement
	gammaRate := 0
	epsilonRate := 0
	// assuming all the strings are the same size
	for i := 0; i < len(lines[0]); i++ {
		numZeros := 0
		numOnes := 0
		for _, line := range lines {
			if line[i] == '0' {
				numZeros++
			} else {
				numOnes++
			}
		}

		// gamma rate takes the bigger of the two, while epsilonRate takes the smaller
		if numOnes > numZeros {
			gammaRate = gammaRate<<1 | 0b1
			epsilonRate = epsilonRate<<1 | 0b0
		} else {
			gammaRate = gammaRate<<1 | 0b0
			epsilonRate = epsilonRate<<1 | 0b1
		}
	}
	fmt.Println(gammaRate * epsilonRate)
}

// this could probably do well as a generic! and probably a map too...
func filterList(lines []string, filter func(string) bool) []string {
	filteredLines := make([]string, 0)

	for _, line := range lines {
		if filter(line) {
			filteredLines = append(filteredLines, line)
		}
	}

	return filteredLines
}

// filterTrue is used to flip the predicate in the case of the co2 filer... which is the opposite of the o2 one
func getFilteredNumber(lines []string, filterTrue bool) int64 {
	var filter func(string) bool

	for i := 0; i < len(lines[0]); i++ {
		numZeros := 0
		numOnes := 0
		for _, line := range lines {
			if line[i] == '0' {
				numZeros++
			} else {
				numOnes++
			}
		}
		if numOnes >= numZeros {
			filter = func(s string) bool {
				// `filterTrue` determines if we want to find `true` or `false` values with our filter
				return (s[i] == '1') == filterTrue
			}
		} else {
			filter = func(s string) bool {
				return (s[i] == '0') == filterTrue
			}
		}

		// go is pass-by-value, and a reference is a value
		// so, once we return from this scope the updated reference is dropped
		lines = filterList(lines, filter)

		if len(lines) == 1 {
			result, err := strconv.ParseInt(lines[0], 2, 0)
			if err != nil {
				log.Fatal(err)
			}

			return result
		}
	}
	panic("Didn't end up with a single result at the end")
}

func part2(lines []string) {
	o2Generator := getFilteredNumber(lines, true)
	co2Scrubber := getFilteredNumber(lines, false)

	fmt.Println(o2Generator * co2Scrubber)
}

func main() {
	// test input
	// lines := []string{
	// 	"00100",
	// 	"11110",
	// 	"10110",
	// 	"10111",
	// 	"10101",
	// 	"01111",
	// 	"00111",
	// 	"11100",
	// 	"10000",
	// 	"11001",
	// 	"00010",
	// 	"01010",
	// }
	lines := utils.ReadFile("input.txt")
	part1(lines)
	part2(lines)
}
