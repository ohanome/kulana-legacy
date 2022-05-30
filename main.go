package main

import (
	"kulana/app"
	"kulana/setup"
)

// Main entrypoint
func main() {
	application := app.ProcessArgs()
	setup.EnsureEnvironmentIsReady()

	app.Run(application)
}
