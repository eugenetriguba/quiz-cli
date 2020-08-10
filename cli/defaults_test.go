package cli

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestInputIsGoingToStdout(t *testing.T) {
	expectedIn := os.Stdin
	actualIn := in

	if actualIn != expectedIn {
		t.Errorf("expected %#v, got %#v", expectedIn, actualIn)
	}
}

func TestOutputIsGoingToStdout(t *testing.T) {
	expectedOut := os.Stdout
	actualOut := out

	if actualOut != expectedOut {
		t.Errorf("expected %#v, got %#v", expectedOut, actualOut)
	}
}

func TestErrorOutputIsGoingToStderr(t *testing.T) {
	expectedErrOut := os.Stderr
	actualErrOut := errOut

	if actualErrOut != expectedErrOut {
		t.Errorf("expected %#v, got %#v", expectedErrOut, actualErrOut)
	}
}

func TestCsvFileHasDefault(t *testing.T) {
	expectedDefault := filepath.Join(getProjectBasepath(), "problems.csv")
	actualDefault := defaultCsvPath

	if actualDefault != expectedDefault {
		t.Errorf("expected %s, got %s", expectedDefault, actualDefault)
	}
}

func TestOsExitMethod(t *testing.T) {
	expectedOsExitMethod := reflect.ValueOf(os.Exit)
	actualOsExitMethod := reflect.ValueOf(osExit)

	if expectedOsExitMethod.Pointer() != actualOsExitMethod.Pointer() {
		t.Errorf("expected %v (os.Exit), got %v", expectedOsExitMethod, actualOsExitMethod)
	}
}

func TestGetProjectBasepathRetrievesQuizCliDir(t *testing.T) {
	expectedPath, err := filepath.Abs("..")
	if err != nil {
		t.Errorf("could not get absolute path to quiz-cli because '%v'", err)
	}
	actualPath := getProjectBasepath()

	if strings.Compare(expectedPath, actualPath) != 0 {
		t.Errorf("expected '%s' to equal '%s'\n", expectedPath, actualPath)
	}
}
