package main

import (
	"fmt"
	"log"
	"os"

	"github.com/king-11/bingo/data"
	"github.com/king-11/bingo/play"
	"github.com/king-11/bingo/reader"
)

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	filename := os.Args[1]

	winner := true

	f, err := os.Open(filename)
	logError(err)
	defer f.Close()
	scanner := reader.GetScanner(f)

	order, err := reader.ReadOrder(scanner)
	logError(err)

	boards := make([]*data.Board, 0)
	for scanner.Scan() {
		board, err := reader.ReadBoard(scanner)
		logError(err)
		boards = append(boards, board)
	}

	if err := scanner.Err(); err != nil {
		logError(err)
	}

	if winner {
		fmt.Printf("Winner value is %d", play.FindWinner(order, boards))
		} else {
		fmt.Printf("Loser value is %d", play.FindLastBoard(order, boards))
	}
}
