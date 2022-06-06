package app

import (
	"fmt"
	"kulana/email"
	"kulana/filter"
	"kulana/misc"
	"kulana/ping"
	"kulana/template"
	"net/mail"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Application struct {
	// The command to execute.
	// Every command has its own go file placed in /app and implements a function called "runX" where X is the command.
	Command string

	// The output format.
	// Supported formats must have a template configured in /template.
	OutputFormat string

	// True if the app should run in a loop.
	RunInLoop bool

	// True if status requests should follow redirects.
	// This will send another request and output its stats.
	FollowRedirect bool

	// Filters the output.
	// Every value set to true will be kept, other values will be removed.
	Filter filter.OutputFilter

	// The delay in which requests will be resent.
	Delay int64

	// The url to fetch.
	Url           string
	Host          string
	Port          int
	NotifyMailTo  string
	NotifyViaMail bool
	Timeout       int
	Protocol      string
}

// Build a default app.
var defaultApp = Application{
	Command:        CommandHelp,
	OutputFormat:   "",
	RunInLoop:      false,
	FollowRedirect: false,
	Filter: filter.OutputFilter{
		Url:           true,
		Time:          true,
		Status:        true,
		Destination:   true,
		ContentLength: false,
	},
	Delay:         1000,
	Url:           "",
	Host:          "",
	Port:          80,
	NotifyMailTo:  "",
	NotifyViaMail: false,
	Timeout:       ping.Timeout,
	Protocol:      ping.DefaultProtocol,
}

const CommandHelp = "help"
const CommandHelpShort = "h"
const CommandStatus = "status"
const CommandStatusShort = "s"
const CommandConfig = "config"
const CommandConfigShort = "c"
const CommandPing = "ping"
const CommandPingShort = "p"
const CommandMX = "mx"

func ProcessArgs() Application {
	app := defaultApp

	if len(os.Args) <= 2 {
		misc.Usage(CommandHelp)
		os.Exit(1)
	}

	cmd := os.Args[1]
	switch cmd {
	case CommandHelp:
		misc.Usage(CommandHelp)
		os.Exit(0)
	case CommandHelpShort:
		misc.Usage(CommandHelp)
		os.Exit(0)

	case CommandStatus:
		app.Command = CommandStatus
		break
	case CommandStatusShort:
		app.Command = CommandStatus
		break
	case CommandConfig:
		app.Command = CommandConfig
		break
	case CommandConfigShort:
		app.Command = CommandConfig
		break
	case CommandPing:
		app.Command = CommandPing
		break
	case CommandPingShort:
		app.Command = CommandPing
		break
	case CommandMX:
		app.Command = CommandMX
		break

	default:
		fmt.Println("Unknown command.")
		os.Exit(1)
	}

	if len(os.Args) >= 2 {
		for _, arg := range os.Args[2:] {
			switch arg {
			case "--help":
			case "-h":
				misc.Usage(CommandStatus)
				os.Exit(0)

			case "--json":
				app.OutputFormat = template.FormatJSON
				break

			case "--csv":
				app.OutputFormat = template.FormatCSV
				break

			case "--loop":
				app.RunInLoop = true
				if app.FollowRedirect {
					misc.Die("Cannot follow redirects in a loop")
				}
				break

			case "--include-length":
			case "-l":
				app.Filter.ContentLength = true
				break

			case "--follow-redirect":
			case "-f":
				app.FollowRedirect = true
				if app.RunInLoop {
					misc.Die("Cannot follow redirects in a loop")
				}
				break

			case "--status-only":
				app.Filter = filter.BuildFilterFromNumeric(filter.Status)
				break

			case "--url-only":
				app.Filter = filter.BuildFilterFromNumeric(filter.Url)
				break

			case "--time-only":
				app.Filter = filter.BuildFilterFromNumeric(filter.Time)
				break

			case "--check-env":
				emailSetup := email.CheckMailEnvironment(false)

				if emailSetup {
					fmt.Println("[OK] Email setup")
				}

				os.Exit(0)

			case "--notify":
			case "-n":
				app.NotifyViaMail = true
				break

			default:
				delayMatch, _ := regexp.Match(`^--delay=\d.+$`, []byte(arg))
				if delayMatch {
					d := strings.ReplaceAll(arg, "--delay=", "")
					dInt, _ := strconv.ParseInt(d, 10, 64)
					app.Delay = dInt * 1000000
				}

				urlMatch, _ := regexp.Match(`^http(s|)://\w.+\.\w{2,3}$`, []byte(arg))
				if urlMatch {
					app.Url = arg
				}

				mailMatch, _ := regexp.Match(`^--notify-mail=.+$`, []byte(arg))
				if mailMatch {
					m := strings.ReplaceAll(arg, "--notify-mail=", "")
					_, mErr := mail.ParseAddress(m)
					if mErr == nil {
						app.NotifyMailTo = m
					}
				}

				hostIPMatch, _ := regexp.Match(`^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`, []byte(arg))
				if hostIPMatch {
					parts := strings.Split(arg, ".")
					for _, p := range parts {
						partValue, iErr := strconv.Atoi(p)
						misc.Check(iErr)

						if partValue > 255 {
							misc.Die("IP address parts cannot have a value higher than 255.")
						}
					}
					app.Host = arg
				}

				hostMatch, _ := regexp.Match(`^\w.+\.\w{2,3}$`, []byte(arg))
				if hostMatch {
					app.Host = arg
				}

				portMatch, _ := regexp.Match(`^\d{1,5}$`, []byte(arg))
				if portMatch {
					portValue, pErr := strconv.Atoi(arg)
					misc.Check(pErr)

					if portValue > 65535 {
						misc.Die("The port cannot be bigger than 65535.")
					}
					app.Port = portValue
				}
				break
			}
		}
	}

	app = validate(app)

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

	case CommandPing:
		runPing(application)
		break

	case CommandMX:
		runMX(application)
		break
	}
}

func validate(app Application) Application {
	if app.Command == CommandStatus {
		if app.Url == "" {
			misc.Die("No URL given.")
		}
	}

	if app.Command == CommandPing {
		if app.Host == "" {
			misc.Die("No host given.")
		}

		app.Filter.IpAddress = true
	}

	if app.NotifyViaMail && app.NotifyMailTo == "" {
		fmt.Println("Email address is missing, no email will be sent.")
		app.NotifyViaMail = false
	}

	if app.Delay < 100 {
		app.Delay = 100
	}

	if app.NotifyViaMail && app.Delay < 30000 {
		app.Delay = 30000
		fmt.Println("When sending notification emails, the minimum delay is 30000 (30 seconds).")
	}

	return app
}
