package data

type Board [5][5]Holder

// Returns a new 2D board
func NewBoard() *Board {
	newBoard := Board{}
	return &newBoard
}

// FindNumber finds the given number in the board and returns
// the row and column of the number alongwith true if value was found
// else returns (-1, -1, false)
func (b *Board) FindNumber(num int) (int, int, bool) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b[i][j].Val == num {
				return (i), (j), true
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
	for i := 0; i < 5; i++ {
		a0 := b[i][0].Marked
		a1 := b[i][1].Marked
		a2 := b[i][2].Marked
		a3 := b[i][3].Marked
		a4 := b[i][4].Marked
		if a0 && a1 && a2 && a3 && a4 {
			return true
		}
	}

	for i := 0; i < 5; i++ {
		a0 := b[0][i].Marked
		a1 := b[1][i].Marked
		a2 := b[2][i].Marked
		a3 := b[3][i].Marked
		a4 := b[4][i].Marked
		if a0 && a1 && a2 && a3 && a4 {
			return true
		}
	}
	return false
}

// AddUnmarked adds the unmarked values in the board
func (b *Board) AddUnmarked() int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b[i][j].Marked {
				sum += b[i][j].Val
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
