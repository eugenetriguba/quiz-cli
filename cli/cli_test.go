package cli

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func createTempFile(content []byte, t *testing.T) *os.File {
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

func cleanTempFile(f *os.File, t *testing.T) {
	filename := f.Name()

	if err := f.Close(); err != nil {
		t.Errorf("could not close the temp file %s", filename)
	}
	os.Remove(filename)
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
	expectedDefault := "../problems.csv"
	actualDefault := defaultCsvPath

	if actualDefault != expectedDefault {
		t.Errorf("expected %s, got %s", expectedDefault, actualDefault)
	}
}

func TestAbsolutePathCanBeRetrieved(t *testing.T) {
	expectedPath, err := filepath.Abs(defaultCsvPath)
	if err != nil {
		t.Errorf("could not get filepath to expected file")
	}
	actualPath := getAbsolutePath(defaultCsvPath)

	if expectedPath != actualPath {
		t.Errorf("expected %s, got %s", expectedPath, actualPath)
	}
}

func TestOsExitMethod(t *testing.T) {
	expectedOsExitMethod := reflect.ValueOf(os.Exit)
	actualOsExitMethod := reflect.ValueOf(osExit)

	if expectedOsExitMethod.Pointer() != actualOsExitMethod.Pointer() {
		t.Errorf("expected %v (os.Exit), got %v", expectedOsExitMethod, actualOsExitMethod)
	}
}

func TestRunReturnsZero(t *testing.T) {
	expectedReturnValue := 0
	actualReturnValue := Run([]string{"./quiz"})

	if expectedReturnValue != actualReturnValue {
		t.Errorf("expected %d, got %d", expectedReturnValue, actualReturnValue)
	}
}

func TestReadFileRetrievesContents(t *testing.T) {
	expectedContents := []byte("test contents")

	tmpfile := createTempFile(expectedContents, t)
	defer cleanTempFile(tmpfile, t)

	actualContents := readFile(tmpfile.Name())

	if bytes.Compare(expectedContents, actualContents) != 0 {
		t.Errorf("expected %v, got %v", expectedContents, actualContents)
	}
}

// We have to replace the real os.Exit and filepath.Abs
// calls here since testing os.Exit is tricky to do and
// filepath.Abs seems to be hard to get to return an error.
//
// However, we still want to ensure our code is working as it
// should in the event there is one.
//
// This test and TestReadFileRaisesOsExitOnError may be able
// to be merged into one test to reduce the boilerplate code.
func TestAbsolutePathRaisesOsExitOnError(t *testing.T) {
	// Save current functions and restore at the end
	oldOsExit := osExit
	oldFilepathAbs := filepathAbs
	defer func() {
		osExit = oldOsExit
		filepathAbs = oldFilepathAbs
	}()

	expectedExitCode := 1
	var actualExitCode int
	testExit := func(code int) {
		actualExitCode = code
	}
	testAbs := func(path string) (string, error) {
		return "", errors.New("testing error")
	}

	osExit = testExit
	filepathAbs = testAbs
	getAbsolutePath("")

	if expectedExitCode != actualExitCode {
		t.Errorf("expected exit code %d, got %d", expectedExitCode, actualExitCode)
	}
}

func TestReadFileRaisesOsExitOnError(t *testing.T) {
	// Save current functions and restore at the end
	oldOsExit := osExit
	defer func() { osExit = oldOsExit }()

	expectedExitCode := 1
	var actualExitCode int
	testExit := func(code int) {
		actualExitCode = code
	}

	osExit = testExit
	readFile(".")

	if expectedExitCode != actualExitCode {
		t.Errorf("expected exit code %d, got %d", expectedExitCode, actualExitCode)
	}
}
