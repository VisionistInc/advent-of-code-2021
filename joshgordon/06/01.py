#!/usr/bin/env python3

with open("input") as infile:
    starting_state = [int(x) for x in infile.read().strip().split(",")]

state = []
for i in range(9):
    state.append(0)

# state is the count of each lanternfix, the index represents the number of days until
# reproduction

for st in starting_state:
    state[st] += 1

for _ in range(80):
    new_fish = state[0]
    state = state[1:]
    state.append(new_fish)
    state[6] += new_fish

print(sum(state))
