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

	var loserBoardScore int
	loserBoardTurns := -1

	for _, board := range boards {
		isBingoed := false
		for turn, val := range order {
			board.MarkValue(val)
			if !isBingoed && board.CheckBingo() {
				if turn > loserBoardTurns {
					loserBoardTurns = turn
					loserBoardScore = val * board.AddUnmarked()
				}
				isBingoed = true
			}
		}
	}

	return loserBoardScore
}
