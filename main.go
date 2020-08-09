package main

import (
	"fmt"
	"os"

	"github.com/eugenetriguba/quiz-cli/cli"
)

func main() {
	cl := cli.NewCommandLine()
	_, err := cl.Parse(os.Args[0], os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not parse flags because '%v'\n", err)
		os.Exit(1)
	}

	exitCode := cl.Run()
	os.Exit(exitCode)
}
