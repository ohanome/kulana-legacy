package app

import (
	"fmt"
	"kulana/fetcher"
	"kulana/misc"
	"os"
)

type Application struct {
	Command        string
	OutputFormat   string
	RunInLoop      bool
	FollowRedirect bool
	Filter         fetcher.ResponseFilter
	Delay          int64
	Url            string
	NotifyMailTo   string
	NotifyViaMail  bool
}

// Build a default app.
var defaultApp = Application{
	Command:        CommandHelp,
	OutputFormat:   "",
	RunInLoop:      false,
	FollowRedirect: false,
	Filter: fetcher.ResponseFilter{
		Url:           true,
		Time:          true,
		Status:        true,
		Destination:   true,
		ContentLength: false,
	},
	Delay:         1000,
	Url:           "",
	NotifyMailTo:  "",
	NotifyViaMail: false,
}

const CommandHelp = "help"
const CommandHelpShort = "h"
const CommandStatus = "status"
const CommandStatusShort = "s"
const CommandConfig = "config"

func ProcessArgs() Application {
	app := defaultApp

	if len(os.Args) == 1 {
		misc.Usage(CommandHelp)
		os.Exit(1)
	}

	if len(os.Args) >= 1 {
		switch os.Args[1] {
		case CommandHelp:
			misc.Usage(CommandHelp)
			os.Exit(0)
		case CommandHelpShort:
			misc.Usage(CommandHelp)
			os.Exit(0)

		case CommandStatus:
			app.Command = CommandStatus
			app = processStatusArgs(app)
			break
		case CommandStatusShort:
			app.Command = CommandStatus
			app = processStatusArgs(app)
			break

		case CommandConfig:
			app.Command = CommandConfig
			app = processConfigArgs(app)
			break

		default:
			fmt.Println("Unknown command.")
			os.Exit(1)
		}
	}

	return app
}

func Run(application Application) {
	switch application.Command {
	case CommandStatus:
		runStatus(application)
		break

	case CommandConfig:
		runConfig(application)
		break
	}
}
