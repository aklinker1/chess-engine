package board

import (
	"fmt"
	"strings"
)

func Marshal(board *Board) string {
	fen := ""

	var x uint
	var y int
	positions := make([]*rune, 64)
	for y = 7; y >= 0; y-- {
		for x = 0; x < 8; x++ {
			i := x + 8*uint(y)
			if piece, ok := board.Pieces[i]; ok {
				positions[i] = &piece.Value
			}
		}
	}
	ranks := []string{}
	for y = 7; y >= 0; y-- {
		nilCount := 0
		rank := ""
		for x = 0; x < 8; x++ {
			i := x + 8*uint(y)
			char := positions[i]
			if char == nil {
				nilCount++
			} else {
				if nilCount > 0 {
					rank += fmt.Sprint(nilCount)
					nilCount = 0
				}
				rank += string(*char)
			}
		}
		if nilCount > 0 {
			rank += fmt.Sprint(nilCount)
		}
		ranks = append(ranks, rank)
	}

	fen += strings.Join(ranks, "/")

	fen += " " + board.ToMove

	if board.Castling.BlackKingside || board.Castling.BlackQueenside || board.Castling.WhiteKingside || board.Castling.WhiteQueenside {
		fen += " "
		if board.Castling.WhiteKingside {
			fen += "K"
		}
		if board.Castling.WhiteQueenside {
			fen += "Q"
		}
		if board.Castling.BlackKingside {
			fen += "k"
		}
		if board.Castling.BlackQueenside {
			fen += "q"
		}
	} else {
		fen += "-"
	}

	if board.EnPassant != nil {
		fen += " " + *board.EnPassant
	} else {
		fen += "-"
	}

	fen += fmt.Sprintf(" %d", board.HalfmoveClock)
	fen += fmt.Sprintf(" %d", board.FullmoveNumber)

	return fen
}
