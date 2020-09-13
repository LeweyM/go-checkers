package checkers

type Match interface {
	Play()
}

type match struct {
	game Game
 	playerOne Player
	playerTwo Player
}

func NewMatch(game Game, playerOne Player, playerTwo Player) *match {
	return &match{game: game, playerOne: playerOne, playerTwo: playerTwo}
}

func (m *match) Play() Player {
	for {
		if m.game.HasWinner() {
			return m.winner()
		} else {
			i1, j1, i2, j2 := m.currentPlayer().GetMove(m.game)
			println(i1, j1, i2, j2)
			m.game.Move(i1, j1, i2, j2)
		}
	}
}

func (m *match) winner() Player {
	if m.game.Winner() == RED {
		return m.playerOne
	} else {
		return m.playerTwo
	}
}

func (m *match) currentPlayer() Player {
	return m.playerOne
}