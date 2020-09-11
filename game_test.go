package checkers

import (
	"fmt"
	"testing"
)

func Test_game(t *testing.T) {
	type args struct {
		i1, j1, i2, j2 int
	}
	tests := []struct {
		args args
		want bool
	}{
		{args{0,0, 7, 7}, false},
		{args{0,0, 1, 1}, false},
		{args{0,2, 1, 3}, true},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%d,%d to %d,%d should be %v", tt.args.i1, tt.args.j1, tt.args.i1, tt.args.j1, tt.want)
		t.Run(testName, func(t *testing.T) {
			oldRow, oldCol, newRow, newCol := tt.args.i1, tt.args.j1, tt.args.i2, tt.args.j2

			board := NewBoard()
			game := NewGame(board, mockPlayer{}, mockPlayer{})
			movingPiece := board.Get(oldRow, newRow)

			result := game.Move(oldRow, oldCol, newRow, newCol); if tt.want != result {
				t.Errorf("%d,%d to %d,%d, should be %v, got %v",
					oldRow, oldCol, newRow, newCol, tt.want, result)
			}

			if tt.want == false {
				return
			}

			expectSquare(oldRow, oldCol, EMPTY, board, t)
			expectSquare(newRow, newCol, movingPiece, board, t)
		})
	}
}

func expectSquare(row int, col int, expectedPiece Piece, board *checkersBoard, t *testing.T) {
	expectSquarePiece(t, row, col, expectedPiece, board.Get(row, col))
}

func expectSquarePiece(t *testing.T, row int, col int, expected Piece, actual Piece) {
	if expected != actual {
		t.Errorf("Expected (%d,%d: %v). Got (%d,%d: %v)",
			row, col, expected, row, col, actual)
	}
}

type mockPlayer struct {}