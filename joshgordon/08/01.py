#!/usr/bin/env python3

total = 0
with open("input") as infile:
    for line in infile:
        output = line.split(" | ")[1].strip()
        output = output.split(" ")
        for value in output:
            if len(value) in [2, 3, 4, 7]:
                total += 1


print(total)
