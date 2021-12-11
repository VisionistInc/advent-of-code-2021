package main

import (
	"aoc2021/aoc"
	"fmt"
	"strings"
)

type Point struct {
	x, y int
}

type Grid struct {
	w, h int
	data []int
}

func  (g *Grid) each(cf func(x, y, v int) int) {
	for y := 0; y < g.h; y++ {
		for x := 0; x < g.w; x++ {
			g.data[y * g.w + x] = cf(x, y, 	g.data[y * g.w + x])
		}
	}
}

func (g *Grid) set(x, y, v int) {
	// silently let bounds check fail. 
	if x < 0 || x >= g.w {
		return
	}
	if y < 0 || y >= g.h {
		return
	}
	g.data[y * g.w + x] = v
}

func (g *Grid) inc(x, y, v int) {
	// silently let bounds check fail. 
	if x < 0 || x >= g.w {
		return
	}
	if y < 0 || y >= g.h {
		return
	}
	g.data[y * g.w + x] += v
}

func (g *Grid) get(x, y int) int {
	// silently let bounds check fail. 
	if x < 0 || x >= g.w {
		return -1
	}
	if y < 0 || y >= g.h {
		return -1
	}
	return g.data[y * g.w + x]
}

// Check if _all_ values equal given value
func (g *Grid) eq(v int) bool {
	for _, iv := range g.data {
		if iv != v {
			return false
		}
	}
	return true
}

func (g Grid) String() string {
	var out strings.Builder

	for y := 0; y < g.h; y++ {
		for x := 0; x < g.w; x++ {
			fmt.Fprintf(&out, "%1d", g.data[y * g.w + x])
		}
		out.WriteString("\n")
	}

	return out.String()
}

func main() {

	grid := &Grid{10, 10, make([]int, 100)}

	lines := aoc.ReadLines("input.txt")

	for y, line := range lines {
		for x, c := range line {
			grid.set(x, y, int(c) - 48)
		}
	}


	totalFlashes := 0
	for step := 0; step < 1000; step++ {
	
		// Increment all
		grid.each(func(x, y, v int) int {
			return v + 1
		})

		// Process flashes
		var visted = make(map[Point]bool)

		for {
			var newFlash = 0
			for y := 0; y < grid.h; y++ {
				for x := 0; x < grid.w; x++ {
					v := grid.get(x, y)

					if v <= 9 {
						continue
					}

					p := Point{x,y}
					_, ok := visted[p]

					if ok {
						continue
					}

					newFlash++
					visted[p] = true
					// increment all neighbors
					grid.inc(x - 1, y - 1, 1)
					grid.inc(x,     y - 1, 1)
					grid.inc(x + 1, y - 1, 1)
					grid.inc(x - 1, y    , 1)
					grid.inc(x + 1, y    , 1)
					grid.inc(x - 1, y + 1, 1)
					grid.inc(x,     y + 1, 1)
					grid.inc(x + 1, y + 1, 1)
				}
			}

			if newFlash == 0 {
				break
			}
		}

		totalFlashes += len(visted)

		for p, _ := range visted {
			grid.set(p.x, p.y, 0)
		}

		if step == 99 {
			fmt.Printf("Step %d, flashes=%d\n", step+1, totalFlashes)
			fmt.Println(grid)
		}

		if grid.eq(0) {
			fmt.Printf("All flashed on step %d\n", step+1)
			break
		}
	}
}
