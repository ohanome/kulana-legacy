package main

import (
	_ "kulana/command"
	_ "kulana/l"
	"kulana/misc"
	"kulana/setup"
)

// Main entrypoint
func main() {
	misc.Welcome()
	setup.EnsureEnvironmentIsReady()
}
