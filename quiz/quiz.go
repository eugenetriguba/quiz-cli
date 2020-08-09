package quiz

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

// Quiz holds all the problems for
// the quiz we're taking.
type Quiz struct {
	Problems []*Problem
}

// NewDefaultQuiz creates a Quiz type with an
// empty Problems array.
func NewDefaultQuiz() *Quiz {
	var q Quiz
	return &q
}

// Parse parses the contents of a csv file
// to set its problems.
//
// The csv file should be in the following format:
// "question", answer
// "question", answer
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
