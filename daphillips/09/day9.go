package main

import (
	"aoc2021/utils"
	"fmt"
	"sort"
)

type Point struct {
	row int
	col int
}

// recursively check the current point and adjacent neighbors, summing the larger points (up until 9)
func traverse(x, y int, grid [][]int, visitedPoints utils.Set[Point]) int {
	if (visitedPoints.Contains(Point{x, y})) {
		// duplicate! return -1 to offset the +1 that the recursive call adds
		return -1
	}

	// add this point to the visited set
	visitedPoints.Add(Point{x, y})
	sum := 0

	if x != 0 && grid[x-1][y] != 9 && grid[x][y] < grid[x-1][y] {
		sum += traverse(x-1, y, grid, visitedPoints) + 1
	}

	if y != 0 && grid[x][y-1] != 9 && grid[x][y] < grid[x][y-1] {
		sum += traverse(x, y-1, grid, visitedPoints) + 1
	}

	if x != len(grid)-1 && grid[x+1][y] != 9 && grid[x][y] < grid[x+1][y] {
		sum += traverse(x+1, y, grid, visitedPoints) + 1
	}

	if y != len(grid[x])-1 && grid[x][y+1] != 9 && grid[x][y] < grid[x][y+1] {
		sum += traverse(x, y+1, grid, visitedPoints) + 1
	}
	return sum
}

func main() {
	lines := utils.ReadFile("input.txt")

	grid := make([][]int, 0)

	for _, line := range lines {
		inner := make([]int, 0)
		for _, num := range line {
			inner = append(inner, int(num-'0'))
		}
		grid = append(grid, inner)
	}

	// part 1
	lowRiskSum := 0
	for i := range grid {
		for j, num := range grid[i] {
			// each level of the if statement checks if there is a valid point next to the current one and then if we have a larger value (if it's smaller, than this point can't be a lowest point!)
			if i == 0 || grid[i-1][j] > num {
				if j == 0 || grid[i][j-1] > num {
					if i == len(grid)-1 || grid[i+1][j] > num {
						if j == len(grid[i])-1 || grid[i][j+1] > num {
							lowRiskSum += num + 1
						}
					}
				}
			}
		}
	}
	fmt.Println(lowRiskSum)

	// part2
	lowestPoints := make([]Point, 0)
	for i := range grid {
		for j, num := range grid[i] {
			// each level of the if statement checks if there is a valid point next to the current one and then if we have a larger value (if it's smaller, than this point can't be a lowest point!)
			if i == 0 || grid[i-1][j] > num {
				if j == 0 || grid[i][j-1] > num {
					if i == len(grid)-1 || grid[i+1][j] > num {
						if j == len(grid[i])-1 || grid[i][j+1] > num {
							lowestPoints = append(lowestPoints, Point{i, j})
						}
					}
				}
			}
		}
	}

	basinSizes := make([]int, 0)
	for _, point := range lowestPoints {
		fmt.Println(point)
		basinSizes = append(basinSizes, traverse(point.row, point.col, grid, make(utils.Set[Point], 0))+1)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	fmt.Println(basinSizes[0] * basinSizes[1] * basinSizes[2])

}
