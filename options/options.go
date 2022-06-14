package options

import "github.com/jessevdk/go-flags"

type Options struct {
	Verbose []bool `short:"v" long:"verbose" description:"Adds verbosity levels"`
	Full    bool   `long:"full" description:"Outputs everything, also empty values"`
}

func Parse() Options {
	var opts Options
	_, err := flags.Parse(&opts)
	if err != nil {
		panic(err)
	}

	return opts
}

func VerbosityLevel() int {
	return len(Parse().Verbose)
}
