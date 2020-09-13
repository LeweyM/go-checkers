package main

import "checkers"

func main() {
	board := checkers.NewBoard()
	playerOne := checkers.NewHumanPlayer()
	playerTwo := checkers.NewHumanPlayer()
	game := checkers.NewGame(board)
	match := checkers.NewMatch(game, playerOne, playerTwo)

	match.Play()
}