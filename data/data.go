package data

type Board [5][5]Holder

// Returns a new 2D board
func NewBoard() *Board {
	board := Board{}
	return &board
}

// FindNumber finds the given number in the board and returns
// the row and column of the number alongwith true if value was found
// else returns (-1, -1, false)
func (b *Board) FindNumber(num int) (int, int, bool) {
	i := 0
	for i < 5 {
		j := 0
		for j < 5 {
			if b[i][j].Val == num {
				return i, j, true
			}
			j = j + 1
		}
		i = i + 1
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

// Let's check for a row 
func (b *Board) CheckBingoRow(row int) bool {
	col := 0
	for col < 5 {
		if !b[row][col].Marked {
			return false
		}
		col = col + 1
	}
	return true
}

// Let's check for a column
func (b *Board) CheckBingoCol(col int) bool {
	row := 0
	for row < 5 {
		if !b[row][col].Marked {
			return false
		}
		row = row + 1
	}
	return true
}

// CheckBingo checks if the board is bingo or not
func (b *Board) CheckBingo() bool {
	i := 0
	for i < 5 {
		if b.CheckBingoRow(i) || b.CheckBingoCol(i) {
			return true
		}
		i = i + 1
	}
	return false
}

// AddUnmarked adds the unmarked values in the board
func (b *Board) AddUnmarked() int {
	sum := 0
	i := 0
	for i < 5 {
		j := 0
		for j < 5 {
			if !b[i][j].Marked {
				sum = sum + b[i][j].Val
			}
			j = j + 1
		}
		i = i + 1
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
