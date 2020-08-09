package quiz

// Problem is a quiz problem.
type Problem struct {
	Question string
	Answer   string
}

// NewDefaultProblem creates a new Problem type
// with empty string question and answer.
func NewDefaultProblem() *Problem {
	var p Problem
	return &p
}

// NewProblem creates a new Problem type with question
// and answer set using the given arguments.
func NewProblem(question string, answer string) *Problem {
	var p Problem

	p.Question = question
	p.Answer = answer

	return &p
}
