package main

import (
	"aoc2021/aoc"
	"fmt"
	"math"
	"strings"
)

func iterate(polymer string, lookup map[string]string, steps int) int64 {

	// Break polymer into pairs
	pairs := make(map[string]int64)
	hist := make(map[string]int64)

	for i := 0; i < len(polymer)-1; i++ {
		pairs[polymer[i:i+2]] = 1
		hist[polymer[i:i+1]]++
	}
	hist[polymer[len(polymer)-1:]]++

	for s := 0; s < steps; s++ {
		newpairs := make(map[string]int64)

		// ugly copy
		for k, v := range pairs {
			newpairs[k] = v
		}

		for pair, count := range pairs {
			// each pair maps to two pairs
			insert := lookup[pair]
			a := pair[0:1] + insert
			b := insert + pair[1:2]

			// reduce count of source pair :: this line took forever to find
			newpairs[pair] -= count
			// increment counts of dest pairs
			newpairs[a] += count
			newpairs[b] += count

			// track individual letters because it's impossible to disambiguate pairs later
			hist[insert] += count
		}
		pairs = newpairs
	}

	min := int64(math.MaxInt64)
	max := int64(math.MinInt64)
	for _, c := range hist {
		if c > max {
			max = c
		}
		if c < min {
			min = c
		}
	}

	return max - min
}

func main() {
	lines := aoc.ReadLines("input.txt")

	polymer := lines[0]

	insm := make(map[string]string)
	for _, ins := range lines[2:] {
		inss := strings.Split(ins, " -> ")
		insm[inss[0]] = inss[1]
	}

	fmt.Println(iterate(polymer, insm, 10))
	fmt.Println(iterate(polymer, insm, 40))
}
