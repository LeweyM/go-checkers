package BoardPrinter

import (
	"checkers"
	"checkers/Board"
	"checkers/Model"
	"testing"
)

func Test(t *testing.T) {
	board := Board.NewBoard()
	game := checkers.NewGame(board)

	var originalPrinted =
	    "7  |   b   b   b   b |\n" +
		"6  | b   b   b   b   |\n" +
		"5  |   b   b   b   b |\n" +
		"4  |                 |\n" +
		"3  |   .   .   .   . |\n" +
		"2  | r   r   r   r   |\n" +
		"1  |   r   r   r   r |\n" +
		"0  | r   r   r   r   |\n" +
		"    ----------------\n" +
		"     0 1 2 3 4 5 6 7"

	printed := NewBoardPrinter(game.GetPieces(), game.AvailableMoves(Model.RED)).Print()

	if originalPrinted != printed {
		t.Fatalf("\nexpected:" +
			"\n%s\ngot:" +
			"\n%s ", originalPrinted, printed)
	}
}
