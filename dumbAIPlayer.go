package checkers

import . "checkers/Model"

type dumbAIPlayer struct {
	color PlayerColor
}

func NewDumbAIPlayer(color PlayerColor) *dumbAIPlayer {
	return &dumbAIPlayer{color: color}
}

func (d *dumbAIPlayer) GetMove(g Game) (i1, j1, i2, j2 int) {
	moves := g.AvailableMoves(d.color)
	origin := moves[0].Origin
	target := moves[0].Target
	return origin.Col, origin.Row, target.Col, target.Row
}

