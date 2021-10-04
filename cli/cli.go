package cli

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	ExitCodeOK             = 0
	ExitCodeParseFlagError = 1
	ExitCodeFail           = 1
)

type CLI struct {
	outStream, errStream io.Writer
}

func NewCli(outStream, errStream io.Writer) *CLI {
	return &CLI{outStream: outStream, errStream: errStream}
}

func (c *CLI) Run(args []string) int {
	//var filename string

	flags := flag.NewFlagSet("fh", flag.ExitOnError)
	flags.SetOutput(c.errStream)

	filename := flags.String("filename", "", "Allowed extensions: .csv")

	fmt.Println(os.Args[1])
	if len(os.Args) < 2 {
		fmt.Println("The file name must be specified.")
		os.Exit(1)
	} else {
		flags.Parse(os.Args[2:])
		fmt.Println("filename:", *filename)
		os.Exit(1)
	}

	return ExitCodeOK
}
