package play

import (
	"github.com/king-11/bingo/data"
)

// FindWinner finds the winner board in bingo and
// returns the iteration value multiplied by addition of
// unmarked values in the board.
func FindWinner(order []int, boards []*data.Board) int {
	for _, val := range order {
		for _, b := range boards {
			b.MarkValue(val)
			bingo := b.CheckBingo()
			if bingo {
				return val * b.AddUnmarked()
			}
		}
	}
	return 0
}

// FindLastBoard finds the last board to bingo and
// returns the iteration value multiplied by addition of
// unmarked values in the board.
func FindLastBoard(order []int, boards []*data.Board) int {
	reqboards := len(boards)
	var mark_boards []bool
	num := 0
	for num < reqboards {
		mark_boards = append(mark_boards, false)
		num += 1
	}

	count := 0
	ind := 0
	for {
		val := order[ind]
		for i, b := range boards {
			b.MarkValue(val)
			bingo := b.CheckBingo()
			if bingo && !mark_boards[i] {
				mark_boards[i] = true
				count = count + 1
				if count == reqboards {
					return val * b.AddUnmarked()
				}
			}
		}
		ind = ind + 1
	}
}
