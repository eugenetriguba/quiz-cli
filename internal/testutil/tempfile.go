package testutil

import (
	"io/ioutil"
	"os"
	"testing"
)

// CreateTempFile creates a temporary file in the
// given OSs tmp dir with contents written to it.
//
// The test is failed with the passed in t if an error
// occurs.
func CreateTempFile(content []byte, t *testing.T) *os.File {
	tmpfile, err := ioutil.TempFile("", "temp-quiz-test-file")
	if err != nil {
		t.Errorf(
			"could not create temp file '%s' because '%v'",
			tmpfile.Name(), err)
	}

	err = ioutil.WriteFile(tmpfile.Name(), content, 0600)
	if err != nil {
		t.Errorf(
			"could not write '%v' to temp file '%s'",
			content, err)
	}

	return tmpfile
}

// CleanTempFile cleans up the created tempfile
// by closing and removing it. It should be called
// or `defer`ed after calling CreateTempFile.
//
// The test is failed with the passed in t if an error
// occurs.
func CleanTempFile(f *os.File, t *testing.T) {
	filename := f.Name()

	if err := f.Close(); err != nil {
		t.Errorf(
			"could not close the temp file '%s' because '%s'\n",
			filename, err)
	}

	if err := os.Remove(filename); err != nil {
		t.Errorf(
			"could not remove temp file '%s' because '%s'\n",
			filename, err)
	}
}
