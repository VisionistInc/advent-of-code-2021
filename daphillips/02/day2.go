package main

import (
	"aoc2021/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type SubPosition struct {
	Depth    int
	Position int
	Aim      int
}

// keep these consts as type string so that life is easier in the switch below
// ideally we would make a separate type for these enum values but I'm lazy
const (
	Forward string = "forward"
	Up             = "up"
	Down           = "down"
)

func part1(lines *[]string) {
	subPosition := SubPosition{0, 0, 0}
	for _, line := range *lines {
		splitLine := strings.Split(line, " ")
		direction := splitLine[0]
		magnitude, err := strconv.Atoi(splitLine[1])

		if err != nil {
			log.Fatal(err)
		}

		switch direction {
		case Forward:
			subPosition.Position += magnitude
		// remember: for a submarine, going up decreases depth, and going down increases it
		case Up:
			subPosition.Depth -= magnitude
		case Down:
			subPosition.Depth += magnitude
		}
	}

	fmt.Println(subPosition.Depth * subPosition.Position)
}

func part2(lines *[]string) {
	subPosition := SubPosition{0, 0, 0}

	for _, line := range *lines {
		splitLine := strings.Split(line, " ")
		direction := splitLine[0]
		magnitude, err := strconv.Atoi(splitLine[1])

		if err != nil {
			log.Fatal(err)
		}

		switch direction {
		case Forward:
			subPosition.Position += magnitude
			subPosition.Depth += subPosition.Aim * magnitude
		// remember: for a submarine, going up decreases depth, and going down increases it
		case Up:
			subPosition.Aim -= magnitude
		case Down:
			subPosition.Aim += magnitude
		}
	}

	fmt.Println(subPosition.Depth * subPosition.Position)
}

func main() {
	lines := utils.ReadFile("input.txt")
	part1(&lines)
	part2(&lines)
}
