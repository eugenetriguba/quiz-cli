package cli

import (
	"io"
	"os"
	"path/filepath"
	"time"
)

var defaultCsvPath string = "../problems.csv"
var defaultTimeLimit = time.Second * 30

// Output and certain methods
// are specified here so we can
// replace them when needed to
// ease testing.
var out io.Writer = os.Stdout
var errOut io.Writer = os.Stderr
var osExit = os.Exit
var filepathAbs = filepath.Abs
