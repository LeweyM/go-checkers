package checkers

type Player interface {
	GetMove(g Game) (i1, j1, i2, j2 int)
}
