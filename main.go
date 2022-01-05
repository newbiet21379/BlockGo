package main

import (
	"github.com/newbiet21379/blockchain/cli"
	"os"
)

func main() {
	defer os.Exit(0)
	commandLine := cli.CommandLine{}
	commandLine.Run()
}
