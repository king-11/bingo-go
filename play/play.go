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
	count := 0
	c_max := len(boards)
	newMap := make(map[int]bool)

	for i := 0; i < c_max; i++ {
		newMap[i] = false
	}

	for _, val := range order {
		for index, b := range boards {
			b.MarkValue(val)
			bingo := b.CheckBingo()
			if bingo && !newMap[index] {
				newMap[index] = true
				count++
				if count == c_max {
					return val * b.AddUnmarked()
				}
			}
		}
	}
	return 0
}
