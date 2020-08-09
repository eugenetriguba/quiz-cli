package quiz

import "testing"

func TestNewDefaultProblem(t *testing.T) {
	expectedQuestion := ""
	expectedAnswer := ""

	p := NewDefaultProblem()

	if expectedQuestion != p.Question {
		t.Errorf("expected %s, got %s", expectedQuestion, p.Question)
	}

	if expectedAnswer != p.Answer {
		t.Errorf("expected %s, got %s", expectedAnswer, p.Answer)
	}
}

func TestNewProblemSetsAttrsWithArgs(t *testing.T) {
	expectedQuestion := "1+1"
	expectedAnswer := "2"

	p := NewProblem(expectedQuestion, expectedAnswer)

	if expectedQuestion != p.Question {
		t.Errorf("expected %s, got %s", expectedQuestion, p.Question)
	}

	if expectedAnswer != p.Answer {
		t.Errorf("expected %s, got %s", expectedAnswer, p.Answer)
	}
}

func TestProblemCanCheckWhetherTheAnswerIsCorrect(t *testing.T) {
	for _, test := range []struct {
		QuizProblem    *Problem
		Input          string
		ExpectedResult bool
	}{
		{
			QuizProblem:    NewProblem("1+1", "2"),
			Input:          "3",
			ExpectedResult: false,
		},
		{
			QuizProblem:    NewProblem("1+1", "2"),
			Input:          "2",
			ExpectedResult: true,
		},
	} {
		actualResult := test.QuizProblem.IsCorrect(test.Input)
		if actualResult != test.ExpectedResult {
			t.Errorf("expected %v, got %v\n", test.ExpectedResult, actualResult)
		}
	}
}
