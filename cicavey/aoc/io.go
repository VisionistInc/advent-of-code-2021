package aoc

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Filter[T any](input []T, pred func(s T) bool) []T {
	var output []T
	for _, v := range input {
		if pred(v) {
			output = append(output, v)
		}
	}
	return output
}

func IsBlank(s string) bool {
	return len(strings.TrimSpace(s)) > 0
}

func StringToInts(s string, sepRE string) []int64 {
	re := regexp.MustCompile(sepRE)
	var output []int64
	for _, v := range re.Split(s, -1) {
		iv, _ := strconv.ParseInt(v, 10, 64)
		output = append(output, iv)
	}
	return output
}

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

func ReadLines(file string) []string {

	f, _ := os.Open(file)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
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

func AbsInt64(n int64) int64 {
	y := n >> 63       // y ← x ⟫ 63
	return (n ^ y) - y // (x ⨁ y) - y
}
