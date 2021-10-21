package calendarSlow

import (
	"fmt"
	"strings"

	"github.com/ryan-ju/calendar-solver/util"
)

type PieceIndex byte
type Orientation int
type Piece [7]byte
type PieceKey struct {
	P PieceIndex
	O Orientation
	R bool
	X uint8
	Y uint8
}

const (
	// **
	// **
	// *
	P_P PieceIndex = 0b00000001
	// **
	// **
	// **
	P_O PieceIndex = 0b00000010
	// **
	// *
	// **
	P_C PieceIndex = 0b00000100
	// *
	// *
	// ***
	P_V PieceIndex = 0b00001000
	// *
	// *
	// *
	// **
	P_L PieceIndex = 0b00010000
	// *
	// *
	// **
	// *
	P_T PieceIndex = 0b00100000
	// *
	// **
	//  *
	//  *
	P_Z PieceIndex = 0b01000000
	// **
	//  *
	//  **
	P_S PieceIndex = 0b10000000

	O_UP Orientation = 0
	O_RT Orientation = 1
	O_DW Orientation = 2
	O_LT Orientation = 3
)

var (
	pieceIndexes = []PieceIndex{
		P_P,
		P_O,
		P_C,
		P_V,
		P_L,
		P_T,
		P_Z,
		P_S,
	}
)

var (
	emptyPiece Piece
	pieceP_UP  Piece = [7]byte{
		0b01100000,
		0b01100000,
		0b01000000,
	}
	pieceP_RT Piece = [7]byte{
		0b01110000,
		0b00110000,
	}
	pieceP_DW Piece = [7]byte{
		0b00100000,
		0b01100000,
		0b01100000,
	}
	pieceP_LT Piece = [7]byte{
		0b01100000,
		0b01110000,
	}
	pieceP_UP_R Piece = [7]byte{
		0b01100000,
		0b01100000,
		0b00100000,
	}
	pieceP_RT_R Piece = [7]byte{
		0b00110000,
		0b01110000,
	}
	pieceP_DW_R Piece = [7]byte{
		0b01000000,
		0b01100000,
		0b01100000,
	}
	pieceP_LT_R Piece = [7]byte{
		0b01110000,
		0b01100000,
	}
	pieceO_UP Piece = [7]byte{
		0b01100000,
		0b01100000,
		0b01100000,
	}
	pieceO_RT Piece = [7]byte{
		0b01110000,
		0b01110000,
	}
	pieceC_UP Piece = [7]byte{
		0b01100000,
		0b01000000,
		0b01100000,
	}
	pieceC_RT Piece = [7]byte{
		0b01110000,
		0b01010000,
	}
	pieceC_DW Piece = [7]byte{
		0b01100000,
		0b00100000,
		0b01100000,
	}
	pieceC_LT Piece = [7]byte{
		0b01010000,
		0b01110000,
	}
	pieceV_UP Piece = [7]byte{
		0b01000000,
		0b01000000,
		0b01110000,
	}
	pieceV_RT Piece = [7]byte{
		0b01110000,
		0b01000000,
		0b01000000,
	}
	pieceV_DW Piece = [7]byte{
		0b01110000,
		0b00010000,
		0b00010000,
	}
	pieceV_LT Piece = [7]byte{
		0b00010000,
		0b00010000,
		0b01110000,
	}
	pieceL_UP Piece = [7]byte{
		0b01000000,
		0b01000000,
		0b01000000,
		0b01100000,
	}
	pieceL_RT Piece = [7]byte{
		0b01111000,
		0b01000000,
	}
	pieceL_DW Piece = [7]byte{
		0b01100000,
		0b00100000,
		0b00100000,
		0b00100000,
	}
	pieceL_LT Piece = [7]byte{
		0b00001000,
		0b01111000,
	}
	pieceL_UP_R Piece = [7]byte{
		0b00100000,
		0b00100000,
		0b00100000,
		0b01100000,
	}
	pieceL_RT_R Piece = [7]byte{
		0b01000000,
		0b01111000,
	}
	pieceL_DW_R Piece = [7]byte{
		0b01100000,
		0b01000000,
		0b01000000,
		0b01000000,
	}
	pieceL_LT_R Piece = [7]byte{
		0b01111000,
		0b00001000,
	}
	pieceT_UP Piece = [7]byte{
		0b01000000,
		0b01000000,
		0b01100000,
		0b01000000,
	}
	pieceT_RT Piece = [7]byte{
		0b01111000,
		0b00100000,
	}
	pieceT_DW Piece = [7]byte{
		0b00100000,
		0b01100000,
		0b00100000,
		0b00100000,
	}
	pieceT_LT Piece = [7]byte{
		0b00010000,
		0b01111000,
	}
	pieceT_UP_R Piece = [7]byte{
		0b00100000,
		0b00100000,
		0b01100000,
		0b00100000,
	}
	pieceT_RT_R Piece = [7]byte{
		0b00100000,
		0b01111000,
	}
	pieceT_DW_R Piece = [7]byte{
		0b01000000,
		0b01100000,
		0b01000000,
		0b01000000,
	}
	pieceT_LT_R Piece = [7]byte{
		0b01111000,
		0b00010000,
	}
	pieceZ_UP Piece = [7]byte{
		0b01000000,
		0b01100000,
		0b00100000,
		0b00100000,
	}
	pieceZ_RT Piece = [7]byte{
		0b00011000,
		0b01110000,
	}
	pieceZ_DW Piece = [7]byte{
		0b01000000,
		0b01000000,
		0b01100000,
		0b00100000,
	}
	pieceZ_LT Piece = [7]byte{
		0b00111000,
		0b01100000,
	}
	pieceZ_UP_R Piece = [7]byte{
		0b00100000,
		0b01100000,
		0b01000000,
		0b01000000,
	}
	pieceZ_RT_R Piece = [7]byte{
		0b01110000,
		0b00011000,
	}
	pieceZ_DW_R Piece = [7]byte{
		0b00100000,
		0b00100000,
		0b01100000,
		0b01000000,
	}
	pieceZ_LT_R Piece = [7]byte{
		0b01100000,
		0b00111000,
	}
	pieceS_UP Piece = [7]byte{
		0b01100000,
		0b00100000,
		0b00110000,
	}
	pieceS_RT Piece = [7]byte{
		0b00010000,
		0b01110000,
		0b01000000,
	}
	pieceS_UP_R Piece = [7]byte{
		0b00110000,
		0b00100000,
		0b01100000,
	}
	pieceS_RT_R Piece = [7]byte{
		0b01000000,
		0b01110000,
		0b00010000,
	}
)

