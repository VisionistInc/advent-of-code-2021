package main

import (
	"aoc2021/aoc"
	"fmt"
	"math/bits"
	"strconv"
	"strings"
)

func binseg(value string) uint8 {
	out := uint8(0)
	for _, c := range value {
		out = out | (1 << (uint8(c) - 97))
	}
	return out
}

func main() {

	p1count := 0
	p2acc := int64(0)

	aoc.ForLine("input.txt", func(line string) {

		s := strings.Split(line, " | ")

		a := strings.Fields(s[0])
		b := strings.Fields(s[1])

		for _, od := range b {
			l := len(od)
			if l == 2 || l == 3 || l == 4 || l == 7 {
				p1count++
			}
		}

		// Part 2

		// Convert alpha strings into bits
		raw := aoc.Map(a, binseg)

		// Map and remove know quantities
		mapping := make(map[string]uint8)

		for _, v := range raw {
			l := bits.OnesCount8(v)
			switch l {
			case 2:
				mapping["1"] = v
			case 3:
				mapping["7"] = v
			case 4:
				mapping["4"] = v
			case 7:
				mapping["8"] = v
			}
		}

		raw = aoc.Filter(raw, func(s uint8) bool {
			l := bits.OnesCount8(s)
			return !(l == 2 || l == 3 || l == 4 || l == 7)
		})

		// Heuristics

		// Combine 4 with 9, 6, 0 to find 9
		val4 := mapping["4"]
		potential := aoc.Filter(raw, func(s uint8) bool {
			l := bits.OnesCount8(s)
			return l == 6
		})
		for _, pv := range potential {
			if val4&pv == val4 {
				// Found 9
				mapping["9"] = pv
				break
			}
		}
		val9 := mapping["9"]
		raw = aoc.Filter(raw, func(s uint8) bool {
			return s != val9
		})
		potential = aoc.Filter(potential, func(s uint8) bool {
			return s != val9
		})
		// potential now contains 0 and 6, check against 1 to
		val1 := mapping["1"]
		if potential[0]&val1 == val1 {
			mapping["0"] = potential[0]
			mapping["6"] = potential[1]
		} else {
			mapping["0"] = potential[1]
			mapping["6"] = potential[0]
		}
		val6 := mapping["6"]
		val0 := mapping["0"]
		raw = aoc.Filter(raw, func(s uint8) bool {
			return s != val6 && s != val0
		})

		// All that remains in raw is 2, 3, 5
		potential = aoc.Filter(raw, func(s uint8) bool {
			return s&val1 == val1
		})
		val3 := potential[0]
		mapping["3"] = val3
		raw = aoc.Filter(raw, func(s uint8) bool {
			return s != val3
		})

		// All that remains in raw is 2, 5, use 6 as a mask
		if bits.OnesCount8(raw[0]&val6) == 5 {
			mapping["5"] = raw[0]
			mapping["2"] = raw[1]
		} else {
			mapping["5"] = raw[1]
			mapping["2"] = raw[0]
		}

		// Reverse mapping for lookup
		lookup := make(map[uint8]string)
		for k, v := range mapping {
			lookup[v] = k
		}

		// convert each final digit, build into string, parse string
		final := aoc.Map(b, binseg)
		finalS := ""
		for _, fv := range final {
			finalS += lookup[fv]
		}

		finalV, _ := strconv.ParseInt(finalS, 10, 64)

		p2acc += finalV
	})

	fmt.Println(p1count)
	fmt.Println(p2acc)
}
