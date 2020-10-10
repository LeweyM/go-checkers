package main

import "checkers"

func main() {
	board := checkers.NewBoard()
	playerOne := checkers.NewHumanPlayer(checkers.RED)
	playerTwo := checkers.NewDumbAIPlayer(checkers.BLUE)
	game := checkers.NewGame(board)
	match := checkers.NewMatch(game, playerOne, playerTwo)

	match.Play()
}
