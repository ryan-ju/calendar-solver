package solve2

import (
	"strconv"

	"github.com/ryan-ju/calendar-solver/calendar2"
	"github.com/ryan-ju/calendar-solver/util"
	"github.com/spf13/cobra"
)

type Command struct {
}

func NewCommand() *cobra.Command {
	cmd := &Command{}
	result := &cobra.Command{
		Use:   "solve2",
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
	solver := calendar2.NewSolver(month, day)
	solver.Solve()
}
