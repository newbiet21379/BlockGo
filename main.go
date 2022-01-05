package main

import (
	"github.com/newbiet21379/blockgo/cli"
	"os"
)

func main() {
	defer os.Exit(0)
	commandLine := cli.CommandLine{}
	commandLine.Run()

	//w := wallet.MakeWallet()
	//w.Address()
}
