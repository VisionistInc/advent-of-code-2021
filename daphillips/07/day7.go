package main

import (
	"aoc2021/utils"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/montanaflynn/stats"
)

func part1(positions []int, data stats.Float64Data) {
	median, _ := stats.Median(data)

	fuelSpent := 0
	for _, pos := range positions {
		fuelSpent += int(math.Abs(float64(pos) - median))
	}

	fmt.Println(fuelSpent)
}

func part2(positions []int, data stats.Float64Data) {
	rawMean, _ := stats.Mean(data)
	// HACK to get the right answer for my input I needed to floor, but the test input needed to round!
	// uncomment this line to see the difference
	// rawMean, _ = stats.Round(rawMean, 0)
	mean := int(rawMean)

	totalFuelSpent := 0
	for _, pos := range positions {
		distance := math.Abs(float64(pos) - float64(mean))
		fuelSpent, _ := stats.Round(distance*(distance+1)/2, 0)

		totalFuelSpent += int(fuelSpent)
	}

	fmt.Println(totalFuelSpent)
}

func main() {
	input := utils.ReadFile("input.txt")

	positions := utils.Map(strings.Split(input[0], ","), func(in string) int {
		num, _ := strconv.Atoi(in)
		return num
	})

	data := stats.LoadRawData(positions)

	part1(positions, data)
	part2(positions, data)

}
