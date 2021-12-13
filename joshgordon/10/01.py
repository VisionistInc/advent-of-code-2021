#!/usr/bin/env python3

matches = {"(": ")", "[": "]", "{": "}", "<": ">"}
scores = {")": 3, "]": 57, "}": 1197, ">": 25137}

score = 0

with open("input") as infile:
    for line in infile:
        tokens = []
        line_valid = True
        for char in line.strip():
            if not line_valid:
                continue
            if char in matches:
                tokens.append(char)
            else:
                match = tokens.pop()
                if char != matches[match]:
                    # we're invalid!
                    score += scores[char]
                    line_valid = False


print(score)
