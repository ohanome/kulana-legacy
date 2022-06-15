package options

import (
	"github.com/jessevdk/go-flags"
	"os"
)

type Options struct {
	Verbose []bool `short:"v" long:"verbose" description:"Adds verbosity levels"`
	Full    bool   `long:"full" description:"Outputs everything, also empty values"`
	Format  string `long:"format" description:"Sets the output format. Allowed values: json, csv" choice:"json" choice:"csv"`
	Json    bool   `long:"json" description:"Sets the output format to JSON"`
	Csv     bool   `long:"csv" description:"Sets the output format to CSV"`
}

func Parse() (Options, error) {
	var opts Options
	_, err := flags.Parse(&opts)
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

	return opts, nil
}

func VerbosityLevel() int {
	o, _ := Parse()
	return len(o.Verbose)
}
