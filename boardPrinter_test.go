package checkers

import "testing"

func Test(t *testing.T) {
	board := NewBoard()
	game := NewGame(board)

	var originalPrinted =
	    "7  |   b   b   b   b |\n" +
		"6  | b   b   b   b   |\n" +
		"5  |   b   b   b   b |\n" +
		"4  |                 |\n" +
		"3  |   x   x   x   x |\n" +
		"2  | r   r   r   r   |\n" +
		"1  |   r   r   r   r |\n" +
		"0  | r   r   r   r   |\n" +
		"    ----------------\n" +
		"     0 1 2 3 4 5 6 7"

	printed := newBoardPrinter(game.getPieces(), game.AvailableMoves(RED)).print()

	if originalPrinted != printed {
		t.Fatalf("\nexpected:" +
			"\n%s\ngot:" +
			"\n%s ", originalPrinted, printed)
	}
}
