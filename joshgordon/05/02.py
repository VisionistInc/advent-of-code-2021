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

    # else it's diagonal
    else:
        # ugh
        x_start = min(line[0][0], line[1][0])
        x_end = max(line[0][0], line[1][0])
        y_start = min(line[0][1], line[1][1])
        y_end = max(line[0][1], line[1][1])

        diff_x = x_end - x_start + 1
        diff_y = y_end - y_start + 1

        x_bigger = line[0][0] > line[1][0]
        y_bigger = line[0][1] > line[1][1]

        backwards_diagonal = x_bigger ^ y_bigger

        for i in range(diff_x):
            x = x_start + i
            if backwards_diagonal:
                y = y_end - i
            else:
                y = y_start + i
            grid[x][y] += 1
            if x == 0 and y == 0:
                print("XY=0")


points = 0
# max_y = max(grid.keys()) + 1
# max_x = max([max(col.keys()) for col in grid.values()]) + 1

# print(max_x)
# print(max_y)

# for x in range(max_x):
#     for y in range(max_y):
#         if grid[y][x] == 0:
#             print(".", end="")
#         else:
#             print(f"{grid[y][x]:1d}", end="")
#     print()


for x in grid.values():
    for y in x.values():
        if y >= 2:
            points += 1

print(points)
