package calendarSlow

import "github.com/ryan-ju/calendar-solver/util"

type Solver struct {
	Target Date
	Board  *Board
}

func NewSolver(target Date) (*Solver, error) {
	b, err := NewBoard(target)
	if err != nil {
		return nil, err
	}
	return &Solver{
		Target: target,
		Board:  b,
	}, nil
}

func (s *Solver) SolveOne() *Board {
	stack := NewStack()
	stack.Push(s.Board)

	var counter int64
	for !stack.IsEmpty() {
		b := stack.Pop()

		nbs := NextBoards(*b)
		for _, nb := range nbs {
			if nb.IsSolved() {
				return nb
			}
		}
		stack.Push(nbs...)
		counter++
	}
	util.Log(util.LevelInfo, "searched times: %d", counter)
	return nil
}

func (s *Solver) Solve() []*Board {
	var solutions []*Board
	stack := NewStack()
	stack.Push(s.Board)

	var counter int64
	for !stack.IsEmpty() {
		b := stack.Pop()

		nbs := NextBoards(*b)
		for _, nb := range nbs {
			if nb.IsSolved() {
				solutions = append(solutions, nb)
			}
		}
		stack.Push(nbs...)
		counter++
	}
	util.Log(util.LevelInfo, "searched times: %d", counter)
	return solutions
}

func NextBoards(b Board) []*Board {
	var result []*Board
	for _, fc := range b.FreeCells {
		pieceIndex := pieceIndexes[b.PieceIndex]
		result = append(result, NextBoardsHelper(b, pieceIndex, fc)...)
	}
	return result
}

func NextBoardsHelper(b Board, pi PieceIndex, cell ShortIndex) []*Board {
	var result []*Board
	// UP is easy, as the top-left corner is none empty
	nb := b.AddPiece(PieceKey{
		P: pi,
		O: O_UP,
		X: cell.X,
		Y: cell.Y,
		R: false,
	})
	if nb != nil {
		result = append(result, nb)
	}
	// RIGHT, the top-left corner may be empty for some pieces
	nb = nil
	switch pi {
	case P_Z, P_S:
		if cell.Y >= 1 {
			nb = b.AddPiece(PieceKey{
				P: pi,
				O: O_RT,
				X: cell.X,
				Y: cell.Y - 1,
				R: false,
			})
		}
	default:
		nb = b.AddPiece(PieceKey{
			P: pi,
			O: O_RT,
			X: cell.X,
			Y: cell.Y,
			R: false,
		})
	}
	if nb != nil {
		result = append(result, nb)
	}
	// DOWN, the top-left corner may be empty for some pieces
	nb = nil
	switch pi {
	case P_P, P_T:
		if cell.Y >= 1 {
			nb = b.AddPiece(PieceKey{
				P: pi,
				O: O_DW,
				X: cell.X,
				Y: cell.Y - 1,
				R: false,
			})
		}
	default:
		nb = b.AddPiece(PieceKey{
			P: pi,
			O: O_DW,
			X: cell.X,
			Y: cell.Y,
			R: false,
		})
	}
	if nb != nil {
		result = append(result, nb)
	}
	// LEFT, the top-left corner may be empty for some pieces
	nb = nil
	switch pi {
	case P_V:
		if cell.Y >= 2 {
			nb = b.AddPiece(PieceKey{
				P: pi,
				O: O_LT,
				X: cell.X,
				Y: cell.Y - 2,
				R: false,
			})
		}
	case P_L, P_T, P_Z, P_S:
		if cell.Y >= 1 {
			nb = b.AddPiece(PieceKey{
				P: pi,
				O: O_LT,
				X: cell.X,
				Y: cell.Y - 1,
				R: false,
			})
		}
	default:
		nb = b.AddPiece(PieceKey{
			P: pi,
			O: O_LT,
			X: cell.X,
			Y: cell.Y,
			R: false,
		})
	}
	if nb != nil {
		result = append(result, nb)
	}

	// ===== Reflected =====
	// UP
	nb = nil
	switch pi {
	case P_V:
		if cell.Y >= 2 {
			nb = b.AddPiece(PieceKey{
				P: pi,
				O: O_UP,
				X: cell.X,
				Y: cell.Y - 2,
				R: true,
			})
		}
	case P_L, P_T, P_Z, P_S:
		if cell.X >= 1 {
			nb = b.AddPiece(PieceKey{
				P: pi,
				O: O_UP,
				X: cell.X - 1,
				Y: cell.Y,
				R: true,
			})
		}
	default:
		nb = b.AddPiece(PieceKey{
			P: pi,
			O: O_UP,
			X: cell.X,
			Y: cell.Y,
			R: true,
		})
	}
	if nb != nil {
		result = append(result, nb)
	}
	// RIGHT
	nb = nil
	switch pi {
	case P_P, P_T:
		if cell.X >= 1 {
			nb = b.AddPiece(PieceKey{
				P: pi,
				O: O_RT,
				X: cell.X - 1,
				Y: cell.Y,
				R: true,
			})
		}
	default:
		nb = b.AddPiece(PieceKey{
			P: pi,
			O: O_RT,
			X: cell.X,
			Y: cell.Y,
			R: true,
		})
	}
	if nb != nil {
		result = append(result, nb)
	}
	// DOWN
	nb = nil
	switch pi {
	case P_Z, P_S:
		if cell.X >= 1 {
			nb = b.AddPiece(PieceKey{
				P: pi,
				O: O_DW,
				X: cell.X - 1,
				Y: cell.Y,
				R: true,
			})
		}
	default:
		nb = b.AddPiece(PieceKey{
			P: pi,
			O: O_DW,
			X: cell.X,
			Y: cell.Y,
			R: true,
		})
	}
	if nb != nil {
		result = append(result, nb)
	}
	// LEFT
	nb = b.AddPiece(PieceKey{
		P: pi,
		O: O_LT,
		X: cell.X,
		Y: cell.Y,
		R: true,
	})
	if nb != nil {
		result = append(result, nb)
	}
	return result
}
