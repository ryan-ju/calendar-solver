package calendar

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/ryan-ju/calendar-solver/util"
)

func getExampleBoard() *Board {
	board, err := NewBoard(Date{
		Month: 10,
		Day:   28,
	})
	util.OnErrorExit1(err, "cannot create board, bug in test")

	board.FreeCells = []ShortIndex{
		{
			X: 1,
			Y: 3,
		},
		{
			X: 1,
			Y: 4,
		},
		{
			X: 1,
			Y: 5,
		},
		{
			X: 2,
			Y: 3,
		},
		{
			X: 3,
			Y: 3,
		},
		{
			X: 4,
			Y: 3,
		},
		{
			X: 5,
			Y: 3,
		},
		{
			X: 4,
			Y: 4,
		},
		{
			X: 5,
			Y: 4,
		},
		{
			X: 4,
			Y: 5,
		},
		{
			X: 5,
			Y: 5,
		},
	}
	board.PiecesUsed = byte(P_T) | byte(P_C) | byte(P_L) | byte(P_P) | byte(P_Z) | byte(P_S)
	board.Pieces = []PieceHolder{
		{
			Key: PieceKey{
				P: P_T,
				O: O_RT,
				R: false,
				X: 0,
				Y: 0,
			},
			Piece: MustNewPiece(P_T, 0, 0, O_RT, false),
		},
		{
			Key: PieceKey{
				P: P_C,
				O: O_LT,
				R: false,
				X: 0,
				Y: 1,
			},
			Piece: MustNewPiece(P_C, 0, 1, O_LT, false),
		},
		{
			Key: PieceKey{
				P: P_L,
				O: O_UP,
				R: false,
				X: 0,
				Y: 3,
			},
			Piece: MustNewPiece(P_L, 0, 3, O_UP, false),
		},
		{
			Key: PieceKey{
				P: P_S,
				O: O_UP,
				R: true,
				X: 3,
				Y: 0,
			},
			Piece: MustNewPiece(P_S, 3, 0, O_UP, true),
		},
		{
			Key: PieceKey{
				P: P_Z,
				O: O_UP,
				R: false,
				X: 5,
				Y: 1,
			},
			Piece: MustNewPiece(P_Z, 5, 1, O_UP, false),
		},
		{
			Key: PieceKey{
				P: P_P,
				O: O_UP,
				R: false,
				X: 2,
				Y: 4,
			},
			Piece: MustNewPiece(P_P, 2, 4, O_UP, false),
		},
	}
	board.Cells = [7]byte{
		0b01111111,
		0b01111111,
		0b01111111,
		0b01000001,
		0b01011001,
		0b01011001,
		0b01111111,
	}

	return board
}

func getDebugBoard() *Board {
	board, err := NewBoard(Date{
		Month: 10,
		Day:   28,
	})
	util.OnErrorExit1(err, "cannot create board, bug in test")

	board.FreeCells = []ShortIndex{
		{
			X: 3,
			Y: 2,
		},
		{
			X: 0,
			Y: 3,
		},
		{
			X: 1,
			Y: 3,
		},
		{
			X: 2,
			Y: 3,
		},
		{
			X: 3,
			Y: 3,
		},
		{
			X: 0,
			Y: 4,
		},
		{
			X: 1,
			Y: 4,
		},
		{
			X: 2,
			Y: 4,
		},
		{
			X: 3,
			Y: 4,
		},
		{
			X: 4,
			Y: 4,
		},
		{
			X: 5,
			Y: 4,
		},
		{
			X: 6,
			Y: 4,
		},
		{
			X: 0,
			Y: 5,
		},
		{
			X: 1,
			Y: 5,
		},
		{
			X: 2,
			Y: 5,
		},
		{
			X: 3,
			Y: 5,
		},
		{
			X: 4,
			Y: 5,
		},
		{
			X: 5,
			Y: 5,
		},
		{
			X: 0,
			Y: 6,
		},
		{
			X: 1,
			Y: 6,
		},
		{
			X: 2,
			Y: 6,
		},
	}
	board.PiecesUsed = byte(P_Z) | byte(P_L) | byte(P_P) | byte(P_S)
	board.Pieces = []PieceHolder{
		{
			Key: PieceKey{
				P: P_Z,
				O: O_DW,
				R: false,
				X: 5,
				Y: 0,
			},
			Piece: MustNewPiece(P_Z, 5, 0, O_DW, false),
		},
		{
			Key: PieceKey{
				P: P_L,
				O: O_UP,
				R: false,
				X: 4,
				Y: 1,
			},
			Piece: MustNewPiece(P_L, 4, 0, O_UP, false),
		},
		{
			Key: PieceKey{
				P: P_P,
				O: O_UP,
				R: false,
				X: 0,
				Y: 3,
			},
			Piece: MustNewPiece(P_P, 0, 0, O_UP, false),
		},
		{
			Key: PieceKey{
				P: P_S,
				O: O_UP,
				R: true,
				X: 1,
				Y: 0,
			},
			Piece: MustNewPiece(P_S, 1, 0, O_UP, true),
		},
	}
	board.Cells = [7]byte{
		0b01110111,
		0b01111111,
		0b01011111,
		0b00010111,
		0b00000000,
		0b00000001,
		0b00001111,
	}

	return board
}

func TestNextBoards(t *testing.T) {
	boards := NextBoards(*getDebugBoard())
	for _, b := range boards {
		fmt.Println(b.String())
	}
}

func TestSolver_Solve_Partial(t *testing.T) {
	g := NewGomegaWithT(t)
	solver := Solver{
		Target: Date{
			Month: 10,
			Day:   28,
		},
		Board: getExampleBoard(),
	}
	solution := solver.Solve()
	g.Expect(solution).ToNot(BeNil())
	fmt.Println(solution.String())
}

func TestSolver_Solve_New(t *testing.T) {
	g := NewGomegaWithT(t)
	solver, err := NewSolver(Date{
		//Month: 10,
		//Day:   28,
		Month: 8,
		Day:   14,
	})
	g.Expect(err).ToNot(HaveOccurred())
	solution := solver.Solve()
	g.Expect(solution).ToNot(BeNil())
	fmt.Println(solution.SolutionPath())
}
