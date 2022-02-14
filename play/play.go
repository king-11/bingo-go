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
	// What I could do is iterate over the values and boards. Once a board goes bingo, pop it off the 'boards' slice,
	// keep repeating it until slice's length is 1.

	var bingoed = make(map[*data.Board]bool)
	var bingoed_count int
	var board_init_count int = len(boards)

	for _, v := range boards {
		bingoed[v] = false
	}

	for _, val := range order {
		for _, board := range boards {
			board.MarkValue(val)
			bingo := board.CheckBingo()

			if bingo && !bingoed[board] {
				bingoed[board] = true
				bingoed_count++
			}

			if bingoed_count == (board_init_count - 1) {
				for key, _ := range bingoed {
					if !bingoed[key] {
						return val * key.AddUnmarked()
					}
				}
			}

			// For some reason the plan didn't work
			// if bingo {
			// 	if (board_i == 0) {
			// 		boards = boards[1:]
			// 	} else if (board_i < len(boards) - 1) {
			// 		boards = append(boards[ :board_i], boards[board_i+1: ]...)
			// 	} else {
			// 		boards = boards[:board_i]
			// 	}
			// 	continue
			// }
			
		}
	}
	return 0
}
