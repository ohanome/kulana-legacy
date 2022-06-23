package command

import (
	"github.com/jessevdk/go-flags"
	"os"
)

type DefaultOptions struct {
	Verbose              []bool   `short:"v" long:"verbose" description:"Adds verbosity levels"`
	Format               string   `long:"format" description:"Sets the output format. Allowed values: json, csv" choice:"json" choice:"csv" value-name:"FORMAT"`
	Json                 bool     `long:"json" description:"Sets the output format to JSON"`
	Csv                  bool     `long:"csv" description:"Sets the output format to CSV"`
	RestoreDefaultConfig bool     `long:"restore-default-config" description:"Restores the default configuration"`
	Hostname             string   `long:"host" description:"The hostname" value-name:"HOSTNAME"`
	Port                 int      `long:"port" description:"The port" value-name:"PORT"`
	Delay                int      `long:"delay" description:"Sets the delay between looping action in milliseconds" value-name:"DELAY"`
	Loop                 bool     `long:"loop" description:"Lets some (not all) action run in a loop"`
	FollowRedirect       bool     `long:"follow-redirect" description:"Sends another request if the status code is a redirect code (3xx)"`
	Notify               bool     `long:"notify" description:"Sends the result as mail. --notify-mails is required"`
	NotifyMails          []string `long:"notify-mails" description:"Sets the receiving mail address for the --notify option."`
	ForeignId            string   `long:"foreign-id" description:"Sets the foreign id"`
	RedirectOutput       bool     `long:"redirect-output" description:"Send the output to another destination than the logs, for example a file or an API"`
	NoColor              bool     `long:"no-color" description:"Disables color output, affects only default format"`
}

var defaultOptions DefaultOptions
var parser *flags.Parser

func init() {
	parser = flags.NewParser(&defaultOptions, flags.Default)

	_, err := parser.AddCommand("add",
		"Add a file",
		"The add command adds a file to the repository. Use -a to add all files.",
		&addCommand)
	if err != nil {
		panic(err)
	}

	_, err = parser.AddCommand("status",
		"Fetches the status of a URL",
		"Fetches the status of a URL",
		&statusCommand)
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

func GetDefaultOptions() DefaultOptions {
	return defaultOptions
}

func GetParser() *flags.Parser {
	return parser
}

func VerbosityLevel() int {
	return len(defaultOptions.Verbose)
}

func GetFormat() string {
	return defaultOptions.Format
}

func GetJson() bool {
	return defaultOptions.Json
}

func GetCsv() bool {
	return defaultOptions.Csv
}

func GetRestoreDefaultConfig() bool {
	return defaultOptions.RestoreDefaultConfig
}

func GetHostname() string {
	return defaultOptions.Hostname
}

func GetPort() int {
	return defaultOptions.Port
}

func GetDelay() int {
	return defaultOptions.Delay
}

func GetLoop() bool {
	return defaultOptions.Loop
}

func GetFollowRedirect() bool {
	return defaultOptions.FollowRedirect
}

func GetNotify() bool {
	return defaultOptions.Notify
}

func GetNotifyMails() []string {
	return defaultOptions.NotifyMails
}

func GetForeignId() string {
	return defaultOptions.ForeignId
}

func GetRedirectOutput() bool {
	return defaultOptions.RedirectOutput
}

func GetNoColor() bool {
	return defaultOptions.NoColor
}
