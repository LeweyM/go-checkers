package Strategy

import (
	. "checkers/Board"
	. "checkers/Model"
)

type PieceStrategy interface {
	GenerateMoves(position Position) []Position
	GenerateTakes(position Position) []Position
	ValidateMove(oldCol, oldRow, newCol, newRow int) bool
	ValidateTake(oldCol, oldRow, newCol, newRow int) bool
}

type PawnStrategy struct {
	board     Board
	direction Direction
}

func NewPawnStrategy(board Board, direction Direction) *PawnStrategy {
	return &PawnStrategy{
		board: board,
		direction: direction,
	}
}

func (p PawnStrategy) GenerateMoves(position Position) []Position {
	newRow := position.Row + int(p.direction)
	return removeNonLegal([]Position{
		{position.Col + 1, newRow},
		{position.Col - 1, newRow},
	})
}

func (p PawnStrategy) GenerateTakes(position Position) []Position {
	newRow := position.Row + int(p.direction) * 2
	return removeNonLegal([]Position{
		{position.Col + 2, newRow},
		{position.Col - 2, newRow},
	})
}

func (p PawnStrategy) ValidateMove(oldCol, oldRow, newCol, newRow int) bool {
	_, destination := p.board.Get(newCol, newRow)
	if !(destination == Empty) {
		return false
	}
	inMovingRange := adjacentRow(newRow, oldRow, 1, p.direction) && adjacentColumn(oldCol, newCol, 1)
	return inMovingRange
}

func (p PawnStrategy) ValidateTake(oldCol, oldRow, newCol, newRow int) bool {
	_, destination := p.board.Get(newCol, newRow)
	if !(destination == Empty) {
		return false
	}
	_, piece := p.board.Get(oldCol, oldRow)
	inTakingRange := adjacentRow(newRow, oldRow, 2, p.direction) && adjacentColumn(oldCol, newCol, 2)
	_, targetPiece := p.board.Get((oldCol+newCol)/2, (oldRow+newRow)/2)
	return inTakingRange && areEnemies(piece, targetPiece)
}

func adjacentColumn(oldCol int, newCol int, distance int) bool {
	return abs(newCol-oldCol) == distance
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func adjacentRow(newRow, oldRow, distance int, direction Direction) bool {
	return newRow == oldRow+(distance*int(direction))
}

func areEnemies(p1, p2 Piece) bool {
	return (p1 == RedPawn && p2 == BluePawn) ||
		(p1 == BluePawn && p2 == RedPawn)
}

func removeNonLegal(moves []Position) []Position {
	var legalMoves []Position
	for _, move := range moves {
		if !illegal(move.Col) && !illegal(move.Row) {
			legalMoves = append(legalMoves, move)
		}
	}
	return legalMoves
}

func illegal(n int) bool { return n < 0 || n > 7 }




