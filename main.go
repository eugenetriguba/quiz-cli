package main

import (
	"os"

	"github.com/eugenetriguba/quiz-cli/cli"
)

func main() {
	cl := cli.NewCommandLine()
	cl.Parse(os.Args[0], os.Args[1:])

	os.Exit(cl.Run())
}
