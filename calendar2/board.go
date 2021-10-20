package calendar2

import (
	"fmt"
	"strings"
)

const (
	LenCells = 49
	LenRows  = 7
	LenCols  = 7
)

var (
	emptyBoard = []string{
		"......x",
		"......x",
		".......",
		".......",
		".......",
		".......",
		"...xxxx",
	}
)

type Board struct {
	Prev       *Board
	PieceIndex int
	MaskIndex  int
	CellIndex  int
	board      []string
	flat       string
}

func NewBoard(month, day int) *Board {
	monthX := (month - 1) % 6
	monthY := (month - 1) / 6
	dayX := (day - 1) % 7
	dayY := (day-1)/7 + 2

	var board []string
	for y, row := range emptyBoard {
		if y == monthY {
			board = append(board, row[:monthX]+"x"+row[monthX+1:])
		} else if y == dayY {
			board = append(board, row[:dayX]+"x"+row[dayX+1:])
		} else {
			board = append(board, row)
		}
	}

	return &Board{
		board: board,
		flat:  strings.Join(board, "\n"),
	}
}

func (b *Board) IsSolved() bool {
	return b.PieceIndex == LenPieces
}

func (b *Board) Next() []*Board {
	if b.PieceIndex >= LenPieces {
		return nil
	}
	var result []*Board
	for maskIndex := PieceToMaskFirst[b.PieceIndex]; maskIndex < PieceToMaskFirst[b.PieceIndex+1]; maskIndex++ {
		mask := Masks[maskIndex]
		indent := MaskIndent[maskIndex]
		height := len(mask.Raw)
		width := len(mask.Raw[0])
		for cellIndex := 0; cellIndex < LenCells; cellIndex++ {
			row := cellIndex / LenCols
			col := cellIndex % LenCols
			if b.board[row][col:col+1] != "." {
				continue
			}

			if row+height > LenRows {
				// No more to look at
				break
			}

			if col+width-indent > LenCols {
				continue
			}

			newRows := make([]string, LenRows, LenRows)
			copy(newRows, b.board)
			valid := true
			for x := 0; x < width && valid; x++ {
				boardX := col + x - indent
				if boardX < 0 {
					valid = false
					continue
				}
				for y := 0; y < height && valid; y++ {
					boardY := row + y
					isBoardFilled := b.board[boardY][boardX:boardX+1] != "."
					isMaskFilled := mask.Raw[y][x:x+1] != "."
					valid = valid && !(isBoardFilled && isMaskFilled)
					if isMaskFilled {
						newRows[boardY] = newRows[boardY][:boardX] + PieceToChar[MaskToPiece[maskIndex]] + newRows[boardY][boardX+1:]
					}
				}
			}

			if valid {
				nb := &Board{
					Prev:       b,
					PieceIndex: b.PieceIndex + 1,
					MaskIndex:  maskIndex,
					CellIndex:  cellIndex,
					board:      newRows,
					flat:       strings.Join(newRows, "\n"),
				}
				if b.PieceIndex == 7 {
					fmt.Printf("new board = \n%s\nmask = \n%s\n", strings.Join(newRows, "\n"), mask.Flat)
				}

				result = append(result, nb)
			}
		}
	}
	return result
}

func (b *Board) Solution() string {
	type PiecePlacement struct {
		MaskIndex int
		CellIndex int
	}

	var solution []PiecePlacement
	for tb := b; tb != nil; tb = tb.Prev {
		solution = append(solution, PiecePlacement{
			MaskIndex: tb.MaskIndex,
			CellIndex: tb.CellIndex,
		})
	}

	var field [LenRows][LenCols]string
	for _, s := range solution {
		mask := Masks[s.MaskIndex]
		height := len(mask.Raw)
		width := len(mask.Raw[0])
		x := (s.CellIndex % LenCols) - MaskIndent[s.MaskIndex]
		y := s.CellIndex / LenCols
		c := PieceToChar[MaskToPiece[s.MaskIndex]]
		for j := 0; j < height; j++ {
			for i := 0; i < width; i++ {
				if mask.Raw[j][i:i+1] != "." {
					field[y+j][x+i] = c
				}
			}
		}
	}

	var sb strings.Builder
	for _, y := range field {
		for _, x := range y {
			if x == "" {
				x = "."
			}
			sb.WriteString("| " + x + " ")
		}
		sb.WriteString("|\n")
	}
	sb.WriteString("\n")
	return sb.String()
}
