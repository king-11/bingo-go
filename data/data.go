package data

type Board [5][5]Holder

// Returns a new 2D board
func NewBoard() *Board {
	var b Board
	return &b
}

// FindNumber finds the given number in the board and returns
// the row and column of the number alongwith true if value was found
// else returns (-1, -1, false)
func (b *Board) FindNumber(num int) (int, int, bool) {
	for i, v := range b {
		for j, h := range v {
			if h.Val == num {
				return i, j, true
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
	horizontal := [5]int8{0, 0, 0, 0, 0}
	vertical := [5]int8{0, 0, 0, 0, 0}

	for i, v := range b {
		for j, h := range v {
			if h.Marked {
				horizontal[i] += 1
				vertical[j] += 1
			}
		}
	}

	for _, row := range horizontal {
		if row == 5 {
			return true
		}
	}

	for _, col := range vertical {
		if col == 5 {
			return true
		}
	}

	return false
}

// AddUnmarked adds the unmarked values in the board
func (b *Board) AddUnmarked() int {
	sum := 0
	for _, v := range b {
		for _, h := range v {
			if !h.Marked {
				sum += h.Val
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
