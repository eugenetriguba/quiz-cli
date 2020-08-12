# Quiz CLI

Quiz CLI is a toy program created for practice and getting
more familiar with various Go standard libraries.

It reads in problems from the `problems.csv` file. That file also
contains the answers. Each problem is prompted to the user
at the CLI and they may answer them, until all problems have been
answered. Then, a score total is shown to the user. The user has by
default a 30 second time limit to answer each question, otherwise the
program will exit. This time limit and the csv file to retrieve problems
from may both be customized with a command line flag.

## Usage

To build the program, we can run `go build .`. We should then see a
`quiz-cli` executable file in our current directory.

To see what flags we can use, we can run `./quiz-cli -h`.
```
Usage of ./quiz-cli:
  -csv string
        Set problems file (default "/home/eugene/Code/quiz-cli/problems.csv")
  -limit duration
        Set the question time limit (default 30s)
  -shuffle
        Shuffle the problems presented
```

The `default` for `-csv` is generated automatically based on wherever the `quiz-cli`
package is. It'll also automatically convert your `-csv` argument to an absolute
path for you based on the current directory you're in, so there is no need to specify
a full absolute path like that default shows.

To run the program, we can run `./quiz-cli`.
```
Problem #1: 5+5 = 10
Problem #2: 7+3 = 10
Problem #3: 1+1 = 2
Problem #4: 8+3 = 11
Problem #5: 1+2 = 3
Problem #6: 8+6 = 14
Problem #7: 3+1 = 4
Problem #8: 1+4 = 5
Problem #9: 5+1 = 6
Problem #10: 2+3 = 5
Problem #11: 3+3 = 6
Problem #12: 2+4 = 6
Problem #13: 5+2 = 7

You scored 13 out of 13.
```
