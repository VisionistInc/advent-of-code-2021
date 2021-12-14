package main

import (
	"aoc2021/utils"
	"fmt"
	"strconv"
	"strings"
)

type direction string

const (
	x direction = "x"
	y direction = "y"
)

type fold struct {
	direction direction
	value     int
}

func foldPoints(points utils.Set[utils.Point[int]], f fold) utils.Set[utils.Point[int]] {
	newPoints := make(utils.Set[utils.Point[int]])
	mag := f.value * 2
	for point := range points {
		switch f.direction {
		case x:
			if point.X < f.value {
				newPoints.Add(point)
			} else {
				newPoints.Add(utils.Point[int]{X: mag - point.X, Y: point.Y})
			}
		case y:
			if point.Y < f.value {
				newPoints.Add(point)
			} else {
				newPoints.Add(utils.Point[int]{X: point.X, Y: mag - point.Y})
			}
		}
	}
	return newPoints
}

func main() {
	lines := utils.ReadFile("input.txt")
	points := make(utils.Set[utils.Point[int]], 0)
	buildingPoints := true
	prefix := "fold along "

	folds := make([]fold, 0)

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			buildingPoints = false
		} else if buildingPoints {
			split := strings.Split(line, ",")
			xStr := split[0]
			yStr := split[1]
			x, _ := strconv.Atoi(xStr)
			y, _ := strconv.Atoi(yStr)
			points.Add(utils.Point[int]{x, y})
		} else {
			f := strings.Split(line[len(prefix):], "=")
			value, _ := strconv.Atoi(f[1])
			folds = append(folds, fold{direction(f[0]), value})
		}
	}

	fmt.Println(len(foldPoints(points, folds[0])))

	for _, fold := range folds {
		points = foldPoints(points, fold)
	}

	// get the max of x and y to define grid bounds
	maxX := 0
	maxY := 0

	for point := range points {
		if maxX < point.X {
			maxX = point.X
		}
		if maxY < point.Y {
			maxY = point.Y
		}
	}

	display := make([][]string, maxY+1)
	for i := range display {
		display[i] = make([]string, maxX+1)
		for j := 0; j < maxX; j++ {
			display[i][j] = "."
		}
	}

	for point := range points {
		display[point.Y][point.X] = "#"
	}

	for _, line := range display {
		fmt.Println(line)
	}
}
