package Model

import "fmt"

type Position struct {
	Col int
	Row int
}

type PiecePosition struct {
	Position Position
	Piece    Piece
}

type Move struct {
	Origin Position
	Target Position
}

func NewMove(originCol, originRow, targetCol, targetRow int) *Move {
	return &Move{Origin: Position{
		Col: originCol,
		Row: originRow,
	}, Target: Position{
		Col: targetCol,
		Row: targetRow,
	}}
}

func (m Move) String() string {
	return fmt.Sprintf("(%d, %d) -> (%d, %d)", m.Origin.Col, m.Origin.Row, m.Target.Col, m.Target.Row)
}

type Piece string
const (
	RedPawn  Piece = "red"
	BluePawn Piece = "blue"
	Empty    Piece = "empty"
)

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
