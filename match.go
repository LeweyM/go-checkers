package checkers

type Match interface {
	Play()
}

type match struct {
	game Game
 	playerOne Player
	playerTwo Player
	currentPlayer Player
}

func NewMatch(game Game, playerOne Player, playerTwo Player) *match {
	return &match{game: game, playerOne: playerOne, playerTwo: playerTwo, currentPlayer: playerOne}
}

func (m *match) Play() Player {
	for {
		if m.game.HasWinner() {
			return m.winner()
		} else {
			i1, j1, i2, j2 := m.getCurrentPlayer().GetMove(m.game)
			ok := m.game.Move(i1, j1, i2, j2)
			if ok {
				m.switchPlayer()
				println(i1, j1, i2, j2)
			}
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

func (m *match) getCurrentPlayer() Player {
	return m.currentPlayer
}

func (m *match) switchPlayer() {
	if m.currentPlayer == m.playerOne {
		m.currentPlayer = m.playerTwo
	} else {
		m.currentPlayer = m.playerOne
	}
}
