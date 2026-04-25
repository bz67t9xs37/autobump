package main

import (
	"fmt"
	"os"

	"github.com/rios0rios0/autobump/internal/cli"
)

// main is the entry point for autobump.
// autobump is a tool that automatically bumps version numbers in your project
// based on conventional commit messages.
func main() {
	if err := cli.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
