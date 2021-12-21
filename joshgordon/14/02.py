#!/usr/bin/env python3

from collections import defaultdict
from copy import deepcopy
from pathlib import Path

template = list()
rules = {}
with open(Path(__file__).parent / "input") as infile:
    for n, line in enumerate(infile):
        line = line.strip()
        if n == 0:
            template = list(line)
        elif n >= 2:
            pair, repl = line.split(" -> ")
            repl = [f"{pair[0]}{repl}", f"{repl}{pair[1]}", repl]
            rules.update({pair: repl})

pairs = defaultdict(lambda: 0)
freqs = defaultdict(lambda: 0)

n = 0
while n < len(template) - 1:
    freqs[template[n]] += 1
    pair = template[n] + template[n + 1]
    pairs[pair] += 1

    n += 1

freqs[template[-1]] += 1

iterations = 0
while iterations < 40:
    # do the thing
    new_pairs = deepcopy(pairs)

    for pair in pairs:
        for rule in rules[pair][:2]:
            new_pairs[rule] += pairs[pair]
        new_pairs[pair] -= pairs[pair]

        freqs[rules[pair][2]] += pairs[pair]

    pairs = new_pairs
    iterations += 1


sorted_frequs = sorted(freqs.items(), key=lambda x: x[1])
maxval = sorted_frequs[-1][1]
minval = sorted_frequs[0][1]

print(maxval - minval)
