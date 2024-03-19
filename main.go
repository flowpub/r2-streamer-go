package main

import (
	"os"

	"github.com/flowpub/r2-streamer-go/cli"
)

const UsageString = "Usage: r2-streamer-go [--server|-s|<filename>]"

func main() {
	if len(os.Args) > 1 {
		arg1 := os.Args[1]
		if arg1 == "--server" || arg1 == "-s" {
			cli.RunServer()
		} else {
			cli.RunParse(arg1)
		}
	}
}
