package printer

import (
	"fmt"
	"strings"

	"github.com/aklinker1/chess-engine/pkg/board"
)

func PrintBoard(board *board.Board) {
	var characters = make([]interface{}, 64)
	for i := range characters {
		characters[i] = " "
	}
	for i, piece := range board.Pieces {
		characters[i] = PieceEnumRunes[piece.Value]
	}

	boardString := strings.Join([]string{
		"   \x1b[40m\x1b[97m╭───┬───┬───┬───┬───┬───┬───┬───╮",
		" 8 \x1b[40m\x1b[97m│ x │ x │ x │ x │ x │ x │ x │ x │",
		"   \x1b[40m\x1b[97m├───┼───┼───┼───┼───┼───┼───┼───┤",
		" 7 \x1b[40m\x1b[97m│ x │ x │ x │ x │ x │ x │ x │ x │",
		"   \x1b[40m\x1b[97m├───┼───┼───┼───┼───┼───┼───┼───┤",
		" 6 \x1b[40m\x1b[97m│ x │ x │ x │ x │ x │ x │ x │ x │",
		"   \x1b[40m\x1b[97m├───┼───┼───┼───┼───┼───┼───┼───┤",
		" 5 \x1b[40m\x1b[97m│ x │ x │ x │ x │ x │ x │ x │ x │",
		"   \x1b[40m\x1b[97m├───┼───┼───┼───┼───┼───┼───┼───┤",
		" 4 \x1b[40m\x1b[97m│ x │ x │ x │ x │ x │ x │ x │ x │",
		"   \x1b[40m\x1b[97m├───┼───┼───┼───┼───┼───┼───┼───┤",
		" 3 \x1b[40m\x1b[97m│ x │ x │ x │ x │ x │ x │ x │ x │",
		"   \x1b[40m\x1b[97m├───┼───┼───┼───┼───┼───┼───┼───┤",
		" 2 \x1b[40m\x1b[97m│ x │ x │ x │ x │ x │ x │ x │ x │",
		"   \x1b[40m\x1b[97m├───┼───┼───┼───┼───┼───┼───┼───┤",
		" 1 \x1b[40m\x1b[97m│ x │ x │ x │ x │ x │ x │ x │ x │",
		"   \x1b[40m\x1b[97m╰───┴───┴───┴───┴───┴───┴───┴───╯",
		"     a   b   c   d   e   f   g   h  ",
		"",
	}, "\x1b[0m\n")
	boardStringFormatted := strings.ReplaceAll(boardString, "x", "%s")
	fmt.Printf(boardStringFormatted, characters...)
}
