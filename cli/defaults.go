package cli

import (
	"path/filepath"
	"runtime"
	"time"
)

var defaultCsvPath string = filepath.Join(getProjectBasepath(), "problems.csv")
var defaultTimeLimit time.Duration = time.Second * 30

func getProjectBasepath() string {
	var (
		_, b, _, _ = runtime.Caller(0)
		basepath   = filepath.Dir(b)
	)

	return filepath.Join(basepath, "..")
}
