#!/usr/bin/env python3

prev = None
incr = 0
with open("input.txt") as infile:
    inlist = [int(line.strip()) for line in infile]
inlist2 = inlist[1:] + [0]
inlist3 = inlist[2:] + [0, 0]
sums = [sum(it) for it in zip(inlist, inlist2, inlist3)][:-2]

sum1 = sums[1:] + [0]
for a, b in zip(sums, sum1):
    if b > a:
        incr += 1
print(incr)
