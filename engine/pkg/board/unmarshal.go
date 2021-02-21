package board

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func validate(fen string) []string {
	row := "([rnbqkpRNBQKP1-8]+)"
	toMove := "(b|w)"
	castling := "(-|[KQkq]+)"
	enPassant := "(-|[a-h][1-8])"
	halfmoveClock := "([0-9]+)"
	fullmoveNumber := "([1-9][0-9]*)"

	regex := fmt.Sprintf(
		"^%s/%s/%s/%s/%s/%s/%s/%s %s %s %s %s %s$",
		row, row, row, row, row, row, row, row,
		toMove,
		castling,
		enPassant,
		halfmoveClock,
		fullmoveNumber,
	)
	validFEN, err := regexp.Compile(regex)
	if err != nil {
		panic(err)
	}

	valid := validFEN.Match([]byte(fen))
	if !valid {
		panic("FEN position was not valid")
	}

	return validFEN.FindStringSubmatch(fen)
}

func Unmarshal(fen string) *Board {
	fmt.Println(strings.Join([]string{
		"Input Board:",
		"",
		"    " + fen,
		"",
	}, "\n"))
	groups := validate(fen)

	ranks := []string{
		groups[8],
		groups[7],
		groups[6],
		groups[5],
		groups[4],
		groups[3],
		groups[2],
		groups[1],
	}
	toMove := groups[9]
	castling := groups[10]
	enPassant := groups[11]
	halfmoveClock, _ := strconv.Atoi(groups[12])
	fullmoveNumber, _ := strconv.Atoi(groups[13])

	pieces := map[uint]Piece{}
	for y, rankData := range ranks {
		x := 0
		for _, char := range rankData {
			if char >= '1' && char <= '8' {
				x += int(char-'1') + 1
			} else {
				position := uint(8*y + x)
				pieces[position] = Piece{
					Value:         char,
					BoardPosition: position,
					Moves:         nil,
				}
				x++
			}
		}
	}

	return &Board{
		FEN:    &fen,
		Pieces: pieces,
		ToMove: toMove,
		Castling: CastlesAvailable{
			WhiteKingside:  strings.ContainsRune(castling, 'K'),
			WhiteQueenside: strings.ContainsRune(castling, 'Q'),
			BlackKingside:  strings.ContainsRune(castling, 'k'),
			BlackQueenside: strings.ContainsRune(castling, 'q'),
		},
		EnPassant:      &enPassant,
		HalfmoveClock:  halfmoveClock,
		FullmoveNumber: uint(fullmoveNumber),
	}
}
