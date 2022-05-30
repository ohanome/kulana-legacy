package app

import (
	"fmt"
	"kulana/fetcher"
	"kulana/misc"
	"kulana/template"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Application struct {
	OutputFormat   string
	RunInLoop      bool
	FollowRedirect bool
	Filter         fetcher.ResponseFilter
	Delay          int64
	Url            string
}

// Build a default app.
var defaultApp = Application{
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
	Delay: 1000,
	Url:   "",
}

func ProcessArgs() Application {
	app := defaultApp
	if len(os.Args) == 1 {
		misc.Usage()
		os.Exit(1)
	}

	if len(os.Args) >= 1 {
		for _, arg := range os.Args[1:] {
			switch arg {
			case "--help":
			case "-h":
				misc.Usage()
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
				app.Filter = fetcher.BuildFilterFromNumeric(fetcher.FilterStatus)
				break

			case "--url-only":
				app.Filter = fetcher.BuildFilterFromNumeric(fetcher.FilterUrl)
				break

			case "--time-only":
				app.Filter = fetcher.BuildFilterFromNumeric(fetcher.FilterTime)
				break

			default:
				delayMatch, _ := regexp.Match(`^--delay=\d.*$`, []byte(arg))
				if delayMatch {
					d := strings.ReplaceAll(arg, "--delay=", "")
					dInt, _ := strconv.ParseInt(d, 10, 64)
					app.Delay = dInt * 1000000
				}

				urlMatch, _ := regexp.Match(`^http(s|)://\w.*\.\w{2,3}$`, []byte(arg))
				if urlMatch {
					app.Url = arg
				}
				break
			}
		}

		if app.Url == "" {
			misc.Die("No URL given.")
		}
	}

	// Correct some values if they're wrong.
	if app.Delay < 100 {
		app.Delay = 100
	}

	return app
}

func Run(application Application) {
	for {
		response, filteredResponse := fetcher.FetchAndFilter(application.Url, application.Filter)
		fmt.Print(template.Render(application.OutputFormat, filteredResponse) + misc.GetNLChar())
		if application.FollowRedirect && !application.RunInLoop && response.Status < 400 && response.Status >= 300 {
			application.Url = response.Destination
			response, filteredResponse = fetcher.FetchAndFilter(application.Url, application.Filter)
			fmt.Print(template.Render(application.OutputFormat, filteredResponse) + misc.GetNLChar())
		}

		if !application.RunInLoop {
			os.Exit(0)
		}
		time.Sleep(time.Duration(application.Delay))
	}
}
