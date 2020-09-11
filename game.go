package checkers

type Game interface {
	hasWinner() bool
}

type game struct {
	board Board
	playerOne Player
	playerTwo Player
}

