package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// gets the data for a file line by line
func readFile(fileName string) []string {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	return lines
}

func part1(lines *[]string) {
	var previous, numIncreasing int
	for i, line := range *lines {
		depth, err := strconv.Atoi(line)

		if err != nil {
			log.Fatal(err)
		}

		if i != 0 && previous < depth {
			numIncreasing++
		}

		previous = depth
	}

	fmt.Println("numIncreasing =", numIncreasing)
}

func part2(lines *[]string) {
	numIncreasing := 0
	var previous [3]int
	for i, line := range *lines {
		depth, err := strconv.Atoi(line)

		if err != nil {
			log.Fatal(err)
		}

		if i > 2 {
			currentSum := depth + previous[2] + previous[1]
			previousSum := previous[2] + previous[1] + previous[0]
			if currentSum > previousSum {
				numIncreasing++
			}
		}
		// shift everything along
		// a deeply-accessible queue would have been perfect here, but I was too lazy to build it and didn't look to see if the standard lib had anything similar
		previous[0] = previous[1]
		previous[1] = previous[2]
		previous[2] = depth
	}
	fmt.Println("numIncreasing = ", numIncreasing)
}

func main() {
	lines := readFile("input.txt")
	part1(&lines)
	part2(&lines)
}
