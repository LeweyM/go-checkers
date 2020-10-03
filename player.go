package checkers

import (
	"bufio"
	"fmt"
	"os"
)

type Player interface {
	GetMove(g Game) (i1, j1, i2, j2 int)
}

type humanPlayer struct {
	color PlayerColor
}

func NewHumanPlayer(color PlayerColor) *humanPlayer {
	return &humanPlayer{color}
}

func (p humanPlayer) GetMove(g Game) (i1, j1, i2, j2 int) {

	printer := newBoardPrinter(g.getPieces())
	fmt.Print(printer.print())
	fmt.Print("\n")

	availableMoves := g.AvailableMoves(p.getColor())
	for i, move := range availableMoves {
		fmt.Printf("Move %d: %v\n", i+1, move)
	}
	for {
		a, b, c, d, ok := p.requestMove()
		if ok {
			return a, b, c, d
		}
	}
}

func (p humanPlayer) getColor() PlayerColor {
	return p.color
}

func (p humanPlayer) requestMove() (int, int, int, int, bool) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Player %v:\nEnter move: ", p.getColor())
	move, _ := reader.ReadString('\n')
	if len(move) != 5 {
		return 0, 0, 0, 0, false
	}
	a := int(move[0] - 48)
	b := int(move[1] - 48)
	c := int(move[2] - 48)
	d := int(move[3] - 48)

	if !moveOk(a, b, c, d) {
		return 0, 0, 0, 0, false
	}

	fmt.Printf("\n(%d,%d) -> (%d,%d)\n", a, b, c, d)
	return a, b, c, d, true
}

//func (g *game) printBoard() string {
//	builder := strings.Builder{}
//	for i := 7; i >= 0; i-- {
//		builder.WriteString(fmt.Sprintf("%d  ", i))
//		builder.WriteByte('|')
//		for j := 0; j <= 7; j++ {
//			builder.WriteByte(' ')
//			ok, piece := g.board.Get(j, i)
//			if !ok || piece == Empty {
//				builder.WriteByte(' ')
//			} else if piece == BluePawn {
//				builder.WriteByte('b')
//			} else if piece == RedPawn {
//				builder.WriteByte('r')
//			}
//		}
//		builder.WriteByte('|')
//		builder.WriteByte('\n')
//	}
//	builder.WriteString(fmt.Sprintf("    ----------------\n    "))
//	for j := 0; j <= 7; j++ {
//		builder.WriteString(fmt.Sprintf(" %d", j))
//	}
//	builder.WriteString(fmt.Sprintf("    \n"))
//	return builder.String()
//}

func moveOk(a int, b int, c int, d int) bool {
	return positionOk(a) && positionOk(b) && positionOk(c) && positionOk(d)
}

func positionOk(a int) bool {
	return a >= 1 || a <= 7
}
