package main

import (
	"aoc2021/aoc"
	"fmt"
	"sort"
)

type Point struct {
	x, y, z  int
}

type Grid struct {
	w, h int
	data []int
}

func (g Grid) boundedLookup(x, y int) int {
	if x < 0 || x >= g.w {
		return 9
	}
	if y < 0 || y >= g.h {
		return 9
	}
	return g.data[y * g.w + x]
}

func main() {
	lines := aoc.ReadLines("input.txt")

	width := len(lines[0])
	height := len(lines)

	grid := Grid{width, height, make([]int, width * height)}

	for y, line := range lines {
		for x, c := range line {
			v := int(c) - 48
			grid.data[y * width + x] = v
		}
	}

	sumRisk := 0

	var lowPoints []*Point

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			p := y * width + x
			v := grid.data[p]

			low := true

			// Consider neighbors
			// Up
			low = low && v < grid.boundedLookup(x, y-1)

			// Down
			low = low && v < grid.boundedLookup(x, y+1)

			// Left
			low = low && v < grid.boundedLookup(x-1, y)

			// Right
			low = low && v < grid.boundedLookup(x+1, y)

			if low {
				sumRisk += (v+1)
				lowPoints = append(lowPoints, &Point{x,y,v})
			}
		}
	}
	fmt.Println(sumRisk)

	// Part 2

	// Iterate low points and find/fill basin
	var basins []int

	for _, lowPoint := range lowPoints {

		var q []*Point
		var cur *Point
		var basin = make(map[Point]bool)

		basin[*lowPoint] = true
		q = append(q, lowPoint)

		for len(q) > 0 {
			cur, q = aoc.Pop(q, nil)
			if cur == nil {
				break
			}

			// Process neighbors
			n := grid.boundedLookup(cur.x, cur.y-1)
			if n != 9 {
				p := Point{cur.x, cur.y-1, n}
				if _, ok := basin[p]; !ok {
					basin[p] = true
					q =  append(q, &p)
				}
			}

			n = grid.boundedLookup(cur.x, cur.y+1)
			if n != 9 {
				p := Point{cur.x, cur.y+1, n}
				if _, ok := basin[p]; !ok {
					basin[p] = true
					q =  append(q, &p)
				}
			}

			n = grid.boundedLookup(cur.x-1, cur.y)
			if n != 9 {
				p := Point{cur.x-1, cur.y, n}
				if _, ok := basin[p]; !ok {
					basin[p] = true
					q =  append(q, &p)
				}
			}

			n = grid.boundedLookup(cur.x+1, cur.y)
			if n != 9 {
				p := Point{cur.x+1, cur.y, n}
				if _, ok := basin[p]; !ok {
					basin[p] = true
					q =  append(q, &p)
				}
			}
		}

		basins = append(basins, len(basin))
	}

	sort.Slice(basins, func(i, j int) bool {
		return basins[i] > basins[j]
	})

	fmt.Println(basins[0] * basins[1] * basins[2])
}