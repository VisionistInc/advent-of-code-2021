#!/opt/homebrew/opt/python@3.10/bin/python

x = 0  # horizontal position
aim = 0
depth = 0


with open("input") as infile:
    for line in infile:
        line = line.strip()
        inst, mag = line.split(" ")
        mag = int(mag)

        match inst:
            case "forward":
                x += mag
                depth += mag * aim
            case "up":
                aim -= mag
            case "down":
                aim += mag

print(x * depth)
