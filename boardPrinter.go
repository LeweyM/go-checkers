package checkers

import (
	"fmt"
	"strings"
)

type boardPrinter struct {
	pieces []PiecePosition
	moves []Move
	shape map[Piece]byte
}

func newBoardPrinter(pieces []PiecePosition, moves []Move) *boardPrinter {
	shape := map[Piece]byte{
		BluePawn: 'b',
		RedPawn: 'r',
		Empty: ' ',
	}

	return &boardPrinter{pieces: pieces, shape: shape, moves: moves}
}

func (bp *boardPrinter) print() string {
	var cols = [8][8]byte{}

	for _, p := range bp.pieces {
		cols[p.position.row][p.position.col] = bp.shape[p.piece]
	}

	for _, move := range bp.moves {
		cols[move.target.row][move.target.col] = 'x'
	}

	builder := strings.Builder{}
	for i, _ := range cols {
		builder.WriteString(fmt.Sprintf("%d  ", len(cols)-1-i))
		builder.WriteByte('|')
		for _, cell := range cols[len(cols)-1-i] {
			builder.WriteByte(' ')
			if cell == 0 {
				builder.WriteByte(' ')
			} else {
				builder.WriteByte(cell)
			}
		}
		builder.WriteByte(' ')
		builder.WriteByte('|')
		builder.WriteByte('\n')
	}

	builder.WriteString(fmt.Sprintf("    ----------------\n    "))
	for j := 0; j <= 7; j++ {
		builder.WriteString(fmt.Sprintf(" %d", j))
	}
	return builder.String()
}
