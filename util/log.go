package util

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
)

type Level string

const (
	LevelDebug Level = "debug"
	LevelInfo  Level = "info"
	LevelWarn  Level = "warn"
	LevelError Level = "error"
)

var (
	stderrLog     *log.Logger
	stdoutLog     *log.Logger
	stdoutLogWarn *log.Logger
	systemLevel   = LevelInfo
)

func init() {
	RestoreLoggers()
}

// Restore is used to restore loggers to their default values.  Used in testing where you want to restore loggers after tests.
func RestoreLoggers() {
	stderrLog = log.New(logWriterColor{output: os.Stderr, clr: color.New(color.FgHiRed)}, "", 0)
	stdoutLog = log.New(logWriter{output: os.Stdout}, "", 0)
	stdoutLogWarn = log.New(logWriterColor{output: os.Stdout, clr: color.New(color.FgHiYellow)}, "", 0)
}

type logWriter struct {
	output io.Writer
}

// Write writes bytes to writer with timestamp.
func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Fprint(writer.output, time.Now().UTC().Format("2006-01-02T15:04:05.000Z")+" "+string(bytes))
}

type logWriterColor struct {
	clr    *color.Color
	output io.Writer
}

// Write writes bytes to writer with timestamp.
func (writer logWriterColor) Write(bytes []byte) (int, error) {
	return writer.clr.Fprint(writer.output, time.Now().UTC().Format("2006-01-02T15:04:05.000Z")+" "+string(bytes))
}

// SetStdoutLogTarget sets global stdout target.  To stop stdout logging, call `SetStdoutLogTarget(ioutil.Discard)`.
func SetStdoutLogTarget(w io.Writer) {
	stdoutLog = log.New(logWriter{output: w}, "", 0)
	stdoutLogWarn = log.New(logWriter{output: w}, "", 0)
}

// SetStderrLogTarget sets global stderr target.  To stop stderr logging, call `SetStderrLogTarget(ioutil.Discard)`.
func SetStderrLogTarget(w io.Writer) {
	stderrLog = log.New(logWriter{output: w}, "", 0)
}

// GetSystemLogLevel returns system log level.
func GetSystemLogLevel() Level {
	return systemLevel
}

// SetSystemLogLevel sets the system log level.
func SetSystemLogLevel(l Level) {
	systemLevel = l
	OnErrorExit1(os.Setenv("PROTOC_DEBUG", "true"), `failed to set env var "PROTOC_DEBUG", bug in api-cli-core`)
}

// Log logs info and warn level to stdout, and error level to stderr.
func Log(level Level, format string, params ...interface{}) {
	f := fmt.Sprintf("[%s] %s", level, format)
	switch level {
	case LevelError:
		stderrLog.Printf(f, params...)
	case LevelWarn:
		stdoutLogWarn.Printf(f, params...)
	case LevelInfo:
		stdoutLog.Printf(f, params...)
	case LevelDebug:
		if systemLevel == LevelDebug {
			stdoutLog.Printf(f, params...)
		}
	}
}
