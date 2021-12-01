package main

import (
	"aoc2021/internal"
	"fmt"
)

func countInc(input []int64) int64 {
	var inc int64

	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			inc++
		}
	}

	return inc
}

func main() {
	input := internal.ReadIntLines("input.txt")

	fmt.Println(countInc(input))

	// Map input to 3-elt windows
	var windowedInput []int64
	for i := 2; i < len(input); i++ {
		w := input[i-2] + input[i-1] + input[i]
		windowedInput = append(windowedInput, w)
	}

	fmt.Println(countInc(windowedInput))
}
