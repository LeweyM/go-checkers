package main

import (
	"checkers"
	"checkers/Board"
)

func main() {
	board := Board.NewBoard()
	playerOne := checkers.NewHumanPlayer(checkers.RED)
	playerTwo := checkers.NewDumbAIPlayer(checkers.BLUE)
	game := checkers.NewGame(board)
	match := checkers.NewMatch(game, playerOne, playerTwo)

	match.Play()
}
