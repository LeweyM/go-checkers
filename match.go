package checkers

type Match interface {
	Play()
}

type match struct {
	game Game
 	playerOne Player
	playerTwo Player
}

func (m *match) Play() Player {
	for {
		if m.game.HasWinner() {
			return m.game.Winner()
		} else {
			i1, j1, i2, j2 := m.currentPlayer().GetMove(m.game)
			m.game.Move(i1, j1, i2, j2)
		}
	}
}

func (m *match) currentPlayer() Player {
	return m.playerOne
}


