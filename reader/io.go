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
	new_scanner := bufio.NewScanner(f)
	return new_scanner
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
	var board_data [5]string

	for i := 0; i < 5; i++ {
		// Read the board as a string
		if scanner.Scan() {
			board_data[i] = scanner.Text()
		}
	}

	// Convert the board_data string into a new board of all the values
	var new_board data.Board
		
	for row := 0; row < 5; row++ {
		var line_data []string = strings.Fields(board_data[row])
		for col := 0; col < 5; col++ {
			new_board[row][col].Val, _ = strconv.Atoi(line_data[col])
			new_board[row][col].Marked = false
		}
	}
	return &new_board, nil
}
