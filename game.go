package checkers

import (
	"checkers/Board"
	. "checkers/Model"
	"checkers/Strategy"
)

type Game interface {
	Move(oldCol, oldRow, newCol, newRow int) bool
	HasWinner() bool
	Winner() PlayerColor
	AvailableMoves(color PlayerColor) []Move
	GetPieces() []PiecePosition
	CurrentPlayer() PlayerColor
}

type game struct {
	board           Board.Board
	playerDirection map[PlayerColor]Direction
	strategy        map[Piece]Strategy.PieceStrategy
	currentPlayer	PlayerColor
}

func NewGame(board Board.Board) *game {
	game := game{
		board:           board,
		strategy:        make(map[Piece]Strategy.PieceStrategy),
		playerDirection: make(map[PlayerColor]Direction),
		currentPlayer:   RED,
	}
	game.playerDirection[BLUE] = DOWN
	game.playerDirection[RED] = UP
	game.strategy[BluePawn] = Strategy.NewPawnStrategy(board, DOWN)
	game.strategy[RedPawn] = Strategy.NewPawnStrategy(board, UP)
	game.strategy[RedKing] = Strategy.NewKingStrategy(board)
	game.strategy[BlueKing] = Strategy.NewKingStrategy(board)
	return &game
}

func (g *game) CurrentPlayer() PlayerColor {
	return g.currentPlayer
}

func (g *game) GetPieces() []PiecePosition {
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

func (g *game) Move(oldCol, oldRow, newCol, newRow int) bool {
	_, piece := g.board.Get(oldCol, oldRow)
	strategy := g.strategy[piece]
	canMove := strategy.ValidateMove(oldCol, oldRow, newCol, newRow); if canMove {
		g.move(oldCol, oldRow, newCol, newRow)
		switchPlayer(g)
		return true
	}
	canTake := strategy.ValidateTake(oldCol, oldRow, newCol, newRow); if canTake {
		g.take(oldCol, oldRow, newCol, newRow)
		if len(generateTakes([]Position{{Col: newCol, Row: newRow}}, g)) == 0 {
			switchPlayer(g)
		}
		return true
	}
	return false
}

func (g *game) move(oldCol int, oldRow int, newCol int, newRow int) {
	_, originalPiece := g.board.Get(oldCol, oldRow)
	g.board.Add(newCol, newRow, originalPiece)
	g.board.Remove(oldCol, oldRow)
	if isOnFinalRow(originalPiece, newRow) {
		g.board.Add(newCol, newRow, Piece.Promote(originalPiece))
	}
}

func isOnFinalRow(piece Piece, row int) bool {
	return true
}

func switchPlayer(g *game) {
	if g.currentPlayer == BLUE { g.currentPlayer = RED
	} else { g.currentPlayer = BLUE}
}

func (g *game) take(col int, row int, col2 int, row2 int) {
	g.move(col, row, col2, row2)
	g.board.Remove((col+col2)/2, (row+row2)/2)
}

func (g *game) AvailableMoves(color PlayerColor) []Move {
	positions := g.board.Pieces(color)
	moves := generateTakes(positions, g)
	if len(moves) > 0 {
		return moves
	}
	takes := generateMoves(positions, g)
	return append(moves, takes...)
}

func generateMoves(positions []Position, g *game) []Move {
	var moves []Move
	for _, position := range positions {
		_, piece := g.board.Get(position.Col, position.Row)
		strategy := g.strategy[piece]
		moves = append(moves, generate(position, strategy.GenerateMoves, strategy.ValidateMove)...)
	}
	return moves
}

func generateTakes(positions []Position, g *game) []Move {
	var moves []Move
	for _, position := range positions {
		_, piece := g.board.Get(position.Col, position.Row)
		strategy := g.strategy[piece]
		moves = append(moves, generate(position, strategy.GenerateTakes, strategy.ValidateTake)...)
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
