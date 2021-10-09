package solve

import (
	"strconv"

	"github.com/ryan-ju/calendar-solver/calendar"
	"github.com/ryan-ju/calendar-solver/util"
	"github.com/spf13/cobra"
)

type Command struct {
}

func NewCommand() *cobra.Command {
	cmd := &Command{}
	result := &cobra.Command{
		Use:   "solve",
		Short: "Solves the calendar puzzle",
		Run:   cmd.run,
	}

	return result
}

func (c *Command) run(_ *cobra.Command, args []string) {
	if len(args) != 1 {
		util.PrintAndExit1("must set date, format mmdd")
	}
	month, err := strconv.Atoi(args[0][:2])
	util.OnErrorExit1(err, "invalid month")
	day, err := strconv.Atoi(args[0][2:])
	util.OnErrorExit1(err, "invalid day")
	util.Log(util.LevelInfo, "solving for month = %d, day = %d", month, day)

	solver, err := calendar.NewSolver(calendar.Date{
		Month: month,
		Day:   day,
	})
	util.OnErrorExit1(err, "solver error")

	board := solver.Solve()
	if board == nil {
		util.Log(util.LevelInfo, "no solution found")
	} else {
		util.Log(util.LevelInfo, "solution found, %s\n", board.String())
	}
}
