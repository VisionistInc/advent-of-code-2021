#!/opt/homebrew/opt/python@3.10/bin/python

x = 0  # horizontal position
y = 0  # depth


with open("input") as infile:
    for line in infile:
        line = line.strip()
        inst, mag = line.split(" ")
        mag = int(mag)

        match inst:
            case "forward":
                x += mag
            case "up":
                y -= mag
            case "down":
                y += mag

print(x * y)