var (
	pieceMap = map[PieceKey]Piece{
		PieceKey{
			P: P_P,
			O: O_UP,
			R: false,
		}: pieceP_UP,
		PieceKey{
			P: P_P,
			O: O_RT,
			R: false,
		}: pieceP_RT,
		PieceKey{
			P: P_P,
			O: O_DW,
			R: false,
		}: pieceP_DW,
		PieceKey{
			P: P_P,
			O: O_LT,
			R: false,
		}: pieceP_LT,
		PieceKey{
			P: P_P,
			O: O_UP,
			R: true,
		}: pieceP_UP_R,
		PieceKey{
			P: P_P,
			O: O_RT,
			R: true,
		}: pieceP_RT_R,
		PieceKey{
			P: P_P,
			O: O_DW,
			R: true,
		}: pieceP_DW_R,
		PieceKey{
			P: P_P,
			O: O_LT,
			R: true,
		}: pieceP_LT_R,

		// O is symmetric, so has fewer options
		PieceKey{
			P: P_O,
			O: O_UP,
			R: false,
		}: pieceO_UP,
		PieceKey{
			P: P_O,
			O: O_RT,
			R: false,
		}: pieceO_RT,

		// C is symmetric, so has fewer options
		PieceKey{
			P: P_C,
			O: O_UP,
			R: false,
		}: pieceC_UP,
		PieceKey{
			P: P_C,
			O: O_RT,
			R: false,
		}: pieceC_RT,
		PieceKey{
			P: P_C,
			O: O_DW,
			R: false,
		}: pieceC_DW,
		PieceKey{
			P: P_C,
			O: O_LT,
			R: false,
		}: pieceC_LT,

		// V is symmetric, so has fewer options
		PieceKey{
			P: P_V,
			O: O_UP,
			R: false,
		}: pieceV_UP,
		PieceKey{
			P: P_V,
			O: O_RT,
			R: false,
		}: pieceV_RT,
		PieceKey{
			P: P_V,
			O: O_DW,
			R: false,
		}: pieceV_DW,
		PieceKey{
			P: P_V,
			O: O_LT,
			R: false,
		}: pieceV_LT,

		PieceKey{
			P: P_L,
			O: O_UP,
			R: false,
		}: pieceL_UP,
		PieceKey{
			P: P_L,
			O: O_RT,
			R: false,
		}: pieceL_RT,
		PieceKey{
			P: P_L,
			O: O_DW,
			R: false,
		}: pieceL_DW,
		PieceKey{
			P: P_L,
			O: O_LT,
			R: false,
		}: pieceL_LT,
		PieceKey{
			P: P_L,
			O: O_UP,
			R: true,
		}: pieceL_UP_R,
		PieceKey{
			P: P_L,
			O: O_RT,
			R: true,
		}: pieceL_RT_R,
		PieceKey{
			P: P_L,
			O: O_DW,
			R: true,
		}: pieceL_DW_R,
		PieceKey{
			P: P_L,
			O: O_LT,
			R: true,
		}: pieceL_LT_R,

		PieceKey{
			P: P_T,
			O: O_UP,
			R: false,
		}: pieceT_UP,
		PieceKey{
			P: P_T,
			O: O_RT,
			R: false,
		}: pieceT_RT,
		PieceKey{
			P: P_T,
			O: O_DW,
			R: false,
		}: pieceT_DW,
		PieceKey{
			P: P_T,
			O: O_LT,
			R: false,
		}: pieceT_LT,
		PieceKey{
			P: P_T,
			O: O_UP,
			R: true,
		}: pieceT_UP_R,
		PieceKey{
			P: P_T,
			O: O_RT,
			R: true,
		}: pieceT_RT,
		PieceKey{
			P: P_T,
			O: O_DW,
			R: true,
		}: pieceT_DW_R,
		PieceKey{
			P: P_T,
			O: O_LT,
			R: true,
		}: pieceT_LT_R,

		PieceKey{
			P: P_Z,
			O: O_UP,
			R: false,
		}: pieceZ_UP,
		PieceKey{
			P: P_Z,
			O: O_RT,
			R: false,
		}: pieceZ_RT,
		PieceKey{
			P: P_Z,
			O: O_DW,
			R: false,
		}: pieceZ_DW,
		PieceKey{
			P: P_Z,
			O: O_LT,
			R: false,
		}: pieceZ_LT,
		PieceKey{
			P: P_Z,
			O: O_UP,
			R: true,
		}: pieceZ_UP_R,
		PieceKey{
			P: P_Z,
			O: O_RT,
			R: true,
		}: pieceZ_RT_R,
		PieceKey{
			P: P_Z,
			O: O_DW,
			R: true,
		}: pieceZ_DW_R,
		PieceKey{
			P: P_Z,
			O: O_LT,
			R: true,
		}: pieceZ_LT_R,

		// S is symmetric, so has fewer options
		PieceKey{
			P: P_S,
			O: O_UP,
			R: false,
		}: pieceS_UP,
		PieceKey{
			P: P_S,
			O: O_RT,
			R: false,
		}: pieceS_RT,
		PieceKey{
			P: P_S,
			O: O_UP,
			R: true,
		}: pieceS_UP_R,
		PieceKey{
			P: P_S,
			O: O_RT,
			R: true,
		}: pieceS_RT_R,
	}
)

