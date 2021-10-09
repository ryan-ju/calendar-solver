package calendar

import (
	"math/rand"

	"github.com/ryan-ju/calendar-solver/util"
)

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

func (s *Solver) Solve() *Board {
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

		if counter%500000 == 0 {
			util.Log(util.LevelInfo, "Tried %d times, last board = \n%s\n", counter, b.String())
		}
	}
	return nil
}

func NextBoards(b Board) []*Board {
	var result []*Board
	unuseds := b.GetUnusedPieces()
	if len(unuseds) > 0 {
		for _, unused := range unuseds {
			for _, fc := range b.FreeCells {
				// UP is easy, as the top-left corner is none empty
				nb := b.AddPiece(PieceKey{
					P: unused,
					O: O_UP,
					X: fc.X,
					Y: fc.Y,
					R: false,
				})
				if nb != nil {
					result = append(result, nb)
				}
				// RIGHT, the top-left corner may be empty for some pieces
				nb = nil
				switch unused {
				case P_Z, P_S:
					if fc.Y >= 1 {
						nb = b.AddPiece(PieceKey{
							P: unused,
							O: O_RT,
							X: fc.X,
							Y: fc.Y - 1,
							R: false,
						})
					}
				default:
					nb = b.AddPiece(PieceKey{
						P: unused,
						O: O_RT,
						X: fc.X,
						Y: fc.Y,
						R: false,
					})
				}
				if nb != nil {
					result = append(result, nb)
				}
				// DOWN, the top-left corner may be empty for some pieces
				nb = nil
				switch unused {
				case P_P, P_T:
					if fc.Y >= 1 {
						nb = b.AddPiece(PieceKey{
							P: unused,
							O: O_DW,
							X: fc.X,
							Y: fc.Y - 1,
							R: false,
						})
					}
				default:
					nb = b.AddPiece(PieceKey{
						P: unused,
						O: O_DW,
						X: fc.X,
						Y: fc.Y,
						R: false,
					})
				}
				if nb != nil {
					result = append(result, nb)
				}
				// LEFT, the top-left corner may be empty for some pieces
				nb = nil
				switch unused {
				case P_V:
					if fc.Y >= 2 {
						nb = b.AddPiece(PieceKey{
							P: unused,
							O: O_LT,
							X: fc.X,
							Y: fc.Y - 2,
							R: false,
						})
					}
				case P_L, P_T, P_Z, P_S:
					if fc.Y >= 1 {
						nb = b.AddPiece(PieceKey{
							P: unused,
							O: O_LT,
							X: fc.X,
							Y: fc.Y - 1,
							R: false,
						})
					}
				default:
					nb = b.AddPiece(PieceKey{
						P: unused,
						O: O_LT,
						X: fc.X,
						Y: fc.Y,
						R: false,
					})
				}
				if nb != nil {
					result = append(result, nb)
				}

				// ===== Reflected =====
				// UP
				nb = nil
				switch unused {
				case P_V:
					if fc.Y >= 2 {
						nb = b.AddPiece(PieceKey{
							P: unused,
							O: O_UP,
							X: fc.X,
							Y: fc.Y - 2,
							R: true,
						})
					}
				case P_L, P_T, P_Z, P_S:
					if fc.X >= 1 {
						nb = b.AddPiece(PieceKey{
							P: unused,
							O: O_UP,
							X: fc.X - 1,
							Y: fc.Y,
							R: true,
						})
					}
				default:
					nb = b.AddPiece(PieceKey{
						P: unused,
						O: O_UP,
						X: fc.X,
						Y: fc.Y,
						R: true,
					})
				}
				if nb != nil {
					result = append(result, nb)
				}
				// RIGHT
				nb = nil
				switch unused {
				case P_P, P_T:
					if fc.X >= 1 {
						nb = b.AddPiece(PieceKey{
							P: unused,
							O: O_RT,
							X: fc.X - 1,
							Y: fc.Y,
							R: true,
						})
					}
				default:
					nb = b.AddPiece(PieceKey{
						P: unused,
						O: O_RT,
						X: fc.X,
						Y: fc.Y,
						R: true,
					})
				}
				if nb != nil {
					result = append(result, nb)
				}
				// DOWN
				nb = nil
				switch unused {
				case P_Z, P_S:
					if fc.X >= 1 {
						nb = b.AddPiece(PieceKey{
							P: unused,
							O: O_DW,
							X: fc.X - 1,
							Y: fc.Y,
							R: true,
						})
					}
				default:
					nb = b.AddPiece(PieceKey{
						P: unused,
						O: O_DW,
						X: fc.X,
						Y: fc.Y,
						R: true,
					})
				}
				if nb != nil {
					result = append(result, nb)
				}
				// LEFT
				nb = b.AddPiece(PieceKey{
					P: unused,
					O: O_LT,
					X: fc.X,
					Y: fc.Y,
					R: true,
				})
				if nb != nil {
					result = append(result, nb)
				}
			}
		}
	}
	rand.Shuffle(len(result), func(i, j int) { result[i], result[j] = result[j], result[i] })
	return result
}
