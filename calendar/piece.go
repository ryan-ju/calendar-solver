package calendar

import (
	"fmt"
	"strings"
)

type PieceIndex int
type Orientation int
type Piece [7]byte
type PieceKey struct {
	P PieceIndex
	O Orientation
	R bool
}

const (
	// **
	// **
	// *
	P_P PieceIndex = 0
	// **
	// **
	// **
	P_O PieceIndex = 1
	// **
	// *
	// **
	P_C PieceIndex = 2
	// *
	// *
	// ***
	P_V PieceIndex = 3
	// *
	// *
	// *
	// **
	P_L PieceIndex = 4
	// *
	// *
	// **
	// *
	P_T PieceIndex = 5
	// *
	// **
	//  *
	//  *
	P_Z PieceIndex = 6
	// **
	//  *
	//  **
	P_S PieceIndex = 7

	O_UP Orientation = 0
	O_RT Orientation = 1
	O_DW Orientation = 2
	O_LT Orientation = 3
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
		0b011100000,
		0b011100000,
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

// NewPiece returns pi at (x, y) with oi and ri.  Bool indicates if the position maybe valid (note it's one-sided, i.e., if false then it's definitely in valid); if invalid, then Piece will be empty.
func NewPiece(pi PieceIndex, x, y int, oi Orientation, reflected bool) (Piece, bool) {
	p, ok := pieceMap[PieceKey{
		P: pi,
		O: oi,
		R: reflected,
	}]
	if !ok {
		return Piece{}, false
	}

	// Check bounding box
	switch pi {
	case P_L, P_T, P_Z:
		switch oi {
		case O_UP, O_DW:
			if y > 3 {
				return Piece{}, false
			}
		default:
			if x > 3 {
				return Piece{}, false
			}
		}
	default:
		switch oi {
		case O_UP, O_DW:
			if y > 4 {
				return Piece{}, false
			}
		default:
			if x > 4 {
				return Piece{}, false
			}
		}
	}

	switch pi {
	case P_V:
		if x > 4 || y > 4 {
			return Piece{}, false
		}
	default:
		switch oi {
		case O_UP, O_DW:
			if x > 5 {
				return Piece{}, false
			}
		default:
			if y > 5 {
				return Piece{}, false
			}
		}
	}

	if y > 0 {
		for i := 6 - y; i >= 0; i-- {
			p[i+y] = p[i]
		}
		for i := 0; i < y; i++ {
			p[i] = 0
		}
	}

	if x > 0 {
		for i := 0; i < 7; i++ {
			p[i] = p[i] >> x
		}
	}

	return p, true
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
