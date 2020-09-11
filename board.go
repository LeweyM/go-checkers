package checkers

type Piece string
const (
	RED Piece = "red"
	BLUE Piece = "blue"
	EMPTY Piece = "empty"
	)

type Board interface {
	Get(int, int) Piece
	Remove(int, int)
	Add(int, int, Piece)
}

//checkers implementation

type checkersBoard struct {
	squares []byte
	pieceMap map[Piece]byte
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

//todo: boundary checks
func (b *checkersBoard) Add(i int, j int, piece Piece) {
	b.set(i, j, b.pieceMap[piece])
}

func (b *checkersBoard) Get(i int, j int) Piece {
	byte := b.squares[index(i, j)]
	if byte == 'b' {
		return BLUE
	} else if byte == 'r' {
		return RED
	} else {
		return EMPTY
	}
}

func (b *checkersBoard) Remove(i int, j int) {
	b.set(i, j, ' ')
}

func (b *checkersBoard) set(i int, j int, p byte) {
	b.squares[index(i, j)] = p
}

func index(i int, j int) int {
	return (i / 2) + (j * 4)
}


