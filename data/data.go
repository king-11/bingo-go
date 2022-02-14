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
	for i := range b {
		for j := range b[i] {
			if b[i][j].Val == num {
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

func (b *Board) checkRowBingo() bool {
	var cnt int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b[i][j].Marked {
				cnt++
			}
		}
		if cnt == 5 {
			return true
		} else {
			cnt = 0
		}
	}
	return false
}

func (b *Board) checkColBingo() bool {
	var cnt int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b[j][i].Marked {
				cnt++
			}
		}
		if cnt == 5 {
			return true
		} else {
			cnt = 0
		}
	}
	return false
}

// CheckBingo checks if the board is bingo or not
func (b *Board) CheckBingo() bool {
	if b.checkColBingo() || b.checkRowBingo() {
		return true
	}

	return false
}

// AddUnmarked adds the unmarked values in the board
func (b *Board) AddUnmarked() int {
	unmarkedCount := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b[i][j].Marked {
				unmarkedCount += b[i][j].Val
			}
		}
	}
	return unmarkedCount
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
