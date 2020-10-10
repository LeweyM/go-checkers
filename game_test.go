package checkers

import (
	"checkers/Board"
	. "checkers/Model"
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

			board := Board.NewBoard()
			game := NewGame(board)
			_, movingPiece := board.Get(oldCol, oldRow)

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
	board := Board.NewBoard()
	game := NewGame(board)

	legalMove(t, game, 0, 2, 1, 3)
	legalMove(t, game, 3, 5, 2, 4)
	legalMove(t, game, 1, 3, 3, 5)
	expectSquare(2, 4, Empty, board, t)
	legalMove(t, game, 2, 6, 4, 4)
	expectSquare(3, 5, Empty, board, t)
}

func Test_available_moves(t *testing.T) {
	board := Board.NewBoard()
	game := NewGame(board)

	checkMoves(t, game, RED, 7, []Move{
		*NewMove(0, 2, 1, 3),
		*NewMove(2, 2, 1, 3),
		*NewMove(2, 2, 3, 3),
		*NewMove(4, 2, 3, 3),
		*NewMove(4, 2, 5, 3),
		*NewMove(6, 2, 5, 3),
		*NewMove(6, 2, 7, 3),
	})

	checkMoves(t, game, BLUE, 7, []Move{
		*NewMove(1, 5, 0, 4),
		*NewMove(1, 5, 2, 4),
		*NewMove(3, 5, 2, 4),
		*NewMove(3, 5, 4, 4),
		*NewMove(5, 5, 4, 4),
		*NewMove(5, 5, 6, 4),
		*NewMove(7, 5, 6, 4),
	})

	game.Move(6, 2, 7, 3)

	checkMoves(t, game, RED, 8, []Move{
		*NewMove(7, 3, 6, 4),
	})

	game.Move(5, 5, 6, 4)

	checkMoves(t, game, RED, 1, []Move{
		*NewMove(7,3, 5,5),
	})
}

func Test_available_moves_only_for_taking_into_free_squares(t *testing.T) {
	board := Board.NewBoard()
	game := NewGame(board)

	game.Move(0, 2, 1, 3)
	game.Move(1, 5, 2, 4)

	checkMoves(t, game, RED, 7, []Move{
		*NewMove(1,3, 0,4),
		*NewMove(1,1, 0,2),
		*NewMove(2,2, 3,3),
		*NewMove(4,2, 3,3),
		*NewMove(4,2, 5,3),
		*NewMove(6,2, 5,3),
		*NewMove(6,2, 7,3),
	})
}

func checkMoves(t *testing.T, game *game, playerColor PlayerColor, totalExpectedMoves int, expectedMoves []Move) {
	availableMoves := game.AvailableMoves(playerColor)
	if len(availableMoves) != totalExpectedMoves {
		t.Fatalf("expected %d availableMoves, got %d. %v", totalExpectedMoves, len(availableMoves), availableMoves)
	}
	assertExpectedMoves(t, expectedMoves, availableMoves)
}

func assertExpectedMoves(t *testing.T, expectedMoves []Move, moves []Move) {
	for _, expectedMove := range expectedMoves {
		assertContains(t, moves, expectedMove)
	}
}

func assertContains(t *testing.T, moves []Move, move Move) {
	if !contains(moves, move) {
		t.Fatalf("Available moves should contain %v. Was %v", move, moves)
	}
}

func contains(moves []Move, move Move) bool {
	for _, m := range moves {
		if m.Origin.Col == move.Origin.Col &&
			m.Origin.Row == move.Origin.Row &&
			m.Target.Col == move.Target.Col &&
			m.Target.Row == move.Target.Row {
			return true
		}
	}
	return false
}

func legalMove(t *testing.T, game *game, oldCol, oldRow, newCol, newRow int) {
	ok := game.Move(oldCol, oldRow, newCol, newRow)
	if !ok {
		t.Fatalf("should be legal")
	}
}

func expectSquare(col int, row int, expectedPiece Piece, board Board.Board, t *testing.T) {
	_, piece := board.Get(col, row)
	expectSquarePiece(t, col, row, expectedPiece, piece)
}

func expectSquarePiece(t *testing.T, col int, row int, expected Piece, actual Piece) {
	if expected != actual {
		t.Errorf("Expected (%d,%d: %v). Got (%d,%d: %v)",
			col, row, expected, col, row, actual)
	}
}
