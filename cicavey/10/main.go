package main

import (
	"aoc2021/aoc"
	"fmt"
	"sort"
)

func isOpen(r rune) bool {
	return r == '(' || r == '{' || r == '[' || r == '<'
}

func isClose(r rune) bool {
	return r == ')' || r == '}' || r == ']' || r == '>'
}

func isPair(open, close rune) bool {
	return (open == '(' && close == ')') || (open == '[' && close == ']') || (open == '{' && close == '}') || (open == '<' && close == '>')
}

func pair(token rune) rune {
	switch token {
	case '(':
		return ')'
	case ')':
		return '('
	case '[':
		return ']'
	case ']':
		return '['
	case '{':
		return '}'
	case '}':
		return '{'
	case '<':
		return '>'
	case '>':
		return '<'
	}
	return ' '
}

func checkerScore(token rune) int {
	switch token {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	default:
		return 0
	}
}

func autocompleteScore(token rune) int {
	switch token {
	case ')':
		return 1
	case ']':
		return 2
	case '}':
		return 3
	case '>':
		return 4
	default:
		return 0
	}
}

func main() {
	total := 0

	var aTotals []int

	aoc.ForLine("input.txt", func(line string) {

		var stack []rune

		for _, c := range line {

			if isOpen(c) {
				stack = append(stack, c)
				continue
			}

			// Pop
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if !isPair(last, c) {
				// invalid
				// fmt.Printf("Expected %c, but found %c instead", pair(last), c)
				total += checkerScore(c)
				stack = nil
				break
			}
		}

		if len(stack) > 0 {
			// fmt.Println("Incomplete", len(stack))
			// We should be able to just unwind the stack finding the matching tokens
			aTotal := 0
			for i := len(stack) - 1; i >= 0; i-- {
				t := pair(stack[i])
				aTotal = (aTotal * 5) + autocompleteScore(t)
				// fmt.Print(string(pair(stack[i])))
			}
			aTotals = append(aTotals, aTotal)
		}

	})

	fmt.Println(total)

	sort.Ints(aTotals)
	fmt.Println(aTotals[len(aTotals)/2])
}
