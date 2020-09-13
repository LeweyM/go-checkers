package checkers

type Piece string

const (
	RedPawn  Piece = "red"
	BluePawn Piece = "blue"
	Empty    Piece = "empty"
)

type Board interface {
	Get(col, row int) Piece
	Remove(int, int)
	Add(int, int, Piece)
}

//checkers implementation

type checkersBoard struct {
	squares  []byte
	pieceMap map[Piece]byte
}

func NewBoard() *checkersBoard {
	pieceMap := make(map[Piece]byte)
	pieceMap[RedPawn] = 'r'
	pieceMap[BluePawn] = 'b'
	return &checkersBoard{
		squares: []byte{
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
	mustBeLegal(i, j)

	byte := b.squares[index(i, j)]
	if byte == 'b' {
		return BluePawn
	} else if byte == 'r' {
		return RedPawn
	} else {
		return Empty
	}
}

func (b *checkersBoard) Remove(i int, j int) {
	b.set(i, j, ' ')
}

func (b *checkersBoard) set(i int, j int, p byte) {
	b.squares[index(i, j)] = p
}

func mustBeLegal(i int, j int) {
	if (isEven(i) && !isEven(j)) ||
		(!isEven(i) && isEven(j)) {
		panic("illegal board request")
	}
}

func isEven(i int) bool {
	return i%2 == 0
}

func index(i int, j int) int {
	row := 7 - j
	return (i / 2) + (row * 4)
}
