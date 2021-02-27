package Strategy

import (
	"checkers/Board"
	. "checkers/Model"
)

type kingStrategy struct {
	board Board.Board
}

func (k kingStrategy) GenerateMoves(position Position) []Position {
	upStrategy := NewPawnStrategy(k.board, UP)
	downStrategy := NewPawnStrategy(k.board, DOWN)
	return append(upStrategy.GenerateMoves(position), downStrategy.GenerateMoves(position)...)
}

func (k kingStrategy) GenerateTakes(position Position) []Position {
	upStrategy := NewPawnStrategy(k.board, UP)
	downStrategy := NewPawnStrategy(k.board, DOWN)
	return append(upStrategy.GenerateTakes(position), downStrategy.GenerateTakes(position)...)
}

func (k kingStrategy) ValidateMove(oldCol, oldRow, newCol, newRow int) bool {
	upStrategy := NewPawnStrategy(k.board, UP)
	downStrategy := NewPawnStrategy(k.board, DOWN)
	return upStrategy.ValidateMove(oldCol, oldRow, newCol, newRow) || downStrategy.ValidateMove(oldCol, oldRow, newCol, newRow)
}

func (k kingStrategy) ValidateTake(oldCol, oldRow, newCol, newRow int) bool {
	upStrategy := NewPawnStrategy(k.board, UP)
	downStrategy := NewPawnStrategy(k.board, DOWN)
	return upStrategy.ValidateTake(oldCol, oldRow, newCol, newRow) || downStrategy.ValidateTake(oldCol, oldRow, newCol, newRow)
}

func NewKingStrategy(board Board.Board) *kingStrategy {
	return &kingStrategy{
		board: board,
	}
}
