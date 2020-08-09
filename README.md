# Quiz CLI

Quiz CLI is a toy program for the purposes of getting more familiar
with various Go standard libraries and practicing/reinforcing TDD.

It reads in problems from the `problems.csv` file. That file also
contains the answers as well. Each problem is prompted to the user
at the CLI and they may answer them, until all problems have been
answered. Then, a score total is shown to the user. The user has by
default 30 seconds to answer each question, otherwise the program
will exit. This time limit and the csv file to retrieve problems from
may both be customized with a command line flag.