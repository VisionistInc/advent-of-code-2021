package main

import (
	"aoc2021/aoc"
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

const (
	X = iota
	Y
)

type Instruction struct {
	dir   int
	value int
}

func PrintGrid(grid map[Point]bool) {
	w := 0
	h := 0
	for p := range grid {
		if p.x > w {
			w = p.x
		}
		if p.y > h {
			h = p.y
		}
	}
	for y := 0; y <= h; y++ {
		for x := 0; x <= w; x++ {
			_, ok := grid[Point{x, y}]
			if ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

}

func main() {

	grid := make(map[Point]bool)
	w := 0
	h := 0

	var instructions []Instruction

	aoc.ForLine("input.txt", func(line string) {
		if len(strings.TrimSpace(line)) == 0 {
			return
		}

		if strings.HasPrefix(line, "fold along") {
			temp := strings.Split(line[11:], "=")
			dir := X
			if temp[0] == "y" {
				dir = Y
			}
			v, _ := strconv.Atoi(temp[1])
			instructions = append(instructions, Instruction{dir, v})
			return
		}

		temp := strings.Split(line, ",")
		x, _ := strconv.Atoi(temp[0])
		y, _ := strconv.Atoi(temp[1])
		grid[Point{x, y}] = true

		if x > w {
			w = x
		}
		if y > h {
			h = y
		}
	})

	fmt.Println()

	for _, ins := range instructions {

		if ins.dir == X {
			newgrid := make(map[Point]bool)
			for p := range grid {
				// Point on other side of fold
				if p.x > ins.value {
					// reflect the point subtracting twice the delta
					delta := p.x - ins.value
					newp := Point{p.x - 2*delta, p.y}
					newgrid[newp] = true

				} else {
					newgrid[p] = true
				}
			}

			grid = newgrid
		} else {
			newgrid := make(map[Point]bool)
			for p := range grid {
				// Point on other side of fold
				if p.y > ins.value {
					// reflect the point subtracting twice the delta
					delta := p.y - ins.value
					newp := Point{p.x, p.y - 2*delta}
					newgrid[newp] = true

				} else {
					newgrid[p] = true
				}
			}

			grid = newgrid
		}
	}

	PrintGrid(grid)
}
