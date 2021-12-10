package main

import (
	"aoc2021/utils"
	"fmt"
	"sort"
)

func getSyntaxErrorScoreFor(r rune) int {
	switch r {
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

func getMatchFor(r rune) rune {
	switch r {
	case ')':
		return '('
	case ']':
		return '['
	case '}':
		return '{'
	case '>':
		return '<'
	case '(':
		return ')'
	case '[':
		return ']'
	case '{':
		return '}'
	case '<':
		return '>'
	default:
		return 0
	}
}

func part1(lines []string) {
	syntaxErrorScore := 0

	// yay for copy-pasted code!
	for _, line := range lines {
		tokenStack := make(utils.Stack[rune], 0)
		for _, char := range line {
			syntaxError := false
			switch char {
			case '[':
				tokenStack.Push('[')
			case '(':
				tokenStack.Push('(')
			case '{':
				tokenStack.Push('{')
			case '<':
				tokenStack.Push('<')
			case ']':
				expected := tokenStack.Peek()
				if expected != getMatchFor(']') {
					// error!
					syntaxErrorScore += getSyntaxErrorScoreFor(']')
					syntaxError = true
				} else {
					tokenStack.Pop()
				}
			case ')':
				expected := tokenStack.Peek()
				if expected != getMatchFor(')') {
					// error!
					syntaxErrorScore += getSyntaxErrorScoreFor(')')
					syntaxError = true
				} else {
					tokenStack.Pop()
				}

			case '}':
				expected := tokenStack.Peek()
				if expected != getMatchFor('}') {
					// error!
					syntaxErrorScore += getSyntaxErrorScoreFor('}')
					syntaxError = true
				} else {
					tokenStack.Pop()
				}

			case '>':
				expected := tokenStack.Peek()
				if expected != getMatchFor('>') {
					// error!
					syntaxErrorScore += getSyntaxErrorScoreFor('>')
					syntaxError = true
				} else {
					tokenStack.Pop()
				}

			}
			// if there was a syntax error, break the loop
			if syntaxError {
				break
			}
		}
	}
	fmt.Println(syntaxErrorScore)
}

func getAutoCompleteScoreFor(r rune) int {
	switch r {
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

func part2(lines []string) {

	autocompleteScores := make([]int64, 0)
	// yay for copy-pasted code!
	for _, line := range lines {
		tokenStack := make(utils.Stack[rune], 0)
		syntaxError := false
		for _, char := range line {
			switch char {
			case '[':
				tokenStack.Push('[')
			case '(':
				tokenStack.Push('(')
			case '{':
				tokenStack.Push('{')
			case '<':
				tokenStack.Push('<')
			case ']':
				expected := tokenStack.Peek()
				if expected != getMatchFor(']') {
					// error!
					syntaxError = true
				} else {
					tokenStack.Pop()
				}
			case ')':
				expected := tokenStack.Peek()
				if expected != getMatchFor(')') {
					// error!
					syntaxError = true
				} else {
					tokenStack.Pop()
				}

			case '}':
				expected := tokenStack.Peek()
				if expected != getMatchFor('}') {
					// error!
					syntaxError = true
				} else {
					tokenStack.Pop()
				}

			case '>':
				expected := tokenStack.Peek()
				if expected != getMatchFor('>') {
					// error!
					syntaxError = true
				} else {
					tokenStack.Pop()
				}

			}
			// if there was a syntax error, break the loop
			if syntaxError {
				break
			}
		}
		if !syntaxError {
			autocompleteScore := int64(0)

			for !tokenStack.IsEmpty() {
				r := tokenStack.Pop()
				autocompleteScore = autocompleteScore*5 + int64(getAutoCompleteScoreFor(getMatchFor(r)))
			}
			autocompleteScores = append(autocompleteScores, autocompleteScore)
		}
	}
	sort.Slice(autocompleteScores, func(a, b int) bool { return autocompleteScores[a] < autocompleteScores[b] })
	fmt.Println(autocompleteScores[len(autocompleteScores)/2])
}

func main() {
	lines := utils.ReadFile("input.txt")
	part1(lines)
	part2(lines)
}
