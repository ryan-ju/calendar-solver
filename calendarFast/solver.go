package calendarFast

import (
	"strings"

	"github.com/ryan-ju/calendar-solver/util"
)

type Solver struct {
	Month int
	Day   int
}

func NewSolver(month, day int) *Solver {
	return &Solver{
		Month: month,
		Day:   day,
	}
}

func (s *Solver) Solve() {
	var solutions []string
	stack := NewStack()
	board := NewBoard(s.Month, s.Day)
	stack.Push(board)

	var counter int64
	for !stack.IsEmpty() {
		b := stack.Pop()

		nbs := b.NextFast()
		for _, nb := range nbs {
			if nb.IsSolved() {
				solutions = append(solutions, nb.Solution())
			} else {
				stack.Push(nb)
			}
		}
		counter++
	}
	util.Log(util.LevelInfo, "searched times: %d", counter)

	if len(solutions) > 0 {
		util.Log(util.LevelInfo, "found %d solutions, \n%s\n", len(solutions), strings.Join(solutions, ""))
		return
	}

	util.Log(util.LevelInfo, "shoot, no solution??? Must be a bug")
}
