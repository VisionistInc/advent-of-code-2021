from collections import defaultdict
import re


class keydefaultdict(defaultdict):
    def __missing__(self, key):
        if self.default_factory is None:
            raise KeyError(key)
        else:
            ret = self[key] = self.default_factory(key)
            return ret


class BoardNumber:
    def __init__(self, num):
        self.called = False
        self.number = int(num)

    def __repr__(self):
        if self.called:
            return f"\033[33m{self.number}\033[0m"
        else:
            return f"\033[41m{self.number}\033[0m"

    def call(self):
        self.called = True


class Board:
    def __init__(self, grid):
        if isinstance(grid, list):
            self.grid == grid
        else:
            self.grid = self._parse_grid(grid)

    def _parse_grid(self, grid):
        return [
            [numbers[int(y.strip())] for y in re.split(r"\s+", x.strip())] for x in grid.strip().split("\n")
        ]

    def __repr__(self):
        res = ""
        for row in self.grid:
            for col in row:
                res += f"{str(col):13s}"
            res += "\n"
        return res

    def check_board(self):
        # iterate through the rows:
        checks = [all([col.called for col in row]) for row in self.grid]
        checks += [all([row.called for row in col]) for col in list(zip(*self.grid))]
        return any(checks)

    def get_unmarked_sum(self):
        return sum([sum([num.number for num in row if not num.called]) for row in self.grid])


numbers = keydefaultdict(lambda x: BoardNumber(x))
