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

func TestQuizReadFillsProblems(t *testing.T) {
	q := NewDefaultQuiz()
	input := "1+1,2\n5+5,10"
	want := []*Problem{NewProblem("1+1", "2"), NewProblem("5+5", "10")}

	err := q.Parse(input)
	if err != nil {
		t.Errorf("could not read '%s' because '%v'\n", input, err)
	}
	got := q.Problems

	testEqProblems(want, got, t)
}

func TestQuizReadReturnsErrOnInvalidCsvFormat(t *testing.T) {
	input := "a,b\nabc,abc,abc"
	q := NewDefaultQuiz()
	err := q.Parse(input)

	if err == nil {
		t.Errorf("expected error on Parse(%q), got nil instead", input)
	}
}

func TestQuizReadReturnsErrOnCsvLineWithoutTwoItems(t *testing.T) {
	input := "abc,abc,abc"
	q := NewDefaultQuiz()
	err := q.Parse(input)

	if err == nil {
		t.Errorf("expected error on Parse(%q), got nil instead", input)
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

func TestQuizShuffle(t *testing.T) {
	input := []*Problem{}
	for i := 0; i < 50; i++ {
		input = append(input, NewProblem(string(i), "n/a"))
	}

	q := NewQuiz(input)
	q.Shuffle()
	shuffled := false

	for i := 0; i < 50; i++ {
		if strings.Compare(q.Problems[i].Question, string(i)) != 0 {
			shuffled = true
			break
		}
	}

	if !shuffled {
		t.Errorf("Shuffle() of 50 problems were all in the correct order, expected to be shuffled\n")
	}
}
