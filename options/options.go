package options

import (
	"github.com/jessevdk/go-flags"
	"kulana/command"
	"os"
)

type Options struct {
	Verbose              []bool   `short:"v" long:"verbose" description:"Adds verbosity levels"`
	Full                 bool     `long:"full" description:"Outputs everything, also empty values"`
	Format               string   `long:"format" description:"Sets the output format. Allowed values: json, csv" choice:"json" choice:"csv" value-name:"FORMAT"`
	Json                 bool     `long:"json" description:"Sets the output format to JSON"`
	Csv                  bool     `long:"csv" description:"Sets the output format to CSV"`
	RestoreDefaultConfig bool     `long:"restore-default-config" description:"Restores the default configuration"`
	Url                  string   `short:"u" long:"url" description:"The URL to get te status from" value-name:"URL"`
	Hostname             string   `long:"host" description:"The hostname" value-name:"HOSTNAME"`
	Port                 int      `long:"port" description:"The port" value-name:"PORT"`
	Delay                int      `long:"delay" description:"Sets the delay between looping action in milliseconds" value-name:"DELAY"`
	Loop                 bool     `long:"loop" description:"Lets some (not all) action run in a loop"`
	Include              []string `long:"include" description:"Includes specific fields" choice:"url" choice:"status" choice:"time" choice:"destination" choice:"content_length" choice:"ip_address" choice:"mx_records" choice:"icmp_code" choice:"hostname" choice:"port" choice:"content" choice:"foreign_id"`
	FollowRedirect       bool     `long:"follow-redirect" description:"Sends another request if the status code is a redirect code (3xx)"`
	Notify               bool     `long:"notify" description:"Sends the result as mail. --notify-mails is required"`
	NotifyMails          []string `long:"notify-mails" description:"Sets the receiving mail address for the --notify option."`
	ForeignId            string   `long:"foreign-id" description:"Sets the foreign id"`
	RedirectOutput       bool     `long:"redirect-output" description:"Send the output to another destination than the logs, for example a file or an API"`
	PrintLogs            bool     `long:"print-logs" description:"Prints the logs into stdout (kulana.log will still be written)"`
	NoColor              bool     `long:"no-color" description:"Disables color output, affects only default format"`
}

var options Options

var parser = flags.NewParser(&options, flags.Default)

func GetParser() *flags.Parser {
	addCommand := command.AddCommand{}
	pingCommand := command.PingCommand{}

	parser.AddCommand("add",
		"Add a file",
		"The add command adds a file to the repository. Use -a to add all files.",
		&addCommand)
	parser.AddCommand("ping",
		"Add a file",
		"The add command adds a file to the repository. Use -a to add all files.",
		&pingCommand)
	return parser
}

func Parse() (Options, []string, error) {
	var opts Options
	remaining, err := GetParser().Parse()
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

	if opts.Csv && opts.Json {
		opts.Csv = false
	}

	if opts.Json {
		opts.Format = "json"
	}

	if opts.Csv {
		opts.Format = "csv"
	}

	if !opts.Json && !opts.Csv {
		switch opts.Format {
		case "json":
			opts.Json = true
			break
		case "csv":
			opts.Csv = true
		}
	}

	return opts, remaining, nil
}

func VerbosityLevel() int {
	o, _, _ := Parse()
	return len(o.Verbose)
}
