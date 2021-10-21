package calendarSlow

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type ShortIndex struct {
	X uint8
	Y uint8
}

type Date struct {
	Month int
	Day   int
}

type PieceHolder struct {
	Key   PieceKey
	Piece Piece
}

type Board struct {
	Parent          *Board
	Date            Date
	DateCoordinates []ShortIndex
	// Each byte is a row, e.g., 01111111 is a full row
	Cells      [7]byte
	PieceIndex int
	// Pieces and their config
	Pieces []PieceHolder

	FreeCells []ShortIndex
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

func NewBoard(date Date) (*Board, error) {
	if date.Month < 1 || date.Month > 12 {
		return nil, errors.Errorf("month %d is wrong", date.Month)
	}

	if date.Day < 1 || date.Day > 31 {
		return nil, errors.Errorf("day %d is wrong", date.Day)
	}

	var monthX, monthY, dayX, dayY uint8
	monthX = uint8((date.Month - 1) % 6)
	if date.Month > 6 {
		monthY = 1
	}

	dayX = uint8((date.Day - 1) % 7)
	dayY = uint8((date.Day-1)/7 + 2)

	fc := freeCells([]ShortIndex{
		{
			X: monthX,
			Y: monthY,
		},
		{
			X: dayX,
			Y: dayY,
		},
	})

	cells := emptyBoard
	cells[monthY] = cells[monthY] | (0b01000000 >> monthX)
	cells[dayY] = cells[dayY] | (0b01000000 >> dayX)

	return &Board{
		Cells: cells,
		Date:  date,
		DateCoordinates: []ShortIndex{
			{
				X: monthX,
				Y: monthY,
			},
			{
				X: dayX,
				Y: dayY,
			},
		},
		FreeCells: fc,
	}, nil
}

// AddPiece returns a new board only if pk can fit in the free cells.  Otherwise returns nil.
func (b *Board) AddPiece(pk PieceKey) *Board {
	p, ok := NewPiece(pk.P, pk.X, pk.Y, pk.O, pk.R)
	if !ok {
		return nil
	}

	var newCells [7]byte
	for i, row := range p {
		boardRow := b.Cells[i]
		// This means the piece doesn't fit the board
		// E.g.,
		// Board: 00111000
		// Piece: 11100000
		// Or   : 11111000
		// XOr  : 11011000
		or := row | boardRow
		if or != row^boardRow {
			return nil
		}
		newCells[i] = or
	}

	newPieces := make([]PieceHolder, len(b.Pieces), len(b.Pieces)+1)
	copy(newPieces, b.Pieces)
	newPieces = append(newPieces, PieceHolder{
		Key:   pk,
		Piece: p,
	})

	newBoard := Board{
		Parent:          b,
		Date:            b.Date,
		DateCoordinates: b.DateCoordinates,
		Cells:           newCells,
		PieceIndex:      b.PieceIndex + 1,
		Pieces:          newPieces,
		FreeCells:       subtractFreeCells(b.FreeCells, cellsOfPiece(p)),
	}

	return &newBoard
}

func (b *Board) IsSolved() bool {
	return b.PieceIndex == 8
}

func (b *Board) SolutionPath() string {
	var sb strings.Builder
	sb.WriteString("Solution path\n")
	var bs []*Board
	for p := b; p != nil; p = p.Parent {
		bs = append([]*Board{p}, bs...)
	}
	for _, p := range bs {
		sb.WriteString(p.String())
		sb.WriteString("\n")
	}
	return sb.String()
}

func (b Board) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Target date: %02d-%02d\n", b.Date.Month, b.Date.Day))
	sb.WriteString(b.StringSimple())
	// Debug
	sb.WriteString("Board cells\n")
	for _, bt := range b.Cells {
		str := fmt.Sprintf("%08b", bt)
		sb.WriteString(strings.Join(strings.Split(str, ""), " "))
		sb.WriteString("\n")
	}
	sb.WriteString("Pieces:\n")
	for _, p := range b.Pieces {
		key := p.Key
		sb.WriteString(fmt.Sprintf("%s, %s, %s, (%d,%d)\n", GetPieceChar(key.P), GetReflected(key.R), GetOrientation(key.O), key.X, key.Y))
	}

	return sb.String()
}

