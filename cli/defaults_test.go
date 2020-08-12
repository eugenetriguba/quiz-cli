package cli

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestGetProjectBasepathRetrievesQuizCliDir(t *testing.T) {
	want, err := filepath.Abs("..")
	if err != nil {
		t.Errorf("could not get absolute path to quiz-cli because '%v'\n", err)
	}
	got := getProjectBasepath()

	if strings.Compare(want, got) != 0 {
		t.Errorf("expected '%s' to equal '%s'\n", want, got)
	}
}
