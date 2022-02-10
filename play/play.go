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
	board_status := make(map[int]bool)
	winning_boards := 0

	for _, val := range order {
		for i, b := range boards {
			_, ok := board_status[i]
			if !ok {
				b.MarkValue(val)
				bingo := b.CheckBingo()
				if bingo {
					winning_boards++
					board_status[i] = true
				}
				if winning_boards == n_boards {
					return val * b.AddUnmarked()
				}
			}
		}
	}
	return 0
}
