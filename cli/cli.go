package cli

import (
	"fmt"
	"io"
	"io/ioutil"
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
func Run(progname string, args []string) int {
	return 0
}

func readFile(path string) []byte {
	contents, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Fprintf(
			errOut,
			"could not retrieve the file contents of '%s' because '%s'\n",
			path, err)
		osExit(1)
	}

	return contents
}

func getAbsolutePath(path string) string {
	absolutePath, err := filepathAbs(path)
	if err != nil {
		fmt.Fprintf(
			errOut,
			"could not retrieve the absolute path to '%s' because '%s'\n",
			path, err)
		osExit(1)
	}
	return absolutePath
}
