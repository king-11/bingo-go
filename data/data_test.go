package data

import (
	"fmt"
	"testing"
)

func TestNewBoard(t *testing.T) {
	board := NewBoard()
	if board == nil {
		t.Error("NewBoard() should not return nil")
	}
}

func TestFindNumber(t *testing.T) {
	board := &Board{
		{
			Holder{1, false},
			Holder{2, false},
			Holder{3, false},
			Holder{4, false},
			Holder{5, false},
		},
		{
			Holder{11, false},
			Holder{12, false},
			Holder{13, false},
			Holder{14, false},
			Holder{15, false},
		},
		{
			Holder{21, false},
			Holder{22, false},
			Holder{23, false},
			Holder{24, false},
			Holder{25, false},
		},
		{
			Holder{16, false},
			Holder{17, false},
			Holder{18, false},
			Holder{19, false},
			Holder{20, false},
		},
		{
			Holder{6, false},
			Holder{7, false},
			Holder{8, false},
			Holder{9, false},
			Holder{10, false},
		},
	}

	tests := []struct {
		i, j, num int
		found     bool
	}{
		{0, 0, 1, true},
		{4, 4, 10, true},
		{1, 1, 12, true},
		{-1, -1, 42, false},
		{2, 3, 24, true},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("find %d", test.num), func(t *testing.T) {
			i, j, found := board.FindNumber(test.num)
			if i != test.i || j != test.j || found != test.found {
				t.Errorf("FindNumber(%d) = (%d, %d, %t), want (%d, %d, %t)", test.num, i, j, found, test.i, test.j, test.found)
			}
		})
	}
}

func TestAddUnmarked(t *testing.T) {
	board := &Board{
		{
			Holder{1, false},
			Holder{2, false},
			Holder{3, false},
			Holder{4, false},
			Holder{5, false},
		},
		{
			Holder{11, false},
			Holder{12, false},
			Holder{13, false},
			Holder{14, false},
			Holder{15, false},
		},
		{
			Holder{21, false},
			Holder{22, false},
			Holder{23, false},
			Holder{24, false},
			Holder{25, false},
		},
		{
			Holder{16, false},
			Holder{17, false},
			Holder{18, false},
			Holder{19, false},
			Holder{20, false},
		},
		{
			Holder{6, false},
			Holder{7, false},
			Holder{8, false},
			Holder{9, false},
			Holder{10, false},
		},
	}

	sum := 25 * 13
	tests := []struct{ i, j int }{
		{0, 1},
		{3, 4},
		{4, 4},
		{2, 1},
		{2, 3},
		{4, 1},
	}

	for _, test := range tests {
		val := board[test.i][test.j].Val
		board.MarkValue(val)
		sum -= val
		if board.AddUnmarked() != sum {
			t.Errorf("AddUnmarked() = %d, want %d", board.AddUnmarked(), sum)
		}
	}
}

func TestBingo(t *testing.T) {
	makeBoard := func() Board {
		return Board{
		{
			Holder{1, false},
			Holder{2, false},
			Holder{3, false},
			Holder{4, false},
			Holder{5, false},
		},
		{
			Holder{11, false},
			Holder{12, false},
			Holder{13, false},
			Holder{14, false},
			Holder{15, false},
		},
		{
			Holder{21, false},
			Holder{22, false},
			Holder{23, false},
			Holder{24, false},
			Holder{25, false},
		},
		{
			Holder{16, false},
			Holder{17, false},
			Holder{18, false},
			Holder{19, false},
			Holder{20, false},
		},
		{
			Holder{6, false},
			Holder{7, false},
			Holder{8, false},
			Holder{9, false},
			Holder{10, false},
		},
	}
	}


	type coord struct {
		i, j int
	}

	tests := []struct {
		coords [5]coord
		bingo  bool
	}{
		{
			[5]coord{{0,0}, {0,1}, {0,2}, {0,3}, {0,4}},
			true,
		},
		{
			[5]coord{{0,0}, {1,1}, {0,2}, {2,3}, {0,4}},
			false,
		},
		{
			[5]coord{{2,0}, {2,1}, {2,2}, {2,3}, {2,4}},
			true,
		},
		{
			[5]coord{{2,0}, {2,1}, {4,2}, {2,3}, {2,4}},
			false,
		},
		{
			[5]coord{{0,2}, {1,2}, {2,2}, {3,2}, {4,2}},
			true,
		},
	}

	for _, test := range tests {
		board := makeBoard()
		for _, coord := range test.coords {
			board.MarkValue(board[coord.i][coord.j].Val)
		}
		if board.CheckBingo() != test.bingo {
			t.Errorf("Bingo() = %t, want %t, for %+v", board.CheckBingo(), test.bingo, test.coords)
			return
		}
	}
}
