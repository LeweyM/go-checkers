package checkers

type Piece string
const (
	RED Piece = "red"
	BLUE Piece = "blue"
	EMPTY Piece = "empty"
	)

type Board interface {
	get(int, int) Piece
	remove(int, int)
}

//checkers implementation

type checkersBoard struct {
	squares []byte
}

func NewBoard() *checkersBoard {
	return &checkersBoard{
		squares: []byte {
			'b', 'b', 'b', 'b',
			'b', 'b', 'b', 'b',
			'b', 'b', 'b', 'b',
			' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ',
			'r', 'r', 'r', 'r',
			'r', 'r', 'r', 'r',
			'r', 'r', 'r', 'r',
		},
	}
}

func (b *checkersBoard) get(i int, j int) Piece {
	byte := b.squares[mapToIndex(i, j)]
	if byte == 'b' {
		return BLUE
	} else if byte == 'r' {
		return RED
	} else {
		return EMPTY
	}
}

func (b *checkersBoard) remove(i int, j int) {
	b.squares[mapToIndex(i, j)] = ' '
}

func mapToIndex(i int, j int) int {
	return (i / 2) + (j * 4)
}


