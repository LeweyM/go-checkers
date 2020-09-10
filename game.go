package checkers

type Game interface {

}

type game struct {
	board Board
	playerOne Player
	playerTwo Player
}

