package board_test

import (
	"testing"

	"github.com/aklinker1/chess-engine/pkg/board"
)

func Test_PieceX(t *testing.T) {
	testCases := [][]uint{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
		{5, 5},
		{6, 6},
		{7, 7},

		{8, 0},
		{9, 1},
		{10, 2},
		{11, 3},
		{12, 4},
		{13, 5},
		{14, 6},
		{15, 7},
	}

	for _, testCase := range testCases {
		p := board.Piece{
			BoardPosition: testCase[0],
		}
		actual := p.X()
		expected := testCase[1]
		if actual != expected {
			t.Errorf("Expected position %d to be x=%d, but was %d", testCase[0], expected, actual)
		}
	}
}

func Test_PieceY(t *testing.T) {
	testCases := [][]uint{
		{0, 0}, {7, 0},
		{8, 1}, {15, 1},
		{16, 2}, {23, 2},
		{24, 3}, {31, 3},
		{32, 4}, {39, 4},
		{40, 5}, {47, 5},
		{48, 6}, {55, 6},
		{56, 7}, {63, 7},
	}

	for i, testCase := range testCases {
		p := board.Piece{
			BoardPosition: testCase[0],
		}
		actual := p.Y()
		expected := testCase[1]
		if actual != expected {
			t.Errorf("%d. Expected position %d to be y=%d, but was %d", i+1, testCase[0], expected, actual)
		}
	}
}

func Test_PieceFile(t *testing.T) {
	type TestCase struct {
		position uint
		file     string
	}
	testCases := []TestCase{
		{0, "a"},
		{1, "b"},
		{2, "c"},
		{3, "d"},
		{4, "e"},
		{5, "f"},
		{6, "g"},
		{7, "h"},
	}

	for _, testCase := range testCases {
		p := board.Piece{
			BoardPosition: testCase.position,
		}
		actual := p.File()
		expected := testCase.file
		if actual != expected {
			t.Errorf("Expected position %d to be file=%s, but was %s", p.BoardPosition, expected, actual)
		}
	}
}

func Test_PieceRank(t *testing.T) {
	type TestCase struct {
		position uint
		rank     string
	}
	testCases := []TestCase{
		{0, "8"},
		{8, "7"},
		{16, "6"},
		{24, "5"},
		{32, "4"},
		{40, "3"},
		{48, "2"},
		{56, "1"},
	}

	for i, testCase := range testCases {
		p := board.Piece{
			BoardPosition: testCase.position,
		}
		actual := p.Rank()
		expected := testCase.rank
		if actual != expected {
			t.Errorf("%d. Expected position %d to be rank=%s, but was %s", i+1, p.BoardPosition, expected, actual)
		}
	}
}
