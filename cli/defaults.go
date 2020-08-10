package cli

import (
	"io"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var defaultCsvPath string = filepath.Join(getProjectBasepath(), "problems.csv")
var defaultTimeLimit time.Duration = time.Second * 30

// Output and certain methods
// are specified here so we can
// replace them when needed to
// ease testing.
var out io.Writer = os.Stdout
var in io.Reader = os.Stdin
var errOut io.Writer = os.Stderr
var osExit = os.Exit
var filepathAbs = filepath.Abs

func getProjectBasepath() string {
	var (
		_, b, _, _ = runtime.Caller(0)
		basepath   = filepath.Dir(b)
	)

	return filepath.Join(basepath, "..")
}
