package main

import (
	"fmt"
	"kulana/config"
	"kulana/l"
	_ "kulana/l"
	"kulana/options"
	"kulana/setup"
)

// Main entrypoint
func main() {
	setup.EnsureEnvironmentIsReady()
	l.Notice("Starting new from scratch...")
	o, _ := options.Parse()
	l.Debug(3, fmt.Sprintf("Parsed arguments: %##v", o))

	mail := config.Get("mail.subject")
	l.Debug(1, fmt.Sprintf("mail: %v", mail))
	str := fmt.Sprintf("%v", mail)
	mail = str + "1"
	config.Set("mail.subject", mail)
	mail = config.Get("mail.subject")
	l.Debug(1, fmt.Sprintf("mail: %v", mail))

}
