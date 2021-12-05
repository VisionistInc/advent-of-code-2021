package main

import (
	"aoc2021/aoc"
	"fmt"
	"strings"
)

type Point struct {
	x, y int64
}

func main() {

	grid := make(map[Point]int64)

	aoc.ForLine("input.txt", func(line string) {
		line = strings.ReplaceAll(line, " -> ", " ")
		line = strings.ReplaceAll(line, ",", " ")
		v := aoc.StringToInts(line, "\\s+")
		p1 := Point{v[0], v[1]}
		p2 := Point{v[2], v[3]}

		dx := p1.x - p2.x
		dy := p1.y - p2.y

		// Non-Vert/Horz line
		// Put pack in for part 1
		// if dx != 0 && dy != 0 {
		// 	return
		// }

		if dx != 0 {
			if dx < 0 {
				dx = 1
			} else {
				dx = -1
			}
		}

		if dy != 0 {
			if dy < 0 {
				dy = 1
			} else {
				dy = -1
			}
		}
		// fmt.Println(p1, p2, dx, dy)

		grid[p1]++
		for p1 != p2 {
			p1.x += dx
			p1.y += dy
			grid[p1]++
		}
	})

	var danger int64
	for _, v := range grid {
		if v > 1 {
			danger++
		}
	}

	fmt.Println(danger)
}
