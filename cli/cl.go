package cli

import (
	"bytes"
	"flag"
	"time"
)

// The CommandLine contains the flags that
// the cli uses and is able to parse them.
type CommandLine struct {
	csvPath   string
	timeLimit time.Duration
}

// NewCommandLine creates a new
// empty CommandLine type.
//
// csvPath is set to an empty string
// and the timeLimit is set to 0.
func NewCommandLine() *CommandLine {
	var cl CommandLine

	cl.csvPath = ""
	cl.timeLimit = 0

	return &cl
}

// Parse parses the flags from the arguments.
//
// Args:
//   progname: The program name (typically os.Args[0])
//   args: The program arguments (typically os.Args[1:])
//
// csvPath is set to the problems.csv path
// in this package by default if a flag is not
// specified.
//
// timeLimit is set to 30 seconds by default if
// a flag is not specified.
func (cl *CommandLine) Parse(progname string, args []string) (output string, err error) {
	flagset := flag.NewFlagSet(progname, flag.ContinueOnError)
	var buf bytes.Buffer
	flagset.SetOutput(&buf)

	flagset.StringVar(
		&cl.csvPath, "csv", defaultCsvPath, "Set problems file")
	flagset.DurationVar(
		&cl.timeLimit, "limit", defaultTimeLimit, "Set the question time limit")

	err = flagset.Parse(args)
	if err != nil {
		return "", err
	}

	cl.csvPath, err = filepathAbs(cl.csvPath)
	if err != nil {
		return "", err
	}

	return "", nil
}

// Run runs our CLI.
func (cl *CommandLine) Run() int {
	return 0
}
