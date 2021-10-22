/*
File to be copied into https://goplay.space/
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Change this to date you want, format mmdd
	Solve("1010")
}

func Solve(date string) {
	month, _ := strconv.Atoi(date[:2])
	day, _ := strconv.Atoi(date[2:])

	var solutions []string
	stack := NewStack()
	board := NewBoard(month, day)
	stack.Push(board)
	now := time.Now()

	var counter int64
	for !stack.IsEmpty() {
		b := stack.Pop()

		nbs := b.Next()
		for _, nb := range nbs {
			if nb.IsSolved() {
				solutions = append(solutions, nb.Solution())
			} else {
				stack.Push(nb)
			}
		}
		counter++
	}
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("searched times: %d, duration = %dms, #solutions = %d\n", counter, time.Since(now)/time.Millisecond, len(solutions)))
	sb.WriteString(fmt.Sprintf("solutions = \n%s\n", strings.Join(solutions, "")))
	fmt.Println(sb.String())
}

// Pieces

type RawPiece []string
type Piece struct {
	Raw  RawPiece
	Flat string
}

var (
	PieceP = []string{
		"xx",
		"xx",
		"x.",
	}
	PieceC = []string{
		"xx",
		"x.",
		"xx",
	}
	PieceL = []string{
		"x.",
		"x.",
		"x.",
		"xx",
	}
	PieceV = []string{
		"x..",
		"x..",
		"xxx",
	}
	PieceO = []string{
		"xx",
		"xx",
		"xx",
	}
	PieceT = []string{
		"x.",
		"xx",
		"x.",
		"x.",
	}
	PieceZ = []string{
		"x.",
		"xx",
		".x",
		".x",
	}
	PieceS = []string{
		"xx.",
		".x.",
		".xx",
	}
	RawPieces = []RawPiece{
		PieceP,
		PieceC,
		PieceL,
		PieceV,
		PieceO,
		PieceT,
		PieceZ,
		PieceS,
	}
)

var (
	Pieces           []Piece
	Masks            []Piece
	MaskToPiece      map[int]int
	PieceToMaskFirst map[int]int
	MaskIndent       map[int]int
	LenMasks         int
	LenPieces        = len(RawPieces)
	PieceToChar      = map[int]string{
		0: "P",
		1: "C",
		2: "L",
		3: "V",
		4: "O",
		5: "T",
		6: "Z",
		7: "S",
	}
)

func init() {
	MaskToPiece = map[int]int{}
	PieceToMaskFirst = map[int]int{}
	MaskIndent = map[int]int{}
	var counter int
	for i, p := range RawPieces {
		PieceToMaskFirst[i] = counter
		piece := Piece{
			Raw:  p,
			Flat: strings.Join(p, "\n"),
		}
		Pieces = append(Pieces, piece)
		dedupe := map[string]Piece{}
		tmp := piece
		flipped := Flip(piece)
		for i := 0; i < 4; i++ {
			dedupe[tmp.Flat] = tmp
			dedupe[flipped.Flat] = flipped
			tmp = Rotate(tmp)
			flipped = Rotate(flipped)
		}
		for _, v := range dedupe {
			MaskToPiece[counter] = i
			MaskIndent[counter] = strings.Index(v.Raw[0], "x")
			Masks = append(Masks, v)
			counter++
		}
	}
	PieceToMaskFirst[len(RawPieces)] = counter
	LenMasks = len(Masks)
}

func Flip(p Piece) Piece {
	nr := len(p.Raw)
	nc := len(p.Raw[0])
	var shape []string
	for y := 0; y < nr; y++ {
		var row []string
		for x := nc - 1; x >= 0; x-- {
			row = append(row, p.Raw[y][x:x+1])
		}
		shape = append(shape, strings.Join(row, ""))
	}
	return Piece{
		Raw:  shape,
		Flat: strings.Join(shape, "\n"),
	}
}

func Rotate(p Piece) Piece {
	nr := len(p.Raw)
	nc := len(p.Raw[0])
	var shape []string
	for x := 0; x < nc; x++ {
		var row []string
		for y := nr - 1; y >= 0; y-- {
			row = append(row, p.Raw[y][x:x+1])
		}
		shape = append(shape, strings.Join(row, ""))
	}
	return Piece{
		Raw:  shape,
		Flat: strings.Join(shape, "\n"),
	}
}

// Board

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
	PiecesUsed []int
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
	return b.PieceIndex == LenPieces || len(b.PiecesUsed) == LenPieces
}

func (b *Board) Next() []*Board {
	var result []*Board
	row := b.CellIndex / LenCols
	col := b.CellIndex % LenCols
	for pi := 0; pi < LenPieces; pi++ {
		var used bool
		for _, pu := range b.PiecesUsed {
			if pu == pi {
				used = true
				break
			}
		}
		if used {
			continue
		}

		for maskIndex := PieceToMaskFirst[pi]; maskIndex < PieceToMaskFirst[pi+1]; maskIndex++ {
			mask := Masks[maskIndex]
			indent := MaskIndent[maskIndex]
			height := len(mask.Raw)
			width := len(mask.Raw[0])

			if row+height > LenRows {
				// No more to look at
				continue
			}

			if col < indent {
				continue
			}

			if col+width-indent > LenCols {
				continue
			}

			newRows := make([]string, LenRows, LenRows)
			copy(newRows, b.board)
			valid := true
			for x := 0; x < width && valid; x++ {
				boardX := col + x - indent
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
				var cellIndex int
				for cellIndex = b.CellIndex; cellIndex < LenCells; cellIndex++ {
					tr := cellIndex / LenCols
					tc := cellIndex % LenCols
					if newRows[tr][tc:tc+1] == "." {
						break
					}
				}

				nb := &Board{
					Prev:       b,
					MaskIndex:  maskIndex,
					CellIndex:  cellIndex,
					PiecesUsed: append([]int{pi}, b.PiecesUsed...), // This way to avoid modifying b.PiecesUsed
					board:      newRows,
					flat:       strings.Join(newRows, "\n"),
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
	maskIndex := b.MaskIndex
	for tb := b.Prev; tb != nil; tb = tb.Prev {
		solution = append(solution, PiecePlacement{
			MaskIndex: maskIndex,
			CellIndex: tb.CellIndex,
		})
		maskIndex = tb.MaskIndex
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

// Stack

type Stack struct {
	queue []*Board
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(bs ...*Board) {
	s.queue = append(s.queue, bs...)
}

func (s *Stack) Pop() *Board {
	end := len(s.queue) - 1
	result := s.queue[end]
	s.queue = s.queue[:end]
	return result
}

func (s *Stack) IsEmpty() bool {
	return len(s.queue) == 0
}