// MustNewPiece is used for testing only.  The same as NewPiece, except this panics if NewPiece returns false.
func MustNewPiece(pi PieceIndex, x, y uint8, oi Orientation, reflected bool) Piece {
	p, ok := NewPiece(pi, x, y, oi, reflected)
	if !ok {
		util.PrintAndExit1("invalid piece, piece = %s, x,y = %d,%d, oi = %+v, reflected = %t", GetPieceChar(pi), x, y, oi, reflected)
	}
	return p
}

// NewPiece returns pi at (x, y) with oi and ri.  Bool indicates if the position maybe valid (note it's one-sided, i.e., if false then it's definitely in valid); if invalid, then Piece will be empty.
func NewPiece(pi PieceIndex, x, y uint8, oi Orientation, reflected bool) (Piece, bool) {
	p, ok := pieceMap[PieceKey{
		P: pi,
		O: oi,
		R: reflected,
	}]
	if !ok {
		return Piece{}, false
	}

	moved := p

	if y > 0 {
		var back [7]byte
		for i := int8(6 - y); i >= 0; i-- {
			moved[i+int8(y)] = p[i]
			back[i] = p[i]
		}
		for i := uint8(0); i < y; i++ {
			moved[i] = 0
		}
		for i := uint8(0); i < 7; i++ {
			if p[i] != back[i] {
				return Piece{}, false
			}
		}
	}

	if x > 0 {
		before := moved
		for i := 0; i < 7; i++ {
			moved[i] = moved[i] >> x &^ emptyBoard[i]
			// Move back and check if the the values are equal.
			// If not, that means moved[i] was out of bounds.
			back := moved[i] << x
			if before[i] != back {
				return Piece{}, false
			}
		}
	}

	return moved, true
}

func (p Piece) String() string {
	var bf strings.Builder
	for i := 0; i < len(p); i++ {
		str := fmt.Sprintf("%08b", p[i])
		bf.WriteString(strings.Join(strings.Split(str, ""), " "))
		bf.WriteString("\n")
	}
	return bf.String()
}

func GetPieceChar(pi PieceIndex) string {
	switch pi {
	case P_P:
		return "p"
	case P_O:
		return "o"
	case P_C:
		return "c"
	case P_V:
		return "v"
	case P_L:
		return "l"
	case P_T:
		return "t"
	case P_Z:
		return "z"
	case P_S:
		return "s"
	default:
		return "x"
	}
}

func GetOrientation(o Orientation) string {
	switch o {
	case O_UP:
		return "up"
	case O_RT:
		return "right"
	case O_DW:
		return "down"
	default:
		return "left"
	}
}

func GetReflected(r bool) string {
	if r {
		return "reflected"
	}
	return "original"
}
