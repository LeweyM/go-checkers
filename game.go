package checkers

import (
	"fmt"
	"strings"
)

type Move struct {
	origin Position
	target Position
}

func (m Move) String() string {
	return fmt.Sprintf("(%d, %d) -> (%d, %d)", m.origin.col, m.origin.row, m.target.col, m.target.row)
}

func NewMove(originCol, originRow, targetCol, targetRow int) *Move {
	return &Move{origin: Position{
		col: originCol,
		row: originRow,
	}, target: Position{
		col: targetCol,
		row: targetRow,
	}}
}

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
	PrintBoard() string
	AvailableMoves(color PlayerColor) []Move
}

type game struct {
	board           Board
	moveStrategy    map[Piece]PieceStrategy
	playerDirection map[PlayerColor]Direction
}

func (g *game) PrintBoard() string {
	builder := strings.Builder{}
	for i := 7; i >= 0; i-- {
		builder.WriteString(fmt.Sprintf("%d  ", i))
		builder.WriteByte('|')
		for j := 0; j <= 7; j++ {
			builder.WriteByte(' ')
			ok, piece := g.board.Get(j, i)
			if !ok || piece == Empty {
				builder.WriteByte(' ')
			} else if piece == BluePawn {
				builder.WriteByte('b')
			} else if piece == RedPawn {
				builder.WriteByte('r')
			}
		}
		builder.WriteByte('|')
		builder.WriteByte('\n')
	}
	builder.WriteString(fmt.Sprintf("    ----------------\n    "))
	for j := 0; j <= 7; j++ {
		builder.WriteString(fmt.Sprintf(" %d", j))
	}
	builder.WriteString(fmt.Sprintf("    \n"))
	return builder.String()
}

func (g *game) HasWinner() bool {
	return false
}

func (g *game) Winner() PlayerColor {
	return BLUE
}

func NewGame(board Board) *game {
	game := game{
		board:        board,
		moveStrategy: make(map[Piece]PieceStrategy),
		playerDirection: make(map[PlayerColor]Direction),
	}
	game.playerDirection[BLUE] = DOWN
	game.playerDirection[RED] = UP
	game.moveStrategy[BluePawn] = game.pawnStrategyFactory(DOWN)
	game.moveStrategy[RedPawn] = game.pawnStrategyFactory(UP)
	game.moveStrategy[Empty] = func(i int, i2 int, i3 int, i4 int) bool { return false }
	return &game
}

func (g *game) Move(oldCol, oldRow, newCol, newRow int) bool {
	ok := g.canMove(oldCol, oldRow, newCol, newRow)
	if ok {
		if isTakingMove(oldCol, oldRow, newCol, newRow) {
			g.take(oldCol, oldRow, newCol, newRow)
		} else {
			g.move(oldCol, oldRow, newCol, newRow)
		}
	}
	return ok
}

func (g *game) AvailableMoves(color PlayerColor) []Move {
	positions := g.board.Pieces(color)
	var moves []Move
	for _, position := range positions {
		possibleMoves := generatePawnMoves(position, g.playerDirection[color])
		legalMoves := removeNonLegal(possibleMoves)
		for _, target := range legalMoves {
			if g.canMove(position.col, position.row, target.col, target.row) {
				moves = append(moves, Move{
					origin: position,
					target: target,
				})
			}
		}
	}
	return moves
}

func removeNonLegal(moves []Position) []Position {
	var legalMoves []Position
	for _, move := range moves {
		if !illegal(move.col) && !illegal(move.row) {
			legalMoves = append(legalMoves, move)
		}
	}
	return legalMoves
}

func illegal(n int) bool { return n < 0 || n > 7 }

func generatePawnMoves(position Position, direction Direction) []Position {
	newRow := position.row + int(direction)
	return []Position{
		{position.col + 1, newRow},
		{position.col - 1, newRow},
	}
}

func (g *game) canMove(oldCol int, oldRow int, newCol int, newRow int) bool {
	_, destination := g.board.Get(newCol, newRow)
	if !(destination == Empty) {
		return false
	}
	_, oldPiece := g.board.Get(oldCol, oldRow)
	return g.moveStrategy[oldPiece](oldCol, oldRow, newCol, newRow)
}

func isTakingMove(col int, row int, col2 int, row2 int) bool {
	return abs(col-col2) == 2 && abs(row-row2) == 2
}

func (g *game) move(oldCol int, oldRow int, newCol int, newRow int) {
	_, originalPiece := g.board.Get(oldCol, oldRow)
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
	return newRow == oldRow+(distance*int(direction))
}

func (g *game) canCapture(oldCol, oldRow, newCol, newRow int, direction Direction) bool {
	inTakingRange := adjacentRow(newRow, oldRow, 2, direction) && adjacentColumn(oldCol, newCol, 2)
	_, targetPiece := g.board.Get((oldCol+newCol)/2, (oldRow+newRow)/2)
	_, piece := g.board.Get(oldCol, oldRow)
	return inTakingRange && g.areEnemies(piece, targetPiece)
}

func (g *game) areEnemies(p1, p2 Piece) bool {
	return (p1 == RedPawn && p2 == BluePawn) ||
		(p1 == BluePawn && p2 == RedPawn)
}
func (g *game) take(col int, row int, col2 int, row2 int) {
	g.move(col, row, col2, row2)
	g.board.Remove((col+col2)/2, (row+row2)/2)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
