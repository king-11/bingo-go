package data

// Holder stores information about a particular
// place on the board
type Holder struct {
	Val    int
	Marked bool
}

type Board [5][5]Holder

// Returns a new 2D board
func NewBoard() *Board {
	return &Board{}
}

// FindNumber finds the given number in the board and returns
// the row and column of the number alongwith true if value was found
// else returns (-1, -1, false)
func (b *Board) FindNumber(num int) (int, int, bool) {

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b[i][j].Val == num {
				return i, j, true
			}
		}
	}

	return -1, -1, false
}

// Mark the holder
func (h *Holder) Mark() {
	h.Marked = true
}

// MarkValue marks the given value in the board if found
func (b *Board) MarkValue(num int) {
	i, j, found := b.FindNumber(num)
	if found {
		b[i][j].Mark()
	}
}

func (b *Board) isBingoCol(col int) bool {

	for row := 0; row < 5; row++ {
		if !b[row][col].Marked {
			return false
		}
	}

	return true
}

func (b *Board) isBingoRow(row int) bool {

	for col := 0; col < 5; col++ {
		if !b[row][col].Marked {
			return false
		}
	}

	return true
}

// CheckBingo checks if the board is bingo or not
func (b *Board) CheckBingo() bool {

	for i := 0; i < 5; i++ {
		if b.isBingoRow(i) {
			return true
		} else if b.isBingoCol(i) {
			return true
		}
	}

	return false
}

// AddUnmarked adds the unmarked values in the board
func (b *Board) AddUnmarked() int {

	sum := 0

	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if !b[row][col].Marked {
				sum += b[row][col].Val
			}
		}
	}

	return sum

}
