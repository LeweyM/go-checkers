package checkers

import (
	"checkers/Board"
	. "checkers/Model"
	"checkers/Strategy"
)

func NewMove(originCol, originRow, targetCol, targetRow int) *Move {
	return &Move{Origin: Position{
		Col: originCol,
		Row: originRow,
	}, Target: Position{
		Col: targetCol,
		Row: targetRow,
	}}
}

type Game interface {
	Move(oldCol, oldRow, newCol, newRow int) bool
	HasWinner() bool
	Winner() PlayerColor
	AvailableMoves(color PlayerColor) []Move
	getPieces() []PiecePosition
}

type game struct {
	board           Board.Board
	playerDirection map[PlayerColor]Direction
	strategy        map[Piece]Strategy.PieceStrategy
}

func (g *game) getPieces() []PiecePosition {
	return append(
		mapPositionToPiecePosition(g.board.Pieces(BLUE), BluePawn),
		mapPositionToPiecePosition(g.board.Pieces(RED), RedPawn)...
	)
}

func mapPositionToPiecePosition(positions []Position, piece Piece) []PiecePosition {
	var piecePositions []PiecePosition
	for _, position := range positions {
		piecePositions = append(piecePositions, PiecePosition{
			Position: position,
			Piece:    piece,
		})
	}
	return piecePositions
}

func (g *game) HasWinner() bool {
	return false
}

func (g *game) Winner() PlayerColor {
	return BLUE
}

func NewGame(board Board.Board) *game {
	game := game{
		board:        board,
		strategy: make(map[Piece]Strategy.PieceStrategy),
		playerDirection: make(map[PlayerColor]Direction),
	}
	game.playerDirection[BLUE] = DOWN
	game.playerDirection[RED] = UP
	game.strategy[BluePawn] = Strategy.NewPawnStrategy(board, DOWN)
	game.strategy[RedPawn] = Strategy.NewPawnStrategy(board, UP)
	return &game
}

func (g *game) Move(oldCol, oldRow, newCol, newRow int) bool {
	_, piece := g.board.Get(oldCol, oldRow)
	canMove := g.strategy[piece].ValidateMove(oldCol, oldRow, newCol, newRow); if canMove {
		g.move(oldCol, oldRow, newCol, newRow); return true
	}
	canTake := g.strategy[piece].ValidateTake(oldCol, oldRow, newCol, newRow); if canTake {
		g.take(oldCol, oldRow, newCol, newRow); return true
	}
	return false
}

func (g *game) AvailableMoves(color PlayerColor) []Move {
	positions := g.board.Pieces(color)
	var moves []Move
	for _, position := range positions {
		_, piece := g.board.Get(position.Col, position.Row)
		strategy := g.strategy[piece]
		moves = append(moves, generate(position, strategy.GenerateTakes, strategy.ValidateTake)...)
	}
	if len(moves) > 0 {
		return moves
	}
	for _, position := range positions {
		_, piece := g.board.Get(position.Col, position.Row)
		strategy := g.strategy[piece]
		moves = append(moves, generate(position, strategy.GenerateMoves, strategy.ValidateMove)...)
	}
	return moves
}

func generate(position Position, generator func(Position) []Position, validator func(int, int, int, int) bool) []Move {
	var moves []Move
	generatedMoves := generator(position)
	for _, target := range generatedMoves {
		if validator(position.Col, position.Row, target.Col, target.Row) {
			moves = append(moves, Move{
				Origin: position,
				Target: target,
			})
		}
	}
	return moves
}

func (g *game) move(oldCol int, oldRow int, newCol int, newRow int) {
	_, originalPiece := g.board.Get(oldCol, oldRow)
	g.board.Add(newCol, newRow, originalPiece)
	g.board.Remove(oldCol, oldRow)
}

func (g *game) take(col int, row int, col2 int, row2 int) {
	g.move(col, row, col2, row2)
	g.board.Remove((col+col2)/2, (row+row2)/2)
}
