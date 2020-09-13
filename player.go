package checkers

import (
	"bufio"
	"fmt"
	"os"
)

type Player interface {
	GetMove(g Game) (i1, j1, i2, j2 int)
}

type humanPlayer struct{}

func NewHumanPlayer() *humanPlayer {
	return &humanPlayer{}
}

func (p humanPlayer) GetMove(g Game) (i1, j1, i2, j2 int) {
	fmt.Print(g.PrintBoard())
	for {
		a, b, c, d, ok := requestMove()
		if ok {
			return a, b, c, d
		}
	}
}

func requestMove() (int, int, int, int, bool) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter move: ")
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

func moveOk(a int, b int, c int, d int) bool {
	return positionOk(a) && positionOk(b) && positionOk(c) && positionOk(d)
}

func positionOk(a int) bool {
	return a >= 1 || a <= 7
}
