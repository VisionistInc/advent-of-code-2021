package internal

import (
	"bufio"
	"os"
	"strconv"
)

func ReadIntLines(file string) []int64 {

	f, _ := os.Open(file)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var input []int64
	for scanner.Scan() {
		value, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		input = append(input, value)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return input
}

func ForLine(fname string, fn func(line string)) {
	f, _ := os.Open(fname)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fn(scanner.Text())
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}
