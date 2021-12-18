#!/usr/bin/env python3

from collections import defaultdict

template = list()
rules = {}
with open("input") as infile:
    for n, line in enumerate(infile):
        line = line.strip()
        if n == 0:
            template = list(line)
        elif n >= 2:
            pair, repl = line.split(" -> ")
            rules.update({pair: repl})


iterations = 0
while iterations < 10:
    # do the thing

    n = 0
    while n < len(template) - 1:
        pair = template[n] + template[n + 1]
        if pair in rules:
            template.insert(n + 1, rules[pair])
            n += 1
        n += 1
    iterations += 1


freqs = defaultdict(lambda: 0)
for item in template:
    freqs[item] += 1

sorted_frequs = sorted(freqs.items(), key=lambda x: x[1])
maxval = sorted_frequs[-1][1]
minval = sorted_frequs[0][1]


print(maxval - minval)
