package main

import (
	"os"

	"github.com/xxmyjk/vcs-test/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
