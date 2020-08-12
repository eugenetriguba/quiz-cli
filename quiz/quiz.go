package quiz

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
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

// Shuffle psuedorandomly shuffles our quiz problems
func (q *Quiz) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(q.Problems), func(i, j int) {
		q.Problems[i], q.Problems[j] = q.Problems[j], q.Problems[i]
	})
}

// Start starts up the quiz.
func (q *Quiz) Start(timeLimit time.Duration) error {
	timer := time.NewTimer(timeLimit)
	answerCh := make(chan string)
	errCh := make(chan error)
	go func() {
		reader := bufio.NewReader(os.Stdin)
		answer, err := reader.ReadString('\n')
		if err != nil {
			errCh <- err
		}
		answerCh <- answer
	}()

timerExpired:
	for i := 0; i < len(q.Problems); i++ {
		fmt.Fprintf(os.Stdout, "Problem #%d: %s = ", i+1, q.Problems[i].Question)

		select {
		case <-timer.C:
			break timerExpired
		case answer := <-answerCh:
			if q.Problems[i].IsCorrect(answer) {
				q.Correct++
			}
		case err := <-errCh:
			return err
		}
	}

	fmt.Fprintf(os.Stdout, "\nYou scored %d out of %d.\n", q.Correct, len(q.Problems))
	return nil
}
