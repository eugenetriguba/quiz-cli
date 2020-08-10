package quiz

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/eugenetriguba/quiz-cli/internal/testutil"
)

// testEqProblems uses the t to fail the test if a and b are not
// equal. Otherwise, does nothing if they are equal.
func testEqProblems(a []*Problem, b []*Problem, t *testing.T) {
	if (&a == nil) != (&b == nil) {
		t.Errorf("expected %v and %v to either be both nil or both not nil", a, b)
		return
	}

	if len(a) != len(b) {
		t.Errorf("expected len of a and b to be equal, got %v != %v", len(a), len(b))
		return
	}

	for i := range a {
		if a[i].Question != b[i].Question && a[i].Answer != b[i].Answer {
			t.Errorf("index %d problem's question and answer are not equal", i)
			return
		}
	}
}

func TestNewDefaultQuiz(t *testing.T) {
	expectedProblems := []*Problem{}

	q := NewDefaultQuiz()

	testEqProblems(q.Problems, expectedProblems, t)
}

func TestNewQuiz(t *testing.T) {
	expectedProblems := []*Problem{NewProblem("1+1", "2")}

	q := NewQuiz(expectedProblems)

	testEqProblems(q.Problems, expectedProblems, t)
}

func TestQuizReadFillsProblems(t *testing.T) {
	csvFile := "1+1,2\n5+5,10"
	expectedProblems := []*Problem{NewProblem("1+1", "2"), NewProblem("5+5", "10")}

	q := NewDefaultQuiz()
	err := q.Parse(csvFile)
	if err != nil {
		t.Errorf("could not read '%s' because '%v'\n", csvFile, err)
	}

	testEqProblems(q.Problems, expectedProblems, t)
}

func TestQuizReadReturnsErrOnInvalidCsvFormat(t *testing.T) {
	csvFile := "a,b\nabc,abc,abc"

	q := NewDefaultQuiz()
	err := q.Parse(csvFile)

	if err == nil {
		t.Errorf("expected error to not be nil, got nil instead")
	}
}

func TestQuizReadReturnsErrOnCsvLineWithoutTwoItems(t *testing.T) {
	csvFile := "abc,abc,abc"

	q := NewDefaultQuiz()
	err := q.Parse(csvFile)

	if err == nil {
		t.Errorf("expected error to not be nil, got nil instead")
	}
}

func TestQuizPlay(t *testing.T) {
	oldStdin := os.Stdin
	oldStdout := os.Stdout
	defer func() { os.Stdin = oldStdin; os.Stdout = oldStdout }()

	inFile := testutil.CreateTempFile("2\n", t)
	outFile := testutil.CreateTempFile("", t)
	defer testutil.CleanTempFile(inFile, t)
	defer testutil.CleanTempFile(outFile, t)
	os.Stdin = inFile
	os.Stdout = outFile

	q := NewQuiz([]*Problem{NewProblem("1+1", "2")})
	err := q.Play()

	if err != nil {
		t.Errorf("expected no error, got '%v'\n", err)
	}

	if 1 != q.Correct {
		t.Errorf("expected %d correct problems, got %d\n", 1, q.Correct)
	}

	outFileContents, err := ioutil.ReadFile(outFile.Name())
	if err != nil {
		t.Errorf("could not retrieve outfile contents because '%v'\n", err)
	}

	lastNum := fmt.Sprintf("#%d", len(q.Problems))
	if !strings.Contains(string(outFileContents), lastNum) {
		t.Errorf("expected '%s' in output, got '%s'\n", lastNum, string(outFileContents))
	}
}
