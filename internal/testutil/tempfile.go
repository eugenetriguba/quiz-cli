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
func CreateTempFile(content string, t *testing.T) *os.File {
	tmpfile, err := ioutil.TempFile("", "temp-quiz-test-file")
	if err != nil {
		t.Errorf(
			"could not create temp file '%s' because '%v'\n",
			tmpfile.Name(), err)
	}

	err = ioutil.WriteFile(tmpfile.Name(), []byte(content), 0644)
	if err != nil {
		t.Errorf(
			"could not write '%v' to temp file '%s'\n",
			content, err)
	}

	_, err = tmpfile.Seek(0, 0)
	if err != nil {
		t.Errorf(
			"could not seek to beginning of '%s' because '%v'\n",
			tmpfile.Name(), err)
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
