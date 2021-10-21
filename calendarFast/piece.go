package calendarFast

import (
	"strings"
)

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
