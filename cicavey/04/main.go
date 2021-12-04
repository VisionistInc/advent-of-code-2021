package main

import (
	"aoc2021/aoc"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type Point struct {
	x, y int64
}

type Entry struct {
	value  int64
	marked bool
}

type Board struct {
	grid   map[Point]*Entry
	width  int64
	height int64
}

func NewBoard(width, height int64) *Board {
	return &Board{make(map[Point]*Entry), width, height}
}

func (b *Board) Mark(call int64) {
	for _, v := range b.grid {
		if v.value == call {
			v.marked = true
		}
	}
}

func (b Board) UnmarkedSum() int64 {
	var sum int64
	for _, v := range b.grid {
		if !v.marked {
			sum += v.value
		}
	}
	return sum
}

func (b Board) HasBingo() bool {
	// Check all rows
	for y := int64(0); y < b.height; y++ {
		bingo := true
		for x := int64(0); x < b.width; x++ {
			e := b.grid[Point{x, y}]
			bingo = bingo && e.marked
		}
		if bingo {
			return true
		}
	}

	for x := int64(0); x < b.width; x++ {
		bingo := true
		for y := int64(0); y < b.height; y++ {
			e := b.grid[Point{x, y}]
			bingo = bingo && e.marked
		}
		if bingo {
			return true
		}
	}

	return false
}

func (b Board) String() string {

	var out strings.Builder

	c := color.New(color.FgHiRed)

	for y := int64(0); y < b.height; y++ {
		for x := int64(0); x < b.width; x++ {
			e := b.grid[Point{x, y}]

			if e.marked {
				c.Fprintf(&out, "%2d ", e.value)
			} else {
				fmt.Fprintf(&out, "%2d ", e.value)
			}
		}
		out.WriteString("\n")
	}

	return out.String()
}

func main() {
	input := aoc.ReadLines("input.txt")
	input = aoc.Filter(input, aoc.IsBlank)

	// input sequence
	seq := aoc.StringToInts(input[0], ",")

	// boards
	var boards []*Board
	for i := 0; i < (len(input)-1)/5; i++ {

		b := NewBoard(5, 5)

		for y := 0; y < 5; y++ {
			line := input[1+(i*5)+y]
			line = strings.TrimSpace(line)
			bl := aoc.StringToInts(line, "\\s+")

			for x, z := range bl {
				p := Point{int64(x), int64(y)}
				b.grid[p] = &Entry{z, false}
			}
		}
		boards = append(boards, b)
	}

	var first *Board

	for _, call := range seq {
		for _, board := range boards {
			board.Mark(call)
		}

		var newBoards []*Board
		for _, board := range boards {
			if board.HasBingo() {
				if first == nil {
					first = board
					fmt.Println(board.UnmarkedSum() * call)
				}
				if len(boards) == 1 {
					fmt.Println(board.UnmarkedSum() * call)
				}
			} else {
				newBoards = append(newBoards, board)
			}
		}
		boards = newBoards

		if len(boards) == 0 {
			break
		}
	}
}
