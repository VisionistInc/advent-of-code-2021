package main

import (
	"aoc2021/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Line struct {
	Start utils.Point[int]
	End   utils.Point[int]
}

func (l Line) Fill(grid map[utils.Point[int]]int) {
	curPoint := l.Start
	for curPoint.X != l.End.X || curPoint.Y != l.End.Y {
		// hopefully it stores the value and not the reference...
		grid[curPoint] = grid[curPoint] + 1

		// because diagnoal is only 45 degrees, consider x first then y
		// call-site is responsible for avoiding diagonals if desired
		if curPoint.X < l.End.X {
			curPoint.X++
		} else if curPoint.X > l.End.X {
			curPoint.X--
		}

		if curPoint.Y < l.End.Y {
			curPoint.Y++
		} else if curPoint.Y > l.End.Y {
			curPoint.Y--
		}
	}

	// make sure to add in the last point!
	grid[curPoint] = grid[curPoint] + 1
}

func part1(lines []Line) {
	overlaps := make(map[utils.Point[int]]int, 0)

	for _, line := range lines {
		// ignore non-diagonals for now
		if line.Start.X == line.End.X || line.Start.Y == line.End.Y {
			line.Fill(overlaps)
		}
	}

	pointsWith2Lines := 0
	for _, v := range overlaps {
		if v >= 2 {
			pointsWith2Lines++
		}
	}

	fmt.Println(pointsWith2Lines)

}

func part2(lines []Line) {
	overlaps := make(map[utils.Point[int]]int, 0)

	for _, line := range lines {
		line.Fill(overlaps)
	}

	pointsWith2Lines := 0
	for _, v := range overlaps {
		if v >= 2 {
			pointsWith2Lines++
		}
	}

	fmt.Println(pointsWith2Lines)
}

func main() {
	input := utils.ReadFile("input.txt")
	// input := utils.ReadFile("test-input.txt")
	lines := make([]Line, 0)

	stringToIntMapper := func(in string) int {
		num, err := strconv.Atoi(in)
		if err != nil {
			log.Panic(err)
		}

		return num
	}

	for _, inputLine := range input {
		points := strings.Split(inputLine, " -> ")
		point1 := utils.Map(strings.Split(points[0], ","), stringToIntMapper)
		point2 := utils.Map(strings.Split(points[1], ","), stringToIntMapper)
		lines = append(lines, Line{utils.Point[int]{X: point1[0], Y: point1[1]}, utils.Point[int]{X: point2[0], Y: point2[1]}})
	}

	part1(lines)
	part2(lines)
}
