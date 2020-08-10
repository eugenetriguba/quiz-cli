package quiz

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

// Quiz holds all the problems for the quiz we're taking
// and how many we've gotten correct.
type Quiz struct {
	Problems []*Problem
	Correct  int
}

// NewDefaultQuiz creates a Quiz type with an
// empty Problems array.
func NewDefaultQuiz() *Quiz {
	var q Quiz
	return &q
}

// NewQuiz creates a Quiz type set to the
// specified problems.
func NewQuiz(problems []*Problem) *Quiz {
	var q Quiz

	q.Problems = problems

	return &q
}

// Parse parses the contents of a csv file
// to set its problems.
//
// The csv file should be in the following format:
// "question", answer
// "question", answer
//
//  Or:
//  question, header
//  "1+1", "2"
//
//  and so on.
func (q *Quiz) Parse(csvFile string) error {
	r := csv.NewReader(strings.NewReader(csvFile))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if len(record) != 2 {
			return fmt.Errorf("each csv line should have two items")
		}

		q.Problems = append(q.Problems, NewProblem(record[0], record[1]))
	}

	return nil
}

// Play starts up the quiz.
func (q *Quiz) Play() error {
	for i := 0; i < len(q.Problems); i++ {
		fmt.Fprintf(os.Stdout, "Problem #%d: %s = ", i+1, q.Problems[i].Question)

		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		if q.Problems[i].IsCorrect(text) {
			q.Correct++
		}
	}

	fmt.Fprintf(os.Stdout, "\nYou scored %d out of %d.\n", q.Correct, len(q.Problems))
	return nil
}
