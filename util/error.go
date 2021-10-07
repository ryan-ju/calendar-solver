package util

import (
	"fmt"
	"log"
	"os"
)

var (
	withStackTrace bool
)

// SetWithStackTrace sets if stack trace should be printed out on errors.
func SetWithStackTrace(v bool) {
	withStackTrace = v
}

// OnErrorExit1 exits with code 1 if error is not nil, otherwise does nothing.
func OnErrorExit1(err error, args ...interface{}) {
	if err != nil {
		PrintAndExit1(err, args...)
	}
}

// PrintAndExit1 prints v and exits with code 1.  Supported types are error and string.  If type is not supported, panic.
func PrintAndExit1(v interface{}, args ...interface{}) {
	if withStackTrace {
		printAndExit1WithStackTrace(v, args...)
	}

	if err, ok := v.(error); ok {
		if len(args) > 0 {
			Log(LevelError, fmt.Sprintf("%s: \n%s", fmt.Sprintf(args[0].(string), args[1:]...), err.Error()))
		} else {
			Log(LevelError, err.Error())
		}
		os.Exit(1)
	}

	msg, ok := v.(string)
	if ok {
		if len(args) > 0 {
			Log(LevelError, fmt.Sprintf("%s: %s", fmt.Sprintf(args[0].(string), args[1:]...), msg))
		} else {
			Log(LevelError, msg)
		}
		os.Exit(1)
	}

	log.Fatalf("unknown error type %T\n", v)
}

// printAndExit1WithStackTrace behaves the same as PrintAndExit1, but logs stack trace.
func printAndExit1WithStackTrace(v interface{}, args ...interface{}) {
	if err, ok := v.(error); ok {
		if len(args) > 0 {
			Log(LevelError, fmt.Sprintf("%s: \n%+v", fmt.Sprintf(args[0].(string), args[1:]...), err))
		} else {
			Log(LevelError, fmt.Sprintf("%+v", err))
		}
		os.Exit(1)
	}

	msg, ok := v.(string)
	if ok {
		if len(args) > 0 {
			Log(LevelError, fmt.Sprintf("%s: %s", fmt.Sprintf(args[0].(string), args[1:]...), msg))
		} else {
			Log(LevelError, msg)
		}
		os.Exit(1)
	}

	panic(v)
}
