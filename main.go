package main

import (
	"os"

	"github.com/eugenetriguba/quiz-cli/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[0], os.Args[1:]))
}
