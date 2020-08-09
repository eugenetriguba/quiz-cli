package cli

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

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
