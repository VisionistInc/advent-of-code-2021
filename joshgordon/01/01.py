#!/usr/bin/env python3

incr = 0
with open("input.txt") as infile:
    inlist = [int(line.strip()) for line in infile]
inlist2 = inlist[1:] + [0]
for a, b in zip(inlist, inlist2):
    if b > a:
        incr += 1
print(incr)
