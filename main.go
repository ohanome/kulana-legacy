package main

import (
	"fmt"
	"kulana/l"
	_ "kulana/l"
	"kulana/options"
)

// Main entrypoint
// TODO: Install and use viper (go get github.com/spf13/viper).
func main() {
	l.Notice("Starting new from scratch...")
	o := options.Parse()
	l.Debug(1, fmt.Sprintf("Full: %v\n", o.Full))
	l.Debug(1, fmt.Sprintf("Verbosity level: %d\n", len(o.Verbose)))
}
