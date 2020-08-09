package cli

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Output and certain methods
// are specified here so we can
// replace them when needed to
// ease testing.
var out io.Writer = os.Stdout
var errOut io.Writer = os.Stderr
var osExit = os.Exit
var filepathAbs = filepath.Abs

var defaultCsvPath string = "../problems.csv"

// Run runs our CLI.
func Run(args []string) int {
	return 0
}

func readFile(path string) []byte {
	return []byte("")
}

func getAbsolutePath(path string) string {
	absoluteCsvPath, err := filepathAbs(path)
	if err != nil {
		fmt.Print(errOut, "could not retrieve the absolute path to %s", path)
		osExit(1)
	}
	return absoluteCsvPath
}
