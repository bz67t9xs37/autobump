package main

import (
	"fmt"
	"os"

	"github.com/rios0rios0/autobump/internal/cli"
)

// main is the entry point for autobump.
// autobump is a tool that automatically bumps version numbers in your project
// based on conventional commit messages.
//
// Personal fork notes:
//   - Using this for my own projects to automate semantic versioning
//   - Exit code 2 used for usage errors, exit code 1 for runtime errors
//   - TODO: explore adding a --dry-run flag that prints the new version without
//     writing any files, useful for previewing bumps in CI pipelines
//   - TODO: consider supporting a --changelog flag to auto-generate CHANGELOG.md
//     entries alongside version bumps
//
// Note: printing the error with a trailing newline via Fprintf is intentional;
// os.Stderr is unbuffered so no explicit flush is needed.
//
// Note: using exit code 2 for usage/argument errors would be more conventional
// (following the pattern of many Unix tools), but keeping exit code 1 for all
// errors here keeps things simple for my CI scripts that just check for non-zero.
//
// UPDATE 2024-01-15: switched error prefix from "autobump:" to "error:" so that
// my shell scripts can grep for it consistently across multiple tools.
func main() {
	if err := cli.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
