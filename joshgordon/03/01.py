#!/opt/homebrew/opt/python@3.10/bin/python

import statistics

indata = []

with open("input") as infile:
    for line in infile:
        line = line.strip()
        indata.append(list(line))


indata_split = list(zip(*indata))

gamma = int("".join([statistics.mode(col) for col in indata_split]), 2)
# just xor the darned thing.
epsilon = gamma ^ (2 ** (len(indata[0])) - 1)

# diagnostics print
print(f"{gamma:012b}")
print(f"{epsilon:012b}")
print(epsilon * gamma)
