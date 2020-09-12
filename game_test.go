package checkers

import (
	"fmt"
	"testing"
)

func Test_game(t *testing.T) {
	type args struct {
		oldCol, oldRow, newCol, newRow int
	}
	tests := []struct {
		args    args
		isLegal bool
	}{
		{args{0,0, 1, 1}, false},
		{args{0,0, 1, 1}, false},
		{args{0,2, 2, 4}, false},
		{args{1,5, 3, 3}, false},
		{args{0,2, 1, 3}, true},
		{args{5,5, 4, 4}, true},
		{args{5,5, 6, 4}, true},
		{args{0,6, 1, 7}, false},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%d,%d to %d,%d", tt.args.oldCol, tt.args.oldRow, tt.args.newCol, tt.args.newRow)
		t.Run(testName, func(t *testing.T) {
			oldCol, oldRow, newCol, newRow := tt.args.oldCol, tt.args.oldRow, tt.args.newCol, tt.args.newRow

			board := NewBoard()
			game := NewGame(board)
			movingPiece := board.Get(oldCol, oldRow)

			board.Remove(1, 7) //remove top left piece

			result := game.Move(oldCol, oldRow, newCol, newRow); if tt.isLegal != result {
				t.Errorf("%d,%d to %d,%d, should be il/legal but isn't.",
					oldCol, oldRow, newCol, newRow)
			}

			if tt.isLegal == false { return }

			expectSquare(oldCol, oldRow, Empty, board, t)
			expectSquare(newCol, newRow, movingPiece, board, t)
		})
	}
}

func Test_taking_pieces(t *testing.T) {
	board := NewBoard()
	game := NewGame(board)

	legalMove(t, game, 0, 2, 1, 3)
	legalMove(t, game, 3, 5, 2, 4)
	legalMove(t, game, 1, 3, 3, 5)
	legalMove(t, game, 2, 6, 4, 4)
}

func legalMove(t *testing.T, game *game, oldCol, oldRow, newCol, newRow int) {
	ok := game.Move(oldCol, oldRow, newCol, newRow)
	if !ok {
		t.Fatalf("should be legal")
	}
}

func expectSquare(col int, row int, expectedPiece Piece, board *checkersBoard, t *testing.T) {
	expectSquarePiece(t, col, row, expectedPiece, board.Get(col, row))
}

func expectSquarePiece(t *testing.T, col int, row int, expected Piece, actual Piece) {
	if expected != actual {
		t.Errorf("Expected (%d,%d: %v). Got (%d,%d: %v)",
			col, row, expected, col, row, actual)
	}
}