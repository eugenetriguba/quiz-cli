package quiz

import (
	"testing"
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
