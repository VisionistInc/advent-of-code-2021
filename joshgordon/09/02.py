#!/usr/bin/env python3

from typing import List, Tuple, Set
from operator import mul
from functools import reduce

grid: List[List[int]] = []
minimums: List[Tuple[int, int]] = []
basins: List[int] = []

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
            minimums.append((row, col))

# let's find the basins.
for basin in minimums:
    explored: Set[Tuple[int, int]] = set()
    to_explore: Set[Tuple[int, int]] = set()
    to_explore.add(basin)

    while to_explore:
        row, col = to_explore.pop()

        for x, y in [
            (-1, 0),
            (0, -1),
            (0, 1),
            (1, 0),
        ]:
            if (row + x, col + y) in explored:
                continue
            if row + x < 0 or row + x >= width:
                continue
            if col + y < 0 or col + y >= height:
                continue
            if grid[row + x][col + y] == 9:
                continue

            to_explore.add((row + x, col + y))
            explored.add((row + x, col + y))

    basins.append(len(explored))

print(reduce(mul, list(reversed(sorted(basins)))[0:3], 1))
