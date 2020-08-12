package cli

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/eugenetriguba/quiz-cli/quiz"
)

// Specified here to ease testing.
var filepathAbs = filepath.Abs

// The CommandLine contains the flags that
// the cli uses and is able to parse them.
type CommandLine struct {
	csvPath   string
	timeLimit time.Duration
	shuffle   bool
}

// NewCommandLine creates a new
// empty CommandLine type.
//
// csvPath is set to an empty string
// and the timeLimit is set to 0.
func NewCommandLine() *CommandLine {
	var cl CommandLine
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
		&cl.timeLimit, "limit", defaultTimeLimit, "Set the quiz time limit")
	flagset.BoolVar(
		&cl.shuffle, "shuffle", false, "Shuffle the problems presented")

	err = flagset.Parse(args)
	if err != nil {
		return buf.String(), err
	}

	cl.csvPath, err = filepathAbs(cl.csvPath)
	if err != nil {
		return buf.String(), err
	}

	return buf.String(), nil
}

// Run runs our CLI to play the quiz.
//
// The exit code that the app should use is returned.
// 1 if we could not read the csv file; 0 otherwise.
func (cl *CommandLine) Run() int {
	q := quiz.NewDefaultQuiz()

	contents, err := ioutil.ReadFile(cl.csvPath)
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"could not read the csv file at '%s' because '%v'\n",
			cl.csvPath, err)
		return 1
	}

	err = q.Parse(string(contents))
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"could not parse the csv file contents because '%v'\n",
			err)
		return 1
	}

	if cl.shuffle {
		q.Shuffle()
	}

	err = q.Start(cl.timeLimit)
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"an error occurred while playing the quiz: '%v'\n",
			err)
		return 1
	}

	return 0
}
