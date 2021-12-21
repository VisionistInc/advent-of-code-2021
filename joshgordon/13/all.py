#!/usr/bin/env python3

from pathlib import Path
from collections import defaultdict
from typing import DefaultDict

from io import StringIO

grid_type = DefaultDict[int, DefaultDict[int, bool]]
# dict of columns
grid: grid_type = defaultdict(lambda: defaultdict(lambda: False))
folds = []


def print_grid(grid: grid_type, nums=False) -> str:
    output = StringIO()
    # find the max_x and max_y:
    max_x = max(grid.keys())
    max_y = max([max(col.keys()) for col in grid.values()])

    if nums:
        print(" ", end="", file=output)
        for x in range(max_x + 1):
            print(x % 10, end="", file=output)
        print(file=output)
    for y in range(max_y + 1):
        print(y % 10, end="", file=output)
        for x in range(max_x + 1):
            if grid[x][y]:
                print("#", end="", file=output)
            else:
                print(".", end="", file=output)
        print(file=output)

    return output.getvalue()


def get_dots(grid: grid_type) -> int:

    max_x = max(grid.keys()) + 1
    max_y = max([max(col.keys()) for col in grid.values()]) + 1

    count = 0
    for y in range(max_y):
        for x in range(max_x):
            if grid[x][y]:
                count += 1

    return count


with open(Path(__file__).parent / "input") as infile:
    for line in infile:
        line = line.strip()

        if "," in line:
            x, y = [int(v) for v in line.split(",")]
            grid[x][y] = True

        if "fold" in line:
            axis, val_s = line.split("=")
            axis = axis[-1]
            val = int(val_s)

            folds.append((axis, val))


def fold(grid: grid_type, axis: str, val: int):
    max_x = max(grid.keys()) + 1
    max_y = max([max(col.keys()) for col in grid.values()]) + 1

    if axis == "x":
        # iterate through columns:
        for n, old_x in enumerate(range(val + 1, max_x)):
            # fold it over
            new_x = val - n - 1

            if new_x < 0:
                try:
                    del grid[old_x]
                except Exception:
                    pass
                continue

            for y in range(max_y):
                if grid[old_x][y]:
                    grid[new_x][y] = True

            try:
                del grid[old_x]
            except Exception:
                pass
        try:
            del grid[val]
        except Exception:
            pass

    elif axis == "y":
        for n, old_y in enumerate(range(val + 1, max_y)):
            new_y = val - n - 1
            if new_y < 0:
                for col in grid.values():
                    try:
                        del col[old_y]
                    except Exception:
                        pass
                continue

            for x in range(max_x):
                if grid[x][old_y]:
                    grid[x][new_y] = True

            for col in grid.values():
                try:
                    del col[old_y]
                except Exception:
                    pass

        for col in grid.values():
            try:
                del col[val]
            except Exception:
                pass

        pass

    else:
        raise Exception("Invalid parameters")


fold(grid, *folds[0])
print(get_dots(grid))

for f in folds[1:]:
    fold(grid, *f)
print(print_grid(grid, nums=True))
