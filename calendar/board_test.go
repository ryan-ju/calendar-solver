package calendar

//
//import (
//	"fmt"
//	"testing"
//
//	"github.com/ryan-ju/calendar-solver/util"
//
//	. "github.com/onsi/gomega"
//)
//
//func TestBoard_ToString(t *testing.T) {
//	g := NewGomegaWithT(t)
//
//	p1, ok := NewPiece(P_P, 3, 3, O_UP, false)
//	g.Expect(ok).To(BeTrue())
//
//	p2, ok := NewPiece(P_L, 0, 3, O_UP, true)
//	g.Expect(ok).To(BeTrue())
//
//	board := Board{
//		Parent: nil,
//		Date: Date{
//			Month: 6,
//			Day:   18,
//		},
//		DateCoordinates: []ShortIndex{
//			{
//				X: 3,
//				Y: 0,
//			},
//			{
//				X: 5,
//				Y: 5,
//			},
//		},
//		Pieces: []PieceHolder{
//			{
//				Key: PieceKey{
//					P: P_P,
//					O: O_UP,
//					X: 3,
//					Y: 3,
//				},
//				Piece: p1,
//			},
//			{
//				Key: PieceKey{
//					P: P_L,
//					O: O_UP,
//					X: 0,
//					Y: 3,
//					R: true,
//				},
//				Piece: p2,
//			},
//		},
//	}
//
//	fmt.Printf("%s\n", board.String())
//}
//
//func TestBoard_GetUnusedPieces(t *testing.T) {
//	b := Board{
//		PiecesUsed: byte(P_P) | byte(P_T),
//	}
//	for _, u := range b.GetUnusedPieces() {
//		fmt.Printf("%08b\n", u)
//	}
//}
//
//func TestBoard_AddPiece(t *testing.T) {
//	b := getDebugBoard2()
//	nbs := NextBoards(*b)
//	for _, nb := range nbs {
//		fmt.Println(nb.String())
//	}
//}
//
///**
//Target date: 10-28
//MM|   |   |   |   |   |   |
//MM|   |   |   | @ |   |   |
//01|   |   |   |   | t |   |   |
//08|   |   | t | t | t | t |   |
//15|   | z | z | z | s |   |   |
//22| z | z | s | s | s |   | @ |
//29|   |   | s |
//Board cells
//0 0 0 0 0 0 0 1
//0 0 0 0 1 0 0 1
//0 0 0 0 0 1 0 0
//0 0 0 1 1 1 1 0
//0 0 1 1 1 1 0 0
//0 1 1 1 1 1 0 1
//0 0 0 1 1 1 1 1
//Pieces:
//s, original, right, (2,4)
//z, original, left, (0,4)
//t, original, left, (2,2)
//Pieces used: 11100000
//*/
//func getDebugBoard2() *Board {
//	board, err := NewBoard(Date{
//		Month: 10,
//		Day:   28,
//	})
//	util.OnErrorExit1(err, "cannot create board, bug in test")
//
//	board.FreeCells = GetShortIndexes([]uint8{
//		0, 0,
//		1, 0,
//		2, 0,
//		3, 0,
//		4, 0,
//		5, 0,
//		0, 1,
//		1, 1,
//		2, 1,
//		4, 1,
//		5, 1,
//		0, 2,
//		1, 2,
//		2, 2,
//		3, 2,
//		5, 2,
//		6, 2,
//		0, 3,
//		1, 3,
//		6, 3,
//		0, 4,
//		5, 4,
//		6, 4,
//		5, 5,
//		0, 7,
//		1, 7,
//	})
//	board.Pieces = []PieceHolder{
//		{
//			Key: PieceKey{
//				P: P_Z,
//				O: O_LT,
//				R: false,
//				X: 0,
//				Y: 4,
//			},
//			Piece: MustNewPiece(P_Z, 0, 4, O_LT, false),
//		},
//		{
//			Key: PieceKey{
//				P: P_T,
//				O: O_LT,
//				R: false,
//				X: 2,
//				Y: 2,
//			},
//			Piece: MustNewPiece(P_T, 2, 2, O_LT, false),
//		},
//		{
//			Key: PieceKey{
//				P: P_S,
//				O: O_RT,
//				R: false,
//				X: 2,
//				Y: 4,
//			},
//			Piece: MustNewPiece(P_S, 2, 4, O_RT, false),
//		},
//	}
//	board.Cells = [7]byte{
//		0b00000001,
//		0b00001001,
//		0b00000100,
//		0b00011110,
//		0b00111100,
//		0b01111101,
//		0b00011111,
//	}
//
//	return board
//}
