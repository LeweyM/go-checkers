package main

import (
	"checkers"
	"checkers/Board"
	. "checkers/Model"
)

func main() {
	board := Board.NewBoard()
	playerOne := checkers.NewHumanPlayer(RED)
	playerTwo := checkers.NewHumanPlayer(BLUE)
	game := checkers.NewGame(board)
	match := checkers.NewMatch(game, playerOne, playerTwo)

	match.Play()
}
