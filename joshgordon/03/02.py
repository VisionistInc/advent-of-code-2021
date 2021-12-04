#!/opt/homebrew/opt/python@3.10/bin/python

indata = []

with open("input") as infile:
    for line in infile:
        line = line.strip()
        indata.append(line)

o2_gen_rating = 1

vals = indata.copy()
for i in range(len(indata[0])):
    zeroes = 0
    ones = 0

    for val in vals:
        if val[i] == "0":
            zeroes += 1
        elif val[i] == "1":
            ones += 1
        else:
            raise Exception("invalid input")

    if zeroes > ones:
        vals = [val for val in vals if val[i] == "0"]
    else:
        vals = [val for val in vals if val[i] == "1"]

    if len(vals) == 1:
        o2_gen_rating = int(vals[0], 2)
        break

# yeah yeah I should make this a function... this is AoC not the mona lisa...
co2_scrubber_rating = 1
vals = indata.copy()
for i in range(len(indata[0])):
    zeroes = 0
    ones = 0

    for val in vals:
        if val[i] == "0":
            zeroes += 1
        elif val[i] == "1":
            ones += 1
        else:
            raise Exception("invalid input")

    if zeroes <= ones:
        vals = [val for val in vals if val[i] == "0"]
    else:
        vals = [val for val in vals if val[i] == "1"]

    if len(vals) == 1:
        co2_scrubber_rating = int(vals[0], 2)
        break

print(f"O2 generator rating: {o2_gen_rating}")
print(f"CO2 scrubber rating: {co2_scrubber_rating}")

print(f"Multiplied: {o2_gen_rating * co2_scrubber_rating}")
