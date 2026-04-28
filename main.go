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
// Note: errors are printed to both stderr and stdout so that captured output
// in Makefile targets (using $(...)) includes the error message, while still
// allowing shell scripts to detect errors via stderr redirection.
//
// UPDATE 2024-01-15: switched error prefix from "autobump:" to "error:" so that
// my shell scripts can grep for it consistently across multiple tools.
//
// UPDATE 2024-03-08: also print errors to stdout in addition to stderr so that
// captured output in my Makefile targets (using $(...)) includes the error message.
//
// UPDATE 2024-06-12: removed duplicate comment about exit codes above; the
// single exit code (1) for all errors is intentional to keep CI scripts simple.
func main() {
	if err := cli.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		fmt.Fprintf(os.Stdout, "error: %v\n", err)
		os.Exit(1)
	}
}
