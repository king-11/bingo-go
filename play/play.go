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
	n_boards := len(boards)
	var marked_boards []bool
	i := 0
	for i < n_boards {
		marked_boards = append(marked_boards, false)
		i += 1
	}

	count := 0
	index := 0
	for {
		val := order[index]
		for i, b := range boards {
			b.MarkValue(val)
			bingo := b.CheckBingo()
			if bingo && !marked_boards[i] {
				marked_boards[i] = true
				count = count + 1
				if count == n_boards {
					return val * b.AddUnmarked()
				}
			}
		}
		index = index + 1
	}
}
