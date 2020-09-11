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
	add(int, int, Piece)
}

//checkers implementation

type checkersBoard struct {
	squares []byte
	pieceMap map[Piece]byte
}

func (b *checkersBoard) add(i int, j int, piece Piece) {
	b.squares[mapToIndex(i, j)] = b.pieceMap[piece]
}

func NewBoard() *checkersBoard {
	pieceMap := make(map[Piece]byte)
	pieceMap[RED] = 'r'
	pieceMap[BLUE] = 'b'
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
		pieceMap: pieceMap,
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


