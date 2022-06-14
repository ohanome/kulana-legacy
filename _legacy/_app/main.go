package _app

import (
	"fmt"
	"kulana/_email"
	"kulana/_f"
	"kulana/_filter"
	"kulana/_misc"
	"kulana/_ping"
	"kulana/_template"
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
	Filter _filter.OutputFilter

	// The delay in which requests will be resent.
	Delay int64

	// The url to fetch.
	Url                  string
	Host                 string
	Port                 int
	NotifyMailTo         string
	NotifyViaMail        bool
	Timeout              int
	Protocol             string
	SkipFilterValidation bool
	ForeignID            string
}

// Build a default app.
var defaultApp = Application{
	Command:        CommandHelp,
	OutputFormat:   "",
	RunInLoop:      false,
	FollowRedirect: false,
	Filter: _filter.OutputFilter{
		Url:           true,
		Time:          true,
		Status:        true,
		Destination:   true,
		ContentLength: false,
		IpAddress:     true,
		MXRecords:     true,
		ICMPCode:      true,
		Hostname:      true,
		Port:          true,
		Content:       false,
		ForeignID:     false,
	},
	Delay:                1000,
	Url:                  "",
	Host:                 "",
	Port:                 -1,
	NotifyMailTo:         "",
	NotifyViaMail:        false,
	Timeout:              _ping.Timeout,
	Protocol:             _ping.DefaultProtocol,
	SkipFilterValidation: false,
	ForeignID:            "",
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
		_misc.Usage(CommandHelp)
		os.Exit(1)
	}

	cmd := os.Args[1]
	switch cmd {
	case CommandHelp:
		_misc.Usage(CommandHelp)
		os.Exit(0)
	case CommandHelpShort:
		_misc.Usage(CommandHelp)
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

	options := _f.Parse()
	if options.Help {
		_misc.Usage(app.Command)
		os.Exit(0)
	}

	if options.Json {
		app.OutputFormat = _template.FormatJSON
	}

	if options.Csv {
		app.OutputFormat = _template.FormatCSV
	}

	if !options.Json && !options.Csv && options.Format != "" {
		switch options.Format {
		case "json":
			app.OutputFormat = _template.FormatJSON
			break

		case "csv":
			app.OutputFormat = _template.FormatCSV
			break
		}
	}

	if options.Full {
		app.Filter = _filter.GetDefault("all")
	}

	if len(os.Args) >= 2 {
		for _, arg := range os.Args[2:] {
			switch arg {

			case "--loop":
				app.RunInLoop = true
				if app.FollowRedirect {
					_misc.Die("Cannot follow redirects in a loop")
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
					_misc.Die("Cannot follow redirects in a loop")
				}
				break

			case "--status-only":
				app.Filter = _filter.BuildFilterFromNumeric(_filter.Status)
				app.SkipFilterValidation = true
				break

			case "--url-only":
				app.Filter = _filter.BuildFilterFromNumeric(_filter.Url)
				app.SkipFilterValidation = true
				break

			case "--time-only":
				app.Filter = _filter.BuildFilterFromNumeric(_filter.Time)
				app.SkipFilterValidation = true
				break

			case "--check-env":
				emailSetup := _email.CheckMailEnvironment(false)

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
						_misc.Check(iErr)

						if partValue > 255 {
							_misc.Die("IP address parts cannot have a value higher than 255.")
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
					_misc.Check(pErr)

					if portValue > 65535 {
						_misc.Die("The port cannot be bigger than 65535.")
					}
					app.Port = portValue
				}

				foreignIdMatch, _ := regexp.Match(`^--foreign-id=.+$`, []byte(arg))
				if foreignIdMatch {
					foreignId := strings.ReplaceAll(arg, "--foreign-id=", "")
					app.ForeignID = foreignId
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
	if !app.SkipFilterValidation {
		app.Filter = _filter.GetDefault(app.Command)
	}

	if app.Command == CommandStatus {
		if app.Url == "" {
			_misc.Die("No URL given.")
		}
	}

	if app.Command == CommandPing {
		if app.Host == "" {
			_misc.Die("No host given.")
		}
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
