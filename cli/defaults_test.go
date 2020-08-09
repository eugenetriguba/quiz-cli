package cli

import (
	"os"
	"reflect"
	"testing"
)

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
	expectedDefault := "../problems.csv"
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
