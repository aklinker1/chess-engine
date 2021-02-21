package board_test

import (
	"strings"
	"testing"

	"github.com/aklinker1/chess-engine/pkg/board"
	"github.com/aklinker1/chess-engine/pkg/printer"
)

func TestMarshalSymmetry(t *testing.T) {
	testCases := []string{
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		"rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2",
		"rnbqkbnr/ppp1pppp/8/3p4/2PP4/8/PP2PPPP/RNBQKBNR b KQkq c3 0 2",
	}
	for i, testCase := range testCases {
		b := board.Unmarshal(testCase)
		printer.PrintBoard(b)
		actual := board.Marshal(b)

		if actual != testCase {
			t.Errorf(strings.Join([]string{
				"%d. Marshalling of boards was unsymmetrical:",
				"\t- Actual:   \"%s\"",
				"\t- Expected: \"%s\"",
			}, "\n"), i+1, actual, testCase)
		}
	}
}
