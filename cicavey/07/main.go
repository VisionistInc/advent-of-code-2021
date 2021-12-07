package main

import (
	"aoc2021/aoc"
	"fmt"
	"math"
	"sort"
	"time"
)

type CostFunc func(int64, int64) float64

func cost1(s, e int64) float64 {
	return math.Abs(float64(e) - float64(s))
}

func cost2(s, e int64) float64 {
	d := math.Abs(float64(e) - float64(s))
	return (d / 2) * (d + 1)
}

type PivotFunc func([]int64) (int64, int64)

func pivot1(v []int64) (int64, int64) {
	// sort a copy to find median
	vCopy := make([]int64, len(v))
	copy(vCopy, v)
	sort.Slice(vCopy, func(i, j int) bool { return vCopy[i] < vCopy[j] })
	// find median
	median := float64(len(vCopy)) / 2.0
	// either side of median
	return vCopy[int64(math.Floor(median))], vCopy[int64(math.Ceil(median))]
}

func pivot2(v []int64) (int64, int64) {
	mean := aoc.Mean(v)
	return int64(math.Floor(mean)), int64(math.Ceil(mean))
}

func pivot0(v []int64) (int64, int64) {
	return aoc.Min(v), aoc.Max(v)
}

func calcCost(initial []int64, cf CostFunc, pf PivotFunc) (int64, float64) {
	// Uses knowledge to
	min, max := pf(initial)

	minPos := int64(0)
	minFuel := float64(math.MaxFloat64)

	for p := min; p <= max; p++ {
		fuel := float64(0)
		for _, v := range initial {
			fuel += cf(p, v)
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

		s := time.Now()
		p1, f1 := calcCost(initial, cost1, pivot0)
		fmt.Println(int64(f1), p1, time.Since(s))

		s = time.Now()
		p1, f1 = calcCost(initial, cost1, pivot1)
		fmt.Println(int64(f1), p1, time.Since(s))

		s = time.Now()
		p2, f2 := calcCost(initial, cost2, pivot0)
		fmt.Println(int64(f2), p2, time.Since(s))

		s = time.Now()
		p2, f2 = calcCost(initial, cost2, pivot2)
		fmt.Println(int64(f2), p2, time.Since(s))
	})
}
