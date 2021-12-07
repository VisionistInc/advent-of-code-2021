#!/usr/bin/env/python3

from collections import defaultdict

grid = defaultdict(lambda: defaultdict(lambda: 0))


with open("input") as infile:
    lines = [
        [tuple([int(x) for x in datapoint.split(",")]) for datapoint in line.strip().split(" -> ")]
        for line in infile
    ]


for line in lines:
    # if it's a straight line:

    # if x == x - vertical line
    if line[0][0] == line[1][0]:
        x = line[0][0]
        y_start = min(line[0][1], line[1][1])
        y_end = max(line[0][1], line[1][1]) + 1
        for y in range(y_start, y_end):
            grid[x][y] += 1

    elif line[0][1] == line[1][1]:
        y = line[0][1]
        x_start = min(line[0][0], line[1][0])
        x_end = max(line[0][0], line[1][0]) + 1
        for x in range(x_start, x_end):
            grid[x][y] += 1

points = 0

for x in grid.values():
    for y in x.values():
        if y >= 2:
            points += 1

print(points)
