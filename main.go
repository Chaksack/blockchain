package main

import (
	"os"

	"github.com/Chaksack/cli"
)

func main() {
	defer os.Exit(0)

	cmd := cli.CommandLine{}
	cmd.Run()
}
