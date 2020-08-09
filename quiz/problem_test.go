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
