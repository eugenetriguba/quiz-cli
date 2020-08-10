package cli

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestCsvFileHasDefault(t *testing.T) {
	expectedDefault := filepath.Join(getProjectBasepath(), "problems.csv")
	actualDefault := defaultCsvPath

	if actualDefault != expectedDefault {
		t.Errorf("expected %s, got %s", expectedDefault, actualDefault)
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
