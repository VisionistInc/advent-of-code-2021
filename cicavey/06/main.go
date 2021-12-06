package main

import (
	"aoc2021/aoc"
	"fmt"
)

func simulate(hist []int64, days int) int64 {

	// Sure, we accept slices ... but only the right kind of slice. See simulate9
	if len(hist) != 9 {
		panic("doh")
	}

	// This is some tension here wrt fixed length vs slices and the logic
	// Could this have been safer as fixed length? But then the generic Sum function wouldn't work... (unless I'm missing something)

	// copy input into work area
	f := make([]int64, len(hist))
	copy(f, hist)

	// use some mod math to rotate array...
	for iter, idx := 0, 0; iter < (days - 1); iter++ {

		// "shift" everthing over aka substract
		idx = (idx + 1) % 9
		// 0s spawn new 8s (offset 9), overwrite is okay since there are no 8s at this point
		f[(idx+9)%9] = f[idx]
		// 0s increment existing count of 6s (offset 7)
		f[(idx+7)%9] += f[idx]
	}

	return aoc.Sum(f)
}

func simulate9(hist [9]int64, days int) int64 {

	// This is some tension here wrt fixed length vs slices and the logic
	// Could this have been safer as fixed length? But then the generic Sum function wouldn't work... (unless I'm missing something)

	// copy input into work area
	f := hist

	// use some mod math to rotate array...
	for iter, idx := 0, 0; iter < (days - 1); iter++ {

		// "shift" everthing over aka substract
		idx = (idx + 1) % 9
		// 0s spawn new 8s (offset 9), overwrite is okay since there are no 8s at this point
		f[(idx+9)%9] = f[idx]
		// 0s increment existing count of 6s (offset 7)
		f[(idx+7)%9] += f[idx]
	}

	return aoc.Sum9(f)
}

func main() {
	aoc.ForLine("input.txt", func(line string) {
		e := aoc.StringToInts(line, ",")

		// There are _only_ 9 possible types of fish. Histogram
		// Use slice instead of fixed length for ease / g
		f := make([]int64, 9)
		var f9 [9]int64
		// Convert input to hist
		for _, v := range e {
			f[v]++
			f9[v]++
		}

		// Why do this twice? To play with generics in-relation to fixed length arrays / types

		fmt.Println(simulate(f, 80))
		fmt.Println(simulate(f, 256))

		fmt.Println(simulate9(f9, 80))
		fmt.Println(simulate9(f9, 256))
	})
}
