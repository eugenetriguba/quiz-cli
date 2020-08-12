package quiz

import "strings"

// Problem is a quiz problem.
type Problem struct {
	Question string
	Answer   string
}

// NewProblem creates a new Problem type with question
// and answer set using the given arguments.
func NewProblem(question string, answer string) *Problem {
	var p Problem

	p.Question = question
	p.Answer = answer

	return &p
}

// IsCorrect checks if the given input is the correct answer
// in a case and whitespace insensitive way.
func (p *Problem) IsCorrect(answer string) bool {
	answer = strings.TrimSpace(answer)
	answer = strings.TrimRight(answer, "\n")

	return strings.EqualFold(p.Answer, answer)
}
