package checkers

type PlayerColor string
const (
	BLUE PlayerColor = "blue"
	RED PlayerColor = "red"
)

type Game interface {
	Move(oldCol, oldRow, newCol, newRow int) bool
	HasWinner() bool
	Winner() PlayerColor
}

type game struct {
	board Board
}

func NewGame(board Board) *game {
	return &game{board: board}
}

func (g *game) Move(oldCol, oldRow, newCol, newRow int) bool {
	ok := g.canMove(oldCol, oldRow, newCol, newRow)
	if ok {
		g.move(oldCol, oldRow, newCol, newRow)
	}
	return ok
}

// private methods

func (g *game) canMove(oldCol int, oldRow int, newCol int, newRow int) bool {
	if !g.squareFree(newCol, newRow) {
		return false
	}

	piece := g.board.Get(oldCol, oldRow)
	if piece == BluePawn {
		return below(newRow, oldRow) && adjacent(oldCol, newCol)
	} else if piece == RedPawn {
		return above(newRow, oldRow) && adjacent(oldCol, newCol)
	} else {
		return false
	}
}

func (g *game) move(oldCol int, oldRow int, newCol int, newRow int) {
	originalPiece := g.board.Get(oldCol, oldRow)
	g.board.Add(newCol, newRow, originalPiece)
	g.board.Remove(oldCol, oldRow)
}

func (g *game) squareFree(newCol int, newRow int) bool {
	return g.board.Get(newCol, newRow) == Empty
}

func adjacent(oldCol int, newCol int) bool {
	return abs(newCol-oldCol) == 1
}

func above(newRow int, oldRow int) bool {
	return newRow == oldRow+1
}

func below(newRow int, oldRow int) bool {
	return newRow == oldRow-1
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}