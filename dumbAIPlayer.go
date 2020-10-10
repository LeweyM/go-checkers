package checkers

type dumbAIPlayer struct {
	color PlayerColor
}

func NewDumbAIPlayer(color PlayerColor) *dumbAIPlayer {
	return &dumbAIPlayer{color: color}
}

func (d *dumbAIPlayer) GetMove(g Game) (i1, j1, i2, j2 int) {
	moves := g.AvailableMoves(d.color)
	origin := moves[0].origin
	target := moves[0].target
	return origin.col, origin.row, target.col, target.row
}

