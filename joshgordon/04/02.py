#!/usr/bin/env python3

from typing import List
import sys
from bingo import Board, numbers


boards: List[Board] = []

with open("input") as infile:
    to_call = [int(x) for x in infile.readline().strip().split(",")]

    infile.readline()

    grid = ""
    for line in infile:
        grid += line
        if line == "\n":
            boards.append(Board(grid))
            grid = ""

    boards.append(Board(grid))

for num in to_call:
    numbers[num].call()
    for board in boards:
        if board.check_board():
            boards.remove(board)
            last_winning_score = board.get_unmarked_sum() * num

print(last_winning_score)
