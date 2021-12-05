package main

import (
	"aoc2021/utils"
	"fmt"
	"strconv"
	"strings"
)

type BingoSquare struct {
	Value  int
	Marked bool
}

type BingoBoardStatus struct {
	boardNum      int
	won           bool
	winningNumber int
}

type BingoBoard [][]BingoSquare

const numRows = 5
const numCols = 5

func makeBingoBoards(lines []string) []BingoBoard {
	bingoBoards := make([]BingoBoard, 0)
	currentBoard := make(BingoBoard, 0)
	for _, line := range lines {
		if line == "" {
			bingoBoards = append(bingoBoards, currentBoard)
			currentBoard = make(BingoBoard, 0)
		} else {
			row := strings.Fields(line)
			currentBoard = append(currentBoard, utils.Map(row, func(in string) BingoSquare {
				value, _ := strconv.Atoi(in)
				return BingoSquare{value, false}
			}))
		}
	}

	bingoBoards = append(bingoBoards, currentBoard)

	return bingoBoards
}

func (bingoBoard *BingoBoard) mark(value int) bool {
	// for i, _ := range *bingoBoard {
	// 	for j, _ := range row {
	// 		if square.Value == value {
	// 			square.Marked = true
	// 			return true
	// 		}
	// 	}
	// }

	// use old-school loops instead of range because we don't want to avoid dereferencing the pointer on a value basis (we need to mutate!)
	for i := 0; i < len(*bingoBoard); i++ {
		for j := 0; j < len((*bingoBoard)[i]); j++ {
			if (*bingoBoard)[i][j].Value == value {
				(*bingoBoard)[i][j].Marked = true
				return true
			}
		}
	}
	return false
}

func (bingoBoard *BingoBoard) calculateScore(winningNumber int) int {
	// first, we sum the unmarked points
	unmarkedRowScores := utils.Map(*bingoBoard, func(b []BingoSquare) int {
		unmarkedSquares := utils.Filter(b, func(s BingoSquare) bool {
			return !s.Marked
		})
		var sum int
		for _, square := range unmarkedSquares {
			sum += square.Value
		}
		return sum
	})

	var unmarkedScore int
	for _, row := range unmarkedRowScores {
		unmarkedScore += row
	}

	return unmarkedScore * winningNumber
}

func (bingoBoard *BingoBoard) checkWin() bool {
	// first, check each row for a win
	for _, row := range *bingoBoard {
		if len(utils.Filter(row, func(s BingoSquare) bool { return s.Marked })) == numCols {
			return true
		}
	}

	// next, check if there's a win along a column
	for i := 0; i < numCols; i++ {
		column := utils.Map(*bingoBoard, func(row []BingoSquare) BingoSquare {
			return row[i]
		})

		if len(utils.Filter(column, func(s BingoSquare) bool { return s.Marked })) == numCols {
			return true
		}
	}

	// not a winner
	return false
}

func part1(lines []string) int {
	bingoInput := utils.Map(strings.Split(lines[0], ","), func(in string) int {
		out, _ := strconv.Atoi(in)
		return out
	})

	bingoBoards := makeBingoBoards(lines[2:])

	for _, nextNumber := range bingoInput {
		// really wanted to try goroutines, since we need to find the first one that completes
		// assuming there won't ever be a race condition here
		turnDoneChan := make(chan BingoBoardStatus)
		for i := 0; i < len(bingoBoards); i++ {
			board := bingoBoards[i]
			// spawn a goroutine to mark as needed, check for a win, and then the status back over the channel
			// mutates the same state... I like living on the edge
			go func(boardNum int, b BingoBoard) {
				markAdded := b.mark(nextNumber)
				var boardWon bool
				if markAdded {
					boardWon = b.checkWin()
				}

				var status BingoBoardStatus
				if boardWon {
					status = BingoBoardStatus{boardNum, boardWon, nextNumber}
				} else {
					status = BingoBoardStatus{boardNum, boardWon, 0}
				}

				turnDoneChan <- status
			}(i, board)
		}

		for i := 0; i < len(bingoBoards); i++ {
			// receive from channel
			boardStatus := <-turnDoneChan
			// if a board won, calculate the score and bail
			if boardStatus.won {
				return bingoBoards[boardStatus.boardNum].calculateScore(boardStatus.winningNumber)
			}
		}
	}
	// no boards won... oops
	return 0
}

func part2(lines []string) int {
	// create a map so that we can keep track of which boards have already won
	// would prefer a set, but go stdlib doesn't have those...
	wonGames := make(map[int]bool, 0)

	bingoInput := utils.Map(strings.Split(lines[0], ","), func(in string) int {
		out, _ := strconv.Atoi(in)
		return out
	})

	bingoBoards := makeBingoBoards(lines[2:])

	for _, nextNumber := range bingoInput {
		turnDoneChan := make(chan BingoBoardStatus)
		for i := 0; i < len(bingoBoards); i++ {
			board := bingoBoards[i]
			go func(boardNum int, b BingoBoard) {
				markAdded := b.mark(nextNumber)
				var boardWon bool
				if markAdded {
					boardWon = b.checkWin()
				}

				var status BingoBoardStatus
				if boardWon {
					status = BingoBoardStatus{boardNum, boardWon, nextNumber}
				} else {
					status = BingoBoardStatus{boardNum, boardWon, 0}
				}

				turnDoneChan <- status
			}(i, board)
		}

		for i := 0; i < len(bingoBoards); i++ {
			// receive from channel
			boardStatus := <-turnDoneChan
			if boardStatus.won {
				// add the won game to the set
				wonGames[boardStatus.boardNum] = true
				// if all games have now won, take the board that just won and get its score!
				if len(wonGames) == len(bingoBoards) {
					return bingoBoards[boardStatus.boardNum].calculateScore(boardStatus.winningNumber)
				}
			}
		}
	}
	// probably should have added a case where we get through all the numbers and not all boards won...
	return 0
}

func main() {
	// lines := utils.ReadFile("test-input.txt")
	lines := utils.ReadFile("input.txt")
	part1Result := part1(lines)
	fmt.Println(part1Result)
	part2Result := part2(lines)
	fmt.Println(part2Result)
}
