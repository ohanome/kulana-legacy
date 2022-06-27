package command

import (
	"github.com/jessevdk/go-flags"
	"os"
)

type DefaultOptions struct {
	Verbose              []bool `short:"v" long:"verbose" description:"Adds verbosity levels"`
	Format               string `long:"format" description:"Sets the output format. Allowed values: json, csv" choice:"json" choice:"csv" value-name:"FORMAT"`
	Json                 bool   `long:"json" description:"Sets the output format to JSON"`
	Csv                  bool   `long:"csv" description:"Sets the output format to CSV"`
	RestoreDefaultConfig bool   `long:"restore-default-config" description:"Restores the default configuration"`
	Delay                int    `long:"delay" description:"Sets the delay between looping action in milliseconds" value-name:"DELAY"`
	Loop                 bool   `long:"loop" description:"Lets some (not all) action run in a loop"`
	ForeignId            string `long:"foreign-id" description:"Sets the foreign id"`
	NoColor              bool   `long:"no-color" description:"Disables color output, affects only default format"`
}

var defaultOptions DefaultOptions
var parser *flags.Parser

func init() {
	parser = flags.NewParser(&defaultOptions, flags.Default)

	_, err := parser.AddCommand("status",
		"Fetches the status of a URL",
		"Fetches the status of a URL",
		&statusCommand)
	if err != nil {
		panic(err)
	}

	_, err = parser.AddCommand("ping",
		"Pings a given host",
		"The ping command pings a given host",
		&pingCommand)
	if err != nil {
		panic(err)
	}

	_, err = parser.AddCommand("mx",
		"Fetches the MX records of a given host",
		"Fetches the MX records of a given host",
		&mxCommand)
	if err != nil {
		panic(err)
	}

	_, err = parser.AddCommand("cname",
		"Fetches the CNAME record of a given host",
		"Feeds the CNAME record of a given host",
		&cnameCommand)
	if err != nil {
		panic(err)
	}

	_, err = parser.AddCommand("cert",
		"Fetches SSL certificate information of a given host",
		"Fetches SSL certificate information of a given host",
		&certCommand)
	if err != nil {
		panic(err)
	}

	_, err = parser.Parse()
	if err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}
}

func SetFormat() {
	if defaultOptions.Format == "" {
		if defaultOptions.Json {
			defaultOptions.Format = "json"
		} else if defaultOptions.Csv {
			defaultOptions.Format = "csv"
		} else {
			defaultOptions.Format = "default"
		}
	}

	if defaultOptions.Format == "json" {
		defaultOptions.Json = true
	} else if defaultOptions.Format == "csv" {
		defaultOptions.Csv = true
	}
}

func VerbosityLevel() int {
	return len(defaultOptions.Verbose)
}
