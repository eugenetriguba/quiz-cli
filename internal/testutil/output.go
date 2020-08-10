package testutil

import (
	"bytes"
	"log"
	"os"
)

// CaptureOutputAndReturnValue captures the logging output of f
// when it logs to logOutput. That output is returned back as a
// string.
//
// A limitation with this function here is that the return value
// of the function must be an integer.
func CaptureOutputAndReturnValue(f func() int, logOutput *os.File) (string, int) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	val := f()
	log.SetOutput(logOutput)
	return buf.String(), val
}

func CaptureOutputAndErr(f func() error, logOutput *os.File) (string, error) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	val := f()
	log.SetOutput(logOutput)
	return buf.String(), val
}
