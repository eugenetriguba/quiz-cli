package main

import (
	"os"

	"github.com/eugenetriguba/quiz/cli"
)

func main() {
	os.Exit(cli.Run(os.Args))
}
