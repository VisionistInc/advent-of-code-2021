package main

import (
	"aoc2021/utils"
	"fmt"
	"strconv"
)

type Point struct {
	x, y int
}

type Grid struct {
	grid       map[Point]int
	maxX, maxY int
}

func (p Point) translate(dx, dy int) Point {
	return Point{p.x + dx, p.y + dy}
}

func (g *Grid) tick(flashed utils.Set[Point], p Point) int64 {
	// first, increment the point
	if flashed.Contains(p) {
		return 0
	}
	g.grid[p]++
	if g.grid[p] > 9 {
		// add this point to the set and zero out its value
		g.grid[p] = 0
		flashed.Add(p)

		numFlashed := int64(1)

		// flash and increment neighbors
		var neighbor Point
		if p.x != 0 && p.y != 0 {
			neighbor = p.translate(-1, -1)
			if !flashed.Contains(neighbor) {
				numFlashed += g.tick(flashed, neighbor)
			}
		}
		if p.x != g.maxX && p.y != g.maxY {
			neighbor = p.translate(1, 1)
			if !flashed.Contains(neighbor) {
				numFlashed += g.tick(flashed, neighbor)
			}
		}
		if p.x != 0 {
			neighbor = p.translate(-1, 0)
			if !flashed.Contains(neighbor) {
				numFlashed += g.tick(flashed, neighbor)
			}
		}
		if p.y != 0 {
			neighbor = p.translate(0, -1)
			if !flashed.Contains(neighbor) {
				numFlashed += g.tick(flashed, neighbor)
			}
		}
		if p.x != g.maxX {
			neighbor = p.translate(1, 0)
			if !flashed.Contains(neighbor) {
				numFlashed += g.tick(flashed, neighbor)
			}
		}
		if p.y != g.maxY {
			neighbor = p.translate(0, 1)
			if !flashed.Contains(neighbor) {
				numFlashed += g.tick(flashed, neighbor)
			}
		}
		if p.x != 0 && p.y != g.maxY {
			neighbor = p.translate(-1, 1)
			if !flashed.Contains(neighbor) {
				numFlashed += g.tick(flashed, neighbor)
			}
		}
		if p.x != g.maxX && p.y != 0 {
			neighbor = p.translate(1, -1)
			if !flashed.Contains(neighbor) {
				numFlashed += g.tick(flashed, neighbor)
			}
		}

		return numFlashed
	}
	return 0
}

func part1(lines []string) {
	maxX := len(lines[0]) - 1
	maxY := len(lines) - 1
	g := Grid{make(map[Point]int, 0), maxX, maxY}
	for y, line := range lines {
		for x, r := range line {
			num, _ := strconv.Atoi(string(r))
			g.grid[Point{x, y}] = num
		}
	}

	numFlashed := int64(0)

	for i := 0; i < 100; i++ {
		flashed := make(utils.Set[Point], 0)
		for p := range g.grid {
			numFlashed += g.tick(flashed, p)
		}
	}
	fmt.Println(numFlashed)
}

func (g Grid) synchronizedFlash() bool {
	for _, energy := range g.grid {
		if energy != 0 {
			return false
		}
	}
	return true
}

func part2(lines []string) {
	maxX := len(lines[0]) - 1
	maxY := len(lines) - 1
	g := Grid{make(map[Point]int, 0), maxX, maxY}
	for y, line := range lines {
		for x, r := range line {
			num, _ := strconv.Atoi(string(r))
			g.grid[Point{x, y}] = num
		}
	}

	numTicks := int64(0)
	for !g.synchronizedFlash() {
		numTicks++
		flashed := make(utils.Set[Point], 0)
		for p := range g.grid {
			g.tick(flashed, p)
		}
	}
	fmt.Println(numTicks)
}

func main() {
	lines := utils.ReadFile("input.txt")
	part1(lines)
	part2(lines)
}
