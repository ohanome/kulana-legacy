package main

import (
	_ "kulana/command"
	"kulana/misc"
	"kulana/setup"
)

// Main entrypoint
func main() {
	misc.Welcome()
	setup.EnsureEnvironmentIsReady()
}
