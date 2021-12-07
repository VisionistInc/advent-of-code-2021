package main

import (
	"aoc2021/aoc"
	"fmt"
	"math"
)

type CostFunc func(int64, int64) float64

func cost1(s, e int64) float64 {
	return math.Abs(float64(e) - float64(s))
}

func cost2(s, e int64) float64 {
	d := math.Abs(float64(e) - float64(s))
	return (d / 2) * (d + 1)
}

func calcCost(initial []int64, costFunc CostFunc) (int64, float64) {
	min := aoc.Min(initial)
	max := aoc.Max(initial)

	minPos := int64(0)
	minFuel := float64(math.MaxFloat64)

	for p := min; p <= max; p++ {
		fuel := float64(0)
		for _, v := range initial {
			fuel += costFunc(p, v)
		}

		if fuel < minFuel {
			minFuel = fuel
			minPos = p
		}
	}

	return minPos, minFuel
}

func main() {

	aoc.ForLine("input.txt", func(line string) {
		initial := aoc.StringToInts(line, ",")

		_, f1 := calcCost(initial, cost1)
		fmt.Println(int64(f1))

		_, f2 := calcCost(initial, cost2)
		fmt.Println(int64(f2))
	})
}
