package board

import "fmt"

type Board struct {
	FEN            *string
	Pieces         map[uint]Piece
	ToMove         string
	Castling       CastlesAvailable
	EnPassant      *string
	HalfmoveClock  int
	FullmoveNumber uint
}

type Piece struct {
	Value         rune
	BoardPosition uint
	Moves         interface{}
}

type CastlesAvailable struct {
	WhiteKingside  bool
	WhiteQueenside bool
	BlackKingside  bool
	BlackQueenside bool
}

// X is the index from left (0) to right (7) on the board. Used by the computer for calculations
func (piece Piece) X() uint {
	return piece.BoardPosition % 8
}

// Y is the index from top (0) to bottom (7) on the board.
func (piece Piece) Y() uint {
	return piece.BoardPosition / 8
}

// Rank is the "row" that the piece is on, from bottom (1) to top (8). Just used for display purposes
func (piece Piece) Rank() string {
	return fmt.Sprint(8 - piece.Y())
}

// Rank is the "row" that the piece is on, from bottom (1) to top (8). Just used for display purposes
func (piece Piece) File() string {
	switch piece.X() {
	default:
		fallthrough
	case 0:
		return "a"
	case 1:
		return "b"
	case 2:
		return "c"
	case 3:
		return "d"
	case 4:
		return "e"
	case 5:
		return "f"
	case 6:
		return "g"
	case 7:
		return "h"
	}
}
