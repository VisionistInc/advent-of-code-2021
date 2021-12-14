package main

import (
	"aoc2021/utils"
	"fmt"
	"math"
	"strings"
)

func main() {
	lines := utils.ReadFile("input.txt")

	pairs := make(map[string]int64, 0)

	mapper := make(map[string][]string, 0)

	start := []rune(lines[0])

	for i := 0; i < len(start)-1; i++ {
		pairs[string(start[i:i+2])]++
	}

	for _, line := range lines[2:] {
		// fmt.Println(line)
		split := strings.Split(line, " -> ")
		generatedPairs := make([]string, 0)
		pair := []rune(split[0])
		newChar := []rune(split[1])

		generatedPairs = append(generatedPairs, string(pair[0])+string(newChar[0]))
		generatedPairs = append(generatedPairs, string(newChar[0])+string(pair[1]))

		mapper[string(pair)] = generatedPairs
	}

	// part 1 is 10 times
	for i := 0; i < 10; i++ {
		newPairs := make(map[string]int64, 0)
		for pair, numToIncrement := range pairs {
			for _, newPair := range mapper[pair] {
				newPairs[newPair] += numToIncrement
			}
		}
		pairs = newPairs
	}

	numChars := make(map[string]int64, 0)
	for pair, num := range pairs {
		for _, char := range pair {
			numChars[string(char)] += num
		}
	}

	// we have to cut in half because the expansion results in double counting each char in the pairs map
	for pair, num := range numChars {
		numChars[pair] = int64(math.Ceil(float64(num) / 2))
	}

	maxCount := int64(0)
	minCount := int64(math.MaxInt64)

	for _, count := range numChars {
		if count > maxCount {
			maxCount = count
		} else if count < minCount {
			minCount = count
		}
	}

	fmt.Println(maxCount - minCount)

	// part 2 is 40, but we're keeping what we have for the first 10 so only do 30 more iterations
	for i := 0; i < 30; i++ {
		newPairs := make(map[string]int64, 0)
		for pair, numToIncrement := range pairs {
			for _, newPair := range mapper[pair] {
				newPairs[newPair] += numToIncrement
			}
		}
		pairs = newPairs
	}

	numChars = make(map[string]int64, 0)
	for pair, num := range pairs {
		for _, char := range pair {
			numChars[string(char)] += num
		}
	}

	// we have to cut in half because the expansion results in double counting each char in the pairs map
	for pair, num := range numChars {
		numChars[pair] = int64(math.Ceil(float64(num) / 2))
	}

	maxCount = int64(0)
	minCount = int64(math.MaxInt64)

	for _, count := range numChars {
		if count > maxCount {
			maxCount = count
		} else if count < minCount {
			minCount = count
		}
	}

	fmt.Println(maxCount - minCount)

}