func (b Board) StringSimple() string {
	var sb strings.Builder
	for i := uint8(0); i < 7; i++ { // row
		if i < 2 {
			sb.WriteString("MM")
		} else {
			sb.WriteString(fmt.Sprintf("%02d", (i-2)*7+1))
		}
		for j := uint8(0); j < 7; j++ { // column
			if (i <= 1 && j >= 6) || (i == 6 && j >= 3) {
				continue
			}
			var skip bool
			for _, idx := range b.DateCoordinates {
				if i == idx.Y && j == idx.X {
					sb.WriteString("|")
					sb.WriteString(" @ ")
					skip = true
				}
			}
			if skip {
				continue
			}
			var found bool
			for _, p := range b.Pieces {
				var mask byte = 0b01000000 >> j
				if p.Piece[i]&mask > 0 {
					sb.WriteString("|")
					sb.WriteString(" " + GetPieceChar(p.Key.P) + " ")
					found = true
				}
			}
			if !found {
				sb.WriteString("|   ")
			}
		}
		sb.WriteString("|\n")
	}
	return sb.String()
}

func freeCells(skips []ShortIndex) []ShortIndex {
	var result [49]ShortIndex
	var counter int
	for i := uint8(0); i < 7; i++ {
		for j := uint8(0); j < 7; j++ {
			if (i <= 1 && j >= 6) || (i == 6 && j >= 3) {
				continue
			}
			var skip bool
			for _, s := range skips {
				if j == s.X && i == s.Y {
					skip = true
				}
			}
			if skip {
				continue
			}
			result[counter] = ShortIndex{
				X: j,
				Y: i,
			}
			counter++
		}
	}

	return result[:counter]
}

func cellsOfPiece(p Piece) []ShortIndex {
	var result [6]ShortIndex
	var counter int
	for i, row := range p {
		if row > 0 {
			for j := 0; j < 7; j++ {
				var mask byte = 0b01000000 >> j
				if mask&row > 0 {
					result[counter] = ShortIndex{
						X: uint8(j),
						Y: uint8(i),
					}
					counter++
				}
			}
		}
	}
	return result[:counter]
}

func subtractFreeCells(source, subtract []ShortIndex) []ShortIndex {
	result := make([]ShortIndex, 0, 49)
	subMap := map[ShortIndex]interface{}{}
	for _, sub := range subtract {
		subMap[sub] = nil
	}
	for _, src := range source {
		if _, ok := subMap[src]; !ok {
			result = append(result, src)
		}
	}
	return result
}

//func subtractFreeCells2(source, subtract []ShortIndex) []ShortIndex {
//	result := make([]ShortIndex, 0, 49)
//	subMap := map[ShortIndex]interface{}{}
//	for _, sub := range subtract {
//		subMap[sub] = nil
//	}
//	srcMap := map[ShortIndex]interface{}{}
//	for _, src := range source {
//		if _, ok := subMap[src]; !ok {
//			srcMap[src] = nil
//		}
//	}
//	for k := range srcMap {
//		var isLeftFree, isRightFree, isUpFree, isDownFree bool
//		if k.X > 0 {
//			_, isLeftFree = srcMap[ShortIndex{
//				X: k.X - 1,
//				Y: k.Y,
//			}]
//		}
//		if k.X < 6 {
//			_, isRightFree = srcMap[ShortIndex{
//				X: k.X + 1,
//				Y: k.Y,
//			}]
//		}
//		if k.Y > 0 {
//			_, isUpFree = srcMap[ShortIndex{
//				X: k.X,
//				Y: k.Y - 1,
//			}]
//		}
//		if k.Y < 6 {
//			_, isDownFree = srcMap[ShortIndex{
//				X: k.X,
//				Y: k.Y + 1,
//			}]
//		}
//		if isLeftFree || isRightFree || isUpFree || isDownFree {
//			result = append(result, k)
//		}
//	}
//	return result
//}
