package checkers

type Piece string

const (
	RedPawn  Piece = "red"
	BluePawn Piece = "blue"
	Empty    Piece = "empty"
)

type Position struct {
	col int
	row int
}

type Board interface {
	Get(col, row int) (bool, Piece)
	Remove(int, int)
	Add(int, int, Piece)
	Pieces(playerColor PlayerColor) []Position
}

//checkers implementation

type checkersBoard struct {
	squares  []byte
	pieceMap map[Piece]byte
}

//todo: optimise by maintaining data structure of pieces
func (b *checkersBoard) Pieces(playerColor PlayerColor) []Position {
	var playerPiece Piece
	if playerColor == BLUE {
		playerPiece = BluePawn
	} else if playerColor == RED {
		playerPiece = RedPawn
	}

	var piecePositions []Position

	for i := 0; i <= 7; i++ {
		for j := 0; j <= 7; j++ {
			_, piece := b.Get(i, j)
			if piece == playerPiece {
				piecePositions = append(piecePositions, Position{i, j})
			}
		}
	}

	return piecePositions
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

func (b *checkersBoard) Get(i int, j int) (bool, Piece) {
	illegal := isIllegal(i, j)
	if illegal {
		return false, Empty
	}

	byte := b.squares[index(i, j)]
	if byte == 'b' {
		return true, BluePawn
	} else if byte == 'r' {
		return true, RedPawn
	} else {
		return true, Empty
	}
}

func (b *checkersBoard) Remove(i int, j int) {
	b.set(i, j, ' ')
}

func (b *checkersBoard) set(i int, j int, p byte) {
	b.squares[index(i, j)] = p
}

func isIllegal(i int, j int) bool {
	return (isEven(i) && !isEven(j)) ||
		(!isEven(i) && isEven(j))
}

func isEven(i int) bool {
	return i%2 == 0
}

func index(i int, j int) int {
	row := 7 - j
	return (i / 2) + (row * 4)
}
