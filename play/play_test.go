package play

import (
	"os"
	"testing"

	"github.com/king-11/bingo/data"
	"github.com/king-11/bingo/reader"
)

func TestFindWinner(t *testing.T) {
	tests := []struct {
		filename string
		expected int
	}{
		{"test1.txt", 16716},
		{"test2.txt", 69579},
	}

	for _, test := range tests {
		t.Run(test.filename, func(t *testing.T) {
			f, err := os.Open(test.filename)
			if err != nil {
				t.Error(err)
				return
			}
			defer f.Close()

			scanner := reader.GetScanner(f)

			order, err := reader.ReadOrder(scanner)
			if err != nil {
				t.Error(err)
				return
			}

			boards := make([]*data.Board, 0)
			for scanner.Scan() {
				board, err := reader.ReadBoard(scanner)
				if err != nil {
					t.Error(err)
					return
				}
				boards = append(boards, board)
			}

			if err := scanner.Err(); err != nil {
				t.Error(err)
				return
			}

			val := FindWinner(order, boards)
			if test.expected != val {
				t.Errorf("expected %d, got %d", test.expected, val)
			}
		})
	}
}

func TestFindLastBoard(t *testing.T) {
	tests := []struct {
		filename string
		expected int
	}{
		{"test1.txt", 4880},
		{"test2.txt", 14877},
	}

	for _, test := range tests {
		t.Run(test.filename, func(t *testing.T) {
			f, err := os.Open(test.filename)
			if err != nil {
				t.Error(err)
				return
			}
			defer f.Close()

			scanner := reader.GetScanner(f)

			order, err := reader.ReadOrder(scanner)
			if err != nil {
				t.Error(err)
				return
			}

			boards := make([]*data.Board, 0)
			for scanner.Scan() {
				board, err := reader.ReadBoard(scanner)
				if err != nil {
					t.Error(err)
					return
				}
				boards = append(boards, board)
			}

			if err := scanner.Err(); err != nil {
				t.Error(err)
				return
			}

			val := FindLastBoard(order, boards)
			if test.expected != val {
				t.Errorf("expected %d, got %d", test.expected, val)
			}
		})
	}
}
