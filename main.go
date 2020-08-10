package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/eugenetriguba/quiz-cli/cli"
)

func main() {
	cl := cli.NewCommandLine()
	output, err := cl.Parse(os.Args[0], os.Args[1:])
	if err == flag.ErrHelp {
		fmt.Println(output)
		os.Exit(2)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not parse flags because '%v'\n", err)
		os.Exit(1)
	}

	os.Exit(cl.Run())
}
