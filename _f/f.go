package _f

import (
	"fmt"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
	Help    bool   `short:"h" long:"help" description:"Show help page"`
	Format  string `short:"f" long:"format" description:"Sets the output format" choice:"default" choice:"json" choice:"csv"`
	Json    bool   `long:"json" description:"Sets the output format to JSON"`
	Csv     bool   `long:"csv" description:"Sets the output format to CSV"`
	Full    bool   `long:"full" description:"Outputs everything, also empty values"`
}

func Parse() Options {
	var opts Options
	remainingArgs, err := flags.Parse(&opts)
	if err != nil {
		panic(err)
	}

	fmt.Println(remainingArgs)
	return opts
}
