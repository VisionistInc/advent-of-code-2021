package main

import (
	"aoc2021/aoc"
	"fmt"
	"strconv"
	"strings"
)

type cmd struct {
	name  string
	value int64
}

func main() {

	var cmds []cmd

	aoc.ForLine("input.txt", func(line string) {
		temp_cmd := strings.Fields(line)
		v, _ := strconv.ParseInt(temp_cmd[1], 10, 64)
		cmds = append(cmds, cmd{temp_cmd[0], v})
	})

	pos := int64(0)
	depth := int64(0)

	for _, c := range cmds {
		switch c.name {
		case "up":
			depth -= c.value
		case "down":
			depth += c.value
		case "forward":
			pos += c.value
		}
	}

	fmt.Println(pos * depth)

	pos = 0
	depth = 0
	aim := int64(0)

	for _, c := range cmds {
		switch c.name {
		case "up":
			aim -= c.value
		case "down":
			aim += c.value
		case "forward":
			pos += c.value
			depth += (aim * c.value)
		}
	}

	fmt.Println(pos * depth)
}
