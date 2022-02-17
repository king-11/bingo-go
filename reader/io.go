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
	scanner := bufio.NewScanner(f)
	return scanner
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
	newBoard := data.NewBoard()
	for i := 0; i < 5; i++ {
		if scanner.Scan() {
			line := scanner.Text()
			if line == "" || line == "\n" {
				return nil, errors.New("empty line")
			}
			val := strings.Fields(line)
			row := make([]int, len(val))
			for j, v := range val {
				row[j], err = strconv.Atoi(v)
				if err != nil {
					return nil, err
				}
				h := data.Holder{
					Val:    row[j],
					Marked: false,
				}
				newBoard[i][j] = h
			}
		} else {
			if err = scanner.Err(); err != nil {
				return nil, err
			}
		}
	}

	return newBoard, nil
}
