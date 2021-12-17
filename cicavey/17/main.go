package main

import (
	"constraints"
	"fmt"
)

type Number interface {
	constraints.Float | constraints.Integer
}

type Vec[T Number] struct {
	x, y T
}

func (v *Vec[T]) Add(a Vec[T]) {
	v.x += a.x
	v.y += a.y
}

func sign[T constraints.Signed](n T) T {
	if n > 0 {
		return 1
	} else if n < 0 {
		return -1
	} else {
		return 0
	}
}

func main() {

	// target area: x=20..30, y=-10..-5
	// bx := []int{20, 30}
	// by := []int{-5, -10}

	// target area: x=265..287, y=-103..-58
	bx := []int{265, 287}
	by := []int{-58, -103}

	var omaxy int
	var iv int

	// Can't ever by negative or move further than outer bounds in one step
	for vx := 0; vx < 288; vx++ {
		// guess and test...
		for vy := -200; vy < 200; vy++ {

			var p Vec[int]
			v := Vec[int]{vx, vy}

			maxy := 0

			// guess ... there is probably some way to optimally solve this
			for step := 0; step < 250; step++ {
				p.Add(v)
				v.Add(Vec[int]{-sign(v.x), -1})

				if p.y > maxy {
					maxy = p.y
				}

				if p.x >= bx[0] && p.x <= bx[1] && (p.y <= by[0] && p.y >= by[1]) {
					iv++
					if maxy > omaxy {
						omaxy = maxy
					}

					break
				}
			}

		}
	}

	fmt.Println(omaxy)
	fmt.Println(iv)
}
