#!/usr/bin/env python3

from collections import defaultdict
from typing import DefaultDict, List

with open("input") as infile:
    data = [int(x) for x in infile.read().strip().split(",")]

# cost -> list of positions of that cost.
costs: DefaultDict[int, List[int]] = defaultdict(lambda: list())

for i in range(min(data), max(data)):
    # calculate the cost:
    cost = sum([abs(i - x) for x in data])
    costs[cost].append(i)

print(min(costs))
print(costs[min(costs)])
