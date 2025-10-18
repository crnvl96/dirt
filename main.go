// Package main is the entry point for the dirt CLI tool.
package main

import "github.com/crnvl96/dirt/internal"

// main runs the dirt CLI application.
//
// dirt is a cli tool to check if you have uncommited changes in any of your repos

func main() {
	internal.Execute()
}
