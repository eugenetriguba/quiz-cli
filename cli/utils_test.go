package cli

import (
	"bytes"
	"errors"
	"path/filepath"
	"testing"

	"github.com/eugenetriguba/quiz-cli/internal/testutil"
)

func TestReadFileRetrievesContents(t *testing.T) {
	expectedContents := []byte("test contents")

	tmpfile := testutil.CreateTempFile(expectedContents, t)
	defer testutil.CleanTempFile(tmpfile, t)

	actualContents, err := readFile(tmpfile.Name())
	if err != nil {
		t.Errorf("could not read '%s' because '%s'", tmpfile.Name(), err)
	}

	if bytes.Compare(expectedContents, actualContents) != 0 {
		t.Errorf("expected %v, got %v", expectedContents, actualContents)
	}
}

func TestAbsolutePathCanBeRetrieved(t *testing.T) {
	expectedPath, err := filepath.Abs(defaultCsvPath)
	if err != nil {
		t.Errorf("could not get filepath to expected file")
	}

	actualPath, err := getAbsolutePath(defaultCsvPath)
	if err != nil {
		t.Errorf("could not get absolute path '%v'", err)
	}

	if expectedPath != actualPath {
		t.Errorf("expected %s, got %s", expectedPath, actualPath)
	}
}

// We have to replace the real filepath.Abs call here since
// filepath.Abs seems to be hard to get to return an error.
// However, we still want to ensure our code is working as it
// should in the event there is one.
func TestAbsolutePathReturnsError(t *testing.T) {
	// Save current function and restore at the end
	oldFilepathAbs := filepathAbs
	defer func() { filepathAbs = oldFilepathAbs }()

	testAbs := func(path string) (string, error) {
		return "", errors.New("testing error")
	}

	filepathAbs = testAbs
	_, err := getAbsolutePath("")

	if err == nil {
		t.Errorf("expected error to not be nil, got nil")
	}
}

func TestReadFileReturnsError(t *testing.T) {
	_, err := readFile(".")

	if err == nil {
		t.Errorf("expected error to not be nil, got nil")
	}
}
