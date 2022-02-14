package data

type Board [5][5]Holder

// Returns a new 2D board
func NewBoard() *Board {
	return &Board{}
}

// FindNumber finds the given number in the board and returns
// the row and column of the number alongwith true if value was found
// else returns (-1, -1, false)
func (b *Board) FindNumber(num int) (int, int, bool) {
	for i_row, row := range b {
		for i_col, v_col := range row {
			if num == v_col.Val {
				return i_row, i_col, true
			}
		}
	}
	return -1, -1, false
}

// MarkValue marks the given value in the board if found
func (b *Board) MarkValue(num int) {
	i, j, found := b.FindNumber(num)
	if found {
		b[i][j].Mark()
	}
}

// CheckBingo checks if the board is bingo or not
func (b *Board) CheckBingo() bool {
	var any_row, any_col bool= false, false
	
	for row := 0; row < 5; row++ {
		// Checking if a row is bingo
		var row_val int
		for col := 0; col < 5; col++ {
			if !b[row][col].Marked {
				break
			} else {
				row_val++
			}
		}

		if row_val == 5 {
			any_row = true
		}
	}
	
	for col := 0; col < 5; col++ {
		// Checking if a column is bingo
		var col_val int
		for row := 0; row < 5; row++ {
			if !b[row][col].Marked {
				break
			} else { 
				col_val++ 
			}
		}

		if col_val == 5 {
			any_col = true
		}
	}

	return any_col || any_row
}

// AddUnmarked adds the unmarked values in the board
func (b *Board) AddUnmarked() int {
	var sum int
	for _, row := range b {
		for _, v_col := range row {
			if !v_col.Marked {
				sum += v_col.Val
			}
		}
	}

	return sum
}

// Holder stores information about a particular
// place on the board
type Holder struct {
	Val    int
	Marked bool
}

// Mark the holder
func (h *Holder) Mark() {
	h.Marked = true
}
