package cli

import (
	"errors"
	"testing"
	"time"
)

func TestCommandLineRunReturnsZero(t *testing.T) {
	expectedReturnValue := 0
	cl := NewCommandLine()
	cl.Parse("quiz", []string{})
	actualReturnValue := cl.Run()

	if expectedReturnValue != actualReturnValue {
		t.Errorf("expected %d, got %d", expectedReturnValue, actualReturnValue)
	}
}

func TestNewCommandLineCanBeCreated(t *testing.T) {
	expectedCsvPath := ""
	expectedTimeLimit := time.Duration(0)

	cl := NewCommandLine()

	if cl.csvPath != expectedCsvPath {
		t.Errorf("expected csv path %s, got %s", expectedCsvPath, cl.csvPath)
	}

	if cl.timeLimit != expectedTimeLimit {
		t.Errorf("expected time limit %d, got %d", expectedTimeLimit, cl.timeLimit)
	}
}

func TestCommandLineParseHasDefaultsForFlags(t *testing.T) {
	progname := "quiz"
	args := []string{}

	expectedCsvPath, err := getAbsolutePath(defaultCsvPath)
	if err != nil {
		t.Errorf(
			"could not get absolute path of '%s' because '%s'",
			defaultCsvPath, err)
	}
	expectedTimeLimit := defaultTimeLimit

	cl := NewCommandLine()
	cl.Parse(progname, args)

	if cl.csvPath != expectedCsvPath {
		t.Errorf("expected csv path %s, got %s", expectedCsvPath, cl.csvPath)
	}

	if cl.timeLimit != expectedTimeLimit {
		t.Errorf("expected time limit %d, got %d", expectedTimeLimit, cl.timeLimit)
	}
}

func TestCommandLineCanParseFlags(t *testing.T) {
	csvInput := "test-path.csv"
	progname := "quiz"
	args := []string{"-csv", csvInput, "-limit", "15s"}

	expectedCsvPath, err := getAbsolutePath(csvInput)
	if err != nil {
		t.Errorf(
			"could not get absolute path of '%s' because '%s'",
			csvInput, err)
	}
	expectedTimeLimit := 15 * time.Second

	cl := NewCommandLine()
	_, err = cl.Parse(progname, args)
	if err != nil {
		t.Errorf(
			"could not parse the given progname (%s) and args (%v) because %v",
			progname, args, err)
	}

	if cl.csvPath != expectedCsvPath {
		t.Errorf("expected csv path %s, got %s", expectedCsvPath, cl.csvPath)
	}

	if cl.timeLimit != expectedTimeLimit {
		t.Errorf("expected time limit %d, got %d", expectedTimeLimit, cl.timeLimit)
	}
}

func TestCommandLineHasErrOnInvalidParse(t *testing.T) {
	cl := NewCommandLine()
	_, err := cl.Parse("quiz", []string{"-nonexistentflag", "10"})

	if err == nil {
		t.Errorf("expected err to not be nil on invalid parse, got %v", err)
	}
}

func TestCommandLineParseHasErrOnInvalidPath(t *testing.T) {
	// Save current function and restore at the end
	oldFilepathAbs := filepathAbs
	defer func() { filepathAbs = oldFilepathAbs }()
	testAbs := func(path string) (string, error) {
		return "", errors.New("testing error")
	}

	filepathAbs = testAbs
	cl := NewCommandLine()
	_, err := cl.Parse("quiz", []string{})

	if err == nil {
		t.Errorf("expected error to not be nil, got nil")
	}
}
