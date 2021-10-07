package calendar

type Index struct {
	X int
	Y int
}

type Board struct {
	Parent *Board
	// Each byte is a row, e.g., 01111111 is a full row
	Cells [7]byte
}

var (
	emptyBoard = [7]byte{
		0b00000001,
		0b00000001,
		0b00000000,
		0b00000000,
		0b00000000,
		0b00000000,
		0b00001111,
	}
)

func NewBoard() *Board {
	return &Board{
		Cells: emptyBoard,
	}
}
