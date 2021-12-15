package main

import (
	"aoc2021/utils"
	"fmt"
	"log"
	"strconv"

	"github.com/beefsack/go-astar"
)

type Grid struct {
	g          [][]*Point
	maxX, maxY int
}

type Point struct {
	p      utils.Point[int]
	weight int
	g      Grid // back reference to containing map for A* algo
}

func (p *Point) PathNeighbors() []astar.Pather {
	neighbors := make([]astar.Pather, 0)

	if p.p.X-1 >= 0 {
		neighbors = append(neighbors, p.g.g[p.p.Y][p.p.X-1])
	}
	if p.p.Y-1 >= 0 {
		neighbors = append(neighbors, p.g.g[p.p.Y-1][p.p.X])
	}
	if p.p.X+1 < p.g.maxX {
		neighbors = append(neighbors, p.g.g[p.p.Y][p.p.X+1])
	}
	if p.p.Y+1 < p.g.maxY {
		neighbors = append(neighbors, p.g.g[p.p.Y+1][p.p.X])
	}

	return neighbors
}

func (p *Point) PathNeighborCost(to astar.Pather) float64 {
	toPoint := to.(*Point)
	return float64(toPoint.weight)
}

func (p *Point) PathEstimatedCost(to astar.Pather) float64 {
	toPoint := to.(*Point)
	absX := toPoint.p.X - p.p.X
	if absX < 0 {
		absX = -absX
	}

	absY := toPoint.p.Y - p.p.Y
	if absY < 0 {
		absY = -absY
	}

	return float64(absX + absY)
}

func part1(lines []string) {
	maxX := len(lines[0])
	maxY := len(lines)

	g := make([][]*Point, maxY)

	grid := Grid{g, maxX, maxY}

	for i, line := range lines {
		grid.g[i] = make([]*Point, maxX)
		for j, weight := range line {
			w, _ := strconv.Atoi(string(weight))
			grid.g[i][j] = &Point{utils.Point[int]{X: j, Y: i}, w, grid}
		}
	}

	start := grid.g[0][0]
	end := grid.g[maxY-1][maxX-1]

	_, distance, found := astar.Path(start, end)
	if !found {
		log.Fatal("could not find a path!")
	}

	fmt.Println(distance)
}

func part2(lines []string) {
	// beeeg yoshi mode
	tileX := len(lines[0])
	tileY := len(lines)

	maxX := tileX * 5
	maxY := tileY * 5

	g := make([][]*Point, maxY)

	grid := Grid{g, maxX, maxY}

	// 4-nested for loops ftw!
	// k maps to i
	for k := 0; k < 5; k++ {
		// l maps to j
		for l := 0; l < 5; l++ {
			for i, line := range lines {
				y := i + (k * (tileY))
				if len(grid.g[y]) == 0 {
					grid.g[y] = make([]*Point, maxX)
				}
				for j, weight := range line {
					x := j + (l * (tileX))
					w, _ := strconv.Atoi(string(weight))
					w += k + l
					if w > 9 {
						// wrap back around if we're over 9
						w -= 9
					}
					grid.g[y][x] = &Point{utils.Point[int]{X: x, Y: y}, w, grid}
				}
			}
		}
	}

	start := grid.g[0][0]
	end := grid.g[maxY-1][maxX-1]

	_, distance, found := astar.Path(start, end)
	if !found {
		log.Fatal("could not find a path!")
	}

	fmt.Println(distance)
}

func main() {
	lines := utils.ReadFile("input.txt")
	part1(lines)
	part2(lines)
}
