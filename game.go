package checkers

type PlayerColor string

const (
	BLUE PlayerColor = "blue"
	RED  PlayerColor = "red"
)

type Direction int
const (
	UP   Direction = 1
	DOWN Direction = -1
)

type PieceStrategy func(i int, i2 int, i3 int, i4 int) bool

type Game interface {
	Move(oldCol, oldRow, newCol, newRow int) bool
	HasWinner() bool
	Winner() PlayerColor
}

type game struct {
	board        Board
	moveStrategy map[Piece]PieceStrategy
}

func NewGame(board Board) *game {
	game := game{
		board:        board,
		moveStrategy: make(map[Piece]PieceStrategy),
	}
	game.moveStrategy[BluePawn] = game.pawnStrategyFactory(DOWN)
	game.moveStrategy[RedPawn] = game.pawnStrategyFactory(UP)
	game.moveStrategy[Empty] = func(i int, i2 int, i3 int, i4 int) bool { return false }
	return &game
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
	if !(g.board.Get(newCol, newRow) == Empty) {
		return false
	}
	return g.moveStrategy[g.board.Get(oldCol, oldRow)](oldCol, oldRow, newCol, newRow)
}

func (g *game) move(oldCol int, oldRow int, newCol int, newRow int) {
	originalPiece := g.board.Get(oldCol, oldRow)
	g.board.Add(newCol, newRow, originalPiece)
	g.board.Remove(oldCol, oldRow)
}

func (g *game) pawnStrategyFactory(direction Direction) PieceStrategy {
	return func(oldCol int, oldRow int, newCol int, newRow int) bool {
		inMovingRange := adjacentRow(newRow, oldRow, 1, direction) && adjacentColumn(oldCol, newCol, 1)
		return inMovingRange || g.canCapture(oldCol, oldRow, newCol, newRow, direction)
	}
}

func adjacentColumn(oldCol int, newCol int, distance int) bool {
	return abs(newCol-oldCol) == distance
}

func adjacentRow(newRow, oldRow, distance int, direction Direction) bool {
	return newRow == oldRow + (distance * int(direction))
}

func (g *game) canCapture(oldCol, oldRow, newCol, newRow int, direction Direction) bool {
	inTakingRange := adjacentRow(newRow, oldRow, 2, direction) && adjacentColumn(oldCol, newCol, 2)
	targetPiece := g.board.Get((oldCol+newCol)/2, (oldRow+newRow)/2)
	return inTakingRange && g.areEnemies(g.board.Get(oldCol, oldRow), targetPiece)
}

func (g *game) areEnemies(p1, p2 Piece) bool {
	return (p1 == RedPawn && p2 == BluePawn) ||
		(p1 == BluePawn && p2 == RedPawn)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}