package cli

import (
	"errors"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/eugenetriguba/quiz-cli/internal/testutil"
)

func TestCommandLineRunExitCodesAndErrOut(t *testing.T) {
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	for _, test := range []struct {
		QuizFile            string
		InputFile           string
		ExpectedReturnValue int
		ExpectedErrorOutput string
	}{
		{
			QuizFile:            "1+1,2\n",
			InputFile:           "2\n",
			ExpectedReturnValue: 0,
			ExpectedErrorOutput: "",
		},
		{
			QuizFile:            io.EOF.Error(),
			InputFile:           "",
			ExpectedReturnValue: 1,
			ExpectedErrorOutput: "could not parse the csv file",
		},
		{
			QuizFile:            "1,2",
			InputFile:           "2",
			ExpectedReturnValue: 1,
			ExpectedErrorOutput: "an error occurred while playing the quiz",
		},
	} {
		quizFile := testutil.CreateTempFile(test.QuizFile, t)
		inputFile := testutil.CreateTempFile(test.InputFile, t)
		defer testutil.CleanTempFile(quizFile, t)
		defer testutil.CleanTempFile(inputFile, t)

		os.Stdin = inputFile
		cl := NewCommandLine()
		cl.Parse("quiz", []string{"-csv", quizFile.Name()})
		errOut, actualReturnValue := testutil.CaptureOutputAndReturnValue(cl.Run, os.Stderr)

		if test.ExpectedReturnValue != actualReturnValue {
			t.Errorf("expected return value %d, got %d",
				test.ExpectedReturnValue, actualReturnValue)
		}

		if !strings.Contains(test.ExpectedErrorOutput, errOut) {
			t.Errorf(
				"expected '%s' in stderr output, got '%s' for output\n",
				test.ExpectedErrorOutput, errOut)
		}
	}
}

func TestCommandLineRunBadCsvPath(t *testing.T) {
	expectedReturnValue := 1
	expectedOutput := "could not read the csv file at"

	cl := NewCommandLine()
	cl.Parse("quiz", []string{"-csv", "should-not-exist!"})
	errOut, actualReturnValue := testutil.CaptureOutputAndReturnValue(cl.Run, os.Stderr)

	if expectedReturnValue != actualReturnValue {
		t.Errorf("expected return value %d, got %d",
			expectedReturnValue, actualReturnValue)
	}

	if !strings.Contains(expectedOutput, errOut) {
		t.Errorf(
			"expected '%s' in stderr output, got '%s' for output\n",
			expectedOutput, errOut)
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

	wantCsvPath, err := filepathAbs(defaultCsvPath)
	if err != nil {
		t.Errorf(
			"could not get absolute path of '%s' because '%s'",
			defaultCsvPath, err)
	}
	wantTimeLimit := defaultTimeLimit

	cl := NewCommandLine()
	cl.Parse(progname, args)
	gotCsvPath := cl.csvPath
	gotTimeLimit := cl.timeLimit

	if gotCsvPath != wantCsvPath {
		t.Errorf("expected csv path %s, got %s", wantCsvPath, cl.csvPath)
	}

	if gotTimeLimit != wantTimeLimit {
		t.Errorf("expected time limit %d, got %d", wantTimeLimit, cl.timeLimit)
	}
}

func TestCommandLineCanParseFlags(t *testing.T) {
	csvInput := "test-path.csv"
	progname := "quiz"
	args := []string{"-csv", csvInput, "-limit", "15s"}

	expectedCsvPath, err := filepathAbs(csvInput)
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
	output, err := cl.Parse("quiz", []string{"-nonexistentflag", "10"})

	if !strings.Contains(output, "flag provided but not defined: -nonexistentflag") {
		t.Errorf("expected '', got '%s'", output)
	}

	if err == nil {
		t.Errorf("expected err to not be nil on invalid parse, got %v", err)
	}
}

func TestCommandLineShowsHelpErr(t *testing.T) {
	cl := NewCommandLine()
	want := "flag: help requested"
	_, got := cl.Parse("quiz", []string{"-h"})

	if !strings.Contains(got.Error(), want) {
		t.Errorf("expected '%s' on -h flag, got '%v'\n", want, got)
	}
}

func TestCommandLineParseHasErrOnInvalidPath(t *testing.T) {
	oldFilepathAbs := filepathAbs
	defer func() { filepathAbs = oldFilepathAbs }()
	testAbs := func(path string) (string, error) {
		return "", errors.New("testing error")
	}

	filepathAbs = testAbs
	cl := NewCommandLine()
	output, err := cl.Parse("quiz", []string{})

	if output != "" {
		t.Errorf("expected '', got '%s'", output)
	}

	if err == nil {
		t.Errorf("expected error to not be nil, got nil")
	}
}
