package main

import (
	"fmt"
	"os"

	"github.com/aklinker1/chess-engine/pkg/board"
	"github.com/aklinker1/chess-engine/pkg/printer"
)

func main() {
	fen := os.Args[1]
	board := board.Unmarshal(fen)
	printer.PrintBoard(board)
	fmt.Println()
}
