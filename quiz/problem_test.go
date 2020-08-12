package quiz

import "testing"

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
		{
			QuizProblem:    NewProblem("1+1", "2"),
			Input:          "  2  \n",
			ExpectedResult: true,
		},
		{
			QuizProblem:    NewProblem("1+1", "2  "),
			Input:          "2\n",
			ExpectedResult: true,
		},
		{
			QuizProblem:    NewProblem("What programming language is this?", "GO"),
			Input:          "go",
			ExpectedResult: true,
		},
	} {
		actualResult := test.QuizProblem.IsCorrect(test.Input)
		if actualResult != test.ExpectedResult {
			t.Errorf("expected %v, got %v\n", test.ExpectedResult, actualResult)
		}
	}
}
