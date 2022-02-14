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
	cnt := 0
	isBingo := make([]bool, len(boards))
	result := 0

	for _, val := range order {
		cnt = 0
		for idx, b := range boards {
			b.MarkValue(val)
			bingo := b.CheckBingo()
			if bingo {
				if !isBingo[idx] {
					result = val * b.AddUnmarked()
					isBingo[idx] = true
				}
				cnt++

				if cnt == (len(boards)) {
					return result
				}
			}
		}
	}

	return 0
}
