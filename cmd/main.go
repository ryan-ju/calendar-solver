package main

import (
	"io/ioutil"
	"os"

	"github.com/ryan-ju/calendar-solver/cmd/solve"
	"github.com/ryan-ju/calendar-solver/util"
	"github.com/spf13/cobra"
)

var (
	quiet bool
	debug bool
)

func main() {
	util.SetWithStackTrace(true)

	// Initialise values after arguments are parsed but before any command runs
	cobra.OnInitialize(func() {
		util.SetStderrLogTarget(os.Stderr)
		if quiet {
			util.SetStdoutLogTarget(ioutil.Discard)
		} else {
			util.SetStdoutLogTarget(os.Stdout)
		}

		if debug {
			util.SetSystemLogLevel(util.LevelDebug)
			util.SetWithStackTrace(true)
		}
	})

	cmd := &cobra.Command{
		Use:   "cal-sol",
		Short: "Tool for solving calendar puzzle",
	}

	cmd.PersistentFlags().BoolVar(&quiet, "quiet", false, "If set, no info or warn logs will be printed.  Error logs are still printed (to stderr).  Useful for scripting.")
	cmd.PersistentFlags().BoolVar(&debug, "debug", false, "If set, then debug log will be printed.")
	cmd.AddCommand(solve.NewCommand())
	util.OnErrorExit1(cmd.Execute())
}
