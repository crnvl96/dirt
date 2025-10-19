// Package main provides the entry point for the Dirt CLI tool.
//
// Dirt is a command-line utility that scans specified directories for Git repositories
// and reports any that have uncommitted changes or unpushed commits.
package main

import "github.com/crnvl96/dirt/internal"

func main() {
	internal.Execute()
}
