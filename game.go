package checkers

type PlayerColor string
const (
	BLUE PlayerColor = "blue"
	RED PlayerColor = "red" 
)

type Game interface {
	Move(oldRow, oldCol, newRow, newCol int) bool
	HasWinner() bool
	Winner() PlayerColor
}

type game struct {
	board Board
}

func NewGame(board Board) *game {
	return &game{board: board}
}

func (g *game) Move(oldRow, oldCol, newRow, newCol int) bool {
	ok := g.canMove(oldRow, oldCol, newRow, newCol)
	if ok {
		g.move(oldRow, oldCol, newRow, newCol)
	}
	return ok
}

func (g *game) canMove(oldRow int, oldCol int, newRow int, newCol int) bool {
	return isAdjacentAndAbove(oldRow, oldCol, newRow, newCol) && g.squareFree(newRow, newCol)
}

func isAdjacentAndAbove(i1 int, j1 int, i2 int, j2 int) bool {
	above := j2 == j1+1
	adjacent := abs(i2-i1) == 1
	return above && adjacent
}

func (g *game) squareFree(newRow int, newCol int) bool {
	return g.board.Get(newRow, newCol) == Empty
}

func (g *game) move(oldRow int, oldCol int, newRow int, newCol int) {
	originalPiece := g.board.Get(oldRow, oldCol)
	g.board.Add(newRow, newCol, originalPiece)
	g.board.Remove(oldRow, oldCol)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}