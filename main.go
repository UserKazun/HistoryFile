package main

import (
	"os"

	"github.com/UserKazun/HistoryFile/cli"
)

func main() {
	cli := cli.NewCli(os.Stdout, os.Stderr)
	os.Exit(cli.Run(os.Args))
}
