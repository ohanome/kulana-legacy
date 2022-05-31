package main

import (
	"kulana/app"
	"kulana/setup"
)

// Main entrypoint
func main() {
	setup.EnsureEnvironmentIsReady()
	application := app.ProcessArgs()

	app.Run(application)
}
