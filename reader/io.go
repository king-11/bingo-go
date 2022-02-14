package reader

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"

	"github.com/king-11/bingo/data"
)

// GetScanner returns a scanner for the given file.
func GetScanner(f io.Reader) *bufio.Scanner {
	return bufio.NewScanner(f)
}

// ReadOrder reads the order of values to mark in bingo game from the given scanner.
func ReadOrder(scanner *bufio.Scanner) ([]int, error) {
	var err error
	if scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			return nil, errors.New("empty line")
		}

		val := strings.Split(line, ",")
		order := make([]int, len(val))
		for i, v := range val {
			order[i], err = strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
		}
		return order, nil
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return nil, nil
}

// ReadBoard reads a single board from the given scanner.
func ReadBoard(scanner *bufio.Scanner) (*data.Board, error) {
	var err error
	board := data.NewBoard()

	for i := 0; i < 5; i++ {
		if scanner.Scan() {
			line := scanner.Text()

			nums := strings.Fields(line)
			for j := 0; j < 5; j++ {
				board[i][j].Val, _ = strconv.Atoi(nums[j])
			}
		}
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return board, nil
}
