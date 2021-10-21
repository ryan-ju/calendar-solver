package calendarFast

import (
	"fmt"
	"strings"
	"testing"
)

func TestBoard_Next2(t *testing.T) {
	b := NewBoard(5, 24)
	nbs := b.NextFast()
	for i, nb := range nbs {
		fmt.Printf("index = %d, board = \n%s\n", i, strings.Join(nb.board, "\n"))
	}
}

//func TestBoard_Next2_last(t *testing.T) {
//	b := Configure(5, 24, []Action{
//		{
//			Piece: PieceL,
//			X:     0,
//			Y:     0,
//		},
//		{
//			Piece: PieceO,
//			X:     1,
//			Y:     0,
//		},
//		{
//			Piece:  PieceC,
//			Rotate: 3,
//			X:      3,
//			Y:      0,
//		},
//		{
//			Piece:  PieceT,
//			Flip:   true,
//			Rotate: 3,
//			X:      3,
//			Y:      2,
//		},
//		{
//			Piece:  PieceZ,
//			Rotate: 1,
//			X:      0,
//			Y:      3,
//		},
//		{
//			Piece:  PieceS,
//			Rotate: 1,
//			X:      3,
//			Y:      3,
//		},
//		{
//			Piece:  PieceV,
//			Rotate: 3,
//			X:      4,
//			Y:      3,
//		},
//	})
//	fmt.Println(strings.Join(b.board, "\n"))
//
//	nbs := b.NextFast()
//	for i, nb := range nbs {
//		fmt.Printf("index = %d, board = \n%s\n", i, strings.Join(nb.board, "\n"))
//	}
//}

type Action struct {
	Piece  RawPiece
	Flip   bool
	Rotate int
	X      int
	Y      int
}

func Configure(month, day int, actions []Action) *Board {
	b := NewBoard(month, day)
	for _, action := range actions {
		pi := PieceToIndex(action.Piece)
		mask := Piece{
			Raw:  action.Piece,
			Flat: strings.Join(action.Piece, "\n"),
		}
		if action.Flip {
			mask = Flip(mask)
		}
		for i := 0; i < action.Rotate; i++ {
			mask = Rotate(mask)
		}
		newRows := make([]string, LenRows, LenRows)
		copy(newRows, b.board)
		height := len(mask.Raw)
		width := len(mask.Raw[0])
		mi := MaskToIndex(mask)
		indent := MaskIndent[mi]
		for j := 0; j < height; j++ {
			for i := 0; i < width; i++ {
				if mask.Raw[j][i:i+1] == "x" {
					boardX := action.X + i - indent
					newRows[action.Y+j] = newRows[action.Y+j][:boardX] + PieceToChar[pi] + newRows[action.Y+j][boardX+1:]
				}
			}
		}
		b = &Board{
			Prev:       b,
			PieceIndex: pi,
			MaskIndex:  mi,
			CellIndex:  0,
			PiecesUsed: append([]int{pi}, b.PiecesUsed...),
			board:      newRows,
			flat:       strings.Join(newRows, "\n"),
		}
	}
	return b
}

func PieceToIndex(input RawPiece) int {
	flat := strings.Join(input, "\n")
	for i, p := range Pieces {
		if p.Flat == flat {
			return i
		}
	}
	panic("cannot find piece index")
}

func MaskToIndex(input Piece) int {
	flat := strings.Join(input.Raw, "\n")
	for i, m := range Masks {
		if m.Flat == flat {
			return i
		}
	}
	panic("cannot find mask index")
}
