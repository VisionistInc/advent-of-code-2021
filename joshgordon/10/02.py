#!/usr/bin/env python3

from statistics import median

matches = {"(": ")", "[": "]", "{": "}", "<": ">"}
scores = {")": 1, "]": 2, "}": 3, ">": 4}

valid_lines_completion_scores = []

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
                    line_valid = False

        if line_valid:
            # at this point we've got an incomplete line, let's complete it.
            line_score = 0
            for token in reversed(tokens):
                completion = matches[token]
                # print(completion, end="")
                line_score *= 5
                line_score += scores[completion]
            # print(f" -  {line_score}")

            valid_lines_completion_scores.append(line_score)


print(median(valid_lines_completion_scores))
