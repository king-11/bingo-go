package data

type Board [5][5]Holder

// Returns a new 2D board
func NewBoard() *Board {
	return nil
}

// FindNumber finds the given number in the board and returns
// the row and column of the number alongwith true if value was found
// else returns (-1, -1, false)
func (b *Board) FindNumber(num int) (int, int, bool) {
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
	return false
}

// AddUnmarked adds the unmarked values in the board
func (b *Board) AddUnmarked() int {
	return 0
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
