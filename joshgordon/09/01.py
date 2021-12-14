#!/usr/bin/env python3

from typing import List

grid: List[List[int]] = []
minimums: List[int] = []

with open("input") as infile:
    for line in infile:
        line = line.strip()
        grid.append([int(x) for x in line])

height = len(grid)
width = len(grid[0])

for row in range(height):
    for col in range(width):
        surrounding = []
        for x, y in [(-1, -1), (-1, 0), (-1, 1), (0, -1), (0, 1), (1, -1), (1, 0), (1, 1)]:
            if row + x < 0 or row + x >= width:
                continue
            if col + y < 0 or col + y >= height:
                continue
            surrounding.append(grid[row + x][col + y])

        if grid[row][col] < min(surrounding):
            minimums.append(grid[row][col])

sum_risk = sum([1 + x for x in minimums])
print(sum_risk)
