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
				return val*b.AddUnmarked()
			}
		}
	}
	return 0
}

// FindLastBoard finds the last board to bingo and
// returns the iteration value multiplied by addition of
// unmarked values in the board.
func FindLastBoard(order []int, boards []*data.Board) int {
	return 0
}
